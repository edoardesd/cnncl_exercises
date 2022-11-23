package main

import (
    "testing"
)

func TestShred(t *testing.T){
    const name = "random/path"
    got, err := Shred(name)
    want := true

    if got != want {
        t.Errorf("result %t, get %t", got, want )
    }

    if err != nil {
        t.Errorf("Unexpected error %s", err)
    }
}
