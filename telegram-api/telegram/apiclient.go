package telegram

import (
	"bytes"
	"net/http"
	"net/url"
	"io"
	"io/ioutil"
	"encoding/json"
)

// A APIClient manages communication with Telegram API
type APIClient struct {
	httpClient 	*http.Client

	BaseURL *url.URL
}

// NewAPIClient returns a new Telegram API client
func NewAPIClient(httpClient *http.Client, baseURL *url.URL) *APIClient {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &APIClient{httpClient: httpClient, BaseURL: baseURL}

	return client
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.  If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *APIClient) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}


	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Do sends an API request and returns the API response.  The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.  If rate limit is exceeded and reset time is in the future,
// Do returns *RateLimitError immediately without making a network API call.
func (c *APIClient) Do(req *http.Request, v interface{}) (*Response, error) {

	resp, err := c.httpClient.Do(req)

	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return response, err
	}

	// contents, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	println(err)
	// }
	// fmt.Printf("%s\n", string(contents))

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {

			err = json.Unmarshal(response.Result, &v)

			if err != nil {

				println(err.Error())
			}
			// err = json.NewDecoder(resp.Body).Decode(v)
			// if err == io.EOF {
			// 	err = nil // ignore EOF errors caused by empty response body
			// }
		}
	}

	return response, err
}

