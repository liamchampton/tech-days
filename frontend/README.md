# Frontend
The FE for this experimental application is built with Go & [GopherJS](https://github.com/gopherjs/gopherjs). GopherJS generates JavaScript from Go code and outputs it in the `/scripts` directory.

## Pre-requisites 
- Go 1.18
- GopherJs installed `go install github.com/gopherjs/gopherjs@v1.18.0-beta2`

## Environment variables
The FE requires the URL for the backend. This should be exported as:
```
$ export BACKEND_URL=XXX
```
Once this is exported, run `make generatescripts` for the frontend app to compile. 

## Build locally
Run `make buildfrontend`. This command will create the runnable `frontend.exe` in the project root directory.

## Run locally
Run `make runfrontend`. This command will build the frontend and then immediately execute it.

## Generate scripts
Once changes to the `/frontend/scripts/scripts.go` are made, the JS files will need to be regenerated. Run `make generatescripts` to generate these.

## Run in Docker
Ensure that your variables are exported as detailed in [Environment variables](#environment-variables).
```
$ cd frontend
$ docker build -f Dockerfile -t frontend --build-arg BACKEND_URL=$BACKEND_URL .
$ docker run -dt -p 4321:4321/tcp frontend
```

> NOTE: Use `--platform linux/amd64` in the docker build command when pushing to Azure.