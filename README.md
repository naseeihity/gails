# Gails

Gails inspired by Kails is a web demo build with Gin, Gorm, MySQL, Redis. The structure of this project and UI is based on
[Kails](https://github.com/embbnux/kails) (which is An amazing Web App like Ruby on Rails with Koa2, Webpack and Postgres).

The aim of this project is to show the basic architecture of the website application using Golang and for me to be learn how to be a backend engineer.

## Features
- MVC
- Frame: Gin
- Database: MySQL
- ORM: GORM
- Session: Redis based
- Password with bcrypt
- Middlewares: flashMessage, Cors, Auth
- Template: go-view
- Dependency management: Go modules

## Structrue

```
gails                                  
├─ app                                                  
│  ├─ controllers         //Router handlers     
│  ├─ helpers             //Utils
│  ├─ middlewares         //Middlewares for session,auth,flashMessage etc.    
│  ├─ models              //Data model and Querys
│  ├─ pkg                              
│  │  ├─ config           //Init config       
│  │  └─ redis            //Init Redis   
│  ├─ routes                           
│  │  ├─ groups           //Grouped routers       
│  │  └─ router.go                     
│  └─ views               //Templates  
│      └─assets              //Front-end static resources
├─ config                 //Static Config files             
│  ├─ app.ini             //Development             
│  └─ app_prod.ini        //Production             
├─ db                     //DB struct init             
│  └─ gails.sql                        
├─ log                    //Running log             
├─ test                   //Test cases           
└─ main.go                //Entrance             

```

## Develop local
1. install golang, mysql, redis
2. create database and init
    ```
    CREATE DATABASE gails_dev
    CREATE USER 'gails_dev'@'localhost' identified by 'gails';
    GRANT all privileges on gails_dev.* to gails_dev@localhost identified by 'gails';
    flush privileges;
    use gails_dev;
    source ./db/gails.sql
    ```
3. go mod config(in china) `export GOPROXY=http://mirrors.aliyun.com/goproxy/`
4. run on `http://localhost:8001`
   ```
   go run main.go
   // or use fresh for hot reload
   go get github.com/pilu/fresh
   fresh
   ```

## 传统部署过程
Online: http://118.25.22.201/
1. 配置环境，安装mysql、redis、nginx，配置和启动
2. 服务端编译 / 本地编译：`env GOOS=linux GOARCH=amd64 go build .`
3. 静态文件部署到服务器 `cofig，app/assets, app/view`
4. `sudo chomd 777 gail`
5. 使用tmux运行 `tmux new s gails`, `./gails`
6. 配置nginx

## 使用docker部署
`Dockerfile` and `docker-compose`
`wait-for`  控制在数据库启动后再启动web

### 问题
`x509: certificate signed by unknown authority`： http://www.honlyc.com/post/golang-x509-certificate-unknown-authority/

## TODO
  - [x] User signin / logout
  - [x] Article list and details
  - [x] FlashMesssage Alert 
  - [x] Add HTTPS Support
  - [x] Cache Support
  - [x] Deploy with Docker
  - [ ] CI
  - [ ] Markdwon Support
  - [ ] Validations
  - [ ] A new solution for separating front and back ends:React + NextJS(Server-side rendering) + Nginx + Gails(Only work as an API Server)

## PROBLEMS
1. 在mysql5.7下，TIMESTAMP 不能为NULL，（delete_at字段）用非null值建表才能成功。但是如果使用了 `gorm.Model`，delete_at字段就会影响使用gorm（First方法）进程查表。解决办法：1. 不使用gorm.Model。2. 配置mysql使其允许为null：`set @@explicit_defaults_for_timestamp=1;`

2. hacker-news 的api在国内被墙
## References
- [Kails](https://github.com/embbnux/kails)
- [singo](https://github.com/Gourouting/singo)