# github-release-tool
[![Build Status](https://api.travis-ci.org/eloo/github-release-tool.svg?branch=master)](https://travis-ci.org/eloo/github-release-tool)
[![GoReport](https://goreportcard.com/badge/eloo/github-release-tool)](https://goreportcard.com/report/eloo/github-release-tool)
[![GoDoc](https://godoc.org/github.com/eloo/github-release-tool?status.svg)](https://godoc.org/github.com/eloo/github-release-tool)

Simple cli tool for working with Github releases.

## Current features:
* Download of a release (in Work)

## Usage
Download the latest release of Github repository
```
github-release-tool download <:owner/:repo> [flags]
```
e.g. this would download the latest release containing the string arm-7 to a relative download directory
```
github-release-tool download -s "arm-7" eloo/github-release-tool -o download
```
