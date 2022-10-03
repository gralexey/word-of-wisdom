package pow

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"math"
	"math/rand"
	"time"
)

func NewToken() []byte {
	buf := make([]byte, 16)
	binary.BigEndian.PutUint64(buf[:8], math.MaxUint64/1000000)
	rand.Read(buf[8:])
	return buf
}

func Solve(token []byte) (nonce []byte) {
	nonce = make([]byte, 8)
	for i := uint64(0); ; i++ {
		binary.BigEndian.PutUint64(nonce, i)
		if Verify(token, nonce) {
			return
		}
	}
}

func Verify(token, nonce []byte) bool {
	h := sha256.New()
	h.Write(token)
	h.Write(nonce)
	hash := h.Sum(nil)
	return bytes.Compare(hash, token) < 0
}

func init() {
	rand.Seed(time.Now().Unix())
}
