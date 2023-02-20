buildscripts: 
	gopherjs build frontend/scripts/scripts.go -o frontend/scripts/scripts.js
runfrontend:
	go run frontend/main.go