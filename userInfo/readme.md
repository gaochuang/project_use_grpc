# grpc使用
1.编程pb文件user.proto\
2.protoc user.proto --go_out=./protobuf --go-grpc_out=./protobuf --go-grpc_opt=require_unimplemented_servers=false 生成gprc代码 这个配置不加，会存在兼容问题，有个接口需要实现