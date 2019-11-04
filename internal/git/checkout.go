package git

import (
	"errors"
	"os"
	"os/exec"
)

func Checkout(dir string, ref string) error {
	if !HasGit() {
		return errors.New("No git binary found")
	}

	checkoutArgs := []string{
		"checkout",
		"--quiet",
		ref,
	}

	c := exec.Command("git", checkoutArgs...)
	c.Dir = dir

	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()
}
