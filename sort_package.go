// package main

import "fmt"
import "sort"

func sort() {

    // Sort methods are specific to the builtin type;
    // here's an example for strings. Note that sorting is
    // in-place, so it changes the given slice and doesn't
    // return a new one.
    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)

    // An example of sorting `int`s.
    ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

    // We can also use `sort` to check if a slice is
    // already in sorted order.
    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)
    // example := []int{1, 25, 3, 5, 4}
    sort.Sort(sort.Reverse(sort.IntSlice(ints)))
    fmt.Println(ints)
}
