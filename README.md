# sn

A [Notion](https://notion.so) simple notetaking CLI client written in Go

## Table of Contents

1. [Installation](#installation)

## Installation

1. Create a new integration token (used to access your account workspace) via [the official Notion getting started guide](https://developers.notion.com/docs/getting-started). Make sure to share any pages you wish to be accessed with this integration (via sharing on each page). Then copy the secret integration key. This will be compiled into the binary.

2. Compile and install the program.

```sh
make # will prompt for integration token
$ make install
```
