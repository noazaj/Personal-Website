# Personal Website

This is my personal website, built to showcase my work and skills. It uses Go for the backend and traditional frontend technologies like HTML, CSS, and JavaScript. Future enhancements may include the use of HTMX for modern, AJAX-like behavior without writing any JavaScript.

## Dependencies

List of dependencies utilized
```
go get github.com/joho/godotenv
```
```
github.com/go-chi/chi/v5
```

## Scripts
There is one script written to source the .env variables to the binary. Depending on how one might implement the script and its location, permissions may be denied. You will need to run this command in the terminal to allow permission.
```
chmod +x ./path/to/your/script/example.sh
```

This script pulls from the .env file located in the root directory of backend. The server will run off of the makefile ```make all``` command in the terminal. This is possible because the relative path is used for the script to start the server instead of the typical "go run main.go" command. Using this command means being in the working directory where the main.go file is with the appropriate pathing to load the .env file.

## To-Do

+ Add content to home page
+ Style the home page
+ Look into hosting services besides localhost