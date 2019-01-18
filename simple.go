package main

import (
	"fmt"
	"github.com/jfrog-solutiontest/dog/dog"
        "golang.org/x/net/context"
)

func main () {
	fmt.Println ("Package Name: akita")
	dog.dog.AkitaPackageName()
	dog.dog.ColliePackageName()
        type Context = context.Context
        type CancelFunc = context.CancelFunc
}
