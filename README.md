goclearsolr
===========

A simple script written in Go to clear indexed solr data from a solr instance. Be kind, this is my first stab at writing Go. I'm normally a Java guy.

It simply runs some curl commands to clear the data off of a particular instance. 

To use: 

1. pull the code, (obviously)

2. replace the url and instance variables with the url of the core (including /update at the end, of course) and the instance name (used for lookup) in the hashmap

3. run the script in a command line with 'go run goclearsolr.go'

4. enter the instance name you're wanting to clear and the script will do the rest. 



