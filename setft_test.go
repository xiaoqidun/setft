package setft

import (
	"testing"
	"time"
)

func TestSetFileTime(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	atime, _ := time.ParseInLocation(layout, "2021-01-01 00:00:00", time.Local)
	ctime, _ := time.ParseInLocation(layout, "2021-01-01 00:00:00", time.Local)
	mtime, _ := time.ParseInLocation(layout, "2021-01-01 00:00:00", time.Local)
	if err := SetFileTime("setft_test.go", atime, ctime, mtime); err != nil {
		t.Fatal(err)
	}
}
