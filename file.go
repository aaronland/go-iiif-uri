package uri

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type FileURI struct {
	URI
	origin string
	target string
}

func init() {
	ctx := context.Background()
	RegisterURI(ctx, "file", NewFileURI)
}

func NewFileURI(ctx context.Context, str_uri string) (URI, error) {

	u, err := url.Parse(str_uri)

	if err != nil {
		return nil, err
	}

	origin := strings.TrimLeft(u.Path, "/")

	if origin == "" {
		return nil, errors.New("Invalid path")
	}

	q := u.Query()

	target := q.Get("target")

	if target == "" {
		target = origin
	}

	f_u := FileURI{
		origin: origin,
		target: target,
	}

	return &f_u, nil
}

func (u *FileURI) Origin() string {
	return u.origin
}

func (u *FileURI) Target(opts *url.Values) (string, error) {
	return u.target, nil
}

func (u *FileURI) String() string {

	raw_uri := fmt.Sprintf("%s", u.origin)

	if u.target != "" && u.target != u.origin {
		q := url.Values{}
		q.Set("target", u.target)
		raw_uri = fmt.Sprintf("%s?%s", raw_uri, q.Encode())
	}

	return fmt.Sprintf("file:///%s", raw_uri)
}
