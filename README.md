# Capture

A simple, zero dependency, web server to capture images.

## Use Case

There are times where I want to use my mobile phone to take a picture of a
document (or receipt) and get that picture to my computer without going through
the [public cloud](https://en.wikipedia.org/wiki/Cloud_computing).

## Installation

1. Download Capture for your computer or home server from
   [here](https://github.com/samherrmann/capture/releases/).
1. Move the executable to an installation directory of your choice.

## Usage

Start Capture with the following
[CLI](https://en.wikipedia.org/wiki/Command-line_interface) command:

```sh
cd /path/to/installation/directory
./capture
```

By default, the Capture server listens on port `8080` and saves the images in
the directory from which it was started. See below for alternative options.

### Serving on custom address

```sh
./capture -address example.com:80
```

### Save images in an alternate location

```sh
./capture -destination /path/to/my/images
```

Tip: Within a home network, you may want to save images to a network drive where
you can access the images from multiple devices.
