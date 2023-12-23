///////////////////////////////////////////////////////////////////////////////
//
// rmweb/match_files.go
// John Simpson <jms1@jms1.net> 2023-12-22

package main

import (
    "fmt"
    "regexp"
    "strings"
)

///////////////////////////////////////////////////////////////////////////////
//
// Return a list of UUIDs which match a given pattern
// - if the pattern is a UUID, and that UUID exists, return that UUID
// - otherwise match against the files' "find_by" value

func match_files( the_files map[string]DocInfo , look_for string ) []string {

    if flag_debug {
        fmt.Printf( "match_files: looking for '%s'\n" , look_for )
    }

    rv := make( []string , 0 , len( the_files ) )

    ////////////////////////////////////////
    // If we're looking for a UUID, either it exists or it doesnt.

    re_uuid := "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
    is_uuid := regexp.MustCompile( re_uuid )

    if is_uuid.Match( []byte( look_for ) ) {
        if v,present := the_files[look_for] ; present {
            if ! v.folder {
                rv = append( rv , look_for )

                if flag_debug {
                    fmt.Printf( "match_files:   found by UUID '%s'\n" , look_for )
                }
            }
        }
    } else {

    ////////////////////////////////////////
    // Otherwise, search for matching strings in the_files[].find_by

        for k,_ := range the_files {
            if strings.Contains( the_files[k].find_by , look_for ) {
                rv = append( rv , k )

                if flag_debug {
                    fmt.Printf( "match_files:   found by name '%s' '%s'\n" ,
                        k , the_files[k].full_name )
                }
            }
        }
    }

    ////////////////////////////////////////
    // Done

    if flag_debug {
        fmt.Printf( "match_files: returning %d items\n" , len( rv ) )
    }

    return rv
}
