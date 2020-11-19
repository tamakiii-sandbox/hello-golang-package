package main

import (
    "fmt"
    "github.com/tamakiii-sandbox/hello-golang-package/greet"
    "github.com/tamakiii-sandbox/hello-golang-package/internal"
)

func main() {
    fmt.Printf("%s. %s", greet.Hello(), internal.SHIT)
}
