package vo

import (
	// "bytes"
	"fmt"
	"regexp"
	"time"

	"github.com/astaxie/beego/httplib"
	"github.com/everfore/exc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/toukii/goutils"
	"github.com/toukii/jsnm"
)

type Repo struct {
	User string
	Name string // repo name
}

var (
	Command = &cobra.Command{
		Use:   "vo",
		Short: "get version of go file from github",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			Excute(args)
		},
	}

	commitRegx = regexp.MustCompile("commit\\ (\\S{12})")
	dateRegx   = regexp.MustCompile(`Date\:\ \ \ ([\S\ ]+)`)

	tagFmtUrl = "https://api.github.com/repos/%s/%s/tags"
	dateFmt   = "Mon Jan 2 15:04:05 2006 -0700"
	tagFmt    = `"github.com/%s/%s" %s`
)

func init() {
	Command.PersistentFlags().StringP("user", "u", "toukii", "github user")
	Command.PersistentFlags().StringP("repo", "o", "", "user's repo")
	viper.BindPFlag("user", Command.PersistentFlags().Lookup("user"))
	viper.BindPFlag("repo", Command.PersistentFlags().Lookup("repo"))
}

func (r *Repo) Tag() string {
	tag, err := r.LatestTag()
	if err != nil {
		tag, err = r.LatestMasterCommit()
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
	arr := jsnm.BytesFmt(bs).Arr()
	if len(arr) <= 0 {
		return "", fmt.Errorf("no tags")
	}
	return arr[0].Get("name").RawData().String(), nil
}

func (r *Repo) LatestMasterCommit() (string, error) {
	bs, err := exc.NewCMD("git log master -1").Env("GOPATH").Cd("src/github.com/").Cd(r.User).Cd(r.Name).Do()
	if err != nil {
		panic(err)
	}
	log1 := goutils.ToString(bs)
	var commit, date string

	matchCommits := commitRegx.FindStringSubmatch(log1)
	if len(matchCommits) > 0 {
		commit = matchCommits[1]
	}

	matchDates := dateRegx.FindStringSubmatch(log1)
	if len(matchDates) > 0 {
		date = parseDate(matchDates[1])
	}

	ret := fmt.Sprintf("v0.0.0-%s-%s", date, commit)
	if len(ret) < 10 {
		return "", fmt.Errorf("version is missing.")
	}
	return ret, nil
}

func parseDate(date string) string {
	d, err := time.Parse(dateFmt, date)
	if err != nil {
		panic(err)
	}
	return d.Format("20060102150405")
}

func Excute(args []string) error {
	repo := &Repo{
		User: viper.GetString("user"),
		Name: viper.GetString("repo"),
	}

	tag := repo.Tag()
	fmt.Printf(tagFmt, repo.User, repo.Name, tag)
	return nil
}
