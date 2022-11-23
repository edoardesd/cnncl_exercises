package main

import (
    "os"
    "math/rand"
)

const MAX_BYTES = 2000000

func shred(filename string) (bool, error) {
    randomInt := rand.Intn(MAX_BYTES)
    randomBytes := make([]byte, randomInt)
    rand.Read(randomBytes)

    err := os.WriteFile(filename, randomBytes, 0644)

    if err!= nil{
        return false, err
    }

    return true, nil
}

func main() {
    shred("randomfile")

}