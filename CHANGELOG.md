# CHANGELOG

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
