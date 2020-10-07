# Gehirn DNS CLI Tool

![Go](https://github.com/long-in/gdcli/workflows/Go/badge.svg)

## Description

`gdcli` is a cli application for Gehirn DNS ([Gehirn](https://www.gehirn.jp/))

## Installation

    $ go get github.com/long-in/gdcli

## Build the command

    $ go build -o gdcli cmd/gdcli/main.go

## Synopsis

    $ ./gdcli config init

## Available commands

```
  config   Manage config
  record   Manage record
  zone     Manage Zone
```

### Subcommands

#### `zone`

```
     ls  List zones
```


#### `record`

```
     ls      List records
     add     Add a new record
     rm      Remove one record
     update  Update a record
```

## Options

Please see the running result each subcommands with `-h`.

## Usage

Please read `EXAMPLE.md`.
