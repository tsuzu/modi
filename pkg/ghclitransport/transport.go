package ghclitransport

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

type Transport struct {
}

var _ http.RoundTripper = &Transport{}

func (t *Transport) prepareCommand(req *http.Request) *exec.Cmd {
	u := req.URL

	cmds := []string{"gh", "api"}

	cmds = append(cmds, "-X", req.Method)
	for k, vals := range req.Header {
		for _, v := range vals {
			cmds = append(cmds, "-H", fmt.Sprintf("%s: %s", k, v))
		}
	}
	for k, vals := range u.Query() {
		for _, v := range vals {
			cmds = append(cmds, "-f", fmt.Sprintf("%s=%s", k, v))
		}
	}

	if hostname := u.Hostname(); len(hostname) != 0 {
		cmds = append(cmds, "--hostname", hostname)
	}

	cmds = append(cmds, u.Path)

	return exec.Command(cmds[0], cmds[1:]...)
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Method != http.MethodGet {
		return nil, fmt.Errorf("only GET is allowed")
	}

	cmd := t.prepareCommand(req)
	stdout, stderr := bytes.Buffer{}, bytes.Buffer{}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("executing gh failed(%s): %w", stderr.String(), err)
	}

	return &http.Response{
		StatusCode:    http.StatusOK,
		Status:        http.StatusText(http.StatusOK),
		Body:          io.NopCloser(&stdout),
		ContentLength: int64(stdout.Len()),
	}, nil
}
