// [_Command-line arguments_](http://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// are a common way to parameterize execution of programs.
// For example, `go run hello.go` uses `run` and
// `hello.go` arguments to the `go` program.

package main

import "os"
import (
  "fmt"
  "reflect"
   "strconv"
)

func main() {

    // `os.Args` provides access to raw command-line
    // arguments. Note that the first value in this slice
    // is the path to the program, and `os.Args[1:]`
    // holds the arguments to the program.
    arg := os.Args[1:]

    // You can get individual args with normal indexing.

    fmt.Println(arg)
    fmt.Println(reflect.TypeOf(arg))

    var argi = []int{}
    for _, i := range arg {
        j, err := strconv.Atoi(i)
        if err != nil {
            panic(err)
        }
        argi = append(argi, j)
    }
    fmt.Println(argi)


}
