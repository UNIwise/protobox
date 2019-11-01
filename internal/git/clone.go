package git

import (
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func Clone(url string, branch string, commit string, out string) (*git.Repository, error) {
	var conf *git.CloneOptions

	if branch != "" || commit != "" {
		conf = &git.CloneOptions{
			URL: url,
		}
	} else {
		// Make shallow clone
		conf = &git.CloneOptions{
			URL:          url,
			Depth:        1,
			SingleBranch: true,
		}
	}

	r, err := git.PlainClone(out, false, conf)

	if branch != "" || commit != "" {
		err = r.Fetch(&git.FetchOptions{
			RefSpecs: []config.RefSpec{"refs/*:refs/*", "HEAD:refs/heads/HEAD"},
		})

		if err != nil {
			return nil, err
		}

		w, err := r.Worktree()

		if err != nil {
			return nil, err
		}

		err = w.Checkout(&git.CheckoutOptions{
			Branch: plumbing.ReferenceName(branch),
			Force:  true,
		})
	}

	return r, err
}
