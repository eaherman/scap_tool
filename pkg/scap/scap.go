package scap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

type Stig struct {
	Release int    `json:"release"`
	Version int    `json:"version"`
	Name    string `json:"name"`
	Rules   []Rule `json:"rules"`
}

type Rule struct {
	RuleID string  `json:"ruleID"`
	VulnID string  `json:"vulnID"`
	Checks []Check `json:"checks"`
	Fix    `json:"fix"`
}

type Fix struct {
	FixCommand string `json:"command"`
	FixScript  string `json:"script"`
}

type Check struct {
	Type           string `json:"type"`
	CheckCommand   string `json:"command"`
	Contains       string `json:"contains"`
	Regex          string `json:"regex"`
	ExpectedResult string `json:"expectedResult"`
	Script         string `json:"script"`
	CheckFile      string `json:"checkFile"`
	Result         `json:"-"`
}

type Result struct {
	Res  string
	Err  error
	Pass bool
}

type Host struct {
	Name string
	IP   net.IPAddr
	MAC  net.HardwareAddr
	FQDN string
}

type Options struct {
	Log     bool
	InFile  string
	Results string
	Testing bool
	XCCDF   bool
	JSON    bool
	CKL     bool
	Fix     bool
	Checks  bool
}

type ScapScans struct {
	Host
	Options
}

type ScapRunner struct {
	ScapScans `json:"-"`
	Stigs
	//	Stigs     []Stig `json:"stigs"`
}

type Stigs struct {
	Stigs []Stig `json:"stigs"`
}

func (sr *ScapRunner) SetHost() {
	sr.Name = "Host Name Test"
}

func (sr *ScapRunner) ParseFile() {
	file, err := os.Open(sr.InFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bytes, berr := ioutil.ReadAll(file)
	if berr != nil {
		log.Fatal(berr)
	}
	jerr := json.Unmarshal(bytes, &sr.Stigs)
	if jerr != nil {
		fmt.Println("error:", jerr)
	}
	fmt.Printf("%+v\n", sr)
}
