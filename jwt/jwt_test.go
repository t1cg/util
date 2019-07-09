package jwt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestGetToken(t *testing.T) {
	// no auth in header
	header := http.Header{}
	reqWithoutToken := http.Request{Header: header}
	token, err := GetToken(&reqWithoutToken)
	if err == nil {
		t.Errorf("Should not have parsed token")
	}

	// valid header
	header.Add("Authorization", "bearer TEST")
	reqWithToken := http.Request{Header: header}
	token, err = GetToken(&reqWithToken)
	if err != nil || *token != "TEST" {
		t.Errorf("Failed to parse token")
	}

	// invalid header
	invalidHeader := http.Header{}
	invalidHeader.Add("Authorization", "test test test")
	reqWithInvalidHeader := http.Request{Header: invalidHeader}
	token, err = GetToken(&reqWithInvalidHeader)
	if err == nil {
		t.Errorf("Should not have parsed token")
	}
}

func TestIssueToken(t *testing.T) {
	tm, err := time.Parse(time.RFC3339, "2019-01-02T15:04:05Z")
	if err != nil {
		t.Errorf("Error parsing time")
	}
	expiry := tm.Add(time.Hour * 100000)
	signedToken, err := IssueToken("test@t1cg.com", int64(Admin), "kharon", expiry, "test", nil)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(signedToken)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJ0ZXN0QHQxY2cuY29tIiwicm9sZSI6NSwiZGF0YSI6bnVsbCwiZXhwIjoxOTA2NDQxNDQ1LCJpc3MiOiJraGFyb24ifQ.2nltNJ0FJ2i-WuaHqC6tZfN4KcWkgdxTAhtGwmvTCas"
	if token != signedToken {
		t.Errorf("Token did not match")
	}
}

func TestValidate(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJ0ZXN0QHQxY2cuY29tIiwicm9sZSI6NSwiZGF0YSI6bnVsbCwiZXhwIjoxOTA2NDQxNDQ1LCJpc3MiOiJraGFyb24ifQ.2nltNJ0FJ2i-WuaHqC6tZfN4KcWkgdxTAhtGwmvTCas"
	claims, err := Validate(token, "test")
	if err != nil {
		t.Errorf("Could not validate token")
	}
	out, err := json.Marshal(claims)
	if err != nil {
		t.Errorf(err.Error())
	}

	golden := `{"userId":"test@t1cg.com","role":5,"data":null,"exp":1906441445,"iss":"kharon"}`
	if golden != string(out) {
		t.Error("Claims did not match")
	}
}
