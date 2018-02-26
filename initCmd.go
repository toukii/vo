package vo

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/everfore/exc"
	"github.com/everfore/exc/walkexc/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/toukii/goutils"
)

var (
	InitCommand = &cobra.Command{
		Use:   "init",
		Short: "vgo go.mod init",
		Long:  `init`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := InitExcute(args); err != nil {
				fmt.Println(err)
			}
		},
	}
)

func init() {
	Command.PersistentFlags().StringP("exc", "E", "", "exc")
	viper.BindPFlag("exc", Command.PersistentFlags().Lookup("exc"))
}

func InitExcute(args []string) error {
	goutils.ReWriteFile("go.mod", nil)
	fd, err := os.OpenFile("go.mod", os.O_CREATE|os.O_RDWR, 0644)
	defer fd.Close()

	bs, err := exc.NewCMD("go list -json").DoNoTime()
	gopkg, err := pkg.NewPkg(bs)
	if goutils.CheckErr(err) {
		return err
	}
	mwr := io.MultiWriter(fd, os.Stdout)
	fmt.Fprintln(mwr, fmt.Sprintf("module \"%s\"", gopkg.ImportPath))
	fmt.Fprintln(fd)
	fmt.Fprintln(mwr, fmt.Sprint("require ("))

	dep := make(map[string]bool)
	for _, it := range gopkg.NoStdDepOKPkgs {
		dep[it] = true
	}
	for _, it := range gopkg.NoStdDepErrPkgs {
		dep[it.PkgName] = false
	}
	filterRepo(dep)

	for _, it := range gopkg.NoStdDepOKPkgs {
		if v, ex := dep[it]; v && ex {
			fmt.Fprintln(mwr, "\""+it+"\" ", ParseRepo(it).Tag())
		}
	}
	for _, it := range gopkg.NoStdDepErrPkgs {
		if v, ex := dep[it.PkgName]; v && ex {
			fmt.Fprintln(mwr, "\""+it.PkgName+"\" ", ParseRepo(it.PkgName).Tag())
		} else {
			fmt.Fprintln(mwr, "\""+it.PkgName+"\" ", "missing")
		}
	}

	fmt.Fprintln(mwr, ")")

	return nil
}

func filterRepo(dep map[string]bool) {
	for k, v := range dep {
		if !v {
			fmt.Printf("go get %s\n", k)
			continue
		}
		if strings.Contains(k, "/vendor/") {
			delete(dep, k)
			// fmt.Println("ignore:", k)
			continue
		}
		// base := baseGithubRepo(k)
		// if _, ex := dep[base]; ex && base != k {
		// 	delete(dep, k)
		// 	fmt.Println("ignore:", k)
		// }
	}
}

func baseGithubRepo(repo string) string {
	if strings.Count(repo, "/") <= 2 {
		return repo
	}
	i := 0
	idx := strings.IndexFunc(repo, func(r rune) bool {
		if r == rune("/"[0]) {
			i++
		}
		if i >= 3 {
			return true
		}
		return false
	})
	return repo[:idx]
}
