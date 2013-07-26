package main

import (
	"fmt"
	"exec"
)

func main() {

	var instance string


	instance = "local"
	
	fmt.Println("preparing to clear solr indexing for: " + instance)

	//curl http://localhost:8983/solr/openIdeas/update --data '<delete><query>*:*</query></delete>' -H 'Content-type:text/xml; charset=utf-8'
	//curl http://localhost:8983/solr/openIdeas/update --data '<commit/>' -H 'Content-type:text/xml; charset=utf-8'
	


}