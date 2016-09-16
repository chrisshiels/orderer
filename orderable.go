// 'orderable.go'.
// Chris Shiels.


package main


type OrderableBase struct {
    filename string
    name string
    dependencies []string
}


type Orderable interface {
    Process() error
    Filename() string
    Name() string
    Dependencies() []string
}


type NewOrderableFunc func(filename string) Orderable


func NewOrderableFuncFactory(filetype string) NewOrderableFunc {
    switch filetype {
        case "commented":
            return NewCommentedFile
        case "rpm":
            return NewRpmSpecFile
        default:
            return nil
    }
}
