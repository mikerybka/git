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

func (r *Repo) Push() error {
	cmd := exec.Command("git", "push")
	cmd.Dir = r.Dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git push: %s", out)
	}
	return nil
}

func (r *Repo) ForcePush() error {
	cmd := exec.Command("git", "push", "-f")
	cmd.Dir = r.Dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git push -f: %s", out)
	}
	return nil
}

func (r *Repo) AddAll() error {
	cmd := exec.Command("git", "add", "--all")
	cmd.Dir = r.Dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git add --all: %s", out)
	}
	return nil
}

func (r *Repo) Commit(msg string) error {
	cmd := exec.Command("git", "commit", "-m", msg)
	cmd.Dir = r.Dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git commit: %s", out)
	}
	return nil
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

func (r *Repo) Checkout(branch string) (err error) {
	cmd := exec.Command("git", "checkout", branch)
	cmd.Dir = r.Dir
	out, err := cmd.CombinedOutput()
	s := strings.TrimSpace(string(out))
	if err != nil {
		return fmt.Errorf("git checkout: %s", s)
	}
	return nil
}

func (r *Repo) CheckoutNew(branch string) (err error) {
	cmd := exec.Command("git", "checkout", "-b", branch)
	cmd.Dir = r.Dir
	out, err := cmd.CombinedOutput()
	s := strings.TrimSpace(string(out))
	if err != nil {
		return fmt.Errorf("git checkout: %s", s)
	}
	return nil
}
