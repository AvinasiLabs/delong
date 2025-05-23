package schedule

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/docker/docker/client"
)

// TestDockerConnectivity verifies that we can connect to the Docker daemon
func TestDockerConnectivity(t *testing.T) {
	// Create a Docker client using the default options
	dockerClient, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		t.Fatalf("Failed to create Docker client: %v", err)
	}
	defer dockerClient.Close()

	// Create a context with timeout to prevent the test from hanging
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ping the Docker daemon to verify connectivity
	ping, err := dockerClient.Ping(ctx)
	if err != nil {
		t.Fatalf("Failed to connect to Docker daemon: %v", err)
	}

	// Check if the API version is returned
	if ping.APIVersion == "" {
		t.Error("Expected non-empty API version, but got empty string")
	} else {
		t.Logf("Successfully connected to Docker daemon (API Version: %s)", ping.APIVersion)
	}

	// Further validate connectivity by requesting Docker info
	info, err := dockerClient.Info(ctx)
	if err != nil {
		t.Fatalf("Failed to retrieve Docker information: %v", err)
	}

	// Validate that we received meaningful Docker information
	if info.ServerVersion == "" {
		t.Error("Expected non-empty Docker server version, but got empty string")
	} else {
		t.Logf("Docker engine version: %s", info.ServerVersion)
		t.Logf("Number of containers: %d", info.Containers)
		t.Logf("Number of images: %d", info.Images)
	}
}

func TestBuildImage(t *testing.T) {
	dir := t.TempDir()
	dockerfile := filepath.Join(dir, "Dockerfile")
	if err := os.WriteFile(dockerfile, []byte("FROM busybox\nCMD echo hello"), 0644); err != nil {
		t.Fatalf("write Dockerfile: %v", err)
	}

	s, err := NewAlgoScheduler()
	if err != nil {
		t.Fatalf("new scheduler: %v", err)
	}
	defer s.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	if err := s.BuildImage(ctx, dir, "test-build:latest"); err != nil {
		t.Errorf("BuildImage failed: %v", err)
	}
}

func TestRunContainer(t *testing.T) {
	dir := t.TempDir()
	dockerfile := filepath.Join(dir, "Dockerfile")
	if err := os.WriteFile(dockerfile, []byte("FROM busybox\nCMD echo hello"), 0644); err != nil {
		t.Fatalf("write Dockerfile: %v", err)
	}

	s, err := NewAlgoScheduler()
	if err != nil {
		t.Fatalf("new scheduler: %v", err)
	}
	defer s.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	image := "test-run:latest"
	if err := s.BuildImage(ctx, dir, image); err != nil {
		t.Fatalf("BuildImage failed: %v", err)
	}

	out, errmsg, success, err := s.RunContainer(ctx, image, nil, nil)
	if err != nil {
		t.Fatalf("RunContainer failed: %v", err)
	}

	if success {
		if !bytes.Contains(bytes.TrimSpace(out), []byte("hello")) {
			t.Errorf("expected hello, got %s", string(out))
		}
	} else {
		t.Errorf("Execution success but algo failed: %s", string(errmsg))
	}

}
