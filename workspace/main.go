package workspace

import (
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
)

func workSpace(path string) (bool, error) {
	//check if file path exists, returns err = nil if file exists
	_, err := os.Stat(path) 

	if err == nil { 
		fmt.Println("Path already exists")
	}

	// should return true if file doesn't exist
	if os.IsNotExist(err) {
		//prompt user if they'd like dir to be created
		mk := yesNo()

		if mk == true {
			errDir := os.MkdirAll(path, 0755)
			//should return nil once directory is made, if not, throw err
			if errDir != nil {
				log.Fatal(err)
			}
			fmt.Println("Path created")

		} else {
			fmt.Println("Please create a directory for the Builder")
			return true, err
		}
	}

	//check workspace env exists, if not, create it
	val, present := os.LookupEnv("BUILDER_WORKSPACE") 
	if !present {
		os.Setenv("BUILDER_WORKSPACE", path) 
		fmt.Println("BUILDER_WORKSPACE", os.Getenv("BUILDER_WORKSPACE"))
	} else {
		fmt.Println("BUILDER_WORKSPACE", val)
	}
	return true, err
}

func yesNo() bool {
	prompt := promptui.Select{
			Label: "Select[Yes/No]",
			Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
	}
	return result == "Yes"
}

//MakeMainDir does...
func MakeMainDir() {
	workSpace("C:/var/lib/builder/workspace")
}