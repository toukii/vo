package vo

import (
	"fmt"
	"regexp"
	"testing"
)

func TestLog(t *testing.T) {
	log1 := `commit 86cba64d65c84b4e0b2329351d0cf9fd9d0b968f
Author: shiyongbin <shiyongbin@ezbuy.com>
Date:   Fri Feb 23 15:02:38 2018 +0800

    vgo`

	commit := regexp.MustCompile("commit\\ (\\S{11})")
	ma := commit.FindStringSubmatch(log1)
	fmt.Println(ma)

	date := regexp.MustCompile(`Date\:\ \ \ ([\S\ ]+)`)
	dt := date.FindStringSubmatch(log1)
	fmt.Println(dt)
}
