- clean up code

- make sure that a distinct status code is returned for all bad token verifications

- deleted images from test server trial gcp account bucket

- make sure the data you pass into the pg functions in configure_girc.go are set either before or after you pass it into the functions. Right now, User structs are set in configure_girc.go and Incident structs are set in pg.go. Choose one or the other.

- make sure the db.Model().Where(shit) shit is a struct and not "userid = ?" bullshit format since you know how to do the struct way correctly now

- make sure thar gcp test is still only negative 16 dollars.

- log errors before panic or return err

