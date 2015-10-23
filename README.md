go-web-seed
===================


**go-web-seed** is a seed for building golang based web applications quickly. It contains a simple middleware system, routing system using gorilla's mux, sessions using gorilla sessoins and some other core utilities including logging.

----------


Getting Started
-------------

Getting started with go-web-seed is easy.  Just clone the repository and you are ready to start building. You might want to change the "imports" to match the name of the new package if you rename the project. Here is how to get started:

    go get -u github.com/ksubedi/go-web-seed
    cd $GOPATH/src/github.com/ksubedi/go-web-seed

Then you can edit the routes on routes.go file, edit configuration at config.go & add controllers and middlewares at their respective folders.

Auto Reload Script
-------------

The run.sh script located on the root directory of the repository will allow you to use CompileDaemon to auto compile and reload the web application as soon as you make changes. This allows for faster development. To get started, you need to get the CompileDaemon tool first.

    go get -u github.com/githubnemo/CompileDaemon

Then you can run the run.sh script and it will auto detect changes, compile the app and run it.

    sh run.sh

Or

    bash run.sh

Or

    chmod +x run.sh
    ./run.sh
