package main

import "testing"

func TestHello(t*testing.T) {
  t.Run("saying hello to people", func(t *testing.T){
  actual := Hello("1")
  expected := "Hello, 1!"

  if expected != actual{
    t.Errorf("expected %q but got %q", expected, actual)
  }
  })
  t.Run("default to 'world' when an empty string is supplied", func(t *testing.T){
    actual := Hello("")
    expected := "Hello, world!"

    if expected != actual{
      t.Errorf("expected %q but got %q", expected, actual)
    }
  })
}
