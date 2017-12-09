

PREREQUISITES:
The following packages must be installed on the system:
Go Language:       https://golang.org/
NodeJS (with NPM): https://nodejs.org/en/

The npm and go binaries must be available in the path.

STARTING APPLICATION:
1. After the prerequisites are installed, the application prerequesites must
be installed. In this directory, run the following command. (This only needs
to be run once.)
   npm install

2. Start the server
   go run simpleServer.go

The server should print a message that it is running on port 8081.

Access the application at:
http://localhost:8081

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

        a. Unit tests (on both the front-end JavaScript and back-end Go)
        b. Integration tests (Using a selenium-based test runner to validate
           the end to end experience.)
        c. Build framework (such as webpack). This will run the unit tests
           as well as lint code and minify code and produce artifacts



3. Unit tests are also missing. 

4. End to End testing is also missing. A selinium-based framework

TODO:
Fix dates.
Wider display.
Code cleanup
Better gab between chart.
error handling
tests
