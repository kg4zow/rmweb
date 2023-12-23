///////////////////////////////////////////////////////////////////////////////
//
// rmweb/do_download.go
// John Simpson <jms1@jms1.net> 2023-12-22

package main

import (
    "fmt"
    "os"
    "sort"
    "strings"
)

///////////////////////////////////////////////////////////////////////////////
//
// Select and download one or more files

func do_download( args ...string ) {

    ////////////////////////////////////////
    // Make sure we *have* something to look for.

    if len( args ) < 1 {
        fmt.Println( "ERROR: 'download' requires a UUID or filename\n" )
        os.Exit( 1 )
    }

    ////////////////////////////////////////
    // Read the contents of the tablet

    the_files := read_files()

    ////////////////////////////////////////////////////////////
    // Build the list of UUIDs to be downloaded

    get_uuids := make( map[string]bool , len( the_files ) )

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

            lname := the_files[uuid].full_name + ".pdf"
            if ! flag_overwrite {
                lname = safe_filename( lname )
            }

            download_pdf( uuid , lname )
        }
    }

}
