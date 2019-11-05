package git

import (
	"errors"
	"os"
	"os/exec"
)

func Clone(url string, ref string, out string) error {
	if !HasGit() {
		return errors.New("No git binary found")
	}

	cloneArgs := []string{
		"clone",
		"--quiet",
		url,
		out,
	}

	c := exec.Command("git", cloneArgs...)

	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	err := c.Run()
	if err != nil {
		return err
	}

	if ref == "" {
		return nil
	}

	return Checkout(out, ref)
}
