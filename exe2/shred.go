package main

import (
    "fmt"
    "os"
    "math/rand"
    "time"
)

func shredIter(filename string) error{
    const MAX_BYTES = 2000000

    randomInt := rand.Intn(MAX_BYTES)
    fmt.Println("Random length: ", randomInt)
    randomBytes := make([]byte, randomInt)
    rand.Read(randomBytes)

    e := os.WriteFile(filename, randomBytes, 0644)
    return e

}

func shred(filename string) (bool, error) {
    const N_SHRED = 3

    for i := 0; i <N_SHRED; i++ {
        fmt.Println("Iteration ", i+1, " of the shred function.")

        err := shredIter(filename)

        if err != nil {
            return false, nil
        }

    }
    return true, nil
}

func main() {
    rand.Seed(time.Now().UnixNano())

    shred("randomfile")

}