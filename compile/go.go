// takes in code as arg from go
//run go build on code given

package compile

import (
	"fmt"
	"log"
	"os/exec"
)

//Go creates exe from file passed in as arg
func Go(arg string) {	
	fmt.Printf(arg)
	cmd := exec.Command("go", "build", arg)

	stdout, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
}

fmt.Print(string(stdout))
}

