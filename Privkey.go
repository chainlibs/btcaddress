package btcaddress

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"
	"crypto/rand"
)

/*
Description:
PrivKey wraps an ecdsa.PrivKey as a convenience mainly for signing
things with the the priv key without having to directly import the ecdsa
package.
 * Author: architect.bian
 * Date: 2018/11/09 11:45
 */
type PrivKey ecdsa.PrivateKey

/*
Description:
PrivKeyFromBytes returns a priv and public key for `curve' based on the
priv key passed as an argument as a byte slice.
 * Author: architect.bian
 * Date: 2018/11/09 11:49
 */
func NewPrivKeyFromBytes(curve elliptic.Curve, pk []byte) (*PrivKey, *PubKey) {
	x, y := curve.ScalarBaseMult(pk)

	priv := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
			X:     x,
			Y:     y,
		},
		D: new(big.Int).SetBytes(pk),
	}

	return (*PrivKey)(priv), (*PubKey)(&priv.PublicKey)
}

/*
Description:
NewPrivKey is a wrapper for ecdsa.GenerateKey that returns a PrivKey
instead of the normal ecdsa.PrivKey.
 * Author: architect.bian
 * Date: 2018/11/09 11:51
 */
func NewPrivKey(curve elliptic.Curve) (*PrivKey, error) {
	key, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}
	return (*PrivKey)(key), nil
}


// PubKey returns the PubKey corresponding to this private key.
func (p *PrivKey) PubKey() *PubKey {
	return (*PubKey)(&p.PublicKey)
}

// ToECDSA returns the private key as a *ecdsa.PrivKey.
func (p *PrivKey) ToECDSA() *ecdsa.PrivateKey {
	return (*ecdsa.PrivateKey)(p)
}

// Sign generates an ECDSA signature for the provided hash (which should be the result
// of hashing a larger message) using the private key. Produced signature
// is deterministic (same message and same key yield the same signature) and canonical
// in accordance with RFC6979 and BIP0062.
func (p *PrivKey) Sign(hash []byte) (*Signature, error) {
	return signRFC6979(p, hash)
}

// PrivKeyBytesLen defines the length in bytes of a serialized private key.
const PrivKeyBytesLen = 32

// Serialize returns the private key number d as a big-endian binary-encoded
// number, padded to a length of 32 bytes.
func (p *PrivKey) Serialize() []byte {
	b := make([]byte, 0, PrivKeyBytesLen)
	return paddedAppend(PrivKeyBytesLen, b, p.ToECDSA().D.Bytes())
}