package main

import (
    "container/list"
)

func selection(r *relation, pred func(t *tuple) bool) *relation {
    ret := new(relation)

    ret.attrs = make([]string, len(r.attrs))
    copy(ret.attrs, r.attrs)
    ret.tuples = list.New()
    for node := r.tuples.Front(); node != nil; node = node.Next() {
        t := node.Value.(*tuple)
        if (pred(t)) {
            ret.tuples.PushBack(t)
        }
    }

    return ret
}

func projection(r *relation, attrs []string) *relation {
    ret := new(relation)

    ret.attrs = make([]string, len(attrs))
    copy(ret.attrs, attrs)
    ret.tuples = list.New()
    for node := r.tuples.Front(); node != nil; node = node.Next() {
        t := node.Value.(*tuple)

        tp := new(tuple)
        tp.values = make(map[string]string)
        for i := range attrs {
            tp.values[attrs[i]] = t.values[attrs[i]]
        }

        ret.tuples.PushBack(tp)
    }

    return ret
}

func binaryCartesianProduct(ra *relation, rb *relation) *relation {
    ret := new(relation)

    ret.attrs = append(ret.attrs, ra.attrs...)
    ret.attrs = append(ret.attrs, rb.attrs...)
    ret.tuples = list.New()
    for nodeI := ra.tuples.Front(); nodeI != nil; nodeI = nodeI.Next() {
        for nodeJ := rb.tuples.Front(); nodeJ != nil; nodeJ = nodeJ.Next() {
            ti := nodeI.Value.(*tuple)
            tj := nodeJ.Value.(*tuple)

            t := new(tuple)
            t.values = make(map[string]string)
            for k := range ti.values {
                t.values[k] = ti.values[k]
            }
            for k := range tj.values {
                t.values[k] = tj.values[k]
            }

            ret.tuples.PushBack(t)
        }
    }

    return ret
}

func cartesianProduct(rels ...*relation) *relation {
    var ret *relation

    for i, r := range rels {
        if i == 0 {
            ret = r
        } else {
            ret = binaryCartesianProduct(ret, r)
        }
    }

    return ret
}

func join(ra *relation, rb *relation, cond func(t *tuple) bool) *relation {
    var ret *relation

    ret = cartesianProduct(ra, rb)
    ret = selection(ret, cond)

    return ret
}

