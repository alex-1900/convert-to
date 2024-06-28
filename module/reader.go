package module

import (
	"fmt"
	"io"
	"net/http"
)

type Reader struct {
	URL     string
	Content []byte
	client  *http.Client
}

func NewReader(url string, c []byte) *Reader {
	return &Reader{url, c, new(http.Client)}
}

// Read 复制远程文件内容
func (r Reader) Read(p []byte) (int, error) {
	body, err := r.GetContent()
	if err != nil {
		return 0, err
	}
	return copy(p, body), nil
}

// GetContent 获取远程文件内容
func (r Reader) GetContent() ([]byte, error) {
	if r.Content != nil {
		return r.Content, nil
	}

	req, err := http.NewRequest("GET", r.URL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.5 Safari/605.1.15")

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(r io.ReadCloser) {
		if closeErr := r.Close(); closeErr != nil {
			fmt.Println("Error closing connection:", closeErr)
		}
	}(resp.Body)

	r.Content, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return r.Content, nil
}
