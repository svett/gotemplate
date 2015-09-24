# Go templates

This repository contains a set of templates that can be substitute by [gotemplate](https://github.com/ncw/gotemplate).
command line tool.

## Requirments

- Go (at least version 1.4)
- [Go template](https://github.com/ncw/gotemplate)

## Installation

```
$ go get github.com/ncw/gotemplate/...
$ go get github.com/svett/gotemplate/...
```

## Usage

```
//go:generate gotemplate "github.com/svett/gotemplate/stack" StudentStack(*Student)
type Student struct {
	FirstName string
	LastName  string
	BirthDate time.Time
}
```

```
$ go generate
```

## Author

[Svett Ralchev](http://blog.ralch.com)

