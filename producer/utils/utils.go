package utils

import (
	"encoding/json"
	"fmt"
)

func StructPrettyPrint(s interface{}) {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))
}
