package web_net

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestTicket(t *testing.T) {
	s := NewServer(&DummyServerHandler{}, nil)

	jsb, _ := json.Marshal(map[string]string{"test_data": "test"})

	ticketReq, err := http.NewRequest("POST", "/ticket", bytes.NewReader(jsb))
	if err != nil {
		t.Error(err)
		return
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.ServeHTTP)

	handler.ServeHTTP(rr, ticketReq)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	if ok, err := regexp.Match(`{"ticket":"(.*)"}`, rr.Body.Bytes()); !ok || err != nil {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), `{"ticket": "<some ticket id>"}`)
	}
}