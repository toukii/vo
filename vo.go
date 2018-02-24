package vo

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Command = &cobra.Command{
		Use:   "vo",
		Short: "get version of go file from github",
		Long:  `vo toukii/goutils:dev -e v0.1.1v0.1.0v0.0.1v0.1`,
		Run: func(cmd *cobra.Command, args []string) {
			Excute(args)
		},
	}
)

func init() {
	Command.PersistentFlags().StringP("exclude", "e", "", "exclude")
	viper.BindPFlag("exclude", Command.PersistentFlags().Lookup("exclude"))
}

func ParseRepo(repoStr string) *Repo {
	repo := &Repo{Exclude: make(map[string]bool)}
	inputs := strings.Split(repoStr, "/")
	repo.User = inputs[0]
	input_1 := inputs[1]

	if strings.Contains(input_1, ":") {
		input_1s := strings.Split(input_1, ":")
		repo.Name = input_1s[0]
		repo.Branch = input_1s[1]
	} else {
		if strings.Contains(input_1, "@") {
			input_1s := strings.Split(input_1, "@")
			repo.Name = input_1s[0]
			repo.Commit = input_1s[1]
		} else {
			repo.Name = input_1
			repo.Branch = "master"
		}
	}

	exclude := viper.GetString("exclude")
	excs := strings.Split(exclude, "v")
	for _, exc := range excs {
		if exc == "" {
			continue
		}
		repo.Exclude["v"+exc] = true
	}

	fmt.Printf("%+v\n", repo)
	return repo
}

func Excute(args []string) error {
	repoStr := args[0]
	if !strings.Contains(repoStr, "/") {
		tips := "[user/]repo[:branch]  > $"
		fmt.Print(tips)
		fmt.Scanf("%s", &repoStr)
	}
	repo := ParseRepo(repoStr)

	fmt.Printf(repo.Require())
	return nil
}
