package client

import "testing"

func Test_GCPTranslator(t *testing.T) {
	sut := NewGCPTranslator()
	texts := []string{
		"hello world",
		"my name is george",
	}

	res, err := sut.TranslateText(texts)
	if err != nil {
		t.Fatalf("Error translate texts: %v", err)
	}

	if len(res) != len(texts) {
		t.Fatalf("Expected same length")
	}
}
