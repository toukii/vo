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

func TestRepo(t *testing.T) {
	args := []string{
		"toukii/goutils:dev@1eb9",
		"toukii/goutils",
		"toukii/goutils:dev",
		"toukii/goutils@1eb9",
		"toukii/goutils:",
		"toukii/goutils@",
		"toukii/goutils:@",
		"everfore/exc/walkexc/pkg",
	}

	for _, it := range args {
		user := userRegx.FindStringSubmatch(it)
		repo := repoRegx.FindStringSubmatch(it)
		branch := branchRegx.FindStringSubmatch(it)
		commitSha := commitShaRegx.FindStringSubmatch(it)
		fmt.Printf("%s \n  user:%+v\n  repo:%+v\n  branch:%+v\n  commit:%+v\n", it, user, repo, branch, commitSha)
	}
}

func TestBaseGithub(t *testing.T) {
	args := []string{
		"github.com/toukii/vo",
		"github.com/toukii/vo/vo",
		"github.com/toukii",
	}

	for _, it := range args {
		fmt.Printf("%s, %s\n", it, baseGithubRepo(it))
	}
}
