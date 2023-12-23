///////////////////////////////////////////////////////////////////////////////
//
// rmweb/do_list.go
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

func do_list() {

    the_files := read_files()

    ////////////////////////////////////////
    // Build a list of filenames
    // - the keys in the_files are UUIDs

    var l_name  int = 4     // length of "Name" header
    var l_size  int = 4     // length of "Size" header
    var l_pages int = 5     // length of "Pages" header

    files_by_name := make( []string , 0 , len( the_files ) )
    for uuid := range the_files {
        files_by_name = append( files_by_name , uuid )

        ////////////////////////////////////////
        // Find the length of the longest full_name

        l := len( the_files[uuid].full_name )
        if the_files[uuid].folder {
            l ++
        }

        if l > l_name {
            l_name = l
        }

        ////////////////////////////////////////
        // Find the length of the longest "size"

        l = len( fmt.Sprintf( "%d" , the_files[uuid].size ) )
        if l > l_size {
            l_size = l
        }

        ////////////////////////////////////////
        // Find the length of the longest page count

        l = len( fmt.Sprintf( "%d" , the_files[uuid].pages ) )
        if l > l_pages {
            l_pages = l
        }

    }

    ////////////////////////////////////////
    // Sort the list by fullname

    sortby_name := func( a int , b int ) bool {
        a_name := the_files[files_by_name[a]].full_name
        b_name := the_files[files_by_name[b]].full_name
        return a_name < b_name
    }
    sort.SliceStable( files_by_name , sortby_name )

    ////////////////////////////////////////
    // Print entries

    fmt.Printf( "%-36s %*s %*s %s\n" ,
        "UUID" ,
        l_size  , "Size" ,
        l_pages , "Pages" ,
        "Name" )
    fmt.Printf( "%s %s %s %s\n" ,
        strings.Repeat( "-" , 36      ) ,
        strings.Repeat( "-" , l_size  ) ,
        strings.Repeat( "-" , l_pages ) ,
        strings.Repeat( "-" , l_name  ) )


    for _,uuid := range files_by_name {
        if the_files[uuid].folder {
            fmt.Printf( "%-36s %*s %*s %s/\n" ,
                uuid ,
                l_size  , "" ,
                l_pages , "" ,
                the_files[uuid].full_name )

        } else {
            fmt.Printf( "%-36s %*d %*d %s\n" ,
                uuid ,
                l_size  , the_files[uuid].size  ,
                l_pages , the_files[uuid].pages ,
                the_files[uuid].full_name )
        }
    }

}
