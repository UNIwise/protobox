package git

import (
	"os"
	"os/exec"
)

func Clone(url string, branch string, commit string, out string) error {
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

	checkoutArgs := []string{
		"checkout",
		"--quiet",
		checkoutTarget,
	}

	c = exec.Command("git", checkoutArgs...)
	c.Dir = out

	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()
}
