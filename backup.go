///////////////////////////////////////////////////////////////////////////////
//
// rmweb/backup.go
// John Simpson <jms1@jms1.net> 2023-12-17

package main

import (
///     "fmt"
///     "io/fs"
///     "os"
    "sort"
)

///////////////////////////////////////////////////////////////////////////////
//
// Download PDFs of all documents on a tablet

func backup( the_files map[string]DocInfo ) {

    ////////////////////////////////////////
    // Build and sort a list of filenames
    // - the keys in the_files are UUIDs

    files_by_name := make( []string , 0 , len( the_files ) )
    for uuid := range the_files {
        files_by_name = append( files_by_name , uuid )
    }

    sortby_name := func( a int , b int ) bool {
        a_name := the_files[files_by_name[a]].full_name
        b_name := the_files[files_by_name[b]].full_name
        return a_name < b_name
    }
    sort.SliceStable( files_by_name , sortby_name )

    ////////////////////////////////////////////////////////////
    // Process entries

    for _,uuid := range files_by_name {
        if ! the_files[uuid].folder {

            ////////////////////////////////////////
            // Download the file

            lname := the_files[uuid].full_name + ".pdf"
            if ! do_overwrite {
                lname = safe_filename( lname )
            }

            download_pdf( uuid , lname )
        }
    }


}
