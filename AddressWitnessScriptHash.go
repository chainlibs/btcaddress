package btcaddress

import (
	"errors"
	"strings"
)

// AddressWitnessScriptHash is an Address for a pay-to-witness-script-hash
// (P2WSH) output. See BIP 173 for further details regarding native segregated
// witness address encoding:
// https://github.com/bitcoin/bips/blob/master/bip-0173.mediawiki
type AddressWitnessScriptHash struct {
	hrp            string
	witnessVersion byte
	witnessProgram [32]byte
}

// NewAddressWitnessScriptHash returns a new AddressWitnessPubKeyHash.
func NewAddressWitnessScriptHash(witnessProg []byte, net *Params) (*AddressWitnessScriptHash, error) {
	return newAddressWitnessScriptHash(net.Bech32HRPSegwit, witnessProg)
}

// newAddressWitnessScriptHash is an internal helper function to create an
// AddressWitnessScriptHash with a known human-readable part, rather than
// looking it up through its parameters.
func newAddressWitnessScriptHash(hrp string, witnessProg []byte) (*AddressWitnessScriptHash, error) {
	// Check for valid program length for witness version 0, which is 32
	// for P2WSH.
	if len(witnessProg) != 32 {
		return nil, errors.New("witness program must be 32 " +
			"bytes for p2wsh")
	}

	addr := &AddressWitnessScriptHash{
		hrp:            strings.ToLower(hrp),
		witnessVersion: 0x00,
	}

	copy(addr.witnessProgram[:], witnessProg)

	return addr, nil
}

// EncodeAddress returns the bech32 string encoding of an
// AddressWitnessScriptHash.
// Part of the Address interface.
func (a *AddressWitnessScriptHash) EncodeAddress() string {
	str, err := encodeSegWitAddress(a.hrp, a.witnessVersion,
		a.witnessProgram[:])
	if err != nil {
		return ""
	}
	return str
}

// ScriptAddress returns the witness program for this address.
// Part of the Address interface.
func (a *AddressWitnessScriptHash) ScriptAddress() []byte {
	return a.witnessProgram[:]
}

// IsForNet returns whether or not the AddressWitnessScriptHash is associated
// with the passed bitcoin network.
// Part of the Address interface.
func (a *AddressWitnessScriptHash) IsForNet(net *Params) bool {
	return a.hrp == net.Bech32HRPSegwit
}

// String returns a human-readable string for the AddressWitnessScriptHash.
// This is equivalent to calling EncodeAddress, but is provided so the type
// can be used as a fmt.Stringer.
// Part of the Address interface.
func (a *AddressWitnessScriptHash) String() string {
	return a.EncodeAddress()
}

// Hrp returns the human-readable part of the bech32 encoded
// AddressWitnessScriptHash.
func (a *AddressWitnessScriptHash) Hrp() string {
	return a.hrp
}

// WitnessVersion returns the witness version of the AddressWitnessScriptHash.
func (a *AddressWitnessScriptHash) WitnessVersion() byte {
	return a.witnessVersion
}

// WitnessProgram returns the witness program of the AddressWitnessScriptHash.
func (a *AddressWitnessScriptHash) WitnessProgram() []byte {
	return a.witnessProgram[:]
}
