package main

import (
	"fmt"

	"github.com/dedekrnwan/go-solid/dip"
	"github.com/dedekrnwan/go-solid/isp"
	"github.com/dedekrnwan/go-solid/lsp"
	"github.com/dedekrnwan/go-solid/ocp"
	"github.com/dedekrnwan/go-solid/srp"
)

func main() {
	fmt.Println("Learn solid principle in go")

	srp.Run()
	ocp.Run()
	lsp.Run()
	isp.Run()
	dip.Run()
}
