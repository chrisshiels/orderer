// 'main.go'.
// Chris Shiels.


package main


import (
    "flag"
    "fmt"
    "io"
    "os"
)


const exitsuccess = 0
const exitfailure = 1


func _main(stdin io.Reader,
           stdout io.Writer,
           stderr io.Writer,
           args []string) (exitstatus int) {
    flagset := flag.NewFlagSet(args[0], flag.ExitOnError)
    flagset.Usage = func() {
        fmt.Fprintln(stderr,
                     "Usage:  orderer [ -v ] --filetype filetype file ...")
        flagset.PrintDefaults()
    }
    flagfiletype := flagset.String("filetype",
                                   "",
                                   "File type:  commented, rpm")
    flagverbose := flagset.Bool("v",
                                false,
                                "Verbose")

    // Note flagset.Parse() will also handle '-h' and '--help' and will exit
    // with exit status 2.
    flagset.Parse(args[1:])

    if len(flagset.Args()) == 0 {
        flagset.Usage()
        return exitfailure
    }


    neworderablefunc := NewOrderableFuncFactory(*flagfiletype)
    if neworderablefunc == nil {
        flagset.Usage()
        return exitfailure
    }

    var verbosefunc VerboseFunc
    if *flagverbose {
        verbosefunc = func(message string) {
            fmt.Fprintf(stderr, message)
        }
    }


    ordered, err := order(verbosefunc, neworderablefunc, flagset.Args())
    if err != nil {
        fmt.Fprintf(stderr, "orderer: %s\n", err)
        return exitfailure
    }

    for _, v := range ordered {
        fmt.Fprintf(stdout, "%s\n", v)
    }


    return exitsuccess
}


func main() {
    os.Exit(_main(os.Stdin, os.Stdout, os.Stderr, os.Args))
}
