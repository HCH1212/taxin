package utils

import (
	"context"
	"testing"
)

func TestGenerateEmbeddingForLikes(t *testing.T) {
	likes := []string{"apple", "banana", "orange"}
	embedding, err := GenerateEmbeddingForLikes(context.Background(), likes)
	if err != nil {
		t.Error(err)
	}
	t.Log(embedding)
}
