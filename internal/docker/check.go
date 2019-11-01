package docker

import "os/exec"

func HasDocker() bool {
	_, err := exec.LookPath("docker")
	return err == nil
}
