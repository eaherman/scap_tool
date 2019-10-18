package scap

import (
	"fmt"
	"net"
)

type Stig struct {
	Release int
	Version int
	Name    string
	Checks  []Check
}

type Check struct {
	CheckCommand   string
	ExpectedResult string
	CheckFile      string
	FileContents   string
	FixCommand     string
	Result
}

type Result struct {
	Result string
	Err    error
	Pass   bool
}

type Host struct {
	Name string
	Ip   net.IPAddr
	Mac  net.HardwareAddr
	Fqdn string
}

type Options struct {
	Log     bool
	Checks  string
	Results string
	Testing bool
	XCCDF   bool
	JSON    bool
	CKL     bool
}

type ScapScans struct {
	Host
	Options
}

type ScapRunner struct {
	ScapScans
	Stigs []Stig
	Test  string
}

func (sr *ScapRunner) SetHost() {
	sr.Test = "Test"
	sr.Name = "Name"
	fmt.Println(sr.ScapScans)
}
