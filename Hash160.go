package btcaddress

import (
	"golang.org/x/crypto/ripemd160"
	"crypto/sha256"
)

// Hash160 calculates the hash ripemd160(sha256(b)).
func Hash160(buf []byte) []byte {
	return calcHash(calcHash(buf, sha256.New()), ripemd160.New())
}