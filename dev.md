- clean up code

- make sure that a distinct status code is returned for all bad token verifications

- deleted images from test server trial gcp account bucket

- make sure the data you pass into the pg functions in configure_girc.go are set either before or after you pass it into the functions. Right now, User structs are set in configure_girc.go and Incident structs are set in pg.go. Choose one or the other.

- make sure the db.Model().Where(shit) shit is a struct and not "userid = ?" bullshit format since you know how to do the struct way correctly now

- make sure thar gcp test is still only negative 16 dollars.

- log errors before panic or return err

- explain to keith about data model and changing forms for veterinarian

- test with updating empty fields

- add logger to database

- see if there are any registered users in old project backend

- get settigs json from gcp for gircapp3@gmail.com/gircapp

- TODO: change imports from github.com/somersbmatthews/gircapp2 to github.com/gircapp/gircapp



swagger generate server -f swagger2.yaml -A girc 


for future version of app

- add email confirmation

- add password reset email 
