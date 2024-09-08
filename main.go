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
//
// NOTE: if new 'flag_dl_xxx' items are needed to support new file types,
// be sure to update the check at the end of download_rmdoc() to check for
// the new file types.

var flag_debug      bool    = false
var flag_overwrite  bool    = false
var flag_collapse   bool    = false
var tablet_addr     string  = "10.11.99.1"
var flag_dl_pdf     bool    = false         // default is set in main() below
var flag_dl_rmdoc   bool    = false         // default is set in main() below

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

    msg := `%s [options] COMMAND [...]

Download files from a reMarkable tablet.

Commands

    list     ___    List all files on tablet.
    download ___    Download one or more documents to PDF file(s).

    version         Show the program's version info
    help            Show this help message.

Options

-p      Download PDF files.

-d      Download RMDOC files. This requires that the tablet have
        software version version 3.10 or later.

-a      Download all available file types.

-c      Collapse filenames, i.e. don't create any sub-directories.
        All files will be written to the current directory.

-f      Overwrite existing files.

-I ___  Specify the tablet's IP address. Default is '10.11.99.1',
        which the tablet uses when connected via USB cable. Note that
        unless you've "hacked" your tablet, the web interface is not
        available via any interface other than the USB cable.

-D      Show debugging messages.

-h      Show this help message.

If no file types are explicitly requested (i.e. no '-a', '-p' or '-d'
options are used), the program will download PDF files only by default.

Commands with "___" after them allow you to specify one or more patterns
to search for. Only matching documents will be (listed, downloaded, etc.)
If a UUID is specified, that *exact* document will be selected. Otherwise,
all documents whose names (as seen in the tablet's UI) contain the pattern
will be selected.

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
//
// Show deprecation messages

func do_backup() {
    msg := `The 'backup' command has been deprecated.
Please use 'download' instead.
`

    fmt.Print( msg )
    os.Exit( 1 )
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

func main() {

    ////////////////////////////////////////
    // Parse command line options

    var helpme      bool    = false
    var want_pdf    bool    = false
    var want_rmdoc  bool    = false
    var want_all    bool    = false

    flag.Usage = usage
    flag.BoolVar  ( &helpme         , "h" , helpme         , "" )
    flag.BoolVar  ( &flag_debug     , "D" , flag_debug     , "" )
    flag.BoolVar  ( &flag_overwrite , "f" , flag_overwrite , "" )
    flag.BoolVar  ( &flag_collapse  , "c" , flag_collapse  , "" )
    flag.StringVar( &tablet_addr    , "I" , tablet_addr    , "" )
    flag.BoolVar  ( &want_pdf       , "p" , want_pdf       , "" )
    flag.BoolVar  ( &want_rmdoc     , "d" , want_rmdoc     , "" )
    flag.BoolVar  ( &want_all       , "a" , want_all       , "" )
    flag.Parse()

    ////////////////////////////////////////
    // If they used '-h', show usage

    if ( helpme ) {
        usage()
    }

    ////////////////////////////////////////
    // Figure out which file type options were requested

    if want_all {
        flag_dl_pdf     = true
        flag_dl_rmdoc   = true
    } else if want_pdf || want_rmdoc {
        flag_dl_pdf     = want_pdf
        flag_dl_rmdoc   = want_rmdoc
    } else {
        flag_dl_pdf     = true
        flag_dl_rmdoc   = false
    }

    ////////////////////////////////////////
    // Figure out which command we're being asked to run

    if len( flag.Args() ) > 0 {
        switch flag.Args()[0] {
            case "help"     : usage()
            case "version"  : do_version()
            case "backup"   : do_backup()
            case "list"     : do_list     ( flag.Args()[1:]... )
            case "download" : do_download ( flag.Args()[1:]... )
            default         : usage()
        }
    } else {
        usage()
    }

}
