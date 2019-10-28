package git

import (
	git "gopkg.in/src-d/go-git.v4"
)

func Clone(url string, branch string, out string) (*git.Repository, error) {
	r, err := git.PlainClone(out, false, &git.CloneOptions{
		URL:          url,
		SingleBranch: true,
	})

	return r, err
}
