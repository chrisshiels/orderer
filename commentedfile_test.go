// 'commentedfile_test.go'.
// Chris Shiels.


package main


import (
    "sort"
    "testing"
)


func Test_CommentedFileWithDependencies(t *testing.T) {
    commentedfile := NewCommentedFile("testdata/commented/a")
    if err := commentedfile.Process(); err != nil {
        t.Errorf("Test_CommentedFileWithDependencies: %s", err)
        return
    }

    name := commentedfile.Name()
    dependencies := commentedfile.Dependencies()
    sort.Strings(dependencies)
    expectequalstrings(t,
                       "Test_CommentedFileWithDependencies",
                       name,
                       "aaa")
    expectequalarraystrings(t,
                            "Test_CommentedFileWithDependencies",
                            dependencies,
                            []string{ "bbb",
                                      "ddd" })
}


func Test_CommentedFileWithoutDependencies(t *testing.T) {
    commentedfile := NewCommentedFile("testdata/commented/e")
    if err := commentedfile.Process(); err != nil {
        t.Errorf("Test_CommentedFileWithoutDependencies: %s", err)
        return
    }

    name := commentedfile.Name()
    dependencies := commentedfile.Dependencies()
    sort.Strings(dependencies)
    expectequalstrings(t,
                       "Test_CommentedFileWithoutDependencies",
                       name,
                       "eee")
    expectequalarraystrings(t,
                            "Test_CommentedFileWithoutDependencies",
                            dependencies,
                            []string{})
}
