package quotes

import (
	"fmt"
	//"rsc.io/quotes"
	"rsc.io/quote/v4"
)

func Print() {
	fmt.Println(quote.Hello())
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
}
