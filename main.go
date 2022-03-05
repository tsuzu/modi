package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/tsuzu/modi/pkg/dotgit"
)

func goModuleURL() (string, error) {
	url, err := dotgit.RemoteRepositoryURL("origin")

	if err == nil {
		rel, err := dotgit.RelativePathFromRoot()
		if err != nil {
			return "", err
		}

		return path.Join(
			url.Hostname(),
			strings.TrimSuffix(url.Path, ".git"),
			rel,
		), nil
	}

	fmt.Fprintf(os.Stderr, ".git is not initialized. falling back to gh CLI: %v\n", err)

	userOrgs, err := listUserOrgs()

	if err != nil {
		return "", err
	}

	idx, err := fuzzyfinder.Find(userOrgs, func(i int) string {
		return userOrgs[i]
	})

	if err != nil {
		return "", err
	}

	user := userOrgs[idx]

	cwd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("github.com/%s/%s", user, filepath.Base(cwd)), nil
}

func main() {
	url, err := goModuleURL()

	if err != nil {
		panic(err)
	}

	goCommand := os.Getenv("MODI_GO_COMMAND")

	if goCommand == "" {
		goCommand = "go"
	}

	cmd := exec.Command(goCommand, "mod", "init", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
