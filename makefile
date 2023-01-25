# Tried getting this to work for both servers but I couldn't get the npm to work
#CLIENT=$(GOPATH)\github.com\gabo0802\Go-AngularTest\client

dev:
#	cd $(CLIENT) npm run start
	go run api/newsfeeder/httpd/main.go
