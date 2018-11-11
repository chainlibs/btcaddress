// +build ignore

package main

import (
	"fmt"
	"encoding/hex"
	"github.com/chainlibs/btcaddress/bech32"
)

func main() {
	//ExampleDecode()
	ExampleEncode()
}

// This example demonstrates how to decode a bech32 encoded string.
func ExampleDecode() {
	encoded := "bc1pw508d6qejxtdg4y5r3zarvary0c5xw7kw508d6qejxtdg4y5r3zarvary0c5xw7k7grplx"
	hrp, decoded, err := bech32.Decode(encoded)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Show the decoded data.
	fmt.Println("Decoded human-readable part:", hrp)
	fmt.Println("Decoded Data:", hex.EncodeToString(decoded))

	// Output:
	// Decoded human-readable part: bc
	// Decoded Data: 010e140f070d1a001912060b0d081504140311021d030c1d03040f1814060e1e160e140f070d1a001912060b0d081504140311021d030c1d03040f1814060e1e16
}

// This example demonstrates how to encode data into a bech32 string.
func ExampleEncode() {
	data := []byte("Test data")
	fmt.Printf("%b \n", data)
	// Convert test data to base32:
	conv, err := bech32.ConvertBits(data, 8, 5, true)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%b \n", conv)
	encoded, err := bech32.Encode("customHrp!xxxxxx", conv)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Show the encoded data.
	fmt.Println("Encoded Data:", encoded)

	// Output:
	//	[1010100 1100101 1110011 1110100 100000 1100100 1100001 1110100 1100001]
	//	[1010 10001 10010 10111 110 11101 1 0 1100 10001 10000 10111 1000 11000 1000]
	//	Encoded Data: customHrp!xxxxxx123jhxapqv3shgcgsu99gm
}
