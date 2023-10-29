package main

import "fmt"

type File struct {
	name string
}

func NewFile(name string) IFileSystem {
	return &File{
		name: name,
	}
}

func (f *File) ls() {
	fmt.Println("File: " + f.name)
}
