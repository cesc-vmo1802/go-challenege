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
- I chose `feature base` approach to manage this project (you can see in `features` folder at root project)
- Inside the `features` folder. I already present follow by `N-Tier` architecture (`storage`, `usecase` and `transports`).
- The layer always applies rules: the `low layer` did not allow calls to a `high level`. And each layer just have only one responsibility (e.g: `storage` layer only take how to persist data from application to database)
- I did try to apply the `single responsibility` for `file level`. Normally, you can see one `struct`, one `func` take care only one responsibility. However, I want to apply it at `file level`. You can see that in `usecase` folder inside each feature. I think that makes it easier to maintain
