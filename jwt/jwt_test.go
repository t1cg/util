package jwt

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"
)


func TestGetToken(t *testing.T) {
	header := http.Header{}
	reqWithoutToken := http.Request{Header: header}
	token, err := GetToken(&reqWithoutToken)
	if err == nil {
		t.Errorf("Should not have parsed token")
	}
	header.Add("Authorization", "bearer TEST")
	reqWithToken := http.Request{Header: header}
	token, err = GetToken(&reqWithToken)
	if err != nil || *token != "TEST" {
		t.Errorf("Failed to parse token")
	}
}

func TestIssueAdminToken(t *testing.T) {
	tm, err := time.Parse(time.RFC3339, "2019-01-02T15:04:05Z")
	if err != nil {
		t.Errorf("Error parsing time")
	}
	signedToken, err := IssueAdminToken("test@t1cg.com", "kharon", "test", 1, tm.Add(time.Hour * 100000))
	if err != nil {
		t.Errorf(err.Error())
	}
	// fmt.Println(signedToken)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJ0ZXN0QHQxY2cuY29tIiwicm9sZSI6MSwiZXhwIjoxOTA2NDQxNDQ1LCJpc3MiOiJraGFyb24ifQ.87a8654JWyTGXITpbtagKpzZEgCbtiRIMkaZHA3lB00"
	if token != signedToken {
		t.Errorf("Token did not match")
	}
}

func TestValidate(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJ0ZXN0QHQxY2cuY29tIiwicm9sZSI6MSwiZXhwIjoxOTA2NDQxNDQ1LCJpc3MiOiJraGFyb24ifQ.87a8654JWyTGXITpbtagKpzZEgCbtiRIMkaZHA3lB00"
	claims, err := Validate(token, "test")
	if err != nil {
		t.Errorf("Could not validate token")
	}
	out, err := json.Marshal(claims)
	if err != nil {
		t.Errorf(err.Error())
	}
	golden := `{"userId":"test@t1cg.com","role":1,"exp":1906441445,"iss":"kharon"}`
	if golden != string(out) {
		t.Error("Claims did not match")
	}
}