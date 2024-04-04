package dataciteapi

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"testing"
)

var (
	MailTo string
)

func TestClient(t *testing.T) {
	api, err := NewDataCiteClient("datacite_test.go", MailTo)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if api.RateLimitLimit != 0 {
		t.Errorf("expected 0, got %d", api.RateLimitLimit)
	}
	if api.RateLimitInterval != 0 {
		t.Errorf("expected 0, got %d", api.RateLimitInterval)
	}
	if fmt.Sprintf("%s", api.LastRequest) == "0001-01-01 00:00:00 +0000" {
		t.Errorf("expected 0001-01-01 00:00:00 +0000, got %s", api.LastRequest)
	}

	// test low level getJSON, 10.22002/D1.868 is
	// a Caltech Library DataCite DOI
	src, err := api.getJSON("/works/10.22002/D1.868")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if api.StatusCode != 200 {
		t.Errorf("Expected StatusCode 200, got %d -> %s", api.StatusCode, api.Status)
	}
	if len(src) == 0 {
		t.Errorf("expected a response body from /works/10.22002/D1.868")
	}
	if api.RateLimitLimit == 0 {
		t.Errorf("expected greater than zero, got %d", api.RateLimitLimit)
	}
	if api.RateLimitInterval == 0 {
		t.Errorf("expected greater than zero, got %d", api.RateLimitInterval)
	}
	if fmt.Sprintf("%s", api.LastRequest) == "0001-01-01 00:00:00 +0000" {
		t.Errorf("expected not equal to 0001-01-01 00:00:00 +0000, got %s", api.LastRequest)
	}

	// Now test WorksJSON()
	doi := "10.22002/D1.868"
	doi_url := "https://dx.doi.org/10.22002/D1.868"

	src, err = api.WorksJSON(doi)
	if err != nil {
		t.Errorf("expected a JSON response, got %s", err)
		t.FailNow()
	}
	if api.StatusCode != 200 {
		t.Errorf("expected status code 200, got %d -> %q", api.StatusCode, api.Status)
		t.FailNow()
	}
	obj1 := make(Object)
	err = json.Unmarshal(src, &obj1)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if obj1 == nil {
		t.Errorf("expected unmarshaled object, got nil")
		t.FailNow()
	}
	obj2, err := api.Works(doi_url)
	if obj2 == nil {
		t.Errorf("expected an non-nil Object from Types(), got nil but no error")
		t.FailNow()
	}
	if len(obj1) != len(obj2) {
		t.Errorf("expected equal lengths for obj1, obj2 ->\n%+v, \n%+v", obj1, obj2)
		t.FailNow()
	}
}

func TestArXivLookup(t *testing.T) {
	api, err := NewDataCiteClient("datacite_test.go", MailTo)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	ids := []string{
		"10.22002/D1.868",
		"arXiv:2402.12335v1", 
		"arXiv:2401.12460v1", 
		"arXiv:2204.13532v2",
		"arXiv:2312.07215",
		"arXiv:2305.06519",
		"arXiv:2312.03791",
		"arXiv:2305.19279",
		"arXiv:2305.05315",
		"arXiv:2305.07673",
		"arXiv:2111.03606",
		"arXiv:2112.06016",
	}
	for _, id := range ids {
		_, err = api.WorksJSON(id)
		if err != nil {
			t.Errorf("Run api.WorksJSON(%q), got unexpected error, %s", id, err)
		}
	}
}

func TestMain(m *testing.M) {
	flag.StringVar(&MailTo, "mailto", "", "set the mailto for testing")
	flag.Parse()
	if MailTo == "" {
		MailTo = "test@example.library.edu"
	}
	log.Printf("mailto: %q", MailTo)
	os.Exit(m.Run())
}
