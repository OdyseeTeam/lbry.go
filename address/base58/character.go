package base58

import (
	"errors"
	"math/big"
)

var b58_characters = [58]byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x41, 0x42, 0x43, 0x44, 0x45,
	0x46, 0x47, 0x48, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55,
	0x56, 0x57, 0x58, 0x59, 0x5a, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69,
	0x6a, 0x6b, 0x6d, 0x6e, 0x6f, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78,
	0x79, 0x7a}

func CharacterIndex(character byte) (*big.Int, error) {
	for i := 0; i < 58; i++ {
		if b58_characters[i] == character {
			return big.NewInt(int64(i)), nil
		}
	}
	return nil, errors.New("invalid character")
}
