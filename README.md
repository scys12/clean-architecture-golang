# Clean Architecture Golang

![Go build](https://github.com/scys12/clean-architecture-golang/workflows/Go%20build/badge.svg)

## Description
This is an example of implementation of Clean Architecture from Uncle Bob in Go. This is a basic implementation of e-commerce

## Installation
1. First, you have to install docker. You can install docker [here](https://docs.docker.com/get-docker).

2. Then clone this repository.
```bash
git clone https://github.com/scys12/clean-architecture-golang
cd clean-architecture-golang
go mod download
```

3. Build the dependency
```bash
docker-compose up -d --build
```

4. Accessible from http://localhost:8080

## Tools Used
You can see all the tools/libraries used in this repository listed in [go.mod](https://github.com/scys12/clean-architecture-golang/blob/master/go.mod) file.


## Container Structure
```bash
├── carikom
├── carikom_redis
└── carikom_db
```

- Images
  - [go](https://hub.docker.com/_/golang):1.14-alpine
  - [redis](https://hub.docker.com/_/redis):lates
  - [postgreSQL](https://hub.docker.com/_/postgres):latest

## API
/categories
* ```GET``` : Get All Categories

/items
* ```GET``` : Get All Items

/items/category/{category_id}
* ```GET``` : Get All Items Based On Cateory

/latest
* ```GET``` : Get Ten Latest Items

/user/{user_id}/items
* ```GET``` : Get All User Items

/user/item
* ```GET``` : Get Item
* ```POST``` : Create Item
* ```PUT``` : Update Item
* ```DELETE``` : Update Item

/auth/signin
* ```POST``` : Login User

/auth/register
* ```POST``` : Register User

/user/profile
* ```PUT``` : Edit User Profile

/user/logout
* ```POST``` : Logout User

/user/profile/{username}
* ```GET``` : Get User Profile