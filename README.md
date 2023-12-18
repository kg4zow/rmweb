# `rmweb`

John Simpson `<jms1@jms1.net>` 2023-12-17

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

### Download from Github

In the Github repo, there are directories under [`out/`](https://github.com/kg4zow/rmweb/tree/main/out) which contain compiled executables for the following platforms:

* `darwin-amd64` = macOS, Intel
* `darwin-arm64` = macOS, ARM (aka Apple Silicon)
* `linux-386` = Linux, 32-bit Intel
* `linux-amd64` = Linux, 64-bit Intel
* `windows-386` = Windows, 32-bit Intel
* `windows-amd64` = Windows, 64-bit Intel

Download whichever binaries you need for the system(s) where you'll be running the program. Store them in a directory in your `PATH`, and make sure their permissions are set to be executable (i.e. "`chmod u=rwx,go=rx rmweb`").

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

Note that you could also run "`make all`" to build binaries for a list of architectures. (This is what I do before committing changes and pushing them to Github.) The list of architectures in set in the "`ALL_ARCHES :=`" line in `Makefile`, and currently includes the list shown above.

You can see a list of all *possible* `GOOS`/`GOARCH` combinations in your installed copy of `go` by running "`go tool dist list`".


## Running the program

### Set up the tablet

* Connect the tablet via USB port.

    The program has the `10.11.99.1` IP address hard-coded into it. I wrote it with the expectation that the IP *could* be change-able in the future, but reMarkable makes it difficult to access the web interface over anything *other than* the USB cable (which is actually a GOOD THING&#x2122; from a security standpoint), it didn't make sense to add a command line option to change the IP it connects to.

* Make sure the tablet is not sleeping.

* Make sure the web interface is enabled. See "Settings &#x2192; Storage" on the tablet.

    Note that you won't be able to turn the web interface on unless the tablet is connected to a computer via the USB cable.

### List files on the tablet

To list the files on the tablet, use "`rmweb list`".

```
$ rmweb list
UUID                                 Name
------------------------------------ -----------------------------------
22e6d931-c205-4d86-b022-04b6a0527b67 /Amateur Radio/
34f04076-4733-40bc-983d-6a92b41a2728 /Amateur Radio/D-STAR
9d5198b0-ce76-4f58-ba0a-a1a058678695 /Amateur Radio/RTL-SDR
89f3cbbd-dba2-46aa-bc66-debdf97b915a /Documentation to write
f67c74d2-7d23-4587-95bd-7a6e8ebaed2c /Ebooks/
cc2135a3-08ea-4ae5-be77-8f455b039451 /Ebooks/The Art of Unix Programming
702ef913-16a0-47b1-806e-1769f251b06b /Ebooks/The Cathedral & the Bazaar
...
9e6891eb-2558-4e70-b6fc-d03b2d75614b /Quick sheets
383dad70-b9db-4a04-a275-be17cfc6bc8c /ReMarkable 2 Info
225a451f-61c9-4ffb-96ad-cbe7b7bb530c /TODO
015deb02-0589-462a-bc98-3034d7d23628 /Work/
0d318d48-d638-4c4c-9e29-98bac57bb658 /Work/2023-11 Daily
66d3acae-9697-4a10-b827-3e619af36fae /Work/2023-12 Daily
```

The UUID values are the internal identifiers for each document. The files within the tablet that make up each document all have this as part of their filename. If you're curious, this is explain in more detail on [this page](https://remarkable.jms1.info/info/filesystem.html). You may never need to know this value, but if you *do* need them, there they are.

### Download documents to PDF files

To download the documents as PDF files, first `cd` into the directory where you want to download the files. The directory doesn't have to be empty, but if it's not, the program will overwrite files which already exist with the same names.

```
$ mkdir ~/rm2-backup
$ cd ~/rm2-backup/
```

Then, run "`rmweb download`".

```
$ rmweb download
Creating    './Amateur Radio' ... ok
Downloading './Amateur Radio/D-STAR.pdf' ... 1792627 ... ok
Downloading './Amateur Radio/RTL-SDR.pdf' ... 120895 ... ok
Downloading './Documentation to write.pdf' ... 674706 ... ok
Creating    './Ebooks' ... ok
Downloading './Ebooks/The Art of Unix Programming.pdf' ... 7856878 ... ok
Downloading './Ebooks/The Cathedral & the Bazaar.pdf' ... 1382013 ... ok
...
Downloading './Quick sheets.pdf' ... 2577411 ... ok
Downloading './ReMarkable 2 Info.pdf' ... 1451907 ... ok
Downloading './TODO.pdf' ... 258263 ... ok
Creating    './Work' ... ok
Downloading './Work/2023-11 Daily.pdf' ... 5114386 ... ok
Downloading './Work/2023-12 Daily.pdf' ... 2464473 ... ok
```

As each file downloads, the program shows a counter of how many bytes have been read from the tablet. For larger files you'll notice a time delay before this counter starts. This is because the program uses the same API used by the [`http://10.11.99.1/`](http://10.11.99.1) web interface, and this is when the tablet is building the PDF file.
