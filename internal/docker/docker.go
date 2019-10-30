package docker

import (
	"errors"
	"os/exec"
)

func Run(cmd string, mount string, image string) error {
	if !hasDocker() {
		return errors.New("No docker binaries found")
	}

	return nil
}

func hasDocker() bool {
	_, err := exec.LookPath("docker")
	return err == nil
}
