package unmarshal

import (
	"fmt"
	"testing"
)

func TestProviderUnmarshal(t *testing.T) {
	fmt.Println("testing providerunmarshal.go")
	t.Run("testing unmarshalprovider", TestUnmarshalProvider)
}

func TestUnmarshalProvider(t *testing.T) {
	//network := []byte{}
	provider := []byte("{\"ProviderName\":\"aaaa\",\"ProviderLedger\":\"50000\",\"ProviderCredit\":\"A\"}")
	providers := UnmarshalProvider(provider)
	fmt.Println(providers)
}
