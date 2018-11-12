package main

import (
	"github.com/chainlibs/btcaddress"
	"fmt"
	"encoding/hex"
	"log"
)

func main() {
	//result, err := btcaddress.DecodeWIF("L3uDZ6tCkKiZa43ir9uNPE7LwQ4oEBerKTXbcoi8YhjD46527RVu")
	////result, err := btcaddress.DecodeWIF("5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//privBuffer := result.PrivKey.Serialize()

	//var privBuffer = btcaddress.NewRandom256Must()
	privBuffer, err := hex.DecodeString("18e14a7b6a307f426a94f8114701e7c8e774e7f9a47e2c2035db29a206321725")
	//privBuffer, err := hex.DecodeString("a966eb6058f8ec9f47074a2faadd3dab42e2c60ed05bc34d39d6c0e1d32b8bdf")
	if err != nil {
		log.Fatal(err)
	}
	priv, pub := btcaddress.NewPrivKeyFromBytes(btcaddress.S256(), privBuffer)
	privWif, err := btcaddress.NewWIF(priv, &btcaddress.MainNetParams, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("privBuffer \t\t\t %v \n", hex.EncodeToString(privBuffer))
	fmt.Printf("priv.Serialize \t\t %v \n", hex.EncodeToString(priv.Serialize()))
	fmt.Printf("privWif \t\t\t %v \n", privWif)
	fmt.Printf("Uncompressed pubkey  %v \n", hex.EncodeToString(pub.SerializeUncompressed()))
	fmt.Printf("IsCompressedPubKey \t %v \n", btcaddress.IsCompressedPubKey(pub.SerializeUncompressed()))
	fmt.Printf("Compressed pubkey \t %v \n", hex.EncodeToString(pub.SerializeCompressed()))
	fmt.Printf("IsCompressedPubKey \t %v \n", btcaddress.IsCompressedPubKey(pub.SerializeCompressed()))
	sig, err := btcaddress.SignCompact(btcaddress.S256(), priv, []byte("AAA"), false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Signature of AAA \t %v \n", hex.EncodeToString(sig))
	h160 := btcaddress.Hash160(pub.SerializeCompressed())
	fmt.Printf("Hash160 \t\t\t %v \n", hex.EncodeToString(h160))
	addressPK, err := btcaddress.NewAddressPubKeyHash(h160, &btcaddress.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AddressPK \t\t\t %v \n", addressPK.EncodeAddress())
	serializedScript, err := btcaddress.NewScriptBuilder().AddOp(btcaddress.OP_HASH160).AddData(h160).
		AddOp(btcaddress.OP_EQUAL).Script()
	if err != nil {
		log.Fatal(err)
	}
	addressScript, err := btcaddress.NewAddressScriptHash(btcaddress.Hash160(serializedScript), &btcaddress.MainNetParams) //wrong?
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AddressScript \t\t %v \n", addressScript.EncodeAddress())
	serializedScriptWPK, err := btcaddress.NewScriptBuilder().AddOp(btcaddress.OP_0).AddData(h160).Script()
	if err != nil {
		log.Fatal(err)
	}
	h160WPK := btcaddress.Hash160(serializedScriptWPK) //P2WPKH is 160(sha256)
	addressScriptWPK, err := btcaddress.NewAddressScriptHashFromHash160(h160WPK, &btcaddress.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}
	//A commonly used script is a P2WPKH (Pay to Witness Public Key Hash): OP_0 0x14 <PubKey Hash>
	fmt.Printf("AddressScriptWPK \t %v \n", addressScriptWPK.EncodeAddress())
	addressWPK, err := btcaddress.NewAddressWitnessPubKeyHash(h160WPK, &btcaddress.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AddressWPK \t\t\t %v \n", addressWPK.EncodeAddress())
	h256 := btcaddress.DoubleHashB(serializedScriptWPK) //P2WSH is sha256(sha256))
	addressWScript, err := btcaddress.NewAddressWitnessScriptHash(h256, &btcaddress.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AddressWScript \t\t %v \n", addressWScript.EncodeAddress())
}