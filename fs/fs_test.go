package fs

import (
	"reflect"
	"testing"
)

func TestFsMkdir(t *testing.T) {
	f := newFileStruct("root", false)
	fs := &FileSystem{
		root: f,
	}
	fs.mkdir("/a/b/c")
	fs.mkdir("/a/b/d")
}

func TestFsLs(t *testing.T) {
	fs := &FileSystem{
		root: newFileStruct("root", false),
	}
	paths := []string{"/a/b/c", "/a/b/d", "/a/e", "/tmp", "/usr"}
	fs.addContentToFile("/usr/a.txt", "usr_a___txt")

	queries := map[string][]string{
		"/a":         {"b", "e"},
		"/":          {"a", "tmp", "usr"},
		"/a/b":       {"c", "d"},
		"/usr/a.txt": {"a.txt"},
	}
	for _, path := range paths {
		fs.mkdir(path)
	}

	for query, expected := range queries {
		lsRes := fs.ls(query)
		if !reflect.DeepEqual(lsRes, expected) {
			t.Errorf("ls failed, expected %v, received %v", expected, lsRes)
		}
	}

}

func TestAddContentToFile(t *testing.T) {
	fs := &FileSystem{
		root: newFileStruct("root", false),
	}
	queries := map[string][]string{
		"/a/b/c":        {"test_1", " test_2", " test_3"},
		"/tmp/test.txt": {"tmp1", " tmp2", " tmp3"},
	}
	expected := map[string]string{
		"/a/b/c":        "test_1 test_2 test_3",
		"/tmp/test.txt": "tmp1 tmp2 tmp3",
	}
	for filepath, contents := range queries {
		for _, content := range contents {
			fs.addContentToFile(filepath, content)
		}
	}
	for filepath, expectedContent := range expected {
		readContentResult := fs.readContentFromFile(filepath)
		if !reflect.DeepEqual(readContentResult, expectedContent) {
			t.Errorf("failed to read content, expected %v, received %v", expectedContent, readContentResult)
		}
	}
}
