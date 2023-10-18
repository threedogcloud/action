package action

import (
	"os"
	"testing"
)

func TestSecrets(t *testing.T) {
	omg := os.Getenv("OMG")
	if omg != "" {
		t.Log(omg)
		return
	}

	t.Log("this is not path")
}
