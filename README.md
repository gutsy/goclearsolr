goclearsolr
===========

A simple script written in Go to clear indexed solr data from a solr instance. Be kind, this is my first stab at writing Go. I'm normally a Java guy.

It simply runs some curl commands to clear the data off of a particular instance. 

To use: 

1. pull the code, (obviously)

2. replace the url and instance variables with the url of the core (including /update at the end, of course) and the instance name (if you care, of course)

3. run the script in a command line with 'go run goclearsolr.go'

If you want to use this for multiple solr instances, simply copy and rename the files (I have mine as 'goclearlocal', 'gocleardev', 'goclearqa' and so on).

TODO: Make this one script that will take an argument for each instance, and store/pull the instance url/name from an array instead of needing multiple scripts. 

