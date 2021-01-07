package fetch

import "net/http"

type FetchOptions struct {
	Headers       map[string]string // Sets Headers
	Authorization string            // Sets the 'Authorization' header
}

func (o FetchOptions) Initiate(req *http.Request) {
	// Add headers
	for header, content := range o.Headers {
		req.Header.Set(header, content)
	}

	// Authorization
	if o.Authorization != "" {
		req.Header.Set("Authorization", o.Authorization)
	}
}
