# Go Template

This is a port of the more useful routes from TFG's rails template.

# Installation

```
go get github.com/jjwebdev/go-template
```

# Usage

```
cd $GOPATH/src/github.com/jjwebdev/go-template
glide install
cd web/
yarn install
cd src/semantic/
gulp build
cd ../..
webpack
cd ..
go run main.go
```

Your server is listening on port: 8080
