package cmds

import (
	"encoding/json"
	"flag"
	"io"

	"git.fractalqb.de/fractalqb/pack/ospath"
	"git.fractalqb.de/fractalqb/yacfg"
)

var (
	Paths ospath.App
)

func Configure(cfg interface{}) error {
	const configFlag = "cfg"
	Paths = ospath.NewApp(ospath.ExeDirTree(nil), "fqb", "bcplus")
	cfgr := &yacfg.EnvWithPrefixThenFlagFiles{
		EnvPrefix:     "EDPC_",
		FilesFlagName: configFlag,
	}
	flag.String(configFlag, "", cfgr.ListFlagDoc())
	if err := cfgr.Configure(cfg); err != nil {
		return err
	}
	return nil
}

func DumpConfig(wr io.Writer, cfg interface{}) {
	enc := json.NewEncoder(wr)
	enc.SetIndent("", "  ")
	enc.Encode(cfg)
}
