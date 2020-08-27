package sfdc

import (
	"net/http"
	"path"
)

type Service struct {
	client   *Client
	basePath string
}

func (s *Service) NewRequest(method, urlPath string, body interface{}) (*http.Request, error) {
	p := path.Join(s.basePath, urlPath)
	return s.client.NewRequest(method, p, body)
}
