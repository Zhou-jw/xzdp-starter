go env -w GOPROXY=https://goproxy.cn
go install github.com/cloudwego/hertz/cmd/hz@latest
go get -u gorm.io/gorm