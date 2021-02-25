package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	// Initial declaration
 m := map[string]interface{}{
	"key": "value",
}
// Dynamically add a sub-map
 m["sub"] = map[string]interface{}{
	"deepKey": "deepValue",
}
// returns map
 var f interface{}


	//yaml-exec file1.yaml
	//takes second os arg and reads file
	source, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	//print raw yaml
	fmt.Printf("File contents: %s", source)

	//unpacks yaml file in a map 
	err = yaml.Unmarshal([]byte(source), &f) 
	if err != nil {
		log.Printf("error: %v", err)
	}

	//print map of yaml
	fmt.Printf("%+v\n", f)
}
