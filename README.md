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
"github.com/astaxie/beego/httplib"  v1.9.2
"github.com/everfore/exc"  vgo.0.1
"github.com/everfore/exc/walkexc"  vgo.0.1
"github.com/everfore/exc/walkexc/pkg"  vgo.0.1
"github.com/fatih/color"  v1.6.0
"github.com/fsnotify/fsnotify"  v1.4.7
"github.com/hashicorp/hcl"  v0.0.0-20170505105837-392dba7d905e
"github.com/hashicorp/hcl/hcl/ast"  v0.0.0-20170505105837-392dba7d905e
"github.com/hashicorp/hcl/hcl/parser"  v0.0.0-20170505105837-392dba7d905e
"github.com/hashicorp/hcl/hcl/scanner"  v0.0.0-20170505105837-392dba7d905e
"github.com/hashicorp/hcl/hcl/strconv"  v0.0.0-20170505105837-392dba7d905e
"github.com/hashicorp/hcl/hcl/token"  v0.0.0-20170505105837-392dba7d905e
"github.com/hashicorp/hcl/json/parser"  v0.0.0-20170505105837-392dba7d905e
"github.com/hashicorp/hcl/json/scanner"  v0.0.0-20170505105837-392dba7d905e
"github.com/hashicorp/hcl/json/token"  v0.0.0-20170505105837-392dba7d905e
"github.com/magiconair/properties"  v1.7.6
"github.com/mitchellh/mapstructure"  v0.0.0-20170522200023-d0303fe80992
"github.com/pelletier/go-buffruneio"  v0.2.0
"github.com/pelletier/go-toml"  v1.1.0
"github.com/spf13/afero"  v1.0.2
"github.com/spf13/afero/mem"  v1.0.2
"github.com/spf13/cast"  v1.2.0
"github.com/spf13/cobra"  v0.0.1
"github.com/spf13/jwalterweatherman"  v0.0.0-20170523113943-0efa5202c046
"github.com/spf13/pflag"  v1.0.0
"github.com/spf13/viper"  v1.0.0
"github.com/toukii/goutils"  v0.1.1
"github.com/toukii/jsnm"  v0.0.0-20180224181733-38735b07ca23
"github.com/toukii/pull/command"  v0.0.1
"golang.org/x/sys/unix"  v0.0.0-20171109145042-1e2299c37cc9
"golang.org/x/text/transform"  v0.0.0-20170328164024-ab6d1c143672
"golang.org/x/text/unicode/norm"  v0.0.0-20170328164024-ab6d1c143672
"gopkg.in/yaml.v2"  v2.1.1
)
```