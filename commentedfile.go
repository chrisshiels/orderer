// 'commentedfile.go'.
// Chris Shiels.


// Process files commented with orderer directives, e.g.
//
// // orderer: name dependencies: [ name ... ]
//
// or:
//
// /*
//     orderer: name dependencies: [ name ... ]
// */


package main


import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
)


type CommentedFile struct {
    OrderableBase
}


func NewCommentedFile(filename string) Orderable {
    return &CommentedFile{
        OrderableBase: OrderableBase{ filename: filename,
                                      name: "",
                                      dependencies: make([]string, 0) },
    }
}


func (c *CommentedFile) Process() error {
    file, err := os.Open(c.filename)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    regexpcomment :=
        regexp.MustCompile(`^.*orderer:\s*([^\s]+)\s+dependencies:\s*(.*)$`)

    for scanner.Scan() {
        resultcomment := regexpcomment.FindStringSubmatch(scanner.Text())
        if resultcomment == nil {
            continue
        }

        c.name = resultcomment[1]
        c.dependencies = strings.Fields(resultcomment[2])
        break
    }

    if c.name == "" {
        return fmt.Errorf("Unable to parse orderer comment in %s", c.filename)
    }

    return nil
}


func (c *CommentedFile) Filename() string {
    return c.filename
}


func (c *CommentedFile) Name() string {
    return c.name
}


func (c *CommentedFile) Dependencies() []string {
    return c.dependencies
}
