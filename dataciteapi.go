package dataciteapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"

	// Caltech Library Packages
	"github.com/caltechlibrary/doitools"
)

const (
	Version = `v0.0.2`
)

type DataCiteClient struct {
	AppName           string
	MailTo            string `json:"mailto"`
	API               string `json:"api"`
	RateLimitLimit    int    `json:"limit"`
	RateLimitInterval int    `json:"interval"`
	LimtCount         int    `json:"limit"`
	Status            string
	StatusCode        int
	LastRequest       time.Time `json:"last_request"`
}

// Object is the general holder of what get back after unmarshaling json
type Object = map[string]interface{}

// NewDataCiteClient creates a client and makes a request
// and returns the JSON source as a []byte or error if their is
// a problem.
func NewDataCiteClient(appName string, mailTo string) (*DataCiteClient, error) {
	if strings.TrimSpace(mailTo) == "" {
		return nil, fmt.Errorf("An mailto value is required for politeness")
	}
	client := new(DataCiteClient)
	client.AppName = appName
	client.API = `https://api.datacite.org`
	client.MailTo = mailTo
	return client, nil
}

func (c *DataCiteClient) calcDelay() time.Duration {
	if c.RateLimitLimit == 0 {
		return time.Duration(0)
	}
	return time.Duration(int64(math.Ceil(float64(c.RateLimitInterval) / float64(c.RateLimitLimit))))
}

// getJSON retrieves the path from the DataCite API maintaining politeness.
// It returns a []byte of JSON source or an error
func (c *DataCiteClient) getJSON(p string) ([]byte, error) {
	u, err := url.Parse(c.API)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("mailto", c.MailTo)
	u.RawQuery = q.Encode()
	u.Path = p

	client := http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", fmt.Sprintf("%s, based on dataciteapi/%s (github.com/caltechlibrary/dataciteapi/; mailto: %s), A golang cli based on https://support.datacite.org/docs/api", c.AppName, Version, c.MailTo))

	// NOTE: Next request can be made based on last request time plus
	// the duration suggested by X-Rate-Limit-Interval / X-Rate-Limit-Limit
	if delay := c.calcDelay(); delay.Seconds() > 0 {
		time.Sleep(delay)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	// Save the response status
	c.Status = res.Status
	c.StatusCode = res.StatusCode
	// Process the body buffer
	src, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// NOTE: we want to track the current values for any limits
	// `X-Rate-Limit-Limit` and `X-Rate-Limit-Interval` as well
	// as LastRequest time
	if s := res.Header.Get("X-Rate-Limit-Limit"); s != "" {
		if i, err := strconv.Atoi(s); err == nil {
			c.RateLimitLimit = i
		}
	} else if c.RateLimitLimit == 0 {
		c.RateLimitLimit = 1
	}
	if s := res.Header.Get("X-Rate-Limit-Interval"); s != "" {
		if i, err := strconv.Atoi(strings.TrimSuffix(s, "s")); err == nil {
			c.RateLimitInterval = i
		}
	} else if c.RateLimitInterval == 0 {
		c.RateLimitInterval = 1
	}
	c.LastRequest = time.Now()

	return src, nil
}

// WorksJSON return the work JSON source or error for a client and DOI
func (c *DataCiteClient) WorksJSON(doi string) ([]byte, error) {
	s, err := doitools.NormalizeDOI(doi)
	if err != nil {
		return nil, err
	}
	return c.getJSON(path.Join("works", s))
}

// Works return the Work unmarshaled into a Object (i.e. map[string]interface{})
func (c *DataCiteClient) Works(doi string) (Object, error) {
	src, err := c.WorksJSON(doi)
	if err != nil {
		return nil, err
	}
	object := make(Object)
	err = json.Unmarshal(src, &object)
	if err != nil {
		return nil, err
	}
	return object, nil
}