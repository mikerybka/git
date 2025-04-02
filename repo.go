package git

import (
	"fmt"
	"os/exec"
	"strings"
)

func NewRepo(dir string) *Repo {
	return &Repo{
		Dir: dir,
	}
}

type Repo struct {
	Dir string `json:"dir"`
}

func (r *Repo) Branch() (string, error) {
	cmd := exec.Command("git", "status")
	cmd.Dir = r.Dir
	out, err := cmd.CombinedOutput()
	s := string(out)
	if err != nil {
		return "", fmt.Errorf("git status: %s", strings.TrimSpace(s))
	}
	lines := strings.Split(s, "\n")
	if len(lines) == 0 {
		return "", fmt.Errorf("no git output")
	}
	firstLine := strings.TrimSpace(lines[0])
	if !strings.HasPrefix(firstLine, "On branch ") {
		return "", fmt.Errorf("unexpected git output")
	}
	branch := strings.TrimPrefix(firstLine, "On branch ")
	return branch, nil
}

func (r *Repo) Pull() (updated bool, err error) {
	cmd := exec.Command("git", "pull")
	cmd.Dir = r.Dir
	out, err := cmd.CombinedOutput()
	s := strings.TrimSpace(string(out))
	if err != nil {
		return false, fmt.Errorf("git pull: %s", s)
	}
	if s == "Already up to date." {
		return false, nil
	}
	return true, nil
}
