package pow

import (
	"math/rand"
	"testing"
)

func TestSolve(t *testing.T) {
	rand.Seed(0)
	token := NewToken()

	nonce := Solve(token)

	if !Verify(token, nonce) {
		t.Errorf("doesn't solve")
	}
}
