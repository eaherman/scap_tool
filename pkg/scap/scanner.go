package scap

import (
	"fmt"
	"log"
	"os/exec"
)

type ScapScan struct {
	Log string
}

func (ss *ScapScan) Run(sr *ScapRunner) {
	stigs := sr.Stigs.Stigs
	for _, stig := range stigs {
		for _, rule := range stig.Rules {
			if sr.Fix {
				fmt.Println(rule.FixCommand)
				_, err := exec.Command("bash", "-c", rule.FixCommand).Output()
				if err != nil {
					fmt.Println(err)
				}
			}
			if sr.Checks {
				for _, check := range rule.Checks {
					temp, err := exec.Command("bash", "-c", check.CheckCommand).Output()
					fmt.Println(string(temp))
					if err != nil {
						fmt.Println("In fatal")
						log.Fatal(err)
					}
					check.Result.Res = string(temp)
					fmt.Println(check.Result.Res)

					if check.Result.Res == check.ExpectedResult {
						fmt.Println("In true")
						check.Pass = true
					} else {
						check.Pass = false
					}

				}
			}
		}

	}
}
