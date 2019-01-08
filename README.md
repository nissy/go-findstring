# go-findstring
Output string used in package of Go

## Setup
```
$ go get -u github.com/nissy/go-findstring
```


## Usage
```bash
$ go-findstring /src/github.com/nissy/go-findstring/
/src/github.com/nissy/go-findstring/main.go:14:29:   "r"
/src/github.com/nissy/go-findstring/main.go:14:41:   "Search directory recursive."
/src/github.com/nissy/go-findstring/main.go:22:40:   "Error: %s\n"
/src/github.com/nissy/go-findstring/main.go:79:19:   "%s:\t%s\n"
```

## Help
```
Usage: go-findstring [options] directory
  -h    This help
  -r    Search directory recursive
```
