# CHANGELOG

## v0.8 - 2024-09-23

Handle directory list JSON not containing file size or page count.
(Looks like reMarkable removed this between 3.10 and 3.14 software.)
https://github.com/kg4zow/rmweb/issues/1

* `read_files()` check for JSON keys before trying to read them.
* `do_list()` Dondt show size or page count if the directory list JSON
  didn't contain them.

## v0.7 - 2024-09-07

Added "download rmdoc" functionality

* added download_rmdoc()
* replaced '/', '\', and ':' in visible names, with underscores
* moved PassThru() (prints byte counts while reading data from a file)
  from download_pdf.go to do_download.go, so other download_xxx() can
  use it
* included '-D' in usage() message

## v0.06 - 2023-12-22

* Added `download` command
    * by UUID or by substring match in "VissibleName"
    * handles multiple substring match patterns
    * same "-1" handling for duplicate output files
* Added `-c` option to "flatten" folder structure
* Added "pattern searching" (UUID or filename substring) to `list` command
* Deprecated `backup` command, use `download` with no pattern instead
* Updated `list` command to do the same UUID/filename search as `download`
* Added `-I` option to set tablet IP address
    * only useful for tablets which have been "hacked" to make the web
      interface available over wifi or other interfaces

## v0.05 - 2023-12-19

* Set things up to automate "publishing" new binaries to Keybase
    * `Makefile`: added `push` target to "publish" new executables
* Handle duplicate output filenames by adding "-1", "-2", etc.
    * added `-f` option to skip this and overwrite existing files
* `download_pdf()` now creates directories as needed
* removed `download_backup_all.go` (forgot to do this in v0.03)
* Updated `README.md`
    * updated info about where/how to download
    * updated examples with new `list` output format

## v0.04 - 2023-12-18

Re-thinking how the executables are distributed.

* Started a new git repo to remove the executables from the repo.
* Changed targets from `out/GOOS-GOARCH/rmweb` to `out/rmweb-GOOS-GOARCH`
* Updated format of `version` message to show the new executable names.

Executables can be downloaded from:

* /keybase/public/jms1/rmweb/
* https://jms1.pub/rmweb/

## v0.03 - 2023-12-18

* Added `TODO.md`
* Renamed `download` command to `backup`
* Changed `-V` option to `version` command
* Updated `list`: add file size, page count

## v0.02 - 2023-12-17

* Updated `README.md` file
* Fix program name in all files

## v0.01 - 2023-12-17

* Initial commit, seems to be working
