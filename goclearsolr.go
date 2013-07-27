package main

import (
	"fmt"
	"os/exec"
	"strings"
)

/**
 * One simple function to call a couple of curl commands
 */
func main() {

	
	var url string

	//url is the url of the core you want to clear the indexing data for, ex: http://localhost:8983/solr/corename/update
	//make sure to include /update at the end of the url
	url = "http://localhost:8983/solr/corename/update"

	//the instance variable is really only for logging purposes, ex: local, QA, UAT, Prod
	var instance string
	fmt.Printf("Which SOLR instance are we clearing? ")
	fmt.Scan(&instance)

	instance = strings.ToLower(instance)
	
	fmt.Println("preparing to clear solr indexing for: " + instance)

	//first we clear
	fmt.Println("deleting data")
	exec.Command("curl", url, "--data", "<delete><query>*:*</query></delete>", "-H", "Content-type:text/xml; charset=utf-8").Run()

	fmt.Println("data deleted")

	//then we commit
	fmt.Println("Now committing")
	exec.Command("curl", url, "--data", "<commit/>", "-H", "Content-type:text/xml; charset=utf-8").Run()

	//then we celebrate
	fmt.Println("All done! Go have a beer!")
}