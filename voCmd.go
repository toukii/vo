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

	userRegx      = regexp.MustCompile("(\\S+?)/")
	repoRegx      = regexp.MustCompile("/(\\S+?)(:|@|$)") // ? 非贪婪
	branchRegx    = regexp.MustCompile(":(\\S+?)(@|$)")
	commitShaRegx = regexp.MustCompile("@(\\S+)")

	gopkgUserRegx    = regexp.MustCompile("/(\\S+?)/")
	gopkgRepoRegx    = regexp.MustCompile("/(\\S+?)\\.v")
	gopkgVersionRegx = regexp.MustCompile("\\.v(\\S+?)(/|$)")
)

func init() {
	Command.PersistentFlags().StringP("exclude", "e", "", "exclude")
	Command.PersistentFlags().BoolP("init", "i", false, "init")
	viper.BindPFlag("exclude", Command.PersistentFlags().Lookup("exclude"))
	viper.BindPFlag("init", Command.PersistentFlags().Lookup("init"))

	if TOKEN == "" {
		panic("access_token is missing.")
	}
}

// gopkg.in/pkg.v3      → github.com/go-pkg/pkg (branch/tag v3, v3.N, or v3.N.M)
// gopkg.in/user/pkg.v3 → github.com/user/pkg   (branch/tag v3, v3.N, or v3.N.M)
func ParseRepo(rawRepoStr string) *Repo {
	repoStr := strings.TrimPrefix(rawRepoStr, "github.com/")
	var user, repo, version []string
	if strings.HasPrefix(rawRepoStr, "gopkg.in") {
		repoStr = strings.TrimPrefix(rawRepoStr, "gopkg.in/")
		version = gopkgVersionRegx.FindStringSubmatch(repoStr)
	}
	user = userRegx.FindStringSubmatch(repoStr)
	repo = repoRegx.FindStringSubmatch(repoStr)
	branch := branchRegx.FindStringSubmatch(repoStr)
	commitSha := commitShaRegx.FindStringSubmatch(repoStr)

	empty := []string{"", ""}
	if len(user) < 2 {
		user = empty
	}
	if len(repo) < 2 {
		repo = empty
	}
	if len(branch) < 2 {
		branch = empty
	}
	if len(commitSha) < 2 {
		commitSha = empty
	}
	if len(version) < 2 {
		version = empty
	}
	r := &Repo{
		Raw:     rawRepoStr,
		User:    user[1],
		Repo:    repo[1],
		Branch:  branch[1],
		Commit:  commitSha[1],
		Version: "v" + version[1],
		Exclude: make(map[string]bool),
	}
	if strings.HasPrefix(r.Raw, "gopkg.in") {
		vidx := strings.Index(repoStr, "."+r.Version)

		if vidx < 0 {
			i := 0
			vidx = strings.IndexFunc(repoStr, func(r rune) bool {
				if r == rune("/"[0]) || r == rune(":"[0]) || r == rune("@"[0]) {
					i++
				}
				if i > 1 {
					return true
				}
				return false
			})
			if vidx < 0 {
				vidx = len(repoStr)
			}
		}

		rs := strings.Split(repoStr[:vidx], "/")
		if len(rs) < 2 {
			r.User = "go-" + rs[0]
			r.Repo = rs[0]
		} else {
			r.User = rs[0]
			r.Repo = rs[1]
		}
	} else {
		i := 0
		start := 0
		idx := strings.IndexFunc(repoStr, func(r rune) bool {
			if i == 0 {
				start++
			}
			if r == rune("/"[0]) {
				i++
				if i > 1 {
					return true
				}
			}
			return false
		})
		if idx < 0 {
			idx = len(repoStr)
		}
		r.Repo = repoStr[start:idx]
	}
	if strings.Contains(rawRepoStr, "@") {
		r.Name = strings.Split(rawRepoStr, "@")[0]
	} else {
		r.Name = strings.Split(rawRepoStr, ":")[0]
	}

	exclude := viper.GetString("exclude")
	excs := strings.Split(exclude, "v")
	for _, exc := range excs {
		if exc == "" {
			continue
		}
		r.Exclude["v"+exc] = true
	}
	// fmt.Printf("user:%s, repo:%s, name:%s, version:%s, branch:%s, commit:%s\n", r.User, r.Repo, r.Name, r.Version, r.Branch, r.Commit)
	return r
}

func Excute(args []string) error {
	if viper.GetBool("init") {
		return InitExcute(args)
	}

	var repoStr string
	if len(args) > 0 {
		if args[0] == "init" {
			return InitExcute(args)
		}
		repoStr = args[0]
	}
	if !strings.Contains(repoStr, "/") {
		tips := "#input:# user/repo[:branch][@commit]  $"
		fmt.Print(tips)
		fmt.Scanf("%s", &repoStr)
	}
	repo := ParseRepo(repoStr)

	fmt.Printf(repo.Require())
	return nil
}
