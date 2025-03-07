package commandexec

import (
	"flag"
	"fmt"
	"sort"
	"time"
)

// 命令定义
type CmdDef struct {
	Name string
	Desc string
	Funs []func()
}

// 定义命令集合
type CmdDefs map[string]CmdDef

// 创建一个
func (CmdDefs) New() CmdDefs {
	return make(map[string]CmdDef)
}

var FlagName = "exec"

// 输出帮助信息
func (c *CmdDefs) PrintHelp() {
	k := make([]string, len(*c))
	i := 0
	for n := range *c {
		k[i] = n
		i++
	}
	sort.Strings(k)
	fmt.Printf(`go run main.go -%s [options]
    options: must`, FlagName)
	fmt.Println()
	for _, v := range k {
		cv := (*c)[v]
		fmt.Printf("      \033[1;32;32m%s\033[0m\n", cv.Name)
		fmt.Printf("        \033[1;36;0m%s\033[0m\n", cv.Desc)
		fmt.Println()
	}
}

// 执行命令
func (c *CmdDefs) Exec(s string) {
	if v, ok := (*c)[s]; ok {
		startAt := time.Now()
		for _, f := range v.Funs {
			f()
		}
		fmt.Printf("\n\033[1;33;33m%s\033[0m \033[1;35;35m%s\033[0m\n", "执行时间:", time.Since(startAt))
	} else {
		fmt.Println("无效的命令")
		c.PrintHelp()
	}
}

// 通过ExecFlag执行
func (c *CmdDefs) ExecFlag() {
	exec := flag.String(FlagName, "", "exec operate")
	flag.Parse()
	if *exec != "" {
		c.Exec(*exec)
	} else {
		c.PrintHelp()
	}
}
