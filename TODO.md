# TODO

* Handle duplicate output filenames
    * throw a warning/error?
    * detect if output file already exists
    * add "-1", "-2", etc. to filename until we find one that doesn't already exist
    * add option to overwrite existing files

* Add a new `download` command which can ...
    * download a single file by UUID
        * "`-o`" option to specify output PDF filename
    * download only files whose names match a pattern
        * create filenames using "visible" names from tablet
        * option to duplicate or "flatten" folder structure
        * same duplicate output file handling

* Upload PDF/EPUB files
    * option to specify parent folder (by UUID or name)
