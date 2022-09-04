# Reading remote files and write to local files
This branch is to demonstrate how to read a remote file and stream its content directly to a local files
instead of copying the entire file into memory and writing it to disk. This is faster and more memory effienct as
each chunk that is read is then immediately written to disk.

Checkout the `main.go` file.
