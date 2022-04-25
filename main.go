package main

import (
	"fmt"
	"github.com/spuzirev/gomod/pkg"
	"github.com/spuzirev/gomod/pkg/subpkg"
)

func main() {
	fmt.Println(pkg.V)
	fmt.Println(subpkg.V)
}
