package schedule

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	dockerTypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/stdcopy"
	"golang.org/x/term"

	goarchive "github.com/moby/go-archive"
)

// SchedulerHandler defines the callback interface used by AlgoScheduler
// to handle different types of algorithm execution events.
type SchedulerHandler interface {
	OnResolve(ctx context.Context, exeId uint, algoCid string, resolveAt time.Time)
	OnRun(ctx context.Context, exeId uint)
	OnCompleted(ctx context.Context, exeId uint, success bool, output []byte, errmsg []byte)
	OnError(exeId uint, err error)
}

// AlgoScheduler manages algorithm execution events and Docker-based runtime handling.
// It receives scheduling events through an internal channel and dispatches them
// to the registered handler for further processing.
type AlgoScheduler struct {
	eventCh        chan SchedulerEvent
	handler        SchedulerHandler
	dockerClient   *client.Client
	buildSizeLimit int64
}

// SchedulerOption defines a functional option for configuring AlgoScheduler.
type SchedulerOption func(*AlgoScheduler)

const (
	// DEFAULT_CHANNEL_SIZE defines the default buffer size for algorithm execution channel
	DEFAULT_CHANNEL_SIZE = 10

	// DEFAULT_BUILD_SIZE_LIMIT defines the default build context size limit (100MB)
	DEFAULT_BUILD_SIZE_LIMIT = 100 << 20
)

// WithChannelSize sets the buffer size for the algo ID channel.
func WithChannelSize(size int) SchedulerOption {
	return func(s *AlgoScheduler) {
		if size > 0 {
			// s.AlgoIdCh = make(chan uint, size)
			s.eventCh = make(chan SchedulerEvent, size)
		}
	}
}

// WithBuildSizeLimit sets the maximum size limit for Docker image builds.
func WithBuildSizeLimit(bytes int64) SchedulerOption {
	return func(s *AlgoScheduler) {
		if bytes > 0 {
			s.buildSizeLimit = bytes
		}
	}
}

// NewAlgoScheduler creates a new AlgoScheduler instance with optional configuration.
// It initializes the Docker client and internal event channel.
func NewAlgoScheduler(opts ...SchedulerOption) (*AlgoScheduler, error) {
	dockerClient, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, err
	}
	s := &AlgoScheduler{
		eventCh:      make(chan SchedulerEvent, DEFAULT_CHANNEL_SIZE),
		dockerClient: dockerClient,
		// AlgoIdCh:       make(chan uint, DEFAULT_CHANNEL_SIZE),
		buildSizeLimit: DEFAULT_BUILD_SIZE_LIMIT,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s, nil
}

// Run starts the event dispatch loop of the scheduler.
// It listens for incoming scheduling events and delegates handling to the registered handler.
func (s *AlgoScheduler) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("AlgoScheduler shutting down")
			return
		case evt := <-s.eventCh:
			go s.handleEvent(ctx, evt)
		}
	}
}

// handleEvent processes a single SchedulerEvent based on its type.
// It delegates to the appropriate method in the registered SchedulerHandler.
func (s *AlgoScheduler) handleEvent(ctx context.Context, evt SchedulerEvent) {
	if s.handler == nil {
		log.Printf("No handler set for event %s", evt.Type)
		return
	}

	switch evt.Type {
	case EVENT_TYPE_RUN:
		s.handler.OnRun(ctx, evt.ExecutionId)
	case EVENT_TYPE_RESOLVE:
		// unpack resolve_at
		v, ok := evt.Payload["resolve_at"]
		if !ok {
			log.Printf("Missing resolve_at in payload for execution task %d", evt.ExecutionId)
			return
		}
		resolveAt, ok := v.(time.Time)
		if !ok {
			log.Printf("Invalid resolve_at format for execution task %d", evt.ExecutionId)
			return
		}

		v, ok = evt.Payload["algo_cid"]
		if !ok {
			log.Printf("Missing algo_cid in payload for execution task %d", evt.ExecutionId)
		}
		algoCid, ok := v.(string)
		if !ok {
			log.Printf("Invalid algo_cid format for execution task %d", evt.ExecutionId)
		}

		s.handler.OnResolve(ctx, evt.ExecutionId, algoCid, resolveAt)
	}
}

// SetHandler registers the event handler implementation for algorithm execution callbacks.
func (s *AlgoScheduler) SetHandler(handler SchedulerHandler) {
	s.handler = handler
}

// ScheduleRun enqueues an EVENT_TYPE_RUN event to trigger container execution.
func (s *AlgoScheduler) ScheduleRun(exeId uint) {
	event := SchedulerEvent{
		Type:        EVENT_TYPE_RUN,
		ExecutionId: exeId,
	}
	s.eventCh <- event
	log.Printf("Run event scheduled for execution task %d", exeId)
}

// ScheduleResolve enqueues an EVENT_TYPE_RESOLVE event to trigger delayed resolution logic.
func (s *AlgoScheduler) ScheduleResolve(exeId uint, algoCid string, resolveAt time.Time) {
	event := SchedulerEvent{
		Type:        EVENT_TYPE_RESOLVE,
		ExecutionId: exeId,
		Payload:     map[string]any{"resolve_at": resolveAt},
	}
	s.eventCh <- event
	log.Printf("Resolve event scheduled for algo %s", algoCid)
}

// BuildSizeLimit returns the configured maximum Docker build context size.
func (s *AlgoScheduler) BuildSizeLimit() int64 {
	return s.buildSizeLimit
}

// BuildImage builds a Docker image from the specified source directory.
// The directory must contain a valid Dockerfile.
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

// RunContainer runs a Docker container based on the specified image, environment, and mounts.
// It waits for the container to complete and captures its output.
func (s *AlgoScheduler) RunContainer(ctx context.Context, imageName string, env map[string]string, mounts []mount.Mount) ([]byte, []byte, bool, error) {
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
		&container.HostConfig{
			Mounts: mounts,
			// TODO: setup resource
			Resources: container.Resources{
				Memory:   1 << 30,
				NanoCPUs: 1e9,
			},
			SecurityOpt: []string{"no-new-privileges"},
		},
		nil, nil, "",
	)
	if err != nil {
		return nil, nil, false, fmt.Errorf("failed to create container: %w", err)
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
		return nil, nil, false, fmt.Errorf("failed to start container: %w", err)
	}

	// Wait for exit
	statusCh, errCh := s.dockerClient.ContainerWait(ctx, id, container.WaitConditionNotRunning)
	var code int64
	select {
	case err := <-errCh:
		if err != nil {
			return nil, nil, false, fmt.Errorf("waiting error: %w", err)
		}
	case st := <-statusCh:
		code = st.StatusCode
	}

	// Retrieve logs
	logs, err := s.dockerClient.ContainerLogs(ctx, id, container.LogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return nil, nil, false, fmt.Errorf("failed to get logs: %w", err)
	}
	defer logs.Close()

	var out, errOut bytes.Buffer
	if _, err := stdcopy.StdCopy(&out, &errOut, logs); err != nil {
		return nil, nil, false, fmt.Errorf("failed to read logs: %w", err)
	}

	if code != 0 {
		return out.Bytes(), errOut.Bytes(), false, nil
	}
	return out.Bytes(), errOut.Bytes(), true, nil
}

// Close releases resources used by the scheduler, including the Docker client.
func (s *AlgoScheduler) Close() error {
	if s.dockerClient != nil {
		return s.dockerClient.Close()
	}
	return nil
}
