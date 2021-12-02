package main

import (
    "fmt"
    "io"
    "container/list"
)

type tuple struct {
    values map[string]string
}

type relation struct {
    attrs []string
    tuples *list.List
}

func makeRelation(data [][]string) *relation {
    r := new(relation)

    r.tuples = list.New()
    for i := range data {
        if i == 0 {
            r.attrs = make([]string, len(data[i]))
            copy(r.attrs, data[i])
        } else {
            t := new(tuple)
            t.values = make(map[string]string)
            for j := range data[i] {
                t.values[r.attrs[j]] = data[i][j]
            }
            r.tuples.PushBack(t)
        }
    }

    return r
}

func (r *relation) print(w io.Writer) {
    for i := range r.attrs {
        fmt.Fprintf(w, "( %s )", r.attrs[i])
    }
    fmt.Fprintf(w, "\n")

    for node := r.tuples.Front(); node != nil; node = node.Next() {
        t := node.Value.(*tuple)
        for i := range r.attrs {
            fmt.Fprintf(w, "| %s |", t.values[r.attrs[i]])
        }
        fmt.Fprintf(w, "\n")
    }
}

