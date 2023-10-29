package main

func main() {

	var rootDirectory = &Directory{
		name:          "root",
		directoryList: make([]IFileSystem, 0),
	}
	rootDirectory.AddDirectory(NewFile("file1"))
	rootDirectory.AddDirectory(NewFile("file2"))
	rootDirectory.AddDirectory(NewDirectory("dir1"))
	rootDirectory.AddDirectory(NewDirectory("dir2"))

	rootDirectory.ls()
}
