package main

import (
	"fmt"
	"github.com/jfrog-aparicio/dog/dog"
        "golang.org/x/net/context"
)

func main () {
	fmt.Println ("Package Name: akita")
	dog.AkitaPackageName()
	dog.ColliePackageName()
        type Context = context.Context
        type CancelFunc = context.CancelFunc
}
