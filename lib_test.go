// 'lib_test.go'.
// Chris Shiels.


package main


import (
    "reflect"
    "testing"
)


func expectequalstrings(t *testing.T,
                        name string,
                        actual, expected string) {
    if actual != expected {
        t.Errorf("%s: %v, want %v", name, actual, expected)
        return
    }
}


func expectequalarraystrings(t *testing.T,
                             name string,
                             actual, expected []string) {
    if len(actual) != len(expected) {
        t.Errorf("%s: %v, want %v", name, actual, expected)
        return
    }

    for i := 0; i < len(expected); i++ {
        if actual[i] != expected[i] {
            t.Errorf("%s: %v, want %v", name, actual, expected)
            return
        }
    }
}


func expectinstanceof(t *testing.T,
                      name string,
                      actual interface{},
                      expected string) {
    actualtype := reflect.TypeOf(actual).String()
    if actualtype != expected {
        t.Errorf("%s: %v, want %v", name, actualtype, expected)
        return
    }
}


func expectsubsequentelementsinarraystring(t *testing.T,
                                           name string,
                                           actual []string,
                                           expected string,
                                           expectedsubsequent []string) {
    expectedmap := make(map[string]bool, len(expectedsubsequent))
    for _, e := range expectedsubsequent {
        expectedmap[e] = false
    }

    i := 0
    for i < len(actual) {
        if actual[i] == expected {
            break
        }
        i++
    }
    if i == len(actual) {
        t.Errorf("%s: %v, want %v, %v",
                 name, actual, expected, expectedsubsequent)
        return
    }

    for i < len(actual) {
        if _, ok := expectedmap[actual[i]]; ok {
            expectedmap[actual[i]] = true
        }
        i++
    }

    for _, v := range expectedmap {
        if v == false {
            t.Errorf("%s: %v, want %v, %v",
                     name, actual, expected, expectedsubsequent)
            return
        }
    }
}
