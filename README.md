# sn

A [Notion](https://notion.so)-based notetaking CLI syncing client written in Go

## Table of Contents

1. [Introduction](#introduction)
2. [Features](#features)
3. [What this is not](#what-this-is-not)
4. [Installation](#installation)

## Introduction

Syncing textual data across devices is always difficult. I've never found a good alternative for syncing notes across all the devices I use daily. I have tried many alternatives such as Evernote, Simplenote, iNotes, git, and others but none of them can do exactly what I'm looking for.

I chose Notion because of its new developer API which makes it easy to query its API. The goals of this project are:

* to readily access text notes across all platforms
* to convert and download text notes as markdown
* to upload and sync new notes to the cloud

The program works by using a Notion integration token to connect to specific pages you declare, convert them to markdown, then provide them as text files. These files then can be edited with any text editor (vim) and can upload changes by converting them back to Notion blocks.

## Features

This project supports syncing Notion pages containing the following block components:

* paragraphs
* heading ones
* heading twos
* heading threes
* quotes
* bulleted list items
* numbered list items
* todos
* code blocks

## What this is not

It is important to clarify that this syncing client is not:

* a fully-fledged Notion CLI
* a markdown converter

## Installation

1. Create a new integration token (used to access your account workspace) via [the official Notion getting started guide](https://developers.notion.com/docs/getting-started). Make sure to share any pages you wish to be accessed with this integration (via sharing on each page). Then copy the secret integration key. This will be compiled into the binary.

2. Compile and install the program.

```sh
make # will prompt for integration token
$ make install
```
