package server

import "testing"

func TestRewritePath(t *testing.T) {
	got := rewritePath("/assets/path/to/file.txt", "/assets/", "./assets/public/")
	if got != "./assets/public/path/to/file.txt" {
		t.Errorf("rewritePath(\"/assets/path/to/file.txt\", \"/assets/\", \"./assets/public/\") = %s, want ./assets/public/path/to/file.txt", got)
	}
	got = rewritePath("/assets/assets/path/to/file.txt", "/assets/", "./assets/public/")
	if got != "./assets/public/assets/path/to/file.txt" {
		t.Errorf("rewritePath(\"/assets/assets/path/to/file.txt\", \"/assets/\", \"./assets/public/\") = %s, want ./assets/public/assets/path/to/file.txt", got)
	}
}
