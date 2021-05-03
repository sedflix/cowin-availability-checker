# cowin-availability-checker

List and notify about all any vaccination appointment is present in your District with 7 days of today.

## Setup

```
go get github.com/sedflix/cowin-availability-checker
cd $GOPATH/src/github.com/sedflix/cowin-availability-checker
go build .
```

## Usage

This script requires the name of the State and District to find available appointment

```
$ ./cowin-availability-checker --help

Usage of ./cowin-availability-checker:
  -d string
        Name of the District (default "Lucknow")
  -s string
        Name of the Sate (default "Uttar Pradesh")

```

## Example

```
$ ./cowin-availability-checker -s "Uttar Pradesh" -d "Prayagraj"
Listing availble centres in Prayagraj, Uttar Pradesh
CHC Jasra 18-44 available 
End
```