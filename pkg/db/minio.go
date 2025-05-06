package db

import (
	"bytes"
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioStore struct {
	endpoint string // ip:port
	client   *minio.Client
	ak       string
	sk       string
	useSsl   bool
}

type Option func(*MinioStore)

func WithCredentials(ak, sk string) Option {
	return func(m *MinioStore) {
		m.ak = ak
		m.sk = sk
	}
}

func WithSecure(useSsl bool) Option {
	return func(m *MinioStore) {
		m.useSsl = useSsl
	}
}

func NewMinioStore(endpoint string, opts ...Option) (*MinioStore, error) {
	store := &MinioStore{
		endpoint: endpoint,
	}

	for _, opt := range opts {
		opt(store)
	}

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(store.ak, store.sk, ""),
		Secure: store.useSsl,
	})
	if err != nil {
		return nil, err
	}

	store.client = client

	return store, nil
}

// CreateBucket creates a bucket if it does not exist.
func (m *MinioStore) CreateBucket(ctx context.Context, bucket string) error {
	exists, err := m.client.BucketExists(ctx, bucket)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	return m.client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
}

// Upload uploads a local file to MinIO.
func (m *MinioStore) Upload(ctx context.Context, bucket, objectName, filePath string, contentType string) error {
	_, err := m.client.FPutObject(ctx, bucket, objectName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	return err
}

// UploadBytes uploads a byte array to MinIO.
func (m *MinioStore) UploadBytes(ctx context.Context, bucket, objectName string, data []byte, contentType string) error {
	// Convert byte array to io.Reader
	reader := bytes.NewReader(data)

	// Get the size of the data
	size := int64(len(data))

	// Upload data to MinIO
	_, err := m.client.PutObject(ctx, bucket, objectName, reader, size, minio.PutObjectOptions{
		ContentType: contentType,
	})

	return err
}

// Download downloads an object from MinIO and saves it to a local file.
func (m *MinioStore) Download(ctx context.Context, bucket, objectName, filePath string) error {
	return m.client.FGetObject(ctx, bucket, objectName, filePath, minio.GetObjectOptions{})
}

// DownloadBytes downloads an object from MinIO and returns it as a byte array.
func (m *MinioStore) DownloadBytes(ctx context.Context, bucket, objectName string) ([]byte, error) {
	obj, err := m.client.GetObject(ctx, bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer obj.Close()

	// Read all data into a byte array
	return io.ReadAll(obj)
}
