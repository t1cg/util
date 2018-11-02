package awssns

import (
	"testing"
)

// go test -run SendSnsSms
func TestSendSnsSms(t *testing.T) {
	t.Log("running test...")

	mode := "dev"

	t.Logf("running in %v mode", mode)

	phoneNumber := []string{"4437918616", "4105913055"}
	message := "this is a test test test"

	// Replace 'awsProfile' with your profile identifier
	ae := SendSnsSms(phoneNumber, message, "us-east-1", "awsProfile", true)
	if ae != nil {
		t.Errorf("ERROR[%v]", ae)
		return
	}

	t.Log("existing test...")
}
