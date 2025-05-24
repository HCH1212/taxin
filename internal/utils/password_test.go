package utils

import (
	"testing"
)

func TestPassword(t *testing.T) {
	pass1 := "password1"
	pass2 := "password1"
	hashPaas1, _ := HashPassword(pass1)
	t.Log(VerifyPassword(hashPaas1, pass2))
}
