package services

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/delivercodes/bikemessenger/utils"
)

func execService(cmdName string, cmdArgs []string) {
	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("%s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		os.Exit(1)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		os.Exit(1)
	}
}

func runService() {
	config := utils.Readfile("data.yml")
	image := config.Service.Image

	KillService(image)

	name := "--name=" + image
	args := []string{"run", name, image}

	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting Docker instance on pid %d\n", cmd.Process.Pid)
}

//PullService ...
func PullService() {
	config := utils.Readfile("data.yml")
	image := config.Service.Image
	args := []string{"pull", image}
	out, err := exec.Command("docker", args...).Output()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("%s", out)
	runService()
}

//CheckService ...
func CheckService() []byte {
	args := []string{"--unix-socket", "/var/run/docker.sock", "http://localhost/containers/json"}
	out, err := exec.Command("curl", args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}

//KillService kills a docker service, it takes the id of the container to kill
func KillService(container string) (serviceOut []byte, serviceErr error) {
	killArgs := []string{"kill", container}
	rmArgs := []string{"rm", container}
	out, err := exec.Command("docker", killArgs...).Output()
	if out != nil {
		exec.Command("docker", rmArgs...).Run()
	}
	return out, err
}

//RestartService restarts the service .. holy shit dude
func RestartService(container string) []byte {
	out, err := KillService(container)
	if err != nil {
		log.Fatal(err)
	}
	PullService()
	return out
}
