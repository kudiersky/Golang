CHALLENGE #1
 Create a user/password protected program.

EXAMPLE USER
 username: Jake
 password: 1987

``` 
EXPECTED OUTPUT
 go run main.go
   Usage: [username] [password]

 go run main.go albert
   Usage: [username] [password]

 go run main.go hacker 42
   Access denied for "hacker".

 go run main.go jack 6475
   Invalid password for "Jake".

 go run main.go jack 1887
   Access granted to "Jake".

```
CHALLENGE #2
 Create a Test for the programme.