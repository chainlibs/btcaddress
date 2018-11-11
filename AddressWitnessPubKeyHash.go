package btcaddress

import (
	"errors"
	"strings"
)

// NewAddressWitnessPubKeyHash returns a new AddressWitnessPubKeyHash.
func NewAddressWitnessPubKeyHash(witnessProg []byte, net *Params) (*AddressWitnessPubKeyHash, error) {
	return newAddressWitnessPubKeyHash(net.Bech32HRPSegwit, witnessProg)
}

// newAddressWitnessPubKeyHash is an internal helper function to create an
// AddressWitnessPubKeyHash with a known human-readable part, rather than
// looking it up through its parameters.
func newAddressWitnessPubKeyHash(hrp string, witnessProg []byte) (*AddressWitnessPubKeyHash, error) {
	// Check for valid program length for witness version 0, which is 20
	// for P2WPKH.
	if len(witnessProg) != 20 {
		return nil, errors.New("witness program must be 20 " +
			"bytes for p2wpkh")
	}

	addr := &AddressWitnessPubKeyHash{
		hrp:            strings.ToLower(hrp),
		witnessVersion: 0x00,
	}

	copy(addr.witnessProgram[:], witnessProg)

	return addr, nil
}

// AddressWitnessPubKeyHash is an Address for a pay-to-witness-pubkey-hash
// (P2WPKH) output. See BIP 173 for further details regarding native segregated
// witness address encoding:
// https://github.com/bitcoin/bips/blob/master/bip-0173.mediawiki
type AddressWitnessPubKeyHash struct {
	hrp            string
	witnessVersion byte
	witnessProgram [20]byte
}

// EncodeAddress returns the bech32 string encoding of an
// AddressWitnessPubKeyHash.
// Part of the Address interface.
func (a *AddressWitnessPubKeyHash) EncodeAddress() string {
	str, err := encodeSegWitAddress(a.hrp, a.witnessVersion,
		a.witnessProgram[:])
	if err != nil {
		return ""
	}
	return str
}

// ScriptAddress returns the witness program for this address.
// Part of the Address interface.
func (a *AddressWitnessPubKeyHash) ScriptAddress() []byte {
	return a.witnessProgram[:]
}

// IsForNet returns whether or not the AddressWitnessPubKeyHash is associated
// with the passed bitcoin network.
// Part of the Address interface.
func (a *AddressWitnessPubKeyHash) IsForNet(net *Params) bool {
	return a.hrp == net.Bech32HRPSegwit
}

// String returns a human-readable string for the AddressWitnessPubKeyHash.
// This is equivalent to calling EncodeAddress, but is provided so the type
// can be used as a fmt.Stringer.
// Part of the Address interface.
func (a *AddressWitnessPubKeyHash) String() string {
	return a.EncodeAddress()
}

// Hrp returns the human-readable part of the bech32 encoded
// AddressWitnessPubKeyHash.
func (a *AddressWitnessPubKeyHash) Hrp() string {
	return a.hrp
}

// WitnessVersion returns the witness version of the AddressWitnessPubKeyHash.
func (a *AddressWitnessPubKeyHash) WitnessVersion() byte {
	return a.witnessVersion
}

// WitnessProgram returns the witness program of the AddressWitnessPubKeyHash.
func (a *AddressWitnessPubKeyHash) WitnessProgram() []byte {
	return a.witnessProgram[:]
}

// Hash160 returns the witness program of the AddressWitnessPubKeyHash as a
// byte array.
func (a *AddressWitnessPubKeyHash) Hash160() *[20]byte {
	return &a.witnessProgram
}
