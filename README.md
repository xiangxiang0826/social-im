gorm 模型生成命令

goctl model mysql datasource -url="root:root@tcp(192.168.99.88:3306)/social-im" -table="t_user" -dir="services/model" -cache=true --style=goZero --home ./template


api 生成命令（先进入到 desc 目录下）
goctl api go -api \*.api -dir ../ -style goZero

rpc 生成命令 （先进入到 pb 文件夹下）
goctl rpc protoc \*.proto --go_out=../ --go-grpc_out=../ -style goZero
goctl rpc protoc userRpc.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ -style goZero
goctl rpc protoc roomRpc.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ -style goZero  
