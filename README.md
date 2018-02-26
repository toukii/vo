# vo

vgo 产生的go.mod文件中，require的项目版本

## install

```
go get -u github.com/toukii/vo/vo
```

## config

[config github access_token](https://github.com/settings/tokens)


## usage

__vo user/repo[:branch][@commit]__


__`vo toukii/vo`__

__`vo toukii/vo:dev`__

__`vo toukii/vo:master`__

__`vo toukii/vo@d1a0e830b03ff75f9578e8ab88ee75958f88cde6`__



```
"github.com/toukii/vo" v0.0.1
"github.com/toukii/vo" v0.0.0-20180224151534-5f08cd478e56                                                                             
"github.com/toukii/vo" v0.0.0-20180224151534-5f08cd478e56                                                                             
"github.com/toukii/vo" v0.0.0-20180223133159-d1a0e830b03f
```

__`vo init`__


```
module "github.com/toukii/vo"

require (
"github.com/astaxie/beego/httplib"  v0.0.0-20160922231845-2d87d4feafee
"github.com/everfore/exc"  v0.0.0-20180201220233-45314dca7f0f
"github.com/everfore/exc/walkexc"  v0.0.0-20180201220233-45314dca7f0f
"github.com/everfore/exc/walkexc/pkg"  v0.0.0-20180201220233-45314dca7f0f
"github.com/fatih/color"  v0.0.0-20170926141411-5df930a27be2
"github.com/fsnotify/fsnotify"  v0.0.0-20170328212107-4da3e2cfbabc
"github.com/hashicorp/hcl"  v0.0.0-20170505105837-392dba7d905e
"github.com/hashicorp/hcl/hcl/ast"  v0.0.0-20170505105837-392dba7d905e
"github.com/hashicorp/hcl/hcl/parser"  v0.0.0-20170505105837-392dba7d905e
"github.com/hashicorp/hcl/hcl/scanner"  v0.0.0-20170505105837-392dba7d905e
"github.com/hashicorp/hcl/hcl/strconv"  v0.0.0-20170505105837-392dba7d905e
"github.com/hashicorp/hcl/hcl/token"  v0.0.0-20170505105837-392dba7d905e
"github.com/hashicorp/hcl/json/parser"  v0.0.0-20170505105837-392dba7d905e
"github.com/hashicorp/hcl/json/scanner"  v0.0.0-20170505105837-392dba7d905e
"github.com/hashicorp/hcl/json/token"  v0.0.0-20170505105837-392dba7d905e
"github.com/magiconair/properties"  v0.0.0-20170321103039-51463bfca257
"github.com/mitchellh/mapstructure"  v0.0.0-20170522200023-d0303fe80992
"github.com/pelletier/go-buffruneio"  v0.0.0-20170227140311-c37440a7cf42
"github.com/pelletier/go-toml"  v0.0.0-20170601235532-fe7536c3dee2
"github.com/spf13/afero"  v0.0.0-20170217174146-9be650865eab
"github.com/spf13/afero/mem"  v0.0.0-20170217174146-9be650865eab
"github.com/spf13/cast"  v0.0.0-20170413105028-acbeb36b902d
"github.com/spf13/cobra"  v0.0.0-20170528124206-99ff9334bda2
"github.com/spf13/jwalterweatherman"  v0.0.0-20170523113943-0efa5202c046
"github.com/spf13/pflag"  v0.0.0-20170508204326-e57e3eeb33f7
"github.com/spf13/viper"  v0.0.0-20170417100815-0967fc9aceab
"github.com/toukii/goutils"  v0.0.0-20180223150238-86cba64d65c8
"github.com/toukii/jsnm"  v0.0.0-20180224181733-38735b07ca23
"github.com/toukii/pull/command"  v0.0.0-20171230003049-1375e3512a18
"golang.org/x/sys/unix"  latest
"golang.org/x/text/transform"  latest
"golang.org/x/text/unicode/norm"  latest
"gopkg.in/yaml.v2"  latest
)
```