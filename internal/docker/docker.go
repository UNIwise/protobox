package docker

import "errors"

func Run(cmd string, image string) error {
	if !hasDocker() {
		return errors.New("No docker binaries found")
	}

	return nil
}

func hasDocker() bool {
	return true
}
