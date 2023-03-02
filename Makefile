generatescripts: 
	gopherjs build frontend/scripts/scripts.go -o frontend/scripts/scripts.js
buildfrontend:
	go build -o frontend.exe ./frontend/main.go
runfrontend:
	make buildfrontend && ./frontend.exe
