package apidemo_test

import (
	"testing"

	"github.com/dankegel/go-api-demo"
)

func TestAPI(t *testing.T) {
	if apidemo.Answer() != 43 {
		t.Fatal("wrong answer, expected 43")
	}
}
