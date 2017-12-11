

PREREQUISITES:
The following packages must be installed on the system:
Go Language:       https://golang.org/
NodeJS (with NPM): https://nodejs.org/en/

The npm and go binaries must be available in the path.

BUILD AND RUN:
(All commands assume you are in the root of the project directory)
If bash is available, run the follwoing command to install required modules and launch:
start.sh

-OR-
On other platforms, you may manually build and run:
npm install -g brunch
npm install
brunch build
go run simpleService.go

The server currently runs in the foreground. Keep the window open and use CTRL-C to stop.

To run in the background:
nohup go run simpleServer.go &

The app is available at:
http://localhost:8081

TESTS:
Run the following comand to execute the go tests
go test


IMPLEMENTATION NOTES:
1. The front end UI is created using the Vue front-end framework with Chart.JS
charting library.  I chose these because they were the most popular items
in their class that I had never used before. The were adequate for current
purposes. However, were this a production environment, there would have been
many additional factors that came into play. (Factors include: What is already
being used? What other features will we need? What level of support is 
available in the community?)

2. This has been constructed as a quick POC. Many additional components
would be in place in a prouction. This includes:

        a. More robust tests
        b. A dynamic reload development environment. (Using brunch watch)
        c. End to End tests (most likely via a selenium framework)
        d. More extensive error handling. (back end should deal with bad data, etc.)
        e. More flexible sizing of charts
                
