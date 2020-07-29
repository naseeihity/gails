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
│  ├─ assets              //Front-end static resources        
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

## TODO
  - [x] User signin / logout
  - [x] Article list and details
  - [x] FlashMesssage Alert 
  - [ ] Markdwon Support
  - [ ] Validations
  - [ ] Cache Support
  - [ ] Deploy with Docker
  - [ ] A new solution for separating front and back ends:React + NextJS(Server-side rendering) + Nginx + Gails(Only work as an API Server)

## References
- [Kails](https://github.com/embbnux/kails)
- [singo](https://github.com/Gourouting/singo)