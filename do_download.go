///////////////////////////////////////////////////////////////////////////////
//
// rmweb/do_download.go
// John Simpson <jms1@jms1.net> 2023-12-22

package main

import (
    "fmt"
    "io"
    "os"
    "sort"
    "strings"
)

///////////////////////////////////////////////////////////////////////////////
//
// Passthru wrapper for io.Reader, prints total bytes while reading
// used by download_xxx() functions

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
// Select and download one or more files

func do_download( args ...string ) {

    ////////////////////////////////////////
    // Read the contents of the tablet

    the_files := read_files()

    ////////////////////////////////////////////////////////////
    // Figure out which UUIDs we'll be downloading

    get_uuids := make( map[string]bool , len( the_files ) )

    ////////////////////////////////////////
    // If no pattern, include every UUID

    if len( args ) < 1 {
        for uuid,_ := range the_files {
            get_uuids[uuid] = true
        }

        if flag_debug {
            fmt.Printf( "do_list: including all UUIDs\n" )
        }

    ////////////////////////////////////////
    // Otherwise, process each pattern

    } else {
        for _,pattern := range args {
            look_for := strings.ToLower( pattern )

            ////////////////////////////////////////
            // Figure out which items match the current pattern

            this_match := match_files( the_files , look_for )

            if len( this_match ) > 0 {
                for _,x := range this_match {
                    get_uuids[x] = true
                }
            } else {
                fmt.Printf( "no matching items found for '%s'\n" , pattern )
            }
        }
    }

    ////////////////////////////////////////
    // Make sure we found *something*

    if len( get_uuids ) < 1 {
        fmt.Println( "ERROR: nothing to search for" )
        os.Exit( 1 )
    }

    ////////////////////////////////////////////////////////////
    // Build and sort a list of filenames

    var get_names []string

    for uuid,_ := range get_uuids {
        get_names = append( get_names , uuid )
    }

    sortby_name := func( a int , b int ) bool {
        a_name := the_files[get_names[a]].full_name
        b_name := the_files[get_names[b]].full_name
        return a_name < b_name
    }
    sort.SliceStable( get_names , sortby_name )

    ////////////////////////////////////////////////////////////
    // Process entries

    for _,uuid := range get_names {
        if ! the_files[uuid].folder {

            ////////////////////////////////////////
            // Download the file

            lname_pdf   := the_files[uuid].full_name + ".pdf"
            lname_rmdoc := the_files[uuid].full_name + ".rmdoc"

            if ! flag_overwrite {
                lname_pdf   = safe_filename( lname_pdf   )
                lname_rmdoc = safe_filename( lname_rmdoc )
            }

            if flag_dl_pdf {
                download_pdf( uuid , lname_pdf )
            }

            if flag_dl_rmdoc {
                download_rmdoc( uuid , lname_rmdoc )
            }
        }
    }

}
