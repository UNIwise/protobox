package docker

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
)

func Run(image string, cmd string, args []string, mounts ...string) error {
	if !HasDocker() {
		return errors.New("No docker binary found")
	}

	dockerArgs := []string{
		"run",
		"-i",
	}

	for _, m := range mounts {
		absMount, err := filepath.Abs(m)

		if err != nil {
			return err
		}

		dockerArgs = append(dockerArgs, []string{
			"-v",
			absMount,
		}...)
	}

	dockerArgs = append(dockerArgs, []string{
		image,
		cmd,
	}...)

	dockerArgs = append(dockerArgs, args...)

	c := exec.Command("docker", dockerArgs...)

	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()
}
