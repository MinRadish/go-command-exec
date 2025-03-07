/*
 * @Author: Radish
 * @Date: 2025-03-07 15:49
 * @LastEditors: Radish
 * @LastEditTime: 2025-03-07 16:34
 * @Description:
 * - 测试所有包
 * - go test ./... -v
 * Copyright (c) 2025 by EIZ, All Rights Reserved.
 */

package commandexec

import (
	"flag"
	"testing"
)

func TestCmdDefs_New(t *testing.T) {
	cmds := CmdDefs{}.New()
	if cmds == nil {
		t.Error("Expected non-nil CmdDefs")
	}
}

func TestCmdDefs_PrintHelp(t *testing.T) {
	cmds := CmdDefs{}.New()
	cmds["test"] = CmdDef{
		Name: "test",
		Desc: "This is a test command",
		Funs: []func(){},
	}
	cmds.PrintHelp()
}

func TestCmdDefs_Exec(t *testing.T) {
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
	if !executed {
		t.Error("Expected command to be executed")
	}
}

func TestCmdDefs_Exec_InvalidCommand(t *testing.T) {
	cmds := CmdDefs{}.New()
	cmds.Exec("invalid")
}

func TestCmdDefs_ExecFlag(t *testing.T) {
	cmds := CmdDefs{}.New()
	executed := false
	cmds["test"] = CmdDef{
		Name: "test",
		Desc: "This is a test command",
		Funs: []func(){
			func() { executed = true },
		},
	}

	// Set the flag value
	flag.CommandLine = flag.NewFlagSet("test", flag.ExitOnError)
	flag.CommandLine.String(FlagName, "test", "exec operate")
	cmds.ExecFlag()
	if !executed {
		t.Error("Expected command to be executed")
	}
}
