package main

import (
	"flag"
	"fmt"
	"library/api/model"
	"library/genproto"
	"log"
	"os"
	"os/exec"
	"time"
)

// todo: validation by parameter from yaml
func main() {
	cmd := exec.Command("docker", "build", "-f", "cmd/docker/dockerfile", "-t", "mysql", ".")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf(err.Error())
	}

	finishAndDelete()

	cmd = exec.Command("docker", "run", "--name", "mysql_container", "-d", "-p", "3306:3306", "mysql:latest")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf(err.Error())
	}

	time.Sleep(30 * time.Second)

	srv, err := genproto.StartServer(model.ServerOptions{DBOptions: model.DBOptions{
		User:     "local",
		Password: "local",
		Connect:  "127.0.0.1:3306",
		DBName:   "local",
	},
		Port: flag.Int("port", 50051, "The server port")})
	if err != nil {
		log.Fatalf(err.Error())
	}
	time.Sleep(2 * time.Minute)
	srv.Stop()
	fmt.Println("Finish")
}

func finishAndDelete() {
	cmd := exec.Command("docker", "stop", "mysql_container")
	cmd.Run()

	cmd = exec.Command("docker", "rm", "mysql_container")
	cmd.Run()
}
