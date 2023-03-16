generatescripts: 
	echo ${BACKEND_URL} > ./scripts/config
	gopherjs build scripts/scripts.go -o scripts/scripts.js
buildfrontend:
	make generatescripts && go build -o frontend.exe main.go
runfrontend:
	make buildfrontend && ./frontend.exe
