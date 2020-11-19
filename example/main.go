package main

import (
    "fmt"
    "github.com/tamakiii-sandbox/hello-golang-package/greet"
    "github.com/tamakiii-sandbox/hello-golang-package/internal"
)

func main() {
    fmt.Printf("%s. I HATE %s", greet.Hello(), internal.SHIT)
}
