package hex_to_int

import (
	"encoding/hex"
	"math/big"
)

// convert hex string to number in little endian
func HexToLittleEndian(h string) (*big.Int, error) {
	var result big.Int

	bytes, err := fromHex(h)
	if err != nil {
		return nil, err
	}

	reverse(bytes)
	return result.SetBytes(bytes), nil
}

// convert hex string to number in big endian
func HexToBigEndian(h string) (*big.Int, error) {
	var result big.Int

	bytes, err := fromHex(h)
	if err != nil {
		return nil, err
	}

	return result.SetBytes(bytes), nil
}

// convert number in little endian to hex string
func LittleEndianToHex(le *big.Int, numOfBytes int) string {
	bytes := le.Bytes()
	reverse(bytes)
	bytes = appendBytes(bytes, 0, numOfBytes-len(bytes))
	return "0x" + hex.EncodeToString(bytes)
}

// convert number in big endian to hex string
func BigEndianToHex(be *big.Int, numOfBytes int) string {
	bytes := be.Bytes()
	leadingZeros := appendBytes([]byte{}, 0, numOfBytes-len(bytes))
	bytes = append(leadingZeros, bytes...)
	return "0x" + hex.EncodeToString(bytes)
}

// get number of bytes in hex represented by string
func NumberOfBytes(h string) int {
	h = toCanonicalHex(h)
	return len(h) / 2
}

func fromHex(s string) ([]byte, error) {
	s = toCanonicalHex(s)
	return hex.DecodeString(s)
}

func toCanonicalHex(s string) string {
	if has0xPrefix(s) {
		s = s[2:]
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	return s
}

func has0xPrefix(s string) bool {
	return len(s) >= 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X')
}

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func appendBytes(bytes []byte, b byte, n int) []byte {
	for i := 0; i < n; i++ {
		bytes = append(bytes, b)
	}
	return bytes
}
