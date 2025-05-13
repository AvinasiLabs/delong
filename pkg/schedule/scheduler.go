package schedule

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	dockerTypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/stdcopy"
	"golang.org/x/term"

	goarchive "github.com/moby/go-archive"
)

// AlgoScheduler handles Docker operations for algorithm execution
type AlgoScheduler struct {
	AlgoIdCh       chan uint
	dockerClient   *client.Client
	buildSizeLimit int64
}

type SchedulerOption func(*AlgoScheduler)

const (
	// DEFAULT_CHANNEL_SIZE defines the default buffer size for algorithm execution channel
	DEFAULT_CHANNEL_SIZE = 10

	// DEFAULT_BUILD_SIZE_LIMIT defines the default build context size limit (100MB)
	DEFAULT_BUILD_SIZE_LIMIT = 100 << 20
)

// WithChannelSize sets the buffer size for the algo ID channel
func WithChannelSize(size int) SchedulerOption {
	return func(s *AlgoScheduler) {
		if size > 0 {
			s.AlgoIdCh = make(chan uint, size)
		}
	}
}

func WithBuildSizeLimit(bytes int64) SchedulerOption {
	return func(s *AlgoScheduler) {
		if bytes > 0 {
			s.buildSizeLimit = bytes
		}
	}
}

// NewAlgoScheduler creates a new algo scheduler with the specified channel buffer size
func NewAlgoScheduler(opts ...SchedulerOption) (*AlgoScheduler, error) {
	dockerClient, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, err
	}
	s := &AlgoScheduler{
		dockerClient:   dockerClient,
		AlgoIdCh:       make(chan uint, DEFAULT_CHANNEL_SIZE),
		buildSizeLimit: DEFAULT_BUILD_SIZE_LIMIT,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s, nil
}

// BuildSizeLimit returns the maximum size limit for Docker image builds
func (s *AlgoScheduler) BuildSizeLimit() int64 {
	return s.buildSizeLimit
}

// BuildImage builds a Docker image from a directory containing a Dockerfile
func (s *AlgoScheduler) BuildImage(ctx context.Context, dirPath, imageName string) error {
	log.Printf("Building Docker image %s from %s", imageName, dirPath)

	// Create build context tarball using goarchive (non-deprecated)
	buildContext, err := goarchive.TarWithOptions(dirPath, &goarchive.TarOptions{})
	if err != nil {
		return fmt.Errorf("failed to create build context: %w", err)
	}
	defer buildContext.Close()

	options := dockerTypes.ImageBuildOptions{
		Tags:        []string{imageName},
		Dockerfile:  "Dockerfile",
		Remove:      true,
		ForceRemove: true,
		NoCache:     true,
	}

	resp, err := s.dockerClient.ImageBuild(ctx, buildContext, options)
	if err != nil {
		return fmt.Errorf("failed to build image: %w", err)
	}
	defer resp.Body.Close()

	// Stream build output to stdout using updated signature
	fd := os.Stdout.Fd()
	isTTY := term.IsTerminal(int(fd))
	if err := jsonmessage.DisplayJSONMessagesStream(resp.Body, os.Stdout, fd, isTTY, nil); err != nil {
		return fmt.Errorf("error streaming build output: %w", err)
	}

	log.Printf("Successfully built image: %s", imageName)
	return nil
}

// RunContainer runs a Docker container with the specified image and environment variables
func (s *AlgoScheduler) RunContainer(ctx context.Context, imageName string, env map[string]string) ([]byte, error) {
	log.Printf("Running container with image: %s", imageName)

	// Convert env map to array of KEY=VALUE strings
	envArray := make([]string, 0, len(env))
	for k, v := range env {
		envArray = append(envArray, fmt.Sprintf("%s=%s", k, v))
	}

	// Create the container
	resp, err := s.dockerClient.ContainerCreate(
		ctx,
		&container.Config{Image: imageName, Env: envArray, Tty: false},
		&container.HostConfig{Resources: container.Resources{Memory: 1 << 30, NanoCPUs: 1e9}, SecurityOpt: []string{"no-new-privileges"}},
		nil, nil, "",
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create container: %w", err)
	}
	id := resp.ID
	log.Printf("Created container: %s", id)

	// Ensure cleanup
	defer func() {
		timeout := 10
		if err := s.dockerClient.ContainerStop(context.Background(), id, container.StopOptions{Timeout: &timeout}); err != nil {
			log.Printf("Error stopping container: %v", err)
		}
		if err := s.dockerClient.ContainerRemove(context.Background(), id, container.RemoveOptions{Force: true, RemoveVolumes: true}); err != nil {
			log.Printf("Error removing container: %v", err)
		}
		log.Printf("Cleaned up container: %s", id)
	}()

	// Start it
	if err := s.dockerClient.ContainerStart(ctx, id, container.StartOptions{}); err != nil {
		return nil, fmt.Errorf("failed to start container: %w", err)
	}

	// Wait for exit
	statusCh, errCh := s.dockerClient.ContainerWait(ctx, id, container.WaitConditionNotRunning)
	var code int64
	select {
	case err := <-errCh:
		if err != nil {
			return nil, fmt.Errorf("waiting error: %w", err)
		}
	case st := <-statusCh:
		code = st.StatusCode
	}

	// Retrieve logs
	logs, err := s.dockerClient.ContainerLogs(ctx, id, container.LogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return nil, fmt.Errorf("failed to get logs: %w", err)
	}
	defer logs.Close()

	var out, errOut bytes.Buffer
	if _, err := stdcopy.StdCopy(&out, &errOut, logs); err != nil {
		return nil, fmt.Errorf("failed to read logs: %w", err)
	}

	combined := append(out.Bytes(), errOut.Bytes()...)
	if code != 0 {
		return combined, fmt.Errorf("container exited with %d", code)
	}

	return combined, nil
}

// Close releases resources used by the scheduler
func (s *AlgoScheduler) Close() error {
	if s.dockerClient != nil {
		return s.dockerClient.Close()
	}
	return nil
}
