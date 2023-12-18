///////////////////////////////////////////////////////////////////////////////
//
// rm2-download-pdfs/download_pdf.go
// John Simpson <jms1@jms1.net> 2023-12-17
//
// Download a PDF file from a reMarkable tablet

package main

import (
    "fmt"
    "io"
    "log"
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

    ////////////////////////////////////////
    // Request the file

    url := "http://" + tablet_addr + "/download/" + uuid + "/placeholder"

    resp, err := http.Get( url )
    if err != nil {
        log.Fatal( err )
    }

    defer resp.Body.Close()

    ////////////////////////////////////////
    // Create output file

    dest, err := os.Create( localfile )
    if err != nil {
        log.Fatal( err )
    }

    defer dest.Close()

    ////////////////////////////////////////
    // Copy the output to the file

    var src io.Reader = &PassThru{ Reader: resp.Body }

    n, err := io.Copy( dest , src )
    if err != nil {
        log.Fatal( err )
    }

    ////////////////////////////////////////
    // done

    fmt.Printf( "%d ... ok\n" , n )
}
