- clean up code

- make sure that a distinct status code is returned for all bad token verifications

- deleted images from test server trial gcp account bucket

- make sure the data you pass into the pg functions in configure_girc.go are set either before or after you pass it into the functions. Right now, User structs are set in configure_girc.go and Incident structs are set in pg.go. Choose one or the other.

- make sure thar gcp test is still only negative 23 dollars.

- log errors before panic or return err

- test with updating empty fields

- add logger to database

- see if there are any registered users in old project backend


swagger generate server -f swagger4.yaml -A girc 

swagger generate server -t gen -f ./swagger4.yml --exclude-main -A girc

swagger generate server -f ./swagger4.yml --exclude-main -A girc

steps to deploy for testing:

- add signal server shutdown in main.go

- change code in package pg for gcloud sql

- add go-swagger rate limiting code with tollbooth

- make sure main.go is setup for google cloud app engine 

- make a load balancer with NEB for google app engine

- run api tests

- make sure swagger docs are deployed


for future version of app

- add email confirmation

- add password reset email 
