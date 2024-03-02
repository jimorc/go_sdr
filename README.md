# gosdr
An SDR Receiver written in Go

## Status

__Be Forwarned:__

While I have good intentions here, I am using it to learn Go. As such, and given my track record with large projects, it is
probable that this project may never reach a usable state.

This project is in its very intial stage and there is little that is usable at this point.

## Building gosdr

All development at the moment is being done on an M1 Pro Macbook Pro running MacOS 14, so the only instructions provided 
below are for that system. I spent a limited amount of time setting it up on Kubuntu 22.04 and incomplete instructions for that
system are also provided below. If a workable SDR receiver is developed on MacOS, I intend to port to both Linux and Windows, 
but that is far in the
future. You could help by doing this porting work, but please first see the [Contributing](CONTRIBUTING.md) document.

This project is built mainly in Go using Visual Studio Code, so the instructions are provided for that combination.

The only dongles that I have that I can use in developing this project are RTL-SDR v3 and v4. Therefore, I have only listed
libraries required to access those dongles.

### Building on MacOS

Prior to downloading this project from GitHub, you will need the following software installed:

- The latest version of [Go](https://go.dev/doc/install)
- The latest version of [VSCode](https://code.visualstudio.com/Download)
- The VSCode Go extension
- SoapySDR and related libraries. Install SoapySDR libraries using Homebrew:
    ```
    brew install soapysdr
    ```

#### Building for Apple Silicon

The following commands need to be executed in your terminal before `go` is called:
```
export CGO_CFLAGS="-I/opt/homebrew/opt/soapysdr/include"
export CGO_LDFLAGS="-L/opt/homebrew/opt/soapysdr/lib"
export GOARCH="arm64"
export CGO_ENABLED=1
```
You can execute these commands in your terminal before running:
```
go run .
```
or
```
go build .
```
Alternatively, you can add these to your `.zshrc` file and they will execute automatically every time you open a terminal,
including the terminal in VSCode.

#### Building for Intel Silicon

Note: I do not have access to an Intel silicon Mac, so I am relying on information that I found on the internet regarding the
location that Homebrew stores its packages and symlinks.

The following commands need to be executed in your terminal before `go` is called:
```
export CGO_CFLAGS="-I/usr/local/opt/soapysdr/include"
export CGO_LDFLAGS="-L/usr/local/opt/soapysdr/lib"
export GOARCH="arm64"
export CGO_ENABLED=1
```
You can execute these commands in your terminal before running:
```
go run .
```
or
```
go build .
```
Alternatively, you can add these to your `.zshrc` file and they will execute automatically every time you open a terminal,
including the terminal in VSCode.

### Building on Ubuntu 22.04

The following are incomplete instructions.

- SoapySDR and related libraries. Install the following:
    - apt install libsoapysdr0.8 libsoapysdr-doc libsoapysdr-dev
    - apt install librtlsdr0
 