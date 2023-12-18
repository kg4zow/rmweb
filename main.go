///////////////////////////////////////////////////////////////////////////////
//
// rmweb/main.go
// John Simpson <jms1@jms1.net> 2023-12-16

package main

import (
    "flag"
    "fmt"
    "os"
    "runtime"
)

///////////////////////////////////////////////////////////////////////////////
//

var (
    // Actual values will be filled by -X options in Makefile, these values
    // are here in case somebody uses 'go run .'.

    prog_name       string = "rmweb"
    prog_version    string = "(unset)"
    prog_date       string = "(unset)"
    prog_hash       string = ""
    prog_desc       string = ""

    // Hard-coded, not set by 'make'

    prog_url        string = "https://github.com/kg4zow/rmweb/"
)

///////////////////////////////////////////////////////////////////////////////
//
// Values will be set by command line options

var do_debug        bool    = false
var tablet_addr     string  = "10.11.99.1"

////////////////////////////////////////
// All files and directories on the tablet

type DocInfo struct {
    id              string
    parent          string
    folder          bool
    name            string
    full_name       string
    tablet_name     string
}

///////////////////////////////////////////////////////////////////////////////
//
// usage

func usage( ) {

    msg := `%s [options] COMMAND

Download PDF files from a reMarkable tablet

Commands

    list        List all files on tablet
    download    Download all files on tablet, to PDF files

Options

    -V      Show version info.

    -h      Show this help message.
`

    fmt.Printf( msg , prog_name )

    os.Exit( 0 )
}

///////////////////////////////////////////////////////////////////////////////
//
// Show version info

func do_version( args ...string ) {
    fmt.Printf( "%s\nVersion %s for %s/%s\n" ,
        prog_name , prog_version , runtime.GOOS , runtime.GOARCH )

    if prog_desc != "" {
        fmt.Printf( "Built %s from %s\n" , prog_date , prog_desc )
    } else if prog_hash != "" {
        fmt.Printf( "Built %s from commit %s\n" , prog_date , prog_hash )
    } else {
        fmt.Printf( "Built %s\n" , prog_date )
    }

    fmt.Println( prog_url )
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

func main() {

    ////////////////////////////////////////
    // Parse command line options

    var helpme  = false
    var version = false

    flag.Usage = usage
    flag.BoolVar( &helpme   , "h" , helpme   , "" )
    flag.BoolVar( &do_debug , "D" , do_debug , "" )
    flag.BoolVar( &version  , "V" , version  , "" )
    flag.Parse()

    ////////////////////////////////////////
    // If they used '-h', show usage

    if ( helpme ) {
        usage()
    }

    ////////////////////////////////////////
    // If they used '-V', show version info

    if ( version ) {
        do_version()
        os.Exit( 0 )
    }

    ////////////////////////////////////////
    // Figure out which command we're being asked to run

    if len( flag.Args() ) > 0 {
        switch flag.Args()[0] {
            case "download" : download_all_pdf( read_files() )
            case "list"     : do_list( read_files() )
            default         : usage()
        }
    } else {
        usage()
    }


}
