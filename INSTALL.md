
Installation
============

*dataciteapi* is a command line program run from a shell like Bash. It allows you to retrieve JSON works objects from the DataCite API via DataCites REST API.

This is generalized instructions for a release.  For deployment suggestions see NOTES.md

Quick install with curl or irm
------------------------------

If you are running macOS or Linux you can install released versions of crossrefapi
with the following curl command.

~~~
curl https://caltechlibrary.github.io/crossrefapi/installer.sh
~~~

On Windows the install command in Powershell is

~~~
irm https://caltechlibrary.github.io/crossrefapi/installer.ps1 | iex
~~~


Compiled version
----------------

Compiled versions are available for Mac OS X (amd64 and M1 processor, macos-amd64, macos-arm64), Linux (amd64 process, linux-amd64), 
Windows (amd64 processor, windows-amd64) and Rapsberry Pi (arm7 processor, raspbian-arm7)

VERSION_NUMBER is a [symantic version number](http://semver.org/) (e.g. v0.1.2)


For all the released version go to the project page on Github and click latest release

>    https://github.com/caltechlibrary/dataciteapi/releases/latest


| Platform    | Zip Filename                             |
|-------------|------------------------------------------|
| Windows     | dataciteapi-VERSION_NUMBER-windows-amd64.zip |
| Mac OS X    | dataciteapi-VERSION_NUMBER-macos-amd64.zip  |
| Linux/Intel | dataciteapi-VERSION_NUMBER-linux-amd64.zip   |
| Raspbery Pi | dataciteapi-VERSION_NUMBER-raspbian-arm7.zip |


The basic recipe
----------------

+ Find the Zip file listed matching the architecture you're running and download it
    + (e.g. if you're on a Windows 10 laptop/Surface with a amd64 style CPU you'd choose the Zip file with "windows-amd64" in the name).
+ Download the zip file and unzip the file.
+ Copy the contents of the folder named "bin" to a folder that is in your path 
    + (e.g. "$HOME/bin" is common).
+ Adjust your PATH if needed
    + (e.g. export PATH="$HOME/bin:$PATH")
+ Test


### Mac OS X

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Make sure the new location in in our path
5. Test

Here's an example of the commands run in the Terminal App after downloading the 
zip file.

```shell
    cd Downloads/
    unzip dataciteapi-*-macos-amd64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    dataciteapi -version
```

### Windows

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell on Windows 10 after
downloading the zip file.

```shell
    cd Downloads/
    unzip dataciteapi-*-windows-amd64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    dataciteapi -version
```


### Linux 

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell after
downloading the zip file.

```shell
    cd Downloads/
    unzip dataciteapi-*-linux-amd64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    dataciteapi -version
```


### Raspberry Pi

Released version is for a Raspberry Pi 2 or later use (i.e. requires ARM 7 support).

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell after
downloading the zip file.

```shell
    cd Downloads/
    unzip dataciteapi-*-raspbian-arm7.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    dataciteapi -version
```


Compiling from source
---------------------

_dataciteapi_ is "go gettable".  Use the "go get" command to download the dependant packages
as well as _dataciteapi_'s source code. 


```shell
    go get -u github.com/caltechlibrary/dataciteapi/...
```

Or clone the repstory and then compile

```shell
    cd
    git clone https://github.com/caltechlibrary/dataciteapi src/github.com/caltechlibrary/dataciteapi
    cd src/github.com/caltechlibrary/dataciteapi
    make
    make test
    make install
```

Compilation assumes [go](https://github.com/golang/go) v1.16

