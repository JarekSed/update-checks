package main

import (
    "checks"
    "fmt"
)

func main() {
    outOfDatePrograms := checks.GetOutOfDatePrograms()
    for _, program := range outOfDatePrograms{
        fmt.Printf("%v\n",program)
    }
}
