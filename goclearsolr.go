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

	//urlMap contains mapping of instance names to instance urls
	//make sure to include /update in the solr url
	urlMap := map[string]string{
		"local": "http://localhost:8983/solr/corename/update",
	}

	//the instance variable selects which url we're pulling out of the hashmap
	var instance string
	fmt.Printf("Which SOLR instance are we clearing? ")
	fmt.Scan(&instance)

	//you know, in case you're an all caps type of person. I suppose you could remove this if you gave the instances cased names
	instance = strings.ToLower(instance)

	//assign the url value to a new variable
	url := urlMap[instance]

	if url != "" {
		fmt.Println("preparing to clear solr indexing for: " + instance + " at URL " + url)

		//first we clear
		fmt.Println("deleting data")
		exec.Command("curl", url, "--data", "<delete><query>*:*</query></delete>", "-H", "Content-type:text/xml; charset=utf-8").Run()

		fmt.Println("data deleted")

		//then we commit
		fmt.Println("Now committing")
		exec.Command("curl", url, "--data", "<commit/>", "-H", "Content-type:text/xml; charset=utf-8").Run()

		//then we celebrate
		fmt.Println("All done! Go have a beer!")
	} else {
		//BZZZ ERROR
		fmt.Println("No URL found for that instance, go update this script and try again")
	}

}
