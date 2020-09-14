package server

import (
	"os"
	"sort"
	"strings"
)

// ReadDirModTime reads the directory named by dirname and returns
// a list of directory entries sorted by ModTime.
func ReadDirModTime(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Slice(list, func(i, j int) bool { return list[i].ModTime().Unix() > list[j].ModTime().Unix() })
	return list, nil
}

// FormatFileNameToTitle format a file name to a displayable title name
func FormatFileNameToTitle(name string) string {
	return strings.Title(strings.ReplaceAll(strings.TrimSuffix(name, ".md"), "_", " "))
}

// RewritePath rewrites a URL by remove oldprefix prefix and replacing it with newprefix
func RewritePath(path, oldprefix, newprefix string) string {
	return strings.Replace(path, oldprefix, newprefix, 1)
}

// CalculateMaxPage calculate and return the max number of pages needed to displays all posts
func CalculateMaxPage(n, maxPerPage int) int {
	if n%maxPerPage == 0 {
		return n / maxPerPage
	} else {
		return n/maxPerPage + 1
	}
}
