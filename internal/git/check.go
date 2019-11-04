package git

import "os/exec"

func HasGit() bool {
	_, err := exec.LookPath("git")
	return err == nil
}
