package main

import (
	"fmt"
	"github.com/iancoleman/strcase"
)

func main() {
	fmt.Printf("%q\n", strcase.ToLowerCamel(""))
	fmt.Printf("%q\n", strcase.ToLowerCamel(" "))
	fmt.Printf("%q\n", strcase.ToLowerCamel("\n"))
	fmt.Printf("%q\n", strcase.ToLowerCamel("ASCII"))
	fmt.Printf("%q\n", strcase.ToLowerCamel("🙃"))
	fmt.Printf("%q\n", strcase.ToLowerCamel("Thor"))
	fmt.Printf("%q\n", strcase.ToLowerCamel("Thor, Mímir, Ēostre, & Jörð"))
	fmt.Printf("%q\n", strcase.ToLowerCamel("🙃, Mímir, ēostre, & Jörð"))
	fmt.Printf("%q\n", strcase.ToLowerCamel("thorMímir--Ēostre_Jörð"))
}
