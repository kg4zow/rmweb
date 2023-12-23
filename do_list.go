///////////////////////////////////////////////////////////////////////////////
//
// rmweb/do_list.go
// John Simpson <jms1@jms1.net> 2023-12-17

package main

import (
    "fmt"
    "os"
    "sort"
    "strings"
)

///////////////////////////////////////////////////////////////////////////////
//
// do_list

func do_list( args ...string ) {

    the_files := read_files()

    ////////////////////////////////////////////////////////////
    // Select which UUIDs to show

    show_uuids := make( map[string]bool , len( the_files ) )

    ////////////////////////////////////////
    // If no pattern, include every UUID

    if len( args ) < 1 {
        for uuid,_ := range the_files {
            show_uuids[uuid] = true
        }

        if flag_debug {
            fmt.Printf( "do_list: including all UUIDs\n" )
        }

    ////////////////////////////////////////
    // Otherwise, build a list of matching UUIDs

    } else {
        for _,pattern := range args {
            if flag_debug {
                fmt.Printf( "do_list: pattern '%s'\n" )
            }

            look_for := strings.ToLower( pattern )

            ////////////////////////////////////////
            // Figure out which items match the current pattern

            this_match := match_files( the_files , look_for )

            if len( this_match ) > 0 {
                for _,x := range this_match {
                    show_uuids[x] = true

                    if flag_debug {
                        fmt.Printf( "do_list:   found '%s' '%s'\n" , x ,
                            the_files[x].full_name )
                    }
                }
            } else {
                fmt.Printf( "no matching items found for '%s'\n" , pattern )
            }
        }
    }

    ////////////////////////////////////////
    // Make sure we found *something*

    if len( show_uuids ) < 1 {
        fmt.Println( "ERROR: nothing to search for" )
        os.Exit( 1 )
    }

    ////////////////////////////////////////////////////////////
    // Build a list of filenames
    // - the keys in the_files are UUIDs

    var l_name  int = 4     // length of "Name" header
    var l_size  int = 4     // length of "Size" header
    var l_pages int = 5     // length of "Pages" header

    show_names := make( []string , 0 , len( show_uuids ) )
    for uuid := range show_uuids {
        show_names = append( show_names , uuid )

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
        a_name := the_files[show_names[a]].full_name
        b_name := the_files[show_names[b]].full_name
        return a_name < b_name
    }
    sort.SliceStable( show_names , sortby_name )

    ////////////////////////////////////////////////////////////
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


    for _,uuid := range show_names {
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
