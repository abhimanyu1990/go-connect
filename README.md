# go-connect
A  discussion forum application written in go


**Swagger Integration **

1. Created files as docs/swagger.docs.go
2. To use vendor package : GO111MODULE=on go mod vendor
3.  GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger 
        i. above command will create directory go/bin in home directory
        ii. Set the path in your environment to access "swagger" command from anywhere
4. from home directory run command: swagger generate spec -o ./docs/swagger.yaml --scan-models
5. run command: swagger serve -F=swagger ./docs/swagger.yaml for serving swagger on url