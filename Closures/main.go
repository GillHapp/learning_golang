package main

import "fmt"

func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor // Factor is remembered
    }
}

func main() {
    double := multiplier(2) // Closure remembers factor = 2
    triple := multiplier(3) // Closure remembers factor = 3

    fmt.Println(double(5)) // Output: 10 (5 * 2)
    fmt.Println(triple(5)) // Output: 15 (5 * 3)
}
