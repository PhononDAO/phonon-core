package util

import (
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func TestParseECCPubKey(t *testing.T) {
	//copied from phonon demo CA key
	var uncompressedKey = []byte{
		0x04,
		0x5c, 0xfd, 0xf7, 0x7a, 0x00, 0xb4, 0xb6, 0xb4,
		0xa5, 0xb8, 0xbb, 0x26, 0xb5, 0x49, 0x7d, 0xbc,
		0x7a, 0x4d, 0x01, 0xcb, 0xef, 0xd7, 0xaa, 0xea,
		0xf5, 0xf6, 0xf8, 0xf8, 0x86, 0x59, 0x76, 0xe7,
		0x94, 0x1a, 0xb0, 0xec, 0x16, 0x51, 0x20, 0x9c,
		0x44, 0x40, 0x09, 0xfd, 0x48, 0xd9, 0x25, 0xa1,
		0x7d, 0xe5, 0x04, 0x0b, 0xa4, 0x7e, 0xaf, 0x3f,
		0x5b, 0x51, 0x72, 0x0d, 0xd4, 0x0b, 0x2f, 0x9d,
	}
	var compressedPositiveKey = []byte{
		0x02,
		0xb4, 0x63, 0x2d, 0x08, 0x48, 0x5f, 0xf1, 0xdf,
		0x2d, 0xb5, 0x5b, 0x9d, 0xaf, 0xd2, 0x33, 0x47,
		0xd1, 0xc4, 0x7a, 0x45, 0x70, 0x72, 0xa1, 0xe8,
		0x7b, 0xe2, 0x68, 0x96, 0x54, 0x9a, 0x87, 0x37,
	}

	privKey, err := ethcrypto.GenerateKey()
	if err != nil {
		t.Error("could not generate key")
		return
	}
	generatedCompressedPubKey := ethcrypto.CompressPubkey(&privKey.PublicKey)

	//Took uncompressed key above and changed prefix to 0x05
	invalidKey := []byte{
		0x05,
		0x5c, 0xfd, 0xf7, 0x7a, 0x00, 0xb4, 0xb6, 0xb4,
		0xa5, 0xb8, 0xbb, 0x26, 0xb5, 0x49, 0x7d, 0xbc,
		0x7a, 0x4d, 0x01, 0xcb, 0xef, 0xd7, 0xaa, 0xea,
		0xf5, 0xf6, 0xf8, 0xf8, 0x86, 0x59, 0x76, 0xe7,
		0x94, 0x1a, 0xb0, 0xec, 0x16, 0x51, 0x20, 0x9c,
		0x44, 0x40, 0x09, 0xfd, 0x48, 0xd9, 0x25, 0xa1,
		0x7d, 0xe5, 0x04, 0x0b, 0xa4, 0x7e, 0xaf, 0x3f,
		0x5b, 0x51, 0x72, 0x0d, 0xd4, 0x0b, 0x2f, 0x9d,
	}

	_, err = ParseECCPubKey(uncompressedKey)
	if err != nil {
		t.Error("unable to parse uncompressed pubKey")
		return
	}
	_, err = ParseECCPubKey(compressedPositiveKey)
	if err != nil {
		t.Error("unable to parse compressed pubKey")
		return
	}
	_, err = ParseECCPubKey(generatedCompressedPubKey)
	if err != nil {
		t.Error("unable to parse generated compressed pubKey")
		return
	}
	_, err = ParseECCPubKey(invalidKey)
	if err != ErrInvalidECCPubKeyFormat {
		t.Error("did not detect invalid key correctly")
		return
	}
}
