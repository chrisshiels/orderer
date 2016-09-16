// 'orderable_test.go'.
// Chris Shiels.


package main


import (
    "testing"
)


func Test_NewOrderableFuncFactoryCommented(t *testing.T) {
    expectinstanceof(t,
                     "Test_NewOrderableFuncFactoryCommented",
                     NewOrderableFuncFactory("commented")("f"),
                     "*main.CommentedFile")
}


func Test_NewOrderableFuncFactoryRpm(t *testing.T) {
    expectinstanceof(t,
                     "Test_NewOrderableFuncFactoryRpm",
                     NewOrderableFuncFactory("rpm")("f"),
                     "*main.RpmSpecFile")
}


func Test_NewOrderableFuncFactoryOther(t *testing.T) {
    actual := NewOrderableFuncFactory("other")
    if actual != nil {
        t.Errorf("%s: %v, want %v",
                 "Test_NewOrderableFuncFactoryOther", actual, nil)
        return
    }
}
