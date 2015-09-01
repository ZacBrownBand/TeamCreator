package main

/**
  * A module to create ids for objects
  */

import "math/rand"

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// creates a unique id
func UUID() string {
    b := make([]rune, 24)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}
