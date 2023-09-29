build: 
	@ echo Api starting to build 
	@ go build -o  ./dist/backend.exe ./*.go 
	@ echo API built 

start: build 
	@ echo API starting 
	@ ./dist/backend.exe 
	@ echo API started