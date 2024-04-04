package dataciteapi

import (
	//"fmt"
	//"encoding/json"
	"testing"
)

// TestArXivToDOI tests the conversion of arXiv to DOI
func TestArXivToDOI(t *testing.T) {
	ids := map[string]string{
		"10.22002/D1.868":    "10.22002/D1.868",
		"arXiv:2402.12335v1": "10.48550/arXiv.2402.12335v1",
		"arXiv:2401.12460v1": "10.48550/arXiv.2401.12460v1",
		"arXiv:2204.13532v2": "10.48550/arXiv.2204.13532v2",
		"arXiv:2312.07215":   "10.48550/arXiv.2312.07215",
		"arXiv:2305.06519":   "10.48550/arXiv.2305.06519",
		"arXiv:2312.03791":   "10.48550/arXiv.2312.03791",
		"arXiv:2305.19279":   "10.48550/arXiv.2305.19279",
		"arXiv:2305.05315":   "10.48550/arXiv.2305.05315",
		"arXiv:2305.07673":   "10.48550/arXiv.2305.07673",
		"arXiv:2111.03606":   "10.48550/arXiv.2111.03606",
		"arXiv:2112.06016":   "10.48550/arXiv.2112.06016",
	}
	for id, expected := range ids {
		got := ArXivToDOI(id)
		if expected != got {
			t.Errorf("Expected DOI %q, got %q for %q", expected, got, id)
		}
	}
}
