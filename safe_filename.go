///////////////////////////////////////////////////////////////////////////////
//
// rmweb/safe_filename.go
// John Simpson <jms1@jms1.net> 2023-12-18
//
// Given a filename, return a possibly modified filename which doesn't
// already exist.

package main

import (
    "fmt"
    "log"
    "os"
    "strings"
)


///////////////////////////////////////////////////////////////////////////////
//
// Does a file exist or not?

func file_exists( name string ) bool {
    _,err := os.Stat( name )
    if err == nil {
        return true
    }
    if os.IsNotExist( err ) {
        return false
    }

    log.Fatal( fmt.Sprintf( "file_exists('%s'): %v\n" , name , err ) )
    os.Exit( 1 )

    return false
}

///////////////////////////////////////////////////////////////////////////////
//
// Return a possibly modified filename which doesn't already exist

func safe_filename( name string ) string {

    if do_debug {
        fmt.Printf( "safe_filename('%s') starting\n" , name )
    }

    rv := name

    ////////////////////////////////////////
    // Find extension in filename

    base := rv
    ext  := ""

    dotp := strings.LastIndex( rv , "." )
    if dotp >= 0 {
        base = rv[:dotp]
        ext  = rv[dotp:]
    }

    if do_debug {
        fmt.Printf( "safe_filename()   base='%s' ext='%s'\n" , base , ext )
    }

    ////////////////////////////////////////
    // Try numbers until we find one which doesn't exist yet

    if file_exists( rv ) {
        n := 1
        x := fmt.Sprintf( "%s-%d%s" , base , n , ext )

        if do_debug {
            fmt.Printf( "safe_filename()   x='%s'\n" , x )
        }

        for file_exists( x ) {
            n ++
            x = fmt.Sprintf( "%s-%d%s" , base , n , ext )

            if do_debug {
                fmt.Printf( "safe_filename()   x='%s'\n" , x )
            }
        }

        rv = x
    }

    ////////////////////////////////////////
    // Return what we found

    if do_debug {
        fmt.Printf( "safe_filename('%s') returning '%s'\n" , name , rv )
    }

    return rv
}
