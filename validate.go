package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/coreos/coreos-cloudinit/config/validate"
	ignConfig "github.com/coreos/ignition/config"
)

var (
	flags = struct {
		file string
	}{}
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	flag.StringVar(&flags.file, "file", "/tmp/ignition.json", "file to validate")
}

func main() {
	flag.Parse()
	fmt.Println("Validating " + flags.file)
	code := validateFile()
	os.Exit(code)
}

func validateFile() (int) {
	dat, err := ioutil.ReadFile(flags.file)
	check(err)
	config := bytes.Replace(dat, []byte("\r"), []byte{}, -1)

	_, rpt, err := ignConfig.Parse(config)
	switch err {
	case ignConfig.ErrCloudConfig, ignConfig.ErrEmpty, ignConfig.ErrScript:
		rpt, err := validate.Validate(config)
		check(err)
		for _, entry := range rpt.Entries() {
			fmt.Fprintf(os.Stderr, "problem: %s", entry)
		}
		return -1
	case ignConfig.ErrUnknownVersion:
		os.Stderr.WriteString("Failed to parse config. Is this a valid Ignition Config, Cloud-Config, or script?")
		return -1
	default:
		rpt.Sort()
		fmt.Println(rpt.Entries)
		return len(rpt.Entries)
	}
}
