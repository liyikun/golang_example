package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	str := "wefawefefff"

	encodeString := base64.StdEncoding.EncodeToString([]byte(str))

	fmt.Println(encodeString)

	decodeString, _ := base64.StdEncoding.DecodeString(encodeString)

	fmt.Println(string(decodeString))

}
