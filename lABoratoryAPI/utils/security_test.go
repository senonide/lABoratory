package utils

import (
	"testing"
)

func TestGetPasswordHash(t *testing.T) {
	sp := NewSecurityProvider()
	pw := "hello world"
	expected := "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
	result := sp.GetPasswordHash(pw)
	if result == expected {
		t.Log("GetPasswordHash PASSED")
	} else {
		t.Errorf("GetPasswordHash FAILED. Expected %s, got %s", expected, result)
	}
}

func TestGenJWT(t *testing.T) {
	sp := NewSecurityProvider()
	result, err := sp.GenJWT("subject", true)
	if result == "" || err != nil {
		t.Errorf("GenJWT FAILED. Result: %s. Error %s", result, err)
	}
	result, err = sp.GenJWT("subject", false)
	if result == "" || err != nil {
		t.Errorf("GenJWT FAILED. Result: %s. Error %s", result, err)
	} else {
		t.Log("GenJWT PASSED")
	}
}
