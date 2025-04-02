package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Clone(dir, url string) error {
	parent := filepath.Dir(dir)
	err := os.MkdirAll(parent, os.ModePerm)
	if err != nil {
		return err
	}
	cmd := exec.Command("git", "clone", url)
	cmd.Dir = parent
	out, err := cmd.CombinedOutput()
	s := strings.TrimSpace(string(out))
	if err != nil {
		return fmt.Errorf("git clone: %s", s)
	}
	return nil
}
