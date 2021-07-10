package ghclitransport

import "net/url"

func cloneURL(u *url.URL) *url.URL {
	url := *u
	user := *u.User

	url.User = &user

	return &url
}
