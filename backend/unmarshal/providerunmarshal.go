package unmarshal

import (
	"bytes"
	"fmt"
)

type Provider struct {
	Proname   string
	Promoney  string
	Procredit string
}

func UnmarshalProvider(provider []byte) Provider {
	var provider1 Provider
	if len(provider) != 0 {
		f := func(r rune) bool {
			return bytes.ContainsRune([]byte(":,{}\""), r)
		}
		providerBytes := bytes.FieldsFunc(provider, f)
		fmt.Printf("Fields are: %q\n", providerBytes)
		provider1.Proname = string(providerBytes[1])
		provider1.Promoney = string(providerBytes[3])
		provider1.Procredit = string(providerBytes[5])
		fmt.Printf("provider are: %v\n", provider1)
		return provider1
	} else {
		return provider1
	}
}
