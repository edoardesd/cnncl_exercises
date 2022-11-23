package main

import (
    "fmt"
    "os"
    "math/rand"
    "time"
)

/**
   A Shred overwrite N_SHRED times a file and eventually delete it.
   The file is filled with bytes with a random length.
   Return the a boolean with the result of the shred and the error, if any
@param filename                 the file that needs to be shreded
@return         boolean, error  whether the shred fails or not, error indicating the failure status
*/
func Shred(filename string) (bool, error) {
    const N_SHRED = 3
    const MAX_BYTES = 20000

    for i := 0; i <N_SHRED; i++ {
        randomInt := rand.Intn(MAX_BYTES) // number of the bytes to overwrite the file
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
    fmt.Println(filename, "shred and deleted.")

    return true, nil
}

func main() {
    rand.Seed(time.Now().UnixNano())

    Shred("randomfile")

}