# Spliting a file path
## Overview
This branch is to demonstrate how to split a file path to get the filename and copy a file into the temp directory
with the filename.

Checkout the `main.go` file to see the code.

## The `filepath` package provides the following
* `Base()`: Returns the last element of the path
* `Ext()`: Returns the file extension, if it has one
* `Split()`: Returns the split directory and file

## Concept
A file's final path directory or file is called the **base**. The path your binary is running in  is called the
**working directory**.
