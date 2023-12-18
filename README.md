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

I'm also using this as a way to see if I can figure out how to use Github's "releases" mechanism. I know it ties in with tags, but I haven't actually *used* it before.

If it works as expected, you should see a "Releases" section on the right when you look at [the Github repo](https://github.com/kg4zow/rmweb), with downloadable executables for the operating systems I've compiled the program for.

If that isn't there, there should be executables under the `out/` directory in the repo. Download whichever one(s) you need for the system where you'll be running the program.

### Compiling the Program

If you want to compile the program from source ...

* [Install Go](https://go.dev/doc/install).

* Clone the source code.

    ```
    $ cd ~/git/
    $ git clone https://github.com/kg4zow/rm2-download-pdfs
    ```

    I use `~/git/` as a container for the git repos I have cloned on my machines. Obviously feel free to clone the repo wherever it makes sense for your machine.

* Run `make` in the cloned repo.

    ```
    $ cd ~/git/rm2-download-pdfs/
    $ make
    ```

This will build the correct binary for your computer, under the `out/` directory. It will also create `rm2-download-pdfs` in the current dirctory, as a symbolic link to that binary.

Note that you could also run "`make all`", and build binaries for a list of architectures. The list in set in `Makefile`, and currently includes macOS for both Intel and ARM, as well as 32- and 64-bit versions of Linux and Windows.

You can see a list of all possible `GOOS`/`GOARCH` combinations in your installed copy of `go` by running "`go tool dist list`".


## Running the program

### Set up the tablet

* Connect the tablet via USB port.

    The program has the `10.11.99.1` IP address hard-coded into it. I wrote it with the expectation that the IP *could* be change-able in the future, but reMarkable makes it difficult to access the web interface over anything *other than* the USB cable (which is actually a GOOD THING&#x2122; from a security standpoint), it didn't make sense to add a command line option to change the IP it connects to.

* Make sure the tablet is not sleeping.

* Make sure the web interface is enabled. (See "Settings &#x2192; Storage" on the tablet.)

### Run the program


* `cd` into the directory where you want to download the files. The directory doesn't have to be empty, but if it's not, the program will overwrite files which already exist with the same names.

    ```
    $ mkdir ~/backup
    $ cd ~/backup
    ```

* Run the program.

    ```
    $ rm2-download-pdfs
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
    Downloading './Work/2023-12 Daily.pdf' ... 2464473 ... ok
    ```

As each file downloads, the program shows a counter of how many bytes have been read from the tablet. For larger files you'll notice a time delay before this counter starts. This is because the program uses the same API used by the [`http://10.11.99.1/`](http://10.11.99.1) web interface, and this is when the tablet is building the PDF file.
