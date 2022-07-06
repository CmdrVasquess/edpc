package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"git.fractalqb.de/fractalqb/pack/ospath"
	"github.com/CmdrVasquess/watched"
)

type SysIdent struct {
	ESeq watched.JEventID
	Addr int64
	Name string
}

type Commander struct {
	FID    string
	Name   string
	System SysIdent
}

func (cmdr *Commander) Unknown() bool { return cmdr.FID == "" }

func (cmdr *Commander) SwitchTo(p ospath.AppPaths, fid, name string) error {
	log.Infov("Switch to `cmdr` with `FID`", name, fid)
	if !cmdr.Unknown() {
		filename := p.RoamingData(cmdrFile(cmdr.FID, cmdr.Name))
		err := cmdr.save(filename)
		if err != nil {
			return err
		}
	}
	if fid == "" {
		cmdr.FID = ""
		return nil
	}
	filename := p.RoamingData(cmdrFile(fid, name))
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		cmdr.FID = fid
		cmdr.Name = name
		return nil
	}
	if err := cmdr.load(filename); err != nil {
		return err
	}
	if cmdr.FID != fid {
		return fmt.Errorf("found %s '%s' in '%s' for commander %s",
			cmdr.FID,
			cmdr.Name,
			filename,
			fid,
		)
	}
	return nil
}

func (cmdr *Commander) save(file string) error {
	return wrFile(file, func(wr io.Writer) error {
		enc := json.NewEncoder(wr)
		enc.SetIndent("", "   ")
		return enc.Encode(cmdr)
	})
}

func (cmdr *Commander) load(file string) error {
	rd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer rd.Close()
	dec := json.NewDecoder(rd)
	return dec.Decode(cmdr)
}

func cmdrFile(fid, name string) string {
	return fmt.Sprintf("cmdr-%s.json", fid)
}

func wrFile(name string, do func(io.Writer) error) error {
	tmpname := name + "~"
	wr, err := os.Create(tmpname)
	if err != nil {
		return err
	}
	defer wr.Close()
	if err = do(wr); err != nil {
		return err
	}
	if err = wr.Close(); err != nil {
		return err
	}
	return os.Rename(tmpname, name)
}
