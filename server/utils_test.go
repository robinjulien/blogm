package server

import "testing"

func TestRewritePath(t *testing.T) {
	got := RewritePath("/assets/path/to/file.txt", "/assets/", "./assets/public/")
	if got != "./assets/public/path/to/file.txt" {
		t.Errorf("RewritePath(\"/assets/path/to/file.txt\", \"/assets/\", \"./assets/public/\") = %s, want ./assets/public/path/to/file.txt", got)
	}
	got = RewritePath("/assets/assets/path/to/file.txt", "/assets/", "./assets/public/")
	if got != "./assets/public/assets/path/to/file.txt" {
		t.Errorf("RewritePath(\"/assets/assets/path/to/file.txt\", \"/assets/\", \"./assets/public/\") = %s, want ./assets/public/assets/path/to/file.txt", got)
	}
}

func TestFormatFileNameToTitle(t *testing.T) {
	got := FormatFileNameToTitle("my_page.md")
	if got != "My Page" {
		t.Errorf("FormatFileNameToTitle(\"my_page.md\") = %s, want My Page", got)
	}
	got = FormatFileNameToTitle("my_page._here_is_it")
	if got != "My Page. Here Is It" {
		t.Errorf("FormatFileNameToTitle(\"my_page._here_is_it\") = %s, want My Page. Here Is It", got)
	}
}
