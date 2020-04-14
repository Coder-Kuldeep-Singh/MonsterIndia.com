package main

import (
	"fmt"

	"MonsterIndia.com/companies"
)

func main() {
	items, total, err := companies.FindCompaniesByCharacter("z")
	if err != nil {
		fmt.Println("failed Companies Function..: ", err)
	}
	fmt.Println(total)
	for _, item := range items {
		fmt.Println(item)
	}
	fmt.Println("====================================================================================================")

}
