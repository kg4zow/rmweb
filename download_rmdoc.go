///////////////////////////////////////////////////////////////////////////////
//
// rmweb/download_rmdoc.go
// John Simpson <jms1@jms1.net> 2024-09-07
//
// Download an RMDOC file from a reMarkable tablet

package main

import (
    "fmt"
    "io"
    "io/fs"
    "net/http"
    "os"
)

///////////////////////////////////////////////////////////////////////////////
//
// Is a file a ZIP file or not?

func is_zipfile( filename string ) bool {

    f, err := os.Open( filename )
    if err != nil {
        fmt.Printf( "ERROR: can't read %s: %v\n" , filename , err )
        os.Exit( 1 )
    }
    defer f.Close()

    b := make( []byte , 4 )
    nr, err := f.Read( b )
    if err != nil {
        fmt.Sprintf( "ERROR: can't read 4 bytes from %s: %v\n" , filename , err )
        os.Exit( 1 )
    }

    if nr != 4 {
        fmt.Printf( "ERROR: expected 4 bytes from %s, only got %d" , filename , nr )
        os.Exit( 1 )
    }
    return ( b[0] == 0x50 && b[1] == 0x4b && b[2] == 0x03 && b[3] == 0x04 )
}

///////////////////////////////////////////////////////////////////////////////
//
// Download an RMDOC file

var know_can_rmdoc  bool = false
var can_rmdoc       bool = false

func download_rmdoc( uuid string , localfile string ) {

    ////////////////////////////////////////////////////////////
    // If the output filename contains any directory names,
    // make sure any necessary directories exist.

    for n := 1 ; n < len( localfile ) ; n ++ {
        if localfile[n] == '/' {
            dir := localfile[:n]

            if flag_debug {
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
                    fmt.Printf( "ERROR: %v\n" , err )
                    os.Exit( 1 )
                }

                fmt.Println( " ok" )

            } else if err != nil {
                ////////////////////////////////////////
                // os.Stat() had some other error

                fmt.Printf( "ERROR: os.Stat('%s'): %v\n" , dir , err )
                os.Exit( 1 )

            } else if ( ( s.Mode() & fs.ModeDir ) == 0 ) {
                ////////////////////////////////////////
                // exists and is not a directory

                fmt.Printf( "ERROR: '%s' exists and is not a directory\n" , dir )
                os.Exit( 1 )
            }

        } // if localfile[n] == '/'
    } // for n

    ////////////////////////////////////////////////////////////
    // Download the file

    fmt.Printf( "Downloading '%s' ... " , localfile )

    ////////////////////////////////////////
    // Request the file

    url := "http://" + tablet_addr + "/download/" + uuid + "/rmdoc"

    resp, err := http.Get( url )
    if err != nil {
        fmt.Printf( "ERROR: can't download '%s': %v" , url , err )
        os.Exit( 1 )
    }

    ////////////////////////////////////////
    // Create output file

    dest, err := os.Create( localfile )
    if err != nil {
        fmt.Printf( "ERROR: can't create '%s': %v\n" , localfile , err )
    }

    ////////////////////////////////////////
    // Copy the output to the file

    var src io.Reader = &PassThru{ Reader: resp.Body }

    total, err := io.Copy( dest , src )
    if err != nil {
        fmt.Printf( "ERROR: os.Copy(): %v" , err )
        os.Exit( 1 )
    }

    ////////////////////////////////////////
    // Finished with transfer

    dest.Close()
    resp.Body.Close()
    fmt.Printf( "%d ... ok\n" , total )

    ////////////////////////////////////////
    // Check whether the output file is a ZIP file

    if ! know_can_rmdoc {
        know_can_rmdoc = true

        if is_zipfile( localfile ) {
            can_rmdoc = true
        }

        if ! can_rmdoc {
            ////////////////////////////////////////
            // Remove the file we just downloaded (it isn't an RMDOC file)

            os.Remove( localfile )

            ////////////////////////////////////////
            // Tell the user what's going on

            if ! flag_dl_pdf {
                fmt.Println( "FATAL: this tablet's software cannot download .rmdoc files, cannot continue" )
                os.Exit( 1 )
            } else {
                fmt.Printf( "WARNING: this tablet's software cannot download .rmdoc files\n" )
                flag_dl_rmdoc = false
            }
        }
    }

}
