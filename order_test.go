// 'order_test.go'.
// Chris Shiels.


package main


import (
    "path/filepath"
    "testing"
)


func Test_order(t *testing.T) {
    matches, err := filepath.Glob("testdata/rpm/*.spec")
    if err != nil {
        t.Errorf("Test_order: %s", err)
        return
    }

    neworderablefunc := NewOrderableFuncFactory("rpm")

    actual, err := order(nil, neworderablefunc, matches)
    if err != nil {
        t.Errorf("Test_order: %s", err)
        return
    }

    expectsubsequentelementsinarraystring(
        t,
        "Test_order",
        actual,
        "testdata/rpm/ffmpeg.spec",
        []string{ "testdata/rpm/vlc.spec" })

    expectsubsequentelementsinarraystring(
        t,
        "Test_order",
        actual,
        "testdata/rpm/audacious.spec",
        []string{ "testdata/rpm/audacious-plugins.spec" })

    expectsubsequentelementsinarraystring(
        t,
        "Test_order",
        actual,
        "testdata/rpm/mpg123.spec",
        []string{ "testdata/rpm/audacious-plugins.spec" })
}
