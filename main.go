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

var flag_debug      bool    = false
var flag_overwrite  bool    = false
var flag_collapse   bool    = false

// future
var tablet_addr     string  = "10.11.99.1"

////////////////////////////////////////
// All files and directories on the tablet

type DocInfo struct {
    id              string
    parent          string
    folder          bool
    name            string
    full_name       string
    size            int64
    pages           int64
    find_by         string
}

///////////////////////////////////////////////////////////////////////////////
//
// usage

func usage( ) {

    msg := `%s [options] COMMAND

Download PDF files from a reMarkable tablet

Commands

    list            List all files on tablet
    backup          Download ALL files on tablet, to PDF files
    download ___    Download one or more files, to PDF file(s)

    version         Show the program's version info
    help            Show this help message.

Options

    -c      Collapse output filenames. All output files will be written to
            the current directory, no sub-directories will be created.

    -f      Overwrite existing files.

    -h      Show this help message.

`

    fmt.Printf( msg , prog_name )

    os.Exit( 0 )
}

///////////////////////////////////////////////////////////////////////////////
//
// Show version info

func do_version( args ...string ) {
    fmt.Printf( "%s-%s-%s version %s\n" ,
        prog_name , runtime.GOOS , runtime.GOARCH , prog_version )

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

    flag.Usage = usage
    flag.BoolVar( &helpme         , "h" , helpme         , "" )
    flag.BoolVar( &flag_debug     , "D" , flag_debug     , "" )
    flag.BoolVar( &flag_overwrite , "f" , flag_overwrite , "" )
    flag.BoolVar( &flag_collapse  , "c" , flag_collapse  , "" )
    flag.Parse()

    ////////////////////////////////////////
    // If they used '-h', show usage

    if ( helpme ) {
        usage()
    }

    ////////////////////////////////////////
    // Figure out which command we're being asked to run

    if len( flag.Args() ) > 0 {
        switch flag.Args()[0] {
            case "help"     : usage()
            case "version"  : do_version()
            case "backup"   : do_backup()
            case "list"     : do_list()
            case "download" : do_download( flag.Args()[1:]... )
            default         : usage()
        }
    } else {
        usage()
    }


}
