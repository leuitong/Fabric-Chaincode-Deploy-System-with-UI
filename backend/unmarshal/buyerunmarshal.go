package unmarshal

import (
	"bytes"
	"fmt"
)

type Buyer struct {
	Buyname   string
	Buymoney  string
	Buycredit string
}

func UnmarshalBuyer(buyer []byte) Buyer {
	var buyer1 Buyer
	if len(buyer) != 0 {
		f := func(r rune) bool {
			return bytes.ContainsRune([]byte(": ,{}\""), r)
		}
		buyerBytes := bytes.FieldsFunc(buyer, f)
		fmt.Printf("Fields are: %q\n", buyerBytes)
		buyer1.Buyname = string(buyerBytes[1])
		buyer1.Buymoney = string(buyerBytes[3])
		buyer1.Buycredit = string(buyerBytes[5])
		return buyer1
	} else {
		return buyer1
	}
}
