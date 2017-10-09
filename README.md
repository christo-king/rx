#Standard Deviation Example Project

-----

### And Example of Go services with a React.js interface
This project was built with Docker (docker-compose, specifically) and usees a MySQL database to house standard deviations calculated form user input.  

### Required Software
This project was built with the following technologies, which are therefore also be recommended for its use.
- [Docker](link:https://www.docker.com/)
- [Node.js](link:https://nodejs.org/en/)
- [Npm](link:https://www.npmjs.com/)
- [Percona](link:https://www.percona.com/)
- [Gin](link:https://github.com/gin-gonic/gin)
- [Babel](link:https://babeljs.io/)
- [Webpack](link:https://webpack.github.io/)
- [React.js](link:https://facebook.github.io/react/)
- [React-Bootstrap](link:https://react-bootstrap.github.io/)
- [Postman](link:https://www.getpostman.com/)
- [Ruby](link:https://www.ruby-lang.org/)
- [Sinatra](link:http://www.sinatrarb.com/)

To make full use of this project, you will need to, at minimum, have the following installed:
 - recent, native (non-native could work, but is untested), version of Docker (which includes the docker-compose command line utility)
 - Go SDK (for 'go get' package management)
 - npm 

###Instructions for Usage
After checking out this project, you will have to start a docker-compose instance of the Go server (Gin) and start up webpack.  These will provide you with a base running environment which you can then work with.

```bash
go get github.com/go-sql-driver/mysql
go get github.com/gorilla/mux
git clone https://github.com/christo-king/rx 
cd rx
npm update
docker-compose up -d
npm run dev
```
Once you've completed the above, you shoul dhave a running environment.  You will then be able to fire up your favorite browser and go to 


<a href="http://localhost:3000/" target="_blank">http://localhost:3000/</a>

...and you should see the testing page

![Screenshot][screenshot]

Also provided is a small series of [Postman](link:https://www.getpostman.com/) tests for the web services (/RX.postman.json).

Web services are written in both Go (using gorilla) and Ruby (using sinatra).  The drop-down "Server" can choose which is used to submit new standard deviations, but all go to the same database.  In the future, the list of existing will also come from the selected server.

###Future Additions
Future additions will likely include more extensive client-side unit testing and BDD-style tests in Ruby-Cucumber.


[screenshot]: https://github.com/christo-king/rx/raw/master/screenshot.png "Screenshot of Testing Page"

