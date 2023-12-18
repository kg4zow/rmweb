///////////////////////////////////////////////////////////////////////////////
//
// rm2-download-pdfs/list.go
// John Simpson <jms1@jms1.net> 2023-12-17

package main

import (
    "fmt"
    "sort"
    "strings"
)

///////////////////////////////////////////////////////////////////////////////
//
// do_list

func do_list( the_files map[string]DocInfo ) {

    ////////////////////////////////////////
    // Build and sort a list of filenames
    // - the keys in the_files are UUIDs

    l_name := 4     // length of "Name" header

    files_by_name := make( []string , 0 , len( the_files ) )
    for uuid := range the_files {
        files_by_name = append( files_by_name , uuid )

        ////////////////////////////////////////
        // Also find the length of the longest full_name

        l := len( the_files[uuid].full_name )
        if the_files[uuid].folder {
            l ++
        }

        if l > l_name {
            l_name = l
        }
    }

    sortby_name := func( a int , b int ) bool {
        a_name := the_files[files_by_name[a]].full_name
        b_name := the_files[files_by_name[b]].full_name
        return a_name < b_name
    }
    sort.SliceStable( files_by_name , sortby_name )

    ////////////////////////////////////////
    // Process entries

    h_uuid  := strings.Repeat( "-" , 36 )
    h_name  := strings.Repeat( "-" , l_name )

    fmt.Printf( "%-36s %s\n" , "UUID" , "Name" )
    fmt.Printf( "%s %s\n" , h_uuid , h_name )

    for _,uuid := range files_by_name {
        suffix := ""
        if the_files[uuid].folder {
            suffix = "/"
        }

        fmt.Printf( "%-36s %s%s\n" , uuid , the_files[uuid].full_name , suffix )
    }

}
