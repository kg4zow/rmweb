///////////////////////////////////////////////////////////////////////////////
//
// rmweb/download_pdf.go
// John Simpson <jms1@jms1.net> 2023-12-17
//
// Download a PDF file from a reMarkable tablet

package main

import (
    "fmt"
    "io"
    "io/fs"
    "net/http"
    "os"
    "strings"
)

///////////////////////////////////////////////////////////////////////////////
//
// Passthru wrapper for io.Reader, prints total bytes while reading

type PassThru struct {
    io.Reader
    total       int64
}

func (pt *PassThru) Read( p []byte ) ( int , error ) {
    n, err      := pt.Reader.Read( p )
    pt.total    += int64( n )

    if err == nil {
        x := fmt.Sprintf( "%d" , pt.total )
        b := fmt.Sprintf( strings.Repeat( "\b" , len( x ) ) )

        fmt.Print( x , b )
    }

    return n , err
}

///////////////////////////////////////////////////////////////////////////////
//
// Download a file

func download_pdf( uuid string , localfile string ) {

    ////////////////////////////////////////////////////////////
    // If the output filename contains any directory names,
    // make sure any necessary directories exist.

    for n := 1 ; n < len( localfile ) ; n ++ {
        if localfile[n] == '/' {
            dir := localfile[:n]

            if do_debug {
                fmt.Printf( "checking dir='%s'\n" , dir )
            }

            ////////////////////////////////////////
            // Check the directory

            s,err := os.Stat( dir )
            if os.IsNotExist( err ) {
                ////////////////////////////////////////
                // doesn't exist yet - create it

                fmt.Printf( "Creating    '%s' ..." , dir )

                err := os.Mkdir( dir , 0755 )
                if err != nil {
                    panic( fmt.Sprintf( "ERROR: %v" , err ) )
                }

                fmt.Println( "ok" )
            } else if err != nil {
                ////////////////////////////////////////
                // os.Stat() had some other error

                panic( fmt.Sprintf( "ERROR: os.Stat('%s'): %v\n" , dir , err ) )
            } else if ( ( s.Mode() & fs.ModeDir ) != 0 ) {
                ////////////////////////////////////////
                // exists and is a directory

                if do_debug {
                    fmt.Printf( "DEBUG '%s' exists and is a directory\n" , dir )
                }
            } else {
                ////////////////////////////////////////
                // exists and is not a directory

                panic( fmt.Sprintf( "ERROR: '%s' exists and is not a directory\n" , dir ) )
                os.Exit( 1 )
            }

        } // if localfile[n] == '/'
    } // for n

    ////////////////////////////////////////////////////////////
    // Download the file

    fmt.Printf( "Downloading '%s' ... " , localfile )

    ////////////////////////////////////////
    // Request the file

    url := "http://" + tablet_addr + "/download/" + uuid + "/placeholder"

    resp, err := http.Get( url )
    if err != nil {
        panic( fmt.Sprintf( "ERROR: %v" , err ) )
    }

    defer resp.Body.Close()

    ////////////////////////////////////////
    // Create output file

    dest, err := os.Create( localfile )
    if err != nil {
        panic( fmt.Sprintf( "ERROR: os.Create('%s'): %v" , localfile , err ) )
    }

    defer dest.Close()

    ////////////////////////////////////////
    // Copy the output to the file

    var src io.Reader = &PassThru{ Reader: resp.Body }

    total, err := io.Copy( dest , src )
    if err != nil {
        panic( fmt.Sprintf( "ERROR: os.Copy(): %v" , err ) )
    }

    ////////////////////////////////////////
    // done

    fmt.Printf( "%d ... ok\n" , total )
}
