package aesgcm

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
)

const defaultChunkSize = 128 * 1024

// Writer encrypts data in fixed-size chunks with AES-GCM
// and writes framed ciphertext to the underlying io.Writer.
type Writer struct {
	dst       io.Writer
	aead      cipher.AEAD
	chunkSize int
	sizeBuf   [4]byte
}

// NewWriter returns a Writer that encrypts to dst.
// chunkSize is the max plaintext bytes per frame.
func NewWriter(dst io.Writer, key []byte) (*Writer, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES-GCM AEAD: %w", err)
	}
	return &Writer{dst: dst, aead: aead, chunkSize: defaultChunkSize}, nil
}

// Write encrypts p in chunks and writes each frame as
// [4-byte big-endian length][nonce|ciphertext].
func (w *Writer) Write(p []byte) (int, error) {
	total, off := len(p), 0
	for off < total {
		n := total - off
		n = min(n, w.chunkSize)
		nonce := make([]byte, w.aead.NonceSize())
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			return 0, err
		}
		ct := w.aead.Seal(nil, nonce, p[off:off+n], nil)

		// write length prefix
		length := uint32(len(nonce) + len(ct))
		binary.BigEndian.PutUint32(w.sizeBuf[:], length)
		if _, err := w.dst.Write(w.sizeBuf[:]); err != nil {
			return 0, err
		}
		// write nonce + ciphertext
		if _, err := w.dst.Write(nonce); err != nil {
			return 0, err
		}
		if _, err := w.dst.Write(ct); err != nil {
			return 0, err
		}
		off += n
	}
	return total, nil
}

func (w *Writer) Close() error { return nil }

// Reader decrypts framed ciphertext produced by Writer.
type Reader struct {
	src     io.Reader
	aead    cipher.AEAD
	sizeBuf [4]byte
	buf     []byte
	pos     int
}

// NewReader returns a Reader that decrypts from src.
func NewReader(src io.Reader, key []byte) (*Reader, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return &Reader{src: src, aead: aead}, nil
}

// Read returns decrypted plaintext, handling one frame at a time.
func (r *Reader) Read(p []byte) (int, error) {
	if r.pos < len(r.buf) {
		n := copy(p, r.buf[r.pos:])
		r.pos += n
		return n, nil
	}
	if _, err := io.ReadFull(r.src, r.sizeBuf[:]); err != nil {
		return 0, err
	}
	length := binary.BigEndian.Uint32(r.sizeBuf[:])
	frame := make([]byte, length)
	if _, err := io.ReadFull(r.src, frame); err != nil {
		return 0, err
	}
	ns := r.aead.NonceSize()
	nonce, ct := frame[:ns], frame[ns:]
	pt, err := r.aead.Open(nil, nonce, ct, nil)
	if err != nil {
		return 0, err
	}
	r.buf, r.pos = pt, 0
	return r.Read(p)
}
