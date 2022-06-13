package fs

import (
	"sort"
	"strings"
)

// in-memory fs

type File struct {
	isFile  bool
	child   map[string]*File
	content string
	name    string
}

type FileSystem struct {
	root *File
}

func (f *File) isDirectory() bool {
	return f.isFile == false
}

func (f *File) getChild(name string) *File {
	if _, ok := f.child[name]; ok {
		return f.child[name]
	}
	return nil
}

func (f *File) createNewChildFile(name string, isFile bool) *File {
	f.child[name] = newFileStruct(name, isFile)
	return f.getChild(name)
}

func newFileStruct(name string, isFile bool) *File {
	return &File{
		isFile: isFile,
		name:   name,
		child:  make(map[string]*File),
	}
}

func (fs *FileSystem) mkdir(filepath string) { // []/ [a]/[b/c], /a
	dirs := strings.Split(filepath, "/")
	curr := fs.root
	for i := 1; i < len(dirs); i++ {
		next := curr.getChild(dirs[i])
		if next == nil {
			next = curr.createNewChildFile(dirs[i], false)
		}
		curr = next
	}
}

func (fs *FileSystem) ls(path string) []string { // /a /b/c
	dirs := strings.Split(path, "/")
	curr := fs.root
	if path != "/" {
		for i := 1; i < len(dirs); i++ {
			curr = curr.getChild(dirs[i])
		}
	}

	res := []string{}
	if curr.isDirectory() {
		for filename := range curr.child {
			res = append(res, filename)
		}
		sort.Strings(res)
	} else {
		res = append(res, curr.name)
	}
	return res
}

func (fs *FileSystem) addContentToFile(filepath string, content string) {
	dirs := strings.Split(filepath, "/")
	curr := fs.root
	for i := 1; i < len(dirs); i++ {
		next := curr.getChild(dirs[i])
		if next == nil {
			if i == len(dirs)-1 {
				next = curr.createNewChildFile(dirs[i], true)
			} else {
				next = curr.createNewChildFile(dirs[i], false)
			}
		}
		curr = next
	}
	curr.content += content
}

func (fs *FileSystem) readContentFromFile(filepath string) string {
	dirs := strings.Split(filepath, "/")
	curr := fs.root
	for i, dir := range dirs {
		if i == 0 {
			continue
		}
		curr = curr.getChild(dir)
	}
	return curr.content
}
