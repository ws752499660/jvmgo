package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEnrty struct {
	absDir string
}

func newDirEnrty(path string) *DirEnrty {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEnrty{absDir}
}

func (self *DirEnrty) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEnrty) String() string {
	return self.absDir
}
