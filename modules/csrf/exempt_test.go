package csrf

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eject/eject"
)

func TestExemptPath(t *testing.T) {
	MarkExempt("/Controller/Action")

	resp := httptest.NewRecorder()
	postRequest, _ := http.NewRequest("POST", "http://www.example.com/Controller/Action", nil)
	c := eject.NewController(eject.NewRequest(postRequest), eject.NewResponse(resp))
	c.Session = make(eject.Session)

	testFilters[0](c, testFilters)

	if c.Response.Status == 403 {
		t.Fatal("post to csrf exempt action should pass")
	}
}

func TestExemptPathCaseInsensitive(t *testing.T) {
	MarkExempt("/Controller/Action")

	resp := httptest.NewRecorder()
	postRequest, _ := http.NewRequest("POST", "http://www.example.com/controller/action", nil)
	c := eject.NewController(eject.NewRequest(postRequest), eject.NewResponse(resp))
	c.Session = make(eject.Session)

	testFilters[0](c, testFilters)

	if c.Response.Status == 403 {
		t.Fatal("post to csrf exempt action should pass")
	}
}

func TestExemptAction(t *testing.T) {
	MarkExempt("Controller.Action")

	resp := httptest.NewRecorder()
	postRequest, _ := http.NewRequest("POST", "http://www.example.com/Controller/Action", nil)
	c := eject.NewController(eject.NewRequest(postRequest), eject.NewResponse(resp))
	c.Session = make(eject.Session)
	c.Action = "Controller.Action"

	testFilters[0](c, testFilters)

	if c.Response.Status == 403 {
		t.Fatal("post to csrf exempt action should pass")
	}
}