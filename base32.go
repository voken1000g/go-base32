package base32

import (
    "crypto/sha256"
    B "encoding/base32"
    "errors"
    "strings"
)

const alphabet string = "0123456789abcdefghjkmnpqrstuvwxy"

/*
 * Encode
 */

var encoding = B.NewEncoding(alphabet).WithPadding(B.NoPadding)
var ErrInvalidChecksum = errors.New("Invalid checksum")

func Encode(src []byte) []byte {
    buf := make([]byte, encoding.EncodedLen(len(src)))
    encoding.Encode(buf, src)
    hash32 := sha256.Sum256(buf)

    for i, ch := range buf {
        if ch >= 97 && ch <= 122 && hash32[i%32] > 127 {
            buf[i] -= 32
        }
    }

    return buf
}

func EncodeToString(src []byte) string {
    return string(Encode(src))
}

func EncodeFromString(s string) []byte {
    return Encode([]byte(s))
}

func EncodeString(s string) string {
    return EncodeToString([]byte(s))
}

/*
 * Decode
 */

func Decode(input []byte) ([]byte, error) {
    return DecodeFromString(string(input))
}

func DecodeToString(input []byte) (string, error) {
    decoded, err := Decode(input)
    return string(decoded), err
}

func DecodeFromString(s string) ([]byte, error) {
    decoded, err := encoding.DecodeString(strings.ToLower(s))
    encoded := EncodeToString(decoded)

    if s != encoded {
      return decoded, ErrInvalidChecksum
    }

    return decoded, err
}

func DecodeString(s string) (string, error) {
    return DecodeToString([]byte(s))
}

/*
 * Checksum
 */

func IsBase32String(s string) bool {
    decoded, err := DecodeFromString(s)
    if err != nil {
        return false
    }

    encoded := EncodeToString(decoded)
    return s == encoded
}
