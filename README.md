[![Build Status](https://travis-ci.org/dsoprea/go-gpsbabel.svg?branch=master)](https://travis-ci.org/dsoprea/go-gpsbabel)
[![Coverage Status](https://coveralls.io/repos/github/dsoprea/go-gpsbabel/badge.svg?branch=master)](https://coveralls.io/github/dsoprea/go-gpsbabel?branch=master)
[![GoDoc](https://godoc.org/github.com/dsoprea/go-gpsbabel?status.svg)](https://godoc.org/github.com/dsoprea/go-gpsbabel)

# Overview

This is a simple wrapper project for [gpsbabel](https://www.gpsbabel.org). It provides a Go API for the system command. We'd rather use idiomatic Go than system calls, so this library allows us that abstraction.


# Implementation

We wrap a system call rather than linking to the library because GPSBabel explicitly states that they actively maintain command-line compatibility but not API compatibility.

See [I want GPSBabel to be a library, so I can use it my the program I'm developing. Why can't I do that?](https://www.gpsbabel.org/FAQ.html#library):

    The project itself is also focused on delivering a program and not an API. We change internals without much regard for preserving compatibility, deprecation schedules, API shims, DLL/dylib/shared object concerns or other things reasonably expected of providing a library instead of a finished program

We do not provide constants/enums for formats because we can not be sure which GPSBabel we are using nor do we want to require people to introduce contributions just to use the one new, obscure format they happen to need when a constant provides little value over a string.

We do not provide auto-detection because GPS formats and extensions are a mad science and:

1) we can not guarantee the combinations will lose exclusivity over time (more formats start using extensions that might currently only be used by one format)
2) most formats are obscure and we are not in a position to make any guesses
3) if GPSBabel does not provide this, we are in even less of a position to do so


# Requirements

Install GPSBabel. By default we use whichever "gpsbabel" will be found in the system search-path, but you can also set the GPSBABEL_FILEPATH environment variable or call SetBabelFilepath() with a file-path.


# Getting

```
$ go get github.com/dsoprea/go-gpsbabel
```


# Testable Examples

## [Babel_Convert](https://godoc.org/github.com/dsoprea/go-gpsbabel#example-Babel-Convert)

```go
b := NewBabel("v900", "gpx")

filepath := path.Join(TestAssetPath, "19020501_Portugal2.CSV.head")

f, err := os.Open(filepath)
log.PanicIf(err)

defer f.Close()

buffer := new(bytes.Buffer)

err = b.Convert(f, buffer)
if err != nil {
    if log.Is(err, ErrConversionFailed) == true {
        fmt.Printf("STDOUT:\n\n%s\n", buffer.String())
    }

    log.Panic(err)
}

// We need to parse the GPX so that the current timestamp that's placed into
// the GPX data doesn't destabilize the example.
testPrintGpxData(buffer)

// Output:
// 2019-02-05 08:07:05 +0000 UTC (38.760708,-9.112968)
// 2019-02-05 08:07:08 +0000 UTC (38.760738,-9.112924)
// 2019-02-05 08:07:09 +0000 UTC (38.760749,-9.112930)
// 2019-02-05 08:07:10 +0000 UTC (38.760753,-9.112923)
// 2019-02-05 08:07:14 +0000 UTC (38.760775,-9.112871)
// 2019-02-05 08:07:15 +0000 UTC (38.760775,-9.112861)
// 2019-02-05 08:07:16 +0000 UTC (38.760764,-9.112851)
// 2019-02-05 08:07:17 +0000 UTC (38.760754,-9.112848)
// 2019-02-05 08:07:18 +0000 UTC (38.760751,-9.112848)
```


## [Convert](https://godoc.org/github.com/dsoprea/go-gpsbabel#example-Convert)

```go
filepath := path.Join(TestAssetPath, "19020501_Portugal2.CSV.head")

f, err := os.Open(filepath)
log.PanicIf(err)

defer f.Close()

buffer := new(bytes.Buffer)

err = Convert("v900", "gpx", f, buffer)
if err != nil {
    if log.Is(err, ErrConversionFailed) == true {
        fmt.Printf("STDOUT:\n\n%s\n", buffer.String())
    }

    log.Panic(err)
}

testPrintGpxData(buffer)

// Output:
// 2019-02-05 08:07:05 +0000 UTC (38.760708,-9.112968)
// 2019-02-05 08:07:08 +0000 UTC (38.760738,-9.112924)
// 2019-02-05 08:07:09 +0000 UTC (38.760749,-9.112930)
// 2019-02-05 08:07:10 +0000 UTC (38.760753,-9.112923)
// 2019-02-05 08:07:14 +0000 UTC (38.760775,-9.112871)
// 2019-02-05 08:07:15 +0000 UTC (38.760775,-9.112861)
// 2019-02-05 08:07:16 +0000 UTC (38.760764,-9.112851)
// 2019-02-05 08:07:17 +0000 UTC (38.760754,-9.112848)
// 2019-02-05 08:07:18 +0000 UTC (38.760751,-9.112848)
```


## [ConvertToGpx](https://godoc.org/github.com/dsoprea/go-gpsbabel#example-ConvertToGpx)

```go
filepath := path.Join(TestAssetPath, "19020501_Portugal2.CSV.head")

f, err := os.Open(filepath)
log.PanicIf(err)

defer f.Close()

buffer := new(bytes.Buffer)

err = ConvertToGpx("v900", f, buffer)
if err != nil {
    if log.Is(err, ErrConversionFailed) == true {
        fmt.Printf("STDOUT:\n\n%s\n", buffer.String())
    }

    log.Panic(err)
}

testPrintGpxData(buffer)

// Output:
// 2019-02-05 08:07:05 +0000 UTC (38.760708,-9.112968)
// 2019-02-05 08:07:08 +0000 UTC (38.760738,-9.112924)
// 2019-02-05 08:07:09 +0000 UTC (38.760749,-9.112930)
// 2019-02-05 08:07:10 +0000 UTC (38.760753,-9.112923)
// 2019-02-05 08:07:14 +0000 UTC (38.760775,-9.112871)
// 2019-02-05 08:07:15 +0000 UTC (38.760775,-9.112861)
// 2019-02-05 08:07:16 +0000 UTC (38.760764,-9.112851)
// 2019-02-05 08:07:17 +0000 UTC (38.760754,-9.112848)
// 2019-02-05 08:07:18 +0000 UTC (38.760751,-9.112848)
```
