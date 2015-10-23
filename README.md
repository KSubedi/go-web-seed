go-web-seed
===================


**go-web-seed** is a seed for building golang based web applications quickly. It contains a simple middleware system, routing system using gorilla's mux, sessions using gorilla sessoins and some other core utilities including logging.

----------


Getting Started
-------------

Getting started with go-web-seed is easy.  Just clone the repository and you are ready to start building. You might want to change the "imports" to match the name of the new package if you rename the project. Here is how to get started:

    go get -u https://github.com/ksubedi/go-web-seed
    cd $GOPATH/src/github.com/ksubedi/go-web-seed
    go get -u

Then you can edit the routes on routes.go file, edit configuration at config.go & add controllers and middlewares at their respective folders.