package main

import (
	"flag"
	"fmt"
	"os"

	"git.fractalqb.de/fractalqb/gomk"
	"git.fractalqb.de/fractalqb/gomk/mktask"
)

var (
	must       = gomk.LogMust
	toolUpdate bool

	goBuild = gomk.CmdDef{
		Name: "go",
		Args: []string{"build", "--trimpath", "-ldflags", "-s -w"}, // What about -a
	}

	cmds = []string{"edpc", "edpceh"}
)

func main() {
	flag.BoolVar(&toolUpdate, "u", false, "Get tool update")

	prj := gomk.NewProject(must, &gomk.Config{Env: os.Environ()})

	tStringer := mktask.NewGetStringer(must, prj, toolUpdate)
	tVersioner := mktask.NewGetVersioner(must, prj, toolUpdate)

	tGen := gomk.NewCmdTask(must, prj, "generate", "go", "generate", "./...").
		DependOn(tStringer.Name(), tVersioner.Name())

	tCmds := gomk.NewNopTask(must, prj, "commands")

	for _, cmd := range cmds {
		t := gomk.NewCmdDefTask(must, prj, fmt.Sprintf("cmd:%s", cmd), goBuild).
			WorkDir("cmd", cmd).
			DependOn(tGen.Name())
		tCmds.DependOn(t.Name())
	}

	if len(flag.Args()) == 0 {
		gomk.Build(prj, "commands")
	} else {
		for _, target := range flag.Args() {
			gomk.Build(prj, target)
		}
	}
}
