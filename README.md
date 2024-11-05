# miniTikTok-ECommerce
E-Commerce platform for TikTok

Technical Framework of this e-commerce platform:

![technical framework](./img/technical%20framework.png)



```shell
mkdir hello_world

cd hello_world

go mod init github.com/charleschile/miniTikTok-ECommerce

go get -u github.com/cloudwego/hertz

touch main.go

go mod tidy
```





Hello Hertz:

访问 localhost:8888/hello即可

![hello hertz](./img/hello%20hertz.png)



IDL: interface description language

接口描述语言，通常用于远程调用rpc，两种不同操作系统件通信的桥梁

It is a language that lets a program or object written in one language communicate with another program written in an unknown language. IDLs are usually used to describe data types and interfaces in a language-independent way.



- standardization and consistency
- cross-language support
- version control & compatibility
- simplified development workflow


grpc默认使用Protocol Buffers (protobuf for simplify) 来进行序列化的

在 gRPC 中，服务和消息的定义通常在 .proto 文件中进行，通过该文件定义 RPC 服务接口和消息格式，gRPC 使用 protobuf 编译器（protoc）生成相应的客户端和服务器代码

.proto文件可能内容：

```proto
syntax = "proto3";

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloResponse);
}

message HelloRequest {
  string name = 1;
}

message HelloResponse {
  string message = 1;
}

```

Thrift是另外一种IDL，Apache Thrift开源




[cwgo开源代码生成框架](https://github.com/cloudwego/cwgo)提高开发者体验


thrift在go中生成代码依赖于[thriftgo框架](https://github.com/cloudwego/thriftgo)

[cwgo终端自动补全功能](https://www.cloudwego.io/docs/cwgo/tutorials/auto-completion/)

注意ProtoBuff的话也依赖protoc工具
