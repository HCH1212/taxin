package utils

import "testing"

func TestGenerateUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		id := GenerateUUID()
		t.Log(id)
	}
}
