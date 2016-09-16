// 'order.go'.
// Chris Shiels.


package main


import (
    "fmt"

    "github.com/chrisshiels/golang/directedgraph"
)


// Approach:
// - List of filenames.
// - Each of which when processed has a name and dependences to other names.
// - Some of the dependencies are unknown and should be ignored.
// - Build directed graph of names / dependencies then topological sort.
// - Output filenames.
//
// Psuedocode:
// m := map of name to data structure including filename
// dg := directedgraph
// for k, v := range m
//     dg.addvertex(k)
// for k, v := range m
//     for d := range v.dependencies
//         if vertexes includes d
//             dg.addedge(k, d)
// for _, v := range topologicalsort
//     print v.filename


type VerboseFunc func(message string)


func order(verbosefunc VerboseFunc,
           neworderablefunc NewOrderableFunc,
           args []string) (ordered []string, err error) {
    m := make(map[string]Orderable)
    for _, arg := range args {
        o := neworderablefunc(arg)
        if err = o.Process(); err != nil {
            return nil, err
        }
        if verbosefunc != nil {
            verbosefunc(fmt.Sprintf("Processed %s.\n", o.Name()))
        }
        m[o.Name()] = o
    }

    dg := directedgraph.NewDirectedGraph()


    // Collect vertices.
    for k, _ := range m {
        dg.AddVertex(k)
        if verbosefunc != nil {
            verbosefunc(fmt.Sprintf("Added vertex %s.\n", k))
        }
    }


    // Collect edges.
    for k, v := range m {
        for _, d := range v.Dependencies() {
            // Ignore dependencies for which we don't have vertices.
            if m[d] != nil {
                dg.AddEdge(k, d)
                if verbosefunc != nil {
                    verbosefunc(fmt.Sprintf("Added edge %s to %s.\n", k, d))
                }
            }
        }
    }


    // Return result of topological sort.
    ordered = make([]string, len(args))
    for i, v := range dg.TopologicalSort() {
        ordered[i] = m[v].Filename()
    }
    return ordered, nil
}
