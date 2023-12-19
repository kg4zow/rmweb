///////////////////////////////////////////////////////////////////////////////
//
// rmweb/read_files.go
// John Simpson <jms1@jms1.net> 2023-12-16
//
// List files on a reMarkable tablet

package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
//    "ioutil"
    "log"
    "net/http"
//    "strconv"
)

///////////////////////////////////////////////////////////////////////////////
//
// read_files
//
// Send a series of "POST http://10.11.99.1/documents/" requests to retrieve
// the documents and containers in the tablet.
//
// Returns a map of IDs pointing to DocInfo structures.

func read_files() ( map[string]DocInfo ) {

    ////////////////////////////////////////
    // Start map to store file/dir info

    rv := make( map [string]DocInfo )

    ////////////////////////////////////////////////////////////
    // Process directories until there are no more

    l_dirs := []string{ "" }
    for len( l_dirs ) > 0 {

        ////////////////////////////////////////
        // Get the first directory name from the array

        this_dir    := l_dirs[0]
        l_dirs      =  l_dirs[1:]

        ////////////////////////////////////////
        // Request info about this directory

        url             := "http://" + tablet_addr + "/documents" + this_dir + "/"
        content_type    := "application/json"
        buf             := bytes.NewBufferString( "" )

        if do_debug {
            fmt.Println( "POST " + url )
        }

        resp, err := http.Post( url , content_type , buf )
        if err != nil {
            log.Fatal( err )
        }

        defer resp.Body.Close()

        ////////////////////////////////////////
        // Read the response into memory

        resp_bytes,err := io.ReadAll( resp.Body )
        if ( err != nil ) {
            log.Fatal( err )
        }

        ////////////////////////////////////////
        // Parse the response

        var data []map[string]interface{}

        err = json.Unmarshal( resp_bytes , &data )
        if err != nil {
            log.Fatal(err)
        }

        ////////////////////////////////////////
        // process items within response

        for _,v := range data {

            ////////////////////////////////////////
            // Get info about this item

            var size    int64
            var pages   int

            id          := v["ID"].(string)
            parent      := v["Parent"].(string)
            folder      := bool( v["Type"].(string) == "CollectionType" )
            name        := v["VissibleName"].(string)

            if ! folder {
                fmt.Sscan( v["sizeInBytes"].(string) , &size )
                pages       = int( v["pageCount"].(float64) )
            }

            if do_debug {
                fmt.Printf( "%s  %-5t  %s\n" , id , folder , name )
            }

            ////////////////////////////////////////
            // Build user-facing name for this item

            parent_name := ""
            if parent != "" {
                parent_name = rv[parent].full_name
            }

            full_name := parent_name + "/" + name
            if full_name[0] == '/' {
                full_name = full_name[1:]
            }

            if do_debug {
                fmt.Printf( "id='%s'\n        name='%s'\n      parent='%s'\n parent_name='%s'\n        full_name='%s'\n\n" ,
                    id , name , parent , parent_name , full_name )
            }

            ////////////////////////////////////////
            // Remember this item

            var f DocInfo

            f.id            = id
            f.parent        = parent
            f.folder        = folder
            f.name          = name
            f.full_name     = full_name
            f.size          = int64( size )
            f.pages         = int64( pages )

            rv[f.id] = f

            ////////////////////////////////////////
            // If this item is a folder, add it to the list
            // so it also gets scanned

            if folder {
                l_dirs = append( l_dirs , string( this_dir + "/" + id ) )
            }

        } // for range data
    } // for len( l_dirs ) > 0

    ////////////////////////////////////////
    // Return the files and directories

    return rv
}
