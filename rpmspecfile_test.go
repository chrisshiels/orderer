// 'rpmspecfile_test.go'.
// Chris Shiels.


package main


import (
    "sort"
    "testing"
)


func Test_RpmSpecFile(t *testing.T) {
    rpmspecfile := NewRpmSpecFile("testdata/rpm/vlc.spec")
    if err := rpmspecfile.Process(); err != nil {
        t.Errorf("Test_RpmSpecFile: %s", err)
        return
    }

    name := rpmspecfile.Name()
    dependencies := rpmspecfile.Dependencies()
    sort.Strings(dependencies)
    expectequalstrings(t,
                       "Test_RpmSpecFile",
                       name,
                       "vlc")
    expectequalarraystrings(t,
                            "Test_RpmSpecFile",
                            dependencies,
                            []string{ "alsa-lib-devel",
                                      "ffmpeg",
                                      "libgcrypt-devel",
                                      "libva-devel",
                                      "lua-devel",
                                      "qt-devel",
                                      "xcb-util-keysyms-devel" })
}
