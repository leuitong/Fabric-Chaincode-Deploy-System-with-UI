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

func UnmarshalProvider(provider []byte) []Provider {
	if len(provider) != 0 {
		var providers []Provider
		f := func(r rune) bool {
			return bytes.ContainsRune([]byte(":,[]\n{}\""), r)
		}
		providerBytes := bytes.FieldsFunc(provider, f)
		for i := 0; i < len(providerBytes)/6; i++ {
			var provider1 Provider
			fmt.Printf("Fields are: %q\n", providerBytes)
			provider1.Proname = string(providerBytes[1+i*6])
			provider1.Promoney = string(providerBytes[3+i*6])
			provider1.Procredit = string(providerBytes[5+i*6])
			fmt.Printf("provider are: %v\n", provider1)
			providers = append(providers, provider1)
		}
		return providers
	} else {
		return nil
	}
}
