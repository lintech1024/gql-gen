# gql

## 拉取 schema

安装 gqlfetch

```sh
go install github.com/suessflorian/gqlfetch/gqlfetch@latest
```

```cmd
mkdir schema
mkdir operations
```

拉取 schema

```cmd
gqlfetch --endpoint "http://192.168.1.204:8088/graphql" > schema/schema.gql
```

## 生成 query 和 mutation

```sh
gqlg --schemaFilePath schema/schema.gql --destDirPath operations --depthLimit 5
```

## 生成 Go 代码

```cmd
go mod init gql

go get github.com/Khan/genqlient

go run github.com/Khan/genqlient --init
```

修改 genqlient.yaml 配置

```yaml
schema: schema/schema.gql
operations:
- operations/*/*.gql
generated: gen/genqlient.go
package: gql
bindings:
  Long: 
    type: int64
```

创建 gen.go

```go
//go:generate go run github.com/Khan/genqlient genqlient.yaml
package gen
```

生成代码

```cmd
go generate
```
