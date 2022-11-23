package main

import (
    "math/rand"
    "time"
    "testing"
)

func TestShred(t *testing.T){
    rand.Seed(time.Now().UnixNano())
    got, errorGot := Shred("filename")
    want := true

    if got != want {
        t.Errorf("result %t %t, get %t", got, errorGot, want )
    }
}
