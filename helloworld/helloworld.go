package main

import (
    "fmt"
    "os"
)

func main() {
    for i := 1; i < len(os.Args); i++{
        var args = os.Args[i]
        fmt.Println(args)
    }
}


/* This is invalid:
func main()
{
    fmt.Println("Hello world!");
}
The opening brace has to be on the same line as the function declaration.
*/
