# `rmweb`

John Simpson `<jms1@jms1.net>` 2023-12-17

Last updated 2023-12-19

This program lists all documents on a reMarakble tablet, or downloads them all as PDF files, using the tablet's built-in web interface.

I threw this together after reading [this question](https://www.reddit.com/r/RemarkableTablet/comments/18js4wo/any_way_to_transfer_all_my_files_to_an_ipad_app/) on Reddit.

I figured out the correct web requests by using [Wireshark](https://www.wireshark.org/) to sniff the web traffic between my computer and three different reMarkable tablets (both rM1 and rM2, running a mix of 3.8 and 3.9 software.)

### Go

[Go](https://go.dev/) (or "Golang") is a programming language from Google. I used this language for a few reasons ...

* It's compiled rather than interpreted, which means programs tend to start and run more quickly.

* Users don't have to install any additional libraries or "runtime" packages in order to run the program. All they need is an executable file for the platform where they're going to run it.

* The same code can be compiled for multiple platforms, without having to figure out how to install and configure cross-compiler toolchains. (If you don't know what this means, count yourself lucky that you've never had to think about it.)

* Go is a relatively new language for me, so this was a way to "get some practice" with it.

One of the ideas on my list is a more "general" command line utility, written in Go, which will be able to list, download, upload, and make backups of reMarkable tablets. I'm planning to make this other program use SSH instead of the web interface, however I was going to have to use the web API to download PDF files anyway, so when I get ready to write that other program, I'll be able to copy a few of the functions from this program.

## Installing the Program

### Download

I'm using Github's "releases" mechanism. There should be a "Latest release" section to the right, if you're not reading this through the Github web interface, click [here](https://github.com/kg4zow/rmweb/releases).

Download the appropriate executable for the machine where you plan to run the program. Store it in a directory in your `PATH`, and make sure its permissions are set to be executable. I also recommend renaming it to `rmweb`.

### Compiling the Program

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


## Running the program

The examples below assume that when you installed the executable, you renamed it to "`rmweb`". If not, you'll need to adjust the commands below to use whatever name you gave it.

### Check the version

The "`rmweb version`" command will show you the version number, along with information about the specific code it was built from in the git repo.

```
$ rmweb version
rmweb-darwin-arm64 version 0.04
Built 2023-12-18T17:44:17Z from v0.04-0-g17c101b
https://github.com/kg4zow/rmweb/
```

### Set up the tablet

**Connect the tablet via USB cable.**

The program uses the web interface to talk to the tablet, and reMarkable sets things up so that the web interface is only available over the USB interface.

> &#x2139;&#xFE0F; If your tablet is "hacked" and the web interface is available over some other IP, you can use the "`-I`" (uppercase "i") option to specify that IP, like so:
>
> ```
> $ rmweb -I 192.0.2.7 list
> ```

**Make sure the tablet is not sleeping.**

Fairly obvious.

**Make sure the web interface is enabled.**

See "Settings &#x2192; Storage" on the tablet. Note that you won't be able to turn the web interface on unless the tablet is connected to a computer via the USB cable.

### List files on the tablet

To list the files on the tablet, use "`rmweb list`".

```
$ rmweb list
UUID                                    Size Pages Name
------------------------------------ ------- ----- -----------------------------------
22e6d931-c205-4d86-b022-04b6a0527b67               /Amateur Radio/
34f04076-4733-40bc-983d-6a92b41a2728  311475     7 /Amateur Radio/D-STAR
9d5198b0-ce76-4f58-ba0a-a1a058678695  277447     1 /Amateur Radio/RTL-SDR
89f3cbbd-dba2-46aa-bc66-debdf97b915a 1431776     8 /Documentation to write
f67c74d2-7d23-4587-95bd-7a6e8ebaed2c               /Ebooks/
cc2135a3-08ea-4ae5-be77-8f455b039451 8025642   684 /Ebooks/The Art of Unix Programming
702ef913-16a0-47b1-806e-1769f251b06b 4241185   306 /Ebooks/The Cathedral & the Bazaar
...
9e6891eb-2558-4e70-b6fc-d03b2d75614b 5961550    52 /Quick sheets
383dad70-b9db-4a04-a275-be17cfc6bc8c 2658181    27 /ReMarkable 2 Info
225a451f-61c9-4ffb-96ad-cbe7b7bb530c  597041     3 /TODO
015deb02-0589-462a-bc98-3034d7d23628               /Work/
0d318d48-d638-4c4c-9e29-98bac57bb658 3203221    21 /Work/2023-11 Daily
66d3acae-9697-4a10-b827-3e619af36fae 1692264    10 /Work/2023-12 Daily
```

The UUID values are the internal identifiers for each document. The files within the tablet that make up each document, all have this as part of their filename. If you're curious, [this page](https://remarkable.jms1.info/info/filesystem.html) has a lot more detail.

The size of each document is calculated by the tablet itself. It *looks like* it's the total of the sizes of all files which make up that document, including pen strokes and page thumbnail images. This is not the size you can expect the PDF to be if you download the file, from what I've seen the downloaded PDFs end up being anywhere from half to five times this size.

### Download documents to PDF files

To download one or more individual documents as PDF files, first `cd` into the directory where you want to download the files.

```
$ mkdir ~/rm2-backup
$ cd ~/rm2-backup/
```

Then, run "`rmweb download PATTERN`". The pattern can be either a UUID ...

```
$ rmweb download 9e6891eb-2558-4e70-b6fc-d03b2d75614b
Downloading 'Quick sheets.pdf' ... 2577411 ... ok
```

... or a portion of the filename ...

```
$ rmweb download quick
Downloading 'Quick sheets-1.pdf' ... 2577411 ... ok
```

... or nothing at all, in which case it will download all documents.

```
$ rmweb download
Creating    'Amateur Radio' ... ok
Downloading 'Amateur Radio/D-STAR.pdf' ... 1792627 ... ok
Downloading 'Amateur Radio/RTL-SDR.pdf' ... 120895 ... ok
Downloading 'Documentation to write.pdf' ... 674706 ... ok
Creating    'Ebooks' ... ok
Downloading 'Ebooks/The Art of Unix Programming.pdf' ... 7856878 ... ok
Downloading 'Ebooks/The Cathedral & the Bazaar.pdf' ... 1382013 ... ok
...
Downloading 'Quick sheets.pdf' ... 2577411 ... ok
Downloading 'ReMarkable 2 Info.pdf' ... 1451907 ... ok
Downloading 'TODO.pdf' ... 258263 ... ok
Creating    'Work' ... ok
Downloading 'Work/2023-11 Daily.pdf' ... 5114386 ... ok
Downloading 'Work/2023-12 Daily.pdf' ... 2464473 ... ok
```

Notes about patterns:

* If multiple patterns are specified, all files which match *any* of the patterns will be downloaded.

* If you need to specify a pattern containing spaces, you should quote it. For example ...

    ```
    $ rmweb download quick brown fox
    ```

    This command would download any documents with "quick" in the name, OR with "brown" in the name, OR with "fox" in the name.

    ```
    $ rmweb download "quick brown fox"
    ```

    This command would only download documents with the string "quick brown fox" in their name.

* Filename searches are case-*in*sensitive, i.e. "`test`", "`Test`", "`TEST`", and "`tEsT`" all match each other.

* If a pattern *looks like* a UUID (i.e. has the "8-4-4-4-12 hex digits" structure), the program will look it up by UUID rather than searching for it by filename.

* If no patterns are specified (i.e. just "`rmweb download`" by itself), the program will download ALL documents.

Other notes:

* As each file downloads, the program shows a counter of how many bytes have been read from the tablet. For larger files you'll notice a time delay before this counter starts. This is because the program uses the same API used by the [`http://10.11.99.1/`](http://10.11.99.1) web interface, and this delay is when the tablet is building the PDF file.

* If an output file already exists, the program will add "`-1`", "`-2`", etc. to the filename until it finds a name which doesn't already exist. If you want to *overwrite* existing files, use the "`-f`" option.

    ```
    $ rmweb -f download quick
    Downloading 'Quick sheets.pdf' ... 2577411 ... ok
    ```
