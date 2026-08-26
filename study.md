## 构建 openai
```shell
#安装插件
go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest

# 生成 yaml 文件
protoc --proto_path=. --openapi_out=fq_schema_naming=true,default_response=false:. lesson.proto

linux 系统需要设置 GOPATH
go env -w GO111MODULE=on
```