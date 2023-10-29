package main

import "fmt"

type IDirectory interface {
	AddDirectory(dir IFileSystem)
	RemoveDirectory(dir IFileSystem)
	GetDirectoryList() []IFileSystem
	GetName() string
	SetName(name string)
}

type Directory struct {
	IFileSystem
	name          string
	directoryList []IFileSystem
}

func NewDirectory(name string) IFileSystem {
	return &Directory{
		name:          name,
		directoryList: make([]IFileSystem, 0),
	}
}

func (d *Directory) ls() {
	fmt.Println("Directory: " + d.name)
	for _, dir := range d.directoryList {
		dir.ls()
	}
}

func (d *Directory) AddDirectory(dir IFileSystem) {
	d.directoryList = append(d.directoryList, dir)
}

func (d *Directory) RemoveDirectory(dir IFileSystem) {
	newList := []IFileSystem{}

	for _, directory := range d.directoryList {
		if directory != dir {
			newList = append(newList, directory)
		}
	}
	d.directoryList = newList
}

func (d *Directory) GetDirectoryList() []IFileSystem {
	return d.directoryList
}

func (d *Directory) GetName() string {
	return d.name
}
