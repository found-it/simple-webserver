package main

import "testing"

func TestIncrement(t *testing.T) {

    tables := []struct {
        num  int
        want int
    }{
        {1, 2},
        {2, 3},
        {10, 11},
        {31, 32},
    }

    for _, table := range tables {
        got := Increment(&table.num)
        if got != table.want {
            t.Errorf("Incrementing (%d) incorrect. Got: %d, want: %d.", table.num, got, table.want)
        }
    }
}
