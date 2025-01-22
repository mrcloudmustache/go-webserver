# GO-WEBSEVER

## Initialize Go module
```
go mod init webserver
```

## Create binary for Ubuntu AWS instance
```
GOOS=linux GOARCH=amd64 go build -o webserver-linux-x86
```

## Create release on Github and upload binaries

## Download binary to local server
```
wget https://github.com/mrcloudmustache/go-webserver/releases/download/v.10/webserver-linux-x86
```

## Make binary executable
```
chmod +x webserver-linux-x86
```

## Run webserver
```
./webserver
```

## This is why I love Golang! Binaries...