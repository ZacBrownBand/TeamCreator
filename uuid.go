package main

/**
  * A module to create unique identifiers.
  * @module UUID
  */

import "math/rand"

const (
  length int = 24
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

/**
  * Creates a unique id.
  * @returns {String}
  */
func UUID() string {
    parts := make([]rune, length)
    for i := range parts {
        parts[i] = letters[rand.Intn(len(letters))]
    }
    return string(parts)
}
