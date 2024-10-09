package main

import (
	"fmt"
	"gitlab.com/developerdurp/wingotil/pkg/ostype"
	"log"
)

func main() {
	os, err := ostype.GetOsType()
	if err != nil {
		log.Fatal(err)
	}

	switch os.(type) {
	case *ostype.WinVer:
		fmt.Println("Windows")
	case *ostype.LinVer:
		fmt.Println("Linux")
	default:
		fmt.Println("unable to detect OS version")
	}

}
