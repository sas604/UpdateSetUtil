package main

import (
	"regexp"
	"testing"

	"github.com/spf13/afero"
)

//exits if invalid path
func TestInvalidPath(t *testing.T) {
	_, err := RenameFiles("abcd")
	if err == nil {
		t.Fatalf("The error wasn't trown")
	}
}

// passing valid pass returns 0
func TestValidPath(t *testing.T) {
	count, err := RenameFiles(".")
	if err != nil && count != 0 {
		t.Fatalf("%s", err)
	}
}

// test if handles duplicates
func TestGetNewNameWithDuplicates(t *testing.T) {
	AppFs.MkdirAll("testDir", 0755)
	afero.WriteFile(AppFs, "testDir/test.xml", []byte("file b"), 0644)
	afero.WriteFile(AppFs, "testDir/test-1.xml", []byte("file c"), 0644)
	afero.WriteFile(AppFs, "testDir/test-2.xml", []byte("file c"), 0644)
	afero.WriteFile(AppFs, "testDir/test-3.xml", []byte("file c"), 0644)
	afero.WriteFile(AppFs, "testDir/test-4.xml", []byte("file c"), 0644)
	afero.WriteFile(AppFs, "testDir/test-5.xml", []byte("file c"), 0644)
	want := regexp.MustCompile(`test-6`)
	name := GetNewName("testDir", "test", 0)
	t.Log(name)
	if !want.MatchString(name) {
		t.Fatalf(`GetNewName("test", "test", 0) = %q, want match for %#q`, name, want)
	}

}
