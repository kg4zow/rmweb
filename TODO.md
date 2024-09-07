# TODO

* Option to download `.rmdoc` files
    * only works with 3.10 and later, how to check?
        * if first request succeeds, check the output file to see if it's actually a ZIP file or not
    * download `.rmn` files? (download `.rmdoc` then convert)

* Upload PDF, EPUB, and `.rmdoc` files
    * how to detect input file type
        * does golang have a file-type identifier module?
        * if not, write one?
        * require that each filename ends with `.pdf`, `.epub`, or `.rmdoc`?
    * option to specify parent folder (by UUID or name)
    * upload `.rmn` files (convert to `.rmdoc` first), for [drawj2d](https://drawj2d.sourceforge.io/) users)

## Requested by others

* Match: all files in a given folder
    * [requested](https://old.reddit.com/r/RemarkableTablet/comments/1e2ea01/bulk_exporting_documents/ldgge6f/)

* Option to skip existing files, rather than overwriting or renaming
    * [requested](https://old.reddit.com/r/RemarkableTablet/comments/1e2ea01/bulk_exporting_documents/ldgge6f/)
