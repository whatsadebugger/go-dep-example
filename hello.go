package main

import (
	"fmt"

	"github.com/google/uuid"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	uu, _ := uuid.NewRandom()
	fmt.Println(uu.String())
}
