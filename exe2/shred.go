package main

import (
    "fmt"
    "os"
    "math/rand"
    "time"
)

func shred(filename string) (bool, error) {
    const N_SHRED = 3
    const MAX_BYTES = 20000

    for i := 0; i <N_SHRED; i++ {
        fmt.Println("Iteration ", i+1, " of the shred function.")
        randomInt := rand.Intn(MAX_BYTES)
        fmt.Println("Random length: ", randomInt)
        randomBytes := make([]byte, randomInt)
        rand.Read(randomBytes)

        err := os.WriteFile(filename, randomBytes, 0644)
        if err != nil {
            return false, err
        }
    }

    err := os.Remove(filename)
    if err != nil {
        return false, err
    }

    return true, nil
}

func main() {
    rand.Seed(time.Now().UnixNano())

    shred("randomfile")

}