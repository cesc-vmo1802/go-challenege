## How to run:
- Prerequiste: `docker`, `docker-compose`, `ubuntu:18.04+`
- Run command: `docker-compose up --build` to build and start server
- Access: `http://localhost:8887/api/v1/swagger/docs/index.html` to see api docs

## Technical decision:
- `gin-gonic` for http router
- `https://www.mongodb.com/docs/drivers/go/current/#mongodb-go-driver` for database driver. That help you persist data from application to the mongodb
- `github.com/dgrijalva/jwt-go` for `Authentication` module
- Support `multiple langagues` when response to client


## Architecture
- Project structure
```
├── app_context

│   └── app_context.go

├── cmd

│   ├── http

│   │   ├── http.go

│   │   └── v1

│   │       └── routes.go

│   ├── migration

│   │   └── migration.go

│   └── root.go

├── common

│   ├── app_error.go

│   ├── constants.go

│   ├── random_string.go

│   └── response.go

├── db.json

├── dbmigration

│   └── db.json

├── deployment

│   └── env

│       ├── env.docker.yaml

│       └── env.local.yaml

├── docker-compose.yaml

├── Dockerfile

├── docs

│   ├── docs.go

│   ├── swagger.json

│   └── swagger.yaml

├── features

│   ├── application

│   │   ├── domain

│   │   │   └── application.go

│   │   ├── dto

│   │   │   ├── create_application_request.go

│   │   │   ├── detail_application_response.go

│   │   │   ├── filter.go

│   │   │   ├── list_application_response.go

│   │   │   └── update_application_request.go

│   │   ├── storage

│   │   │   ├── create.go

│   │   │   ├── delete_by_id.go

│   │   │   ├── find_one_by_id.go

│   │   │   ├── find_one_by_name.go

│   │   │   ├── list.go

│   │   │   ├── storage.go

│   │   │   └── update_by_id.go

│   │   ├── transports

│   │   │   └── gin_app

│   │   │       ├── create_application.go

│   │   │       ├── delete_application.go

│   │   │       ├── detail_application.go

│   │   │       ├── list_application.go

│   │   │       └── update_application.go

│   │   └── usecase

│   │       ├── create_application.go

│   │       ├── create_application_test.go

│   │       ├── delete_application.go

│   │       ├── delete_application_test.go

│   │       ├── detail_application.go

│   │       ├── detail_application_test.go

│   │       ├── list_application.go

│   │       ├── list_application_test.go

│   │       ├── update_application.go

│   │       └── update_application_test.go

│   └── auth

│       ├── domain

│       │   └── user.go

│       ├── dto

│       │   ├── create_user_request.go

│       │   ├── login_user_request.go

│       │   └── login_user_response.go

│       ├── storage

│       │   ├── create.go

│       │   ├── find_by_id.go

│       │   ├── find_by_login_id.go

│       │   └── storage.go

│       ├── transports

│       │   └── gin_auth

│       │       ├── login.go

│       │       └── register.go

│       └── usecase

│           ├── login.go

│           └── register.go

├── go.mod

├── go.sum

├── main.go

├── messages

│   ├── message.en.json

│   └── message.vi.json

├── middlewares

│   └── authorize.go

├── pkg

│   ├── database

│   │   ├── base_model.go

│   │   ├── json_date.go

│   │   ├── json_time.go

│   │   ├── mgo

│   │   │   └── mongo.go

│   │   └── storage.go

│   ├── hash

│   │   ├── hasher.go

│   │   └── md5

│   │       └── md5.go

│   ├── httpserver

│   │   ├── middleware

│   │   │   └── recovery.go

│   │   └── router.go

│   ├── i18n

│   │   └── i18n.go

│   ├── logger

│   │   ├── common.go

│   │   ├── logger.go

│   │   ├── log_rotate.go

│   │   ├── log_sink.go

│   │   └── option.go

│   ├── paging

│   │   └── page.go

│   ├── tokenprovider

│   │   ├── jwt

│   │   │   └── jwt.go

│   │   ├── provider.go

│   │   └── token.go

│   └── utils

│       └── random

│           └── random.go

└── README.md
```
- I chose `feature base` approach to manage this project (you can see in `features` folder at root project)
- Inside the `features` folder. I already present follow by `N-Tier` architecture (`storage`, `usecase` and `transports`).
- The layer always applies rules: the `low layer` did not allow calls to a `high level`. And each layer just have only one responsibility (e.g: `storage` layer only take how to persist data from application to database)
- I did try to apply the `single responsibility` for `file level`. Normally, you can see one `struct`, one `func` take care only one responsibility. However, I want to apply it at `file level`. You can see that in `usecase` folder inside each feature. I think that makes it easier to maintain
- API docs is also supported
- Logging
- Error Handle
- Security