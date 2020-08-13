package main

import "testing"

func TestFormatName(t *testing.T) {
	got := formatName("home-page.html")

	if got != "HOME_PAGE_HTML" {
		t.Errorf("formatName(\"home-page.html\") = %s, want HOME_PAGE_HTML", got)
	}
}
