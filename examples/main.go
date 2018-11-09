package main

import (
	"github.com/chainlibs/btcaddress"
	"fmt"
	"encoding/hex"
	"log"
)

func main() {
	//var privBuffer = btcaddress.NewRandom256Must()
	var privBuffer = []byte{
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	}
	priv, pub := btcaddress.NewPrivKeyFromBytes(btcaddress.S256(), privBuffer)
	fmt.Println(hex.EncodeToString(privBuffer))
	fmt.Println(hex.EncodeToString(priv.Serialize()))
	fmt.Println(hex.EncodeToString(pub.SerializeUncompressed()))
	fmt.Println(hex.EncodeToString(pub.SerializeCompressed()))
	fmt.Println(btcaddress.IsCompressedPubKey(pub.SerializeUncompressed()))
	fmt.Println(btcaddress.IsCompressedPubKey(pub.SerializeCompressed()))
	sig, err := btcaddress.SignCompact(btcaddress.S256(), priv, []byte("aaa"), false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(sig))
}