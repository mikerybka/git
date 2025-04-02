package git

import (
	"errors"
	"fmt"
	"os"
)

func CloneOrPull(dir string, url string) (changed bool, err error) {
	fi, err := os.Stat(dir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = Clone(dir, url)
			if err != nil {
				return false, err
			}
			return true, nil
		}
		return false, err
	}

	if !fi.IsDir() {
		return false, fmt.Errorf("%s is a file", dir)
	}
	repo := NewRepo(dir)
	return repo.Pull()
}
