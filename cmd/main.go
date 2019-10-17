package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	//	"strings"
)

func main() {
	cmd2 := "grep -i x11forwarding /etc/ssh/sshd_config | grep -v '^#'"
	cmd := exec.Command("bash", "-c", cmd2)
	//	cmd.Stdin = strings.NewReader("-al")
	//cmd.Stdin = strings.NewReader("logout /etc/dconf/local.d/*")
	var out bytes.Buffer
	var e bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &e
	err := cmd.Run()
	if err != nil {
		fmt.Println("In Error")
		fmt.Printf("Error %q\n", e.String())
		log.Fatal(err)
	}
	fmt.Printf("%q\n", out.String())

}
