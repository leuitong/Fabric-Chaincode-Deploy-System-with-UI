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

func UnmarshalBuyer(buyer []byte) []Buyer {
	if len(buyer) != 0 {
		var buyers []Buyer
		f := func(r rune) bool {
			return bytes.ContainsRune([]byte(": ,[]\n{}\""), r)
		}
		buyerBytes := bytes.FieldsFunc(buyer, f)
		for i := 0; i < len(buyerBytes)/6; i++ {
			var buyer1 Buyer
			fmt.Printf("Fields are: %q\n", buyerBytes)
			buyer1.Buyname = string(buyerBytes[1+i*6])
			buyer1.Buymoney = string(buyerBytes[3+i*6])
			buyer1.Buycredit = string(buyerBytes[5+i*6])
			buyers = append(buyers, buyer1)
		}
		return buyers
	} else {
		return nil
	}
}
