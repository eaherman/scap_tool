package main

import (
	//	"encoding/json"
	"flag"
	"fmt"
	//	"io/ioutil"
	//	"log"
	//	"os"
	"scap_tool/pkg/scap"
)

func main() {

	//	file, err := os.Open("/home/eherman/Documents/checkexample.json")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	defer file.Close()
	//	stigs := scap.Stigs{}
	//	bytes, berr := ioutil.ReadAll(file)
	//	if berr != nil {
	//		log.Fatal(berr)
	//	}
	//	jerr := json.Unmarshal(bytes, &stigs)
	//	if jerr != nil {
	//		fmt.Println("error:", jerr)
	//	}
	//	fmt.Printf("%+v\n", stigs)

	/*
		file, err := os.Open("/Users/EricMacBook/Documents/test-benchmark.xml") // For read access.
		if err != nil {
			log.Fatal(err)
		}

		json, err := xj.Convert(file)
		if err != nil {
			panic("That's embarrassing...")
		}

		fmt.Println(json.String())
	*/
	sr := scap.ScapRunner{}

	inFileHelp := "The json file with the checks. Can be a single json file, a zip file of " +
		"multiple json files, or a tar file with multiple json files"
	flag.BoolVar(&sr.Log, "log", false, "Turn on debug logging")
	flag.StringVar(&sr.InFile, "checkFile", "checkFile.json", inFileHelp)
	flag.BoolVar(&sr.Testing, "testing", false, "Outputs results in json to std out")
	flag.BoolVar(&sr.XCCDF, "xccdf", false, "Creates XCCDF file with the results")
	flag.BoolVar(&sr.JSON, "json", false, "Creates JSON file with the results")
	flag.BoolVar(&sr.CKL, "ckl", false, "Creates CKL file with the results")
	flag.StringVar(&sr.Results, "results", "results", "The name of the file with the results")
	flag.BoolVar(&sr.Checks, "checks", true, "Run the checks against the desired output. Ran after fix if fix is true")
	flag.BoolVar(&sr.Fix, "fix", false, "Run the fix commands")

	flag.Parse()

	sr.ParseFile()

	ss := scap.ScapScan{}

	ss.Run(&sr)

	fmt.Printf("%+v\n", sr)
	//	sr := scap.ScapRunner{}
	//	sr.SetHost()
	//	fmt.Println(sr)

	//	cmd2 := "grep -i x11forwarding /etc/ssh/sshd_config | grep -v '^#'"
	//	cmd := exec.Command("bash", "-c", cmd2)
	//	cmd.Stdin = strings.NewReader("-al")
	//cmd.Stdin = strings.NewReader("logout /etc/dconf/local.d/*")
	//	var out bytes.Buffer
	//	var e bytes.Buffer
	//	cmd.Stdout = &out
	//	cmd.Stderr = &e
	//	err := cmd.Run()
	//	if err != nil {
	//		fmt.Println("In Error")
	//		fmt.Printf("Error %q\n", e.String())
	//		log.Fatal(err)
	//	}
	//	fmt.Printf("%q\n", out.String())

}
