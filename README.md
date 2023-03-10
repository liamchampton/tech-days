# tech-days
This is the repository for the [​"​Microsoft Tech Days: Learn Go on Azure"​](https://msevents.microsoft.com/event?id=4027743502&ocid=AID3052689) event. 

## Frontend
The FE for this experimental application is built with Go & [GopherJS](https://github.com/gopherjs/gopherjs). GopherJS generates JavaScript from Go code and outputs it in the `/frontend/scripts` directory.

## Environment variables
The FE requires the URL for the backend. This should be exported as:
```
$ export BACKEND_URL=XXX
```
Once this is exported, run `make generatescripts` for the frontend app to compile. 

### Build
Run `make buildfrontend`. This command will create the runnable `frontend.exe` in the project root directory.

### Run
Run `make runfrontend`. This command will build the frontend and then immediately execute it.

### Generate scripts
Once changes to the `/frontend/scripts/scripts.go` are made, the JS files will need to be regenerated. Run `make generatescripts` to generate these.