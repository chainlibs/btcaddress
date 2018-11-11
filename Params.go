package btcaddress

type Magic byte

// Params defines a Bitcoin network by its parameters.  These parameters may be
// used by Bitcoin applications to differentiate networks as well as addresses
// and keys for one network from those intended for use on another network.
type Params struct {
	// Address encoding magics
	PubKeyHashMagic        Magic // First byte of a P2PKH address
	ScriptHashMagic        Magic // First byte of a P2SH address
	PrivateKeyMagic        Magic // First byte of a WIF private key
	WitnessPubKeyHashMagic Magic // First byte of a P2WPKH address
	WitnessScriptHashMagic Magic // First byte of a P2WSH address
	Bech32HRPSegwit		   string
}

// main net Address encoding magics
var MainNetParams = Params{
	PubKeyHashMagic:        Magic(0x00), // starts with 1
	ScriptHashMagic:        Magic(0x05), // starts with 3
	PrivateKeyMagic:        Magic(0x80), // starts with 5 (uncompressed) or K (compressed)
	WitnessPubKeyHashMagic: Magic(0x06), // starts with p2
	WitnessScriptHashMagic: Magic(0x0A), // starts with 7Xh
	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: 		"bc", // always bc for main net
}

// regress net Address encoding magics
var RegressionNetParams = Params{
	PubKeyHashMagic:        Magic(0x6f), // starts with m or n
	ScriptHashMagic:        Magic(0xc4), // starts with 2
	PrivateKeyMagic:        Magic(0xef), // starts with 9 (uncompressed) or c (compressed)
	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: 		"bcrt", // always bcrt for reg test net
}

// test net Address encoding magics
var TestNet3Params = Params{
	PubKeyHashMagic:        Magic(0x6f), // starts with m or n
	ScriptHashMagic:        Magic(0xc4), // starts with 2
	PrivateKeyMagic:        Magic(0xef), // starts with 9 (uncompressed) or c (compressed)
	WitnessPubKeyHashMagic: Magic(0x03), // starts with QW
	WitnessScriptHashMagic: Magic(0x28),// starts with T7n
	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: 		"tb", // always tb for test net
}