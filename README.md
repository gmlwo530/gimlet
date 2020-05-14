# Gimlet

Gimlet is CLI Tool for [gin](https://github.com/gin-gonic/gin).

Generate boilerplate code for easy gin development.

## Install

```sh
$ go get -u github.com/gmlwo530/gimlet
```

## Usage

### Generate Controller

Generate controller file. And file content is consisted of CRUD methods.
_Note: If `controllers` folder is not exist, gimlet generate folder too._

```sh
$ gimlet generate controller [controller_name]
```
