package vo

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Command = &cobra.Command{
		Use:   "vo",
		Short: "get version of go file from github",
		Long: `vo user/repo[:branch][@commit]
ex: 
vo toukii/goutils:dev -e v0.1.1v0.1.0v0.0.1v0.1`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := Excute(args); err != nil {
				fmt.Println(err)
			}
		},
	}

	userRegx      = regexp.MustCompile("(\\S+)/")
	repoRegx      = regexp.MustCompile("/(\\S+?)(:|@|$)") // ? 非贪婪
	branchRegx    = regexp.MustCompile(":(\\S+?)(@|$)")
	commitShaRegx = regexp.MustCompile("@(\\S+)")
)

func init() {
	Command.PersistentFlags().StringP("exclude", "e", "", "exclude")
	viper.BindPFlag("exclude", Command.PersistentFlags().Lookup("exclude"))
}

func ParseRepo(repoStr string) *Repo {

	user := userRegx.FindStringSubmatch(repoStr)
	repo := repoRegx.FindStringSubmatch(repoStr)
	branch := branchRegx.FindStringSubmatch(repoStr)
	commitSha := commitShaRegx.FindStringSubmatch(repoStr)

	if len(branch) < 2 {
		branch = []string{"", ""}
	}
	if len(commitSha) < 2 {
		commitSha = []string{"", ""}
	}
	r := &Repo{
		User:    user[1],
		Name:    repo[1],
		Branch:  branch[1],
		Commit:  commitSha[1],
		Exclude: make(map[string]bool),
	}

	exclude := viper.GetString("exclude")
	excs := strings.Split(exclude, "v")
	for _, exc := range excs {
		if exc == "" {
			continue
		}
		r.Exclude["v"+exc] = true
	}
	return r
}

func Excute(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("repo is required.")
	}
	repoStr := args[0]
	if !strings.Contains(repoStr, "/") {
		tips := "user/repo[:branch][@commit]  > $"
		fmt.Print(tips)
		fmt.Scanf("%s", &repoStr)
	}
	repo := ParseRepo(repoStr)

	fmt.Printf(repo.Require())
	return nil
}
