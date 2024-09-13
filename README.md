&#x1F6D1; **"Cloud Archive" files**

The reMarkable cloud has a feature, available to customers with a paid account, which lets you "archive" documents to the cloud. This means that the document's files are in the cloud, but only the documents *metadata* is stored on your tablet.

At first glance, the web API's "list files" request doesn't seem to provide any way to tell if a document is "archived to the cloud" or not. I don't use the reMarkable cloud, so at the moment I have no way to *create* any "cloud archive" files to experiment with.

In addition, attempting to *download* such a document via the web interface can cause the tablet's internal web server to lock up.

If you have a paid cloud account, make sure all of your documents are on your tablet (i.e. not "archived to the cloud") before using this program.

---

# `rmweb`

John Simpson `<jms1@jms1.net>` 2023-12-17

Last updated 2024-09-12

This program lists all documents on a reMarakble tablet, or downloads them all as PDF files, using the tablet's built-in web interface.

I threw this together after reading [this question](https://www.reddit.com/r/RemarkableTablet/comments/18js4wo/any_way_to_transfer_all_my_files_to_an_ipad_app/) on Reddit.

# Background Information

I figured out the correct web requests by using [Wireshark](https://www.wireshark.org/) to sniff the web traffic between my computer and three different reMarkable tablets (both rM1 and rM2, running a mix of 3.8 and 3.9 software.)

### Go (or Golang)

[Go](https://go.dev/) (or "Golang") is a programming language from Google. I used this language for a few reasons ...

* It's compiled rather than interpreted, which means programs tend to start and run more quickly.

* Users don't have to install any additional libraries or "runtime" packages in order to run the program. All they need is an executable file for the platform where they're going to run it.

* The same code can be compiled for multiple platforms, without having to figure out how to install and configure cross-compiler toolchains. (If you don't know what this means, count yourself lucky that you've never had to think about it.)

* Go is a relatively new language for me, so this was a way to "get some practice" with it.

One of the ideas on my list is a more "general" command line utility, written in Go, which will be able to list, download, upload, and make backups of reMarkable tablets. I'm planning to make this other program use SSH instead of the web interface, however I was going to have to use the web API to download PDF files anyway, so when I get ready to write that other program, I'll be able to copy a few of the functions from this program.

### File Formats

The program will download PDF files by default. These can be viewed and printed on a computer, but any pen strokes you may have written will be "burned into" the file, and cannot be edited if you upload them back to the tablet (or to a different tablet).

The program *can* download `.rmdoc` files, which reMarkable calls "Archive files". Your computer can't do a whole lot with these files, but if you upload them to the tablet (or to a different tablet), they will be edit-able, just like the original document.

Note that reMarkable software versions prior to 3.10 were not capable of downloading "Archive files". (There was a bug in 3.9 where the web interface had an option to download Archive files, but the option didn't work - the files it would download were actually PDF files.)

Because of this. when the program tries to download the first `.rmdoc` file, it will check the first few bytes in the downloaded file to be sure it *is* a valid `.rmdoc` file. If not, it will stop trying to download `.rmdoc` files, and if no other file formats were requested, the program will stop.

# Installing the Program

## Download

I'm using Github's "releases" mechanism. There should be a "Latest release" section to the right, if you're not reading this through the Github web interface, click [here](https://github.com/kg4zow/rmweb/releases).

Download the appropriate executable for the machine where you plan to run the program. Store it in a directory in your `PATH`, and make sure its permissions are set to be executable. I also recommend renaming it to `rmweb`.

## Compiling the Program

If you want to compile the program from source ...

* [Install Go](https://go.dev/doc/install).

* Clone the source code.

    ```
    $ cd ~/git/
    $ git clone https://github.com/kg4zow/rmweb
    ```

    I use `~/git/` as a container for the git repos I have cloned on my machines. Obviously feel free to clone the repo wherever it makes sense for your machine.

* Run `make` in the cloned repo.

    ```
    $ cd ~/git/rmweb/
    $ make
    ```

This will build the correct binary for your computer, under the `out/` directory. It will also create `rmweb` in the current dirctory, as a symbolic link to that binary.

Note that you could also run "`make all`" to build binaries for a list of architectures. (This is how I build the executables when creating a release.) The list of architectures is set in the "`ALL_ARCHES :=`" line in `Makefile`, and currently includes the following:

* `darwin/amd64` (Apple Intel 64-bit)
* `darwin/arm64` (Apple M1/M2)
* `linux/386` (Linux Intel 32-bit)
* `linux/amd64` (Linux Intel 64-bit)
* `windows/386` (windows 32-bit)
* `windows/amd64` (windows 64-bit)

You can see a list of all *possible* `GOOS`/`GOARCH` combinations in your installed copy of `go` by running "`go tool dist list`".

# Set up the tablet

### Connect the tablet via USB cable

The program uses the web interface to talk to the tablet, and reMarkable sets things up so that the web interface is only available over the USB interface.

> &#x2139;&#xFE0F; If your tablet is "hacked" and the web interface is available over some other IP, you can use the "`-I`" (uppercase "i") option to specify that IP, like so:
>
> ```
> $ rmweb -I 192.0.2.7 list
> ```

### Make sure the tablet is not sleeping

Fairly obvious.

### Make sure the web interface is enabled

See "Settings &#x2192; Storage" on the tablet. Note that you won't be able to turn the web interface on unless the tablet is connected to a computer via the USB cable.

# Running the program

The examples below assume that when you installed the executable, you renamed it to "`rmweb`". If not, you'll need to adjust the commands below to use whatever name you gave it.

## Available Options

If you use the `-h` option, or run the program without specifying a command, it will show a quick explanation of how to run the program, along with lists of the commands and options it supports.

```
$ rmweb -h
rmweb [options] COMMAND [...]

Download files from a reMarkable tablet.

Commands

    list     ___    List all files on tablet.
    download ___    Download one or more documents to PDF file(s).

    version         Show the program's version info
    help            Show this help message.

Options

-p      Download PDF files.

-d      Download RMDOC files. This requires that the tablet have
        software version version 3.10 or later.

-a      Download all available file types.

-c      Collapse filenames, i.e. don't create any sub-directories.
        All files will be written to the current directory.

-f      Overwrite existing files.

-I ___  Specify the tablet's IP address. Default is '10.11.99.1',
        which the tablet uses when connected via USB cable. Note that
        unless you've "hacked" your tablet, the web interface is not
        available via any interface other than the USB cable.

-D      Show debugging messages.

-h      Show this help message.

If no file types are explicitly requested (i.e. no '-a', '-p' or '-d'
options are used), the program will download PDF files only by default.

Commands with "___" after them allow you to specify one or more patterns
to search for. Only matching documents will be (listed, downloaded, etc.)
If a UUID is specified, that *exact* document will be selected. Otherwise,
all documents whose names (as seen in the tablet's UI) contain the pattern
will be selected.
```

This example is from v0.07. You may see different output if you're using a different version.

## Check the version

The "`rmweb version`" command will show you the version number, along with information about the specific code it was built from in the git repo.

```
$ rmweb version
rmweb-darwin-arm64 version 0.04
Built 2023-12-18T17:44:17Z from v0.04-0-g17c101b
https://github.com/kg4zow/rmweb/
```

## List Documents on the Tablet

To list the files on the tablet, use "`rmweb list`".

```
$ rmweb list
UUID                                      Size Pages Name
------------------------------------ --------- ----- ------------------------------------------
22e6d931-c205-4d86-b022-04b6a0527b67                 Amateur Radio/
ec1989b1-bc41-40d7-a8f7-768b5be42bfd     38048     1 Amateur Radio/GMRS
f7acd9d5-6c98-475d-b7bf-5cbb31501720     95486     1 Amateur Radio/Icom ID-5100
3314ef15-3b23-49dc-86f4-c89696963076    114720     2 Amateur Radio/Quansheng UV-K5(8) aka UV-K6
9d5198b0-ce76-4f58-ba0a-a1a058678695    277447     1 Amateur Radio/RTL-SDR
f67c74d2-7d23-4587-95bd-7a6e8ebaed2c                 Ebooks/
8efcdb0a-891f-40cb-901f-7c5bc0df7ad1  78249189   541 Ebooks/A City on Mars
9800e36b-5a0c-4eb7-aedc-72c1845d2816   5615420   264 Ebooks/Chokepoint Capitalism
cc2135a3-08ea-4ae5-be77-8f455b039451   8025642   684 Ebooks/The Art of Unix Programming
702ef913-16a0-47b1-806e-1769f251b06b   4241185   306 Ebooks/The Cathedral & the Bazaar
...
24f6b013-054e-4706-9248-3d3d97d0d268    606562     8 Quick sheets
225a451f-61c9-4ffb-96ad-cbe7b7bb530c    987530     9 TODO
015deb02-0589-462a-bc98-3034d7d23628                 Work/
6f3d00ae-925b-48a0-83fb-963644bd7747  19915627   380 Work/2024 Daily
...
```

The UUID values are the internal identifiers for each document. The files within the tablet that make up each document, all have this as part of their filename. If you're curious, [this page](https://remarkable.jms1.info/info/filesystem.html) has a lot more detail.

The size of each document is calculated by the tablet itself. It *looks like* it's the total of the sizes of all files which make up that document, including pen strokes and page thumbnail images. This is not the size you can expect the downloaded files to be. From what I've seen ...

* PDF files can be anywhere from half to five times this size.
* `.rmdoc` files can be anywhere from 0.5 to about 1.1 times this size.

### Patterns

You can specify one or more patterns when listing or downloading files. If you do, any documents whose names match one or more of the patterns will be downloaded.

```
$ rmweb list quick unix
UUID                                    Size Pages Name
------------------------------------ ------- ----- ----------------------------------
cc2135a3-08ea-4ae5-be77-8f455b039451 8025642   684 Ebooks/The Art of Unix Programming
24f6b013-054e-4706-9248-3d3d97d0d268  606562     8 Quick sheets
```

#### Notes

* If you need to specify a pattern containing spaces, you should quote it. For example ...

    ```
    $ rmweb list quick brown fox
    ```

    This command would list any documents with "quick" in the name, OR with "brown" in the name, OR with "fox" in the name.

    ```
    $ rmweb list "quick brown fox"
    ```

    This command would only download documents with the string "quick brown fox" in their name.

* Filename searches are case-*in*sensitive, i.e. "`test`", "`Test`", "`TEST`", and "`tEsT`" all match each other.

* If a pattern *looks like* a UUID (i.e. has the "8-4-4-4-12 hex digits" structure), the program will look it up by UUID rather than searching for it by filename.

* If no patterns are specified (i.e. just "`rmweb list`" or "`rmweb download`" by itself), the program will list or download ALL documents.

## Download Documents

To download one or more individual documents, first `cd` into the directory where you want to download the files.

```
$ mkdir ~/rm2-backup
$ cd ~/rm2-backup/
```

Then, run the appropriate "`rmweb download`" command.

### Selecting file formats

* The `-p` option tells the program to download PDF files.

    ```
    $ rmweb -p download
    ```

* The `-d` option tells the program to download RMDOC files.

    ```
    $ rmweb -d download
    ```

* The `-a` option tells the program to download all available file formats. Currently (for v0.07) this means PDF and RMDOC, but if more file formats are added in the future, this option will include those new formats.

    ```
    $ rmweb -a download
    ```

* If none of these options are used, the program will download just PDF files.

### Selecting which files to download

The `rmweb download` command supports the same pattern-matching options as the `rmweb list` command. The notes in the Patterns section above, apply here as well.

Each pattern can be either a UUID ...

```
$ rmweb download 24f6b013-054e-4706-9248-3d3d97d0d268
Downloading 'Quick sheets.pdf' ... 2577411 ... ok
```

... or a portion of the filename ...

```
$ rmweb -d download quick unix
Creating    'Ebooks' ...ok
Downloading 'Ebooks/The Art of Unix Programming.rmdoc' ... 5804725 ... ok
Downloading 'Quick sheets.rmdoc' ... 347002 ... ok
```

... or nothing at all, in which case it will download all documents.

```
$ rmweb -a download
Creating    'Amateur Radio' ...ok
Downloading 'Amateur Radio/GMRS.pdf' ... 27201 ... ok
Downloading 'Amateur Radio/GMRS.rmdoc' ... 31160 ... ok
Downloading 'Amateur Radio/Icom ID-5100.pdf' ... 48021 ... ok
Downloading 'Amateur Radio/Icom ID-5100.rmdoc' ... 55268 ... ok
Downloading 'Amateur Radio/Quansheng UV-K5(8) aka UV-K6.pdf' ... 69276 ... ok
Downloading 'Amateur Radio/Quansheng UV-K5(8) aka UV-K6.rmdoc' ... 78329 ... ok
Downloading 'Amateur Radio/RTL-SDR.pdf' ... 120784 ... ok
Downloading 'Amateur Radio/RTL-SDR.rmdoc' ... 151231 ... ok
Creating    'Ebooks' ...ok
Downloading 'Ebooks/42.pdf' ... 53960677 ... ok
Downloading 'Ebooks/42.rmdoc' ... 52533573 ... ok
Downloading 'Ebooks/A City on Mars.pdf' ... 78018531 ... ok
Downloading 'Ebooks/A City on Mars.rmdoc' ... 60687232 ... ok
Downloading 'Ebooks/Chokepoint Capitalism.pdf' ... 4259286 ... ok
Downloading 'Ebooks/Chokepoint Capitalism.rmdoc' ... 3732748 ... ok
Downloading 'Ebooks/The Art of Unix Programming.pdf' ... 7856878 ... ok
Downloading 'Ebooks/The Art of Unix Programming.rmdoc' ... 5804725 ... ok
Downloading 'Ebooks/The Cathedral & the Bazaar.pdf' ... 1382013 ... ok
Downloading 'Ebooks/The Cathedral & the Bazaar.rmdoc' ... 3962084 ... ok
...
Downloading 'Quick sheets.pdf' ... 309012 ... ok
Downloading 'Quick sheets.rmdoc' ... 347002 ... ok
Downloading 'TODO.pdf' ... 502225 ... ok
Downloading 'TODO.rmdoc' ... 569832 ... ok
Creating    'Work' ...ok
Downloading 'Work/2024 Daily.pdf' ... 12172136 ... ok
Downloading 'Work/2024 Daily.rmdoc' ... 7485556 ... ok
...
```

### Other notes

* As each file downloads, the program shows a counter of how many bytes have been read from the tablet. For larger files you'll notice a time delay before this counter starts. This delay is when the tablet is building the file.

    [This video](https://jms1.pub/reMarkable/rmweb-time-difference.mov) shows this time delay. along with the difference between how long it takes to generate a PDF file (~25 seconds for this 380-page PDF-backed document) and how long it takes to build an `.rmdoc` file (~5 seconds for the same document).

* If an output file already exists, the program will add "`-1`", "`-2`", etc. to the filename until it finds a name which doesn't already exist. If you want to *overwrite* existing files, use the "`-f`" option.

    ```
    $ rmweb -f download quick
    Downloading 'Quick sheets.pdf' ... 2577411 ... ok
    ```
