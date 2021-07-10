package dotgit

import (
	"fmt"
	"net/url"
	"os/exec"
	"strings"

	giturls "github.com/whilp/git-urls"
)

// RemoteRepositoryURL retrives remote git repository from local .git via git command
func RemoteRepositoryURL(name string) (*url.URL, error) {
	b, err := exec.Command("git", "remote", "get-url", name).CombinedOutput()

	if err != nil {
		return nil, fmt.Errorf("failed to get remote URL for %s: %w", name, err)
	}

	url, err := giturls.Parse(strings.TrimSpace(string(b)))

	if err != nil {
		return nil, fmt.Errorf("failed to parse %s: %w", string(b), err)
	}

	return url, nil
}
