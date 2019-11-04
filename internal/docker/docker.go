package docker

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
)

func Run(cmd string, args []string, mount string, image string) error {
	if !HasDocker() {
		return errors.New("No docker binary found")
	}

	absMount, err := filepath.Abs(mount)

	if err != nil {
		return err
	}

	dockerArgs := append([]string{
		"run",
		"-i",
		"-v",
		absMount + ":/mnt",
		image,
		cmd,
	}, args...)

	c := exec.Command("docker", dockerArgs...)

	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()
}
