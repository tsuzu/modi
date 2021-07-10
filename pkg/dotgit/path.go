package dotgit

import (
	"os/exec"
	"strings"
)

// RelativePathFromRoot returns a relative path from .git
func RelativePathFromRoot() (string, error) {
	b, err := exec.Command("git", "rev-parse", "--show-prefix").CombinedOutput()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(b)), nil
}
