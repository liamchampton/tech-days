generatescripts: 
	echo ${BACKEND_URL} > ./frontend/scripts/config
	gopherjs build frontend/scripts/scripts.go -o frontend/scripts/scripts.js
buildfrontend:
	make generatescripts && go build -o frontend.exe ./frontend/main.go
runfrontend:
	make buildfrontend && ./frontend.exe
