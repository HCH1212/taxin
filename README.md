项目实践
1. 安装中间件
  - Redis
  - PostgreSql (需要安装向量插件)
  - jaeger
  所用到的中间件可以本地安装，也可以使用docker或者docker-compose
2. 业务逻辑
  - 使用uber的fx框架搭建grpc server
  - 创建proto文件user.proto，写一个 user service,user service中要写3个方法。
    - 用户注册    ( 需要实现幂等 )
    - 用户登录  （ 需要用到jwt，返回access_token ）
    - 获取用户信息 (需要登录用户才可以操作，只能获取自己的用户信息)
    - user表字段说明:
      - Id: 主键id
      - user_id: 用户分布式id
      - password: 用户密码
      - like: 喜好
      - like_embedding: 喜好的词嵌入向量值 (可以调用任意平台接口去获取embedding)
      - create_at: 创建时间
      - update_at: 更新时间
    - 用户表 user.sql 表结构文件，需要写在代码目录中。
    
  - 创建proto文件system.proto,写一个system service.包含一个方法，SendFile。
    -  读取一个本地文件（可以是音频，视频，文本）以流的形式返回。
3. 单元测试
  - 为SendFile的service写测试用例，可以只写一个断言为发送出的字节流内容和发送文件一致。 
4. 链路追踪
  - 对每一个请求使用jeager实现链路追踪，至少要在tag中记录用户id，方便通过user_id查询。
  -  需要有方法查看你的go执行性能（比如使用 pprof）
  -  写一个部署脚本，可以是shell，python，makefile。能通过脚本直接部署你的程序。（选作）

## 启动

### 1. 本地
- 安装相关中间件后
- 新建.env文件，添加环境变量TOKEN_SECRET和GO_ENV（dev），然后运行：
```
go run cmd/main.go
```
### 2. docker-compose
```
make docker-compose-up
```
### 客户端测试
```
go run test/client.go
```

## 一些可观测

### 1. 链路追踪
```
open http://localhost:16686/
```

### 2. 性能分析
makefile包中pprof相关

## 注意
由于Embedding部分使用的是ollama本地模型，没有本地模型需要先拉取，或修改代码调用api

