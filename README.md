# GPM (Go Package Manager)
GPM tries to be like npm, but for Go, and far more precarious.

## Installation
Having **Go** already installed, of course, run:

~~~~
$ go get github.com/KevinKoshka/gpm
~~~~

## Usage

* To add packages to the config file run:
~~~~
$ gpm add userName/packageName
~~~~
(gpm only supports github hosted packages so far)

* To remove packages from the config file, run:
~~~~
$ gpm remove userName/packageName
~~~~

* To install the packages, run:
~~~~
$ gpm install
~~~~