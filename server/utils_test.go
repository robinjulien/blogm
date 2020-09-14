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

func TestCalculateMaxPage(t *testing.T) {
	got := CalculateMaxPage(50, 10)
	if got != 5 {
		t.Errorf("CalculateMaxPage(50, 10) = %d, want 5", got)
	}
	got = CalculateMaxPage(51, 10)
	if got != 6 {
		t.Errorf("CalculateMaxPage(51, 10) = %d, want 6", got)
	}
	got = CalculateMaxPage(0, 10)
	if got != 6 {
		t.Errorf("CalculateMaxPage(0, 10) = %d, want 0", got)
	}
}
