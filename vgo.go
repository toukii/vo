package vo

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/astaxie/beego/httplib"
	"github.com/everfore/exc"
	"github.com/toukii/goutils"
	"github.com/toukii/jsnm"
)

type Repo struct {
	Raw     string
	rawUser string
	rawRepo string

	User    string
	Repo    string
	Name    string // repo name
	Branch  string
	Commit  string
	Version string

	Exclude map[string]bool
}

var (
	commitRegx = regexp.MustCompile("commit\\ (\\S{12})")
	dateRegx   = regexp.MustCompile(`Date\:\ \ \ ([\S\ ]+)`)

	tagFmtUrl        = "https://api.github.com/repos/%s/%s/tags?access_token=%s"
	commitFmtUrl     = "https://api.github.com/repos/%s/%s/commits/%s?access_token=%s"
	branchLogdateFmt = "Mon Jan 2 15:04:05 2006 -0700"
	branchApidateFmt = "2006-01-02T15:04:05Z"
	tagFmt           = "\"%s\" %s"

	TOKEN = ""
)

func (r *Repo) Require() string {
	return fmt.Sprintf(tagFmt, r.Name, r.Tag())
}

/*
commit非空，获取commit的tag；
branch非空，获取branch的最新提交；
branch为空，先获取tag；没有tag，获取master最新提交；没有master；获取最近提交；
*/
func (r *Repo) Tag() string {
	var tag string
	var err error
	if r.Commit != "" {
		tag, err = r.CommitTag()
		if err == nil {
			return tag
		} else {
			fmt.Println(r.Name, err)
			return ""
		}
	}
	if r.Branch == "" {
		tag, err = r.LatestTag()
		if err != nil {
			fmt.Println(r.Name, err)
		}
		if tag == "" {
			r.Branch = "master"
			tag, err = r.LocalLatestBranchCommit()
			if err != nil {
				fmt.Println(r.Name, err)
			}
		}
	} else {
		tag, err = r.LocalLatestBranchCommit()
		if err != nil {
			fmt.Println(r.Name, err)
		}
	}
	if tag == "" {
		return "latest"
	}
	return tag
}

func (r *Repo) LatestTag() (string, error) {
	if strings.Contains(r.User, "golang.org") {
		return "", fmt.Errorf("Tag api is not supported for golang.org")
	}
	req := httplib.NewBeegoRequest(fmt.Sprintf(tagFmtUrl, r.User, r.Repo, TOKEN), "GET")
	bs, err := req.Bytes()
	if err != nil {
		return "", err
	}
	if len(bs) <= 0 {
		return "", fmt.Errorf("response is nil.")
	}
	if strings.Contains(goutils.ToString(bs), `"message":`) {
		return "", fmt.Errorf("%s", bs)
	}
	tag := ""
	jsnm.BytesFmt(bs).Range(func(i int, ji *jsnm.Jsnm) bool {
		tagTmp := ji.Get("name").RawData().String()
		if len(r.Exclude) > 0 {
			if _, ex := r.Exclude[tagTmp]; !ex && strings.HasPrefix(tagTmp, r.Version) {
				tag = tagTmp
				return true
			}
		} else if strings.HasPrefix(tagTmp, r.Version) {
			tag = tagTmp
			return true
		}
		return false
	})
	if tag == "" {
		return "", fmt.Errorf("No wanted tag.")
	}
	return tag, nil
}

func (r *Repo) LocalLatestBranchCommit() (string, error) {
	glog := exc.NewCMD(fmt.Sprintf("git log %s -1", r.Branch)).Env("GOPATH")
	if strings.Contains(r.User, "golang.org") {
		glog.Env("GOPATH").Cd("src").Cd(r.Name)
	} else {
		glog = glog.Cd("src/github.com/").Cd(r.User).Cd(r.Repo)
	}
	bs, err := glog.Do()
	if err != nil {
		fmt.Printf("%s,%+v\n", bs, err)
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
	if strings.Contains(r.User, "golang.org") {
		return "", fmt.Errorf("Api commits not supported for golang.org")
	}
	req := httplib.NewBeegoRequest(fmt.Sprintf(commitFmtUrl, r.User, r.Repo, r.Commit, TOKEN), "GET")
	bs, err := req.Bytes()
	if err != nil {
		return "", err
	}
	if len(bs) <= 0 {
		return "", fmt.Errorf("response is nil.")
	}
	if strings.Contains(goutils.ToString(bs), `"message":`) {
		return "", fmt.Errorf("%s", bs)
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
