package crm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"sort"
	"strings"
	"time"
)

const (
	DefaultAddress = "https://api.hubapi.com"
)

type Client struct {
	baseURL string
	token  string
	http    *http.Client

	Associations        Associations
	Calls               Calls
	Companies           Companies
	Contacts            Contacts
	Deals               Deals
	Emails              Emails
	FeedbackSubmissions FeedbackSubmissions
	LineItems           LineItems
	Meetings            Meetings
	Notes               Notes
	Owners              Owners
	Pipelines           Pipelines
	Products            Products
	Tasks               Tasks
	Tickets             Tickets
	Quotes              Quotes
}

func NewClient(token string) (*Client, error) {
	if token == "" {
		return nil, fmt.Errorf("hubspot.NewClient(): An API Key must be provided to call the Hubspot API")
	}

	client := &Client{
		baseURL:     DefaultAddress,
		token:      token,
	}

	client.http = &http.Client{
		Timeout: time.Duration(30 * time.Second),
	}

	client.Associations        = &associations{client: client}
	client.Calls               = &calls{client: client}
	client.Companies           = &companies{client: client}
	client.Contacts            = &contacts{client: client}
	client.Deals               = &deals{client: client}
	client.Emails              = &emails{client: client}
	client.FeedbackSubmissions = &feedbackSubmissions{client: client}
	client.LineItems           = &lineItems{client: client}
	client.Meetings            = &meetings{client: client}
	client.Notes               = &notes{client: client}
	client.Owners              = &owners{client: client}
	client.Pipelines           = &pipelines{client: client}
	client.Products            = &products{client: client}
	client.Tasks               = &tasks{client: client}
	client.Tickets             = &tickets{client: client}
	client.Quotes              = &quotes{client: client}
	
	return client, nil
}

func (c *Client) newHttpRequest(method string, endpoint string, v interface{}) (*http.Request, error) {
	var err error
	var body []byte
	var newBody io.Reader
	u, err := c.formatUrlWithapiKey(endpoint)
	if err != nil {
		return nil, fmt.Errorf("client.newHttpRequest(): c.formatUrlWithapiKey(): %v", err)
	}

	reqHeaders := make(http.Header)
	reqHeaders.Set("Content-Type", "application/json")
	reqHeaders.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	
	switch method {
	case "GET", "DELETE":
		if v != nil {
			q, err := query.Values(v)
			if err != nil {
				return nil, fmt.Errorf("client.newHttpRequest(): query.Values(): %v", err)
			}
			u.RawQuery = u.RawQuery + encodeQueryParams(q)
		}
	case "POST", "PUT", "PATCH":
		if v != nil {
			if body, err = json.Marshal(v); err != nil {
				return nil, fmt.Errorf("client.newHttpRequest(): json.Marshal(): %v", err)
			}
			newBody = bytes.NewReader(body)
		}
	}

	req, err := http.NewRequest(method, u.String(), newBody)
	if err != nil {
		return nil, fmt.Errorf("client.newHttpRequest(): http.NewRequest(): %v", err)
	}

	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (error) {
	res, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("client.do(): c.http.Do(): %v", err)
	}
	defer res.Body.Close()

	statusOk := res.StatusCode >= 200 && res.StatusCode < 300

	if !statusOk {
		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("client.do(): ioutil.Readall(): %v", err)
		}
		return fmt.Errorf("client.do(): %s", string(resBody))
	}
	
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("client.do(): ioutil.Readall(): %v", err)
	}

	if len(resBody) > 0 {
		err = json.Unmarshal(resBody, v)
		if err != nil {
			return fmt.Errorf("client.do(): json.Unmarshal(): %v", err)
		}
	}
	
	return nil
}

func (c *Client) formatUrlWithapiKey(endpoint string) (*url.URL, error) {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, fmt.Errorf("hubspot.Client.formatUrlWithapiKey(): url.Parse(): %v", err)
	}

	u.Path = path.Join(u.Path, endpoint)
	q := u.Query()
	u.RawQuery = q.Encode()

	return u, nil
}

func encodeQueryParams(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		if len(vs) > 1 {
			val := strings.Join(vs, ",")
			vs = vs[:0]
			vs = append(vs, val)
		}
		keyEscaped := url.QueryEscape(k)

		for _, v := range vs {
			
			buf.WriteByte('&')
			
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
		}
	}
	return buf.String()
}