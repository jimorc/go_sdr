# gosdr
An SDR Receiver written in Go

## Status

__Be Forwarned:__

While I have good intentions here, I am using it to learn Go. As such, and given my track record with large projects, it is
probable that this project may never reach a usable state.

This project is in its very intial stage and there is little that is usable at this point.

## Building gosdr

All development at the moment is being done on an M1 Macbook Pro laptop, so the only instructions provided below are for
macOS. If a workable SDR receiver is developed on macOS, I intend to port to both Linux and Windows, but that is far in the
future. You could help by doing this porting work, but please first see the [Contributing](CONTRIBUTING.md) document.

This project is built mainly in Go using Visual Studio Code, so the instructions are provided for that combination.

The only dongles that I have that I can use in developing this project are RTL-SDR v3 and v4. Therefore, I have only listed
libraries required to access those dongles.

### Building on macOS

Prior to downloading this project from GitHub, you will need the following software installed:

- The latest version of [Go](https://go.dev/doc/install)
- The latest version of [VSCode](https://code.visualstudio.com/Download)
- The VSCode Go extension
- SoapySDR and related libraries. Install the following libraries using Homebrew:
    - soapysdr
    - soapyrtlsdr
    - librtlsdr
    - libusb