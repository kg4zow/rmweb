///////////////////////////////////////////////////////////////////////////////
//
// rmweb/download_all_pdf.go
// John Simpson <jms1@jms1.net> 2023-12-17

package main

import (
    "fmt"
    "io/fs"
    "log"
    "os"
    "sort"
)

///////////////////////////////////////////////////////////////////////////////
//
// Download PDFs of all documents on a tablet

func download_all_pdf( the_files map[string]DocInfo ) {

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
        if the_files[uuid].folder {
            ////////////////////////////////////////
            // It's a folder, make sure it exists on the computer

            lname := "." + the_files[uuid].full_name

            fmt.Printf( "Creating    '%s' ... " , lname )

            s,err := os.Stat( lname )
            if os.IsNotExist( err ) {
                ////////////////////////////////////////
                // doesn't exist yet - create it

                err := os.Mkdir( lname , 0755 )
                if err != nil {
                    log.Fatal( err )
                }
                fmt.Println( "ok" )
            } else if err != nil {
                ////////////////////////////////////////
                // some other error

                log.Fatal( err )
            } else if ( ( s.Mode() & fs.ModeDir ) != 0 ) {
                ////////////////////////////////////////
                // exists and is a directory

                fmt.Println( "already exists" )
            } else {
                ////////////////////////////////////////
                // exists and is not a directory

                msg := fmt.Sprintf( "ERROR: '%s' already exists and is not a directory\n" , lname )
                log.Fatal( msg )
            }

        } else {
            ////////////////////////////////////////
            // Not a folder, it's a file. Download it.

            lname := "." + the_files[uuid].full_name + ".pdf"

            fmt.Printf( "Downloading '%s' ... " , lname )

            download_pdf( uuid , lname )
        }
    }


}
