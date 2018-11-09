package btcaddress

import "crypto/rand"

/*
Description: 
return a random byte array with 256 bit
 * Author: architect.bian
 * Date: 2018/11/09 18:17
 */
func NewRandom256Must() []byte {
	key := [32]byte{}
	_, err := rand.Read(key[:])
	if err != nil {
		panic(err)
	}
	return key[:]
}