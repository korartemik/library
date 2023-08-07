package test

import (
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"library/api/model"
	"library/genproto"
	"os"
	"os/exec"
	"testing"
	"time"
)

type libraryTester struct {
	serverOpt model.ServerOptions
	srv       *grpc.Server
}

func RunTester(t *testing.T, options model.ServerOptions) *libraryTester {
	cmd := exec.Command("docker", "build", "-f", "docker/dockerfile", "-t", "mysql", ".")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	require.NoError(t, err)

	finishAndDelete()

	cmd = exec.Command("docker", "run", "--name", "mysql_container", "-d", "-p", "3306:3306", "mysql:latest")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	require.NoError(t, err)

	time.Sleep(30 * time.Second)

	srv, err := genproto.StartServer(options)
	require.NoError(t, err)

	return &libraryTester{
		serverOpt: options,
		srv:       srv,
	}
}

func (t libraryTester) Close() {
	t.srv.Stop()
	finishAndDelete()
}

func finishAndDelete() {
	cmd := exec.Command("docker", "stop", "mysql_container")
	cmd.Run()

	cmd = exec.Command("docker", "rm", "mysql_container")
	cmd.Run()
}
