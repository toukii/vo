package vo

import (
	"fmt"
	"regexp"
	"time"

	"github.com/astaxie/beego/httplib"
	"github.com/everfore/exc"

	"github.com/toukii/goutils"
	"github.com/toukii/jsnm"
)

type Repo struct {
	User    string
	Name    string // repo name
	Branch  string
	Commit  string
	Exclude map[string]bool
}

var (
	commitRegx = regexp.MustCompile("commit\\ (\\S{12})")
	dateRegx   = regexp.MustCompile(`Date\:\ \ \ ([\S\ ]+)`)

	tagFmtUrl        = "https://api.github.com/repos/%s/%s/tags"
	commitFmtUrl     = "https://api.github.com/repos/%s/%s/commits/%s"
	branchLogdateFmt = "Mon Jan 2 15:04:05 2006 -0700"
	branchApidateFmt = "2006-01-02T15:04:05Z"
	tagFmt           = `"github.com/%s/%s" %s`
)

func (r *Repo) Require() string {
	return fmt.Sprintf(tagFmt, r.User, r.Name, r.Tag())
}

func (r *Repo) Tag() string {
	tag, err := r.LatestTag()
	if err != nil {
		if r.Commit != "" {
			tag, err = r.CommitTag()
		}
		if err == nil {
			return tag
		}
		fmt.Println(err)
		tag, err = r.LatestBranchCommit()
		if err != nil {
			return "latest"
		}
	}
	return tag
}

func (r *Repo) LatestTag() (string, error) {
	tags_url := fmt.Sprintf(tagFmtUrl, r.User, r.Name)
	bs, err := httplib.Get(tags_url).Bytes()
	if err != nil {
		return "", err
	}
	if len(bs) <= 0 {
		return "", fmt.Errorf("response is nil.")
	}
	tag := ""
	jsnm.BytesFmt(bs).Range(func(i int, ji *jsnm.Jsnm) bool {
		tagTmp := ji.Get("name").RawData().String()
		if len(r.Exclude) > 0 {
			if _, ex := r.Exclude[tagTmp]; !ex {
				tag = tagTmp
				return true
			}
		} else {
			tag = tagTmp
			return true
		}
		return false
	})
	if tag == "" {
		return "", fmt.Errorf("no wanted tag.")
	}
	return tag, nil
}

func (r *Repo) LatestBranchCommit() (string, error) {
	bs, err := exc.NewCMD(fmt.Sprintf("git log %s -1", r.Branch)).Env("GOPATH").Cd("src/github.com/").Cd(r.User).Cd(r.Name).Do()
	if goutils.CheckErr(err) {
		return "", err
	}
	log1 := goutils.ToString(bs)
	var commit, date string

	matchCommits := commitRegx.FindStringSubmatch(log1)
	if len(matchCommits) > 0 {
		commit = matchCommits[1]
	}

	matchDates := dateRegx.FindStringSubmatch(log1)
	if len(matchDates) > 0 {
		date = parseDate(branchLogdateFmt, matchDates[1])
	}

	ret := fmt.Sprintf("v0.0.0-%s-%s", date, commit)
	if len(ret) < 10 {
		return "", fmt.Errorf("version is missing.")
	}
	return ret, nil
}

func (r *Repo) CommitTag() (string, error) {
	bs, err := httplib.Get(fmt.Sprintf(commitFmtUrl, r.User, r.Name, r.Commit)).Bytes()
	if err != nil {
		return "", err
	}
	if len(bs) <= 0 {
		return "", fmt.Errorf("response is nil.")
	}
	js := jsnm.BytesFmt(bs)
	cur := js.Get("sha").RawData().String()
	date := parseDate(branchApidateFmt, js.PathGet("commit", "committer", "date").RawData().String())
	commit := cur[:12]
	return fmt.Sprintf("v0.0.0-%s-%s", date, commit), nil
}

func parseDate(dateFmt, date string) string {
	d, err := time.Parse(dateFmt, date)
	if err != nil {
		panic(err)
	}
	return d.Format("20060102150405")
}
