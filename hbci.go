package hbci

import "crypto/rsa"

type Marshaler interface {
	MarshalHBCI() ([]byte, error)
}

type Unmarshaler interface {
	UnmarshalHBCI([]byte) error
}

// supportedUnmarshaler maps segment IDs to Unmarshalers
var supportedUnmarshaler = map[string]Unmarshaler{}

func MakeCall() string {
	return ""
}

type Client struct {
	rsaKey *rsa.PublicKey
}
