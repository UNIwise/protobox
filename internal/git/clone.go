package git

import (
	"errors"
	"os"
	"os/exec"
)

func Clone(url string, branch string, commit string, out string) error {
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

	if branch == "" && commit == "" {
		return nil
	}

	checkoutTarget := branch
	if commit != "" {
		checkoutTarget = commit
	}

	return Checkout(out, checkoutTarget)
}
