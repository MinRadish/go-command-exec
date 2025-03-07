# go-command-exec

`go-command-exec` 是一个用于在 Go 语言中根据参数执行不同func

## 安装

使用以下命令安装：

```sh
go get github.com/MinRadish/go-command-exec
```

## 使用方法

以下是一个简单的示例，展示如何使用 `go-command-exec` 包来执行系统命令：

```go
package main

import (
    "fmt"
    "log"
    "github.com/MinRadish/go-command-exec"
)

func main() {
    cmds := CmdDefs{}.New()
    executed := false
    cmds["test"] = CmdDef{
        Name: "test",
        Desc: "This is a test command",
        Funs: []func(){
            func() { executed = true },
        },
    }
    cmds.Exec("test")
}
```

## 功能

- 接收参数执行不同的func

## 贡献

欢迎贡献代码！请 fork 本仓库并提交 pull request。

## 许可证

本项目采用 MIT 许可证。详情请参阅 [LICENSE](LICENSE) 文件。
