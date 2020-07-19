package main

import (
	"testing"
)

func TestTextEncoding(t *testing.T) {
	testText := "@startuml\n1 -> 2 : 3\n@enduml"
	correct := "U9noA2v9B2efpStXCbJGjLCmKh1ICEPoICrBAStD0GG00F__U-S8KG00"
	encoded := encodeText(testText)

	if encoded != correct {
		t.Errorf("Correct: %s\nActual:%s", correct, encoded)
	}
}
