package test

import (
	"net/http"
	"testing"
)

func init() {
	// lang.InitLang()
}

func TestWithdrawal(t *testing.T) {

	req, _ := http.NewRequest("GET", "/products", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}
