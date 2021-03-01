// search for parent directory as env var
// mkdir from artifact name

package sub

import (
	"fmt"
	"os"
	"log"
	"io"
)

//MakeSubDir does ...
func MakeSubDir() {
	//get parent directory 
	val, present := os.LookupEnv("BUILDER_WORKSPACE") 
	if present {
		fmt.Println("BUILDER_WORKSPACE", val)
	}

	newDir := val + "/hello"

	// create child directory with project name
	os.Mkdir(newDir, 0755)

	from, err := os.Open("../Builder/hello.exe")
  if err != nil {
    log.Fatal(err)
  }
  defer from.Close()

  to, err := os.OpenFile("C:/var/lib/builder/workspace/hello/hello.exe", os.O_RDWR|os.O_CREATE, 0666)
  if err != nil {
    log.Fatal(err)
  }
  defer to.Close()

  _, err = io.Copy(to, from)
  if err != nil {
    log.Fatal(err)
  }

	defer os.Remove("C:/Users/Ivan/go/src/github.com/ilarocca/Builder/hello.exe")
}

