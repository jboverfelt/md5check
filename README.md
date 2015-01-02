md5check
========

A tool to check md5 sums given an untrusted file and a sum that has been received out of band.
I'm hoping this tool will be useful to others, but originally it was intended as a project to familiarize myself with creating basic command line applications in Go.

Usage
-----

```
go get github.com/jboverfelt/md5check
md5check [file] [known md5 sum]
```

The program will print ```OK``` and exit with a 0 code if the check was successful.
Otherwise, the program will print an error and exit with a non-zero code.

Contributing
-----

Assumptions:
* You're using a *nix system for development

Steps:

1. [Download Go](http://golang.org/doc/install)
2. Fork this repo
3. Run go get as described in the Usage section
4. Follow [these](https://blog.splice.com/contributing-open-source-git-repositories-go/) instructions for setting up your fork as a remote
5. Run ```make```
6. If the output looks clean, go ahead and submit a pull request
