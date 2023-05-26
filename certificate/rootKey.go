package certificate

import "encoding/hex"

// DerPrefix is the DER prefix of IC main net.
const DerPrefix = "308182301d060d2b0601040182dc7c0503010201060c2b0601040182dc7c05030201036100"

// RootKey is the root key of IC main net.
const RootKey = "308182301d060d2b0601040182dc7c0503010201060c2b0601040182dc7c05030201036100814c0e6ec71fab583b08bd81373c255c3c371b2e84863c98a4f1e08b74235d14fb5d9c0cd546d9685f913a0c0b2cc5341583bf4b4392e467db96d65b9bb4cb717112f8472e0d5a4d14505ffd7484b01291091c5f87b98883463f98091a0baaae"

var (
	derPrefix []byte
)

func init() {
	var err error
	if derPrefix, err = hex.DecodeString(DerPrefix); err != nil {
		panic(err)
	}
}
