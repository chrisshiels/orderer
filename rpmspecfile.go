// 'rpmspecfile.go'.
// Chris Shiels.


// Process rpm spec files, e.g.
//
// Name:       name
// BuildRequires:  dependency
// .
// .


package main


import (
    "bufio"
    "fmt"
    "os"
    "regexp"
)


type RpmSpecFile struct {
    OrderableBase
}


func NewRpmSpecFile(filename string) Orderable {
    return &RpmSpecFile{
        OrderableBase: OrderableBase{ filename: filename,
                                      name: "",
                                      dependencies: make([]string, 0) },
    }
}


func (r *RpmSpecFile) Process() error {
    file, err := os.Open(r.filename)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    regexppercent := regexp.MustCompile(`^%`)
    regexptag := regexp.MustCompile(`^\s*([^:]+):\s*(.*)$`)

    for scanner.Scan() {
        resultpercent := regexppercent.MatchString(scanner.Text())
        if resultpercent {
            continue
        }

        resulttag := regexptag.FindStringSubmatch(scanner.Text())
        if resulttag == nil {
            continue
        }

        tagname := resulttag[1]
        tagvalue := resulttag[2]

        if tagname == "Name" {
            r.name = tagvalue
        } else if tagname == "BuildRequires" {
            r.dependencies = append(r.dependencies, tagvalue)
        }
    }

    if r.name == "" {
        return fmt.Errorf("Unable to parse name tag in %s", r.filename)
    }

    return nil
}


func (r *RpmSpecFile) Filename() string {
    return r.filename
}


func (r *RpmSpecFile) Name() string {
    return r.name
}


func (r *RpmSpecFile) Dependencies() []string {
    return r.dependencies
}
