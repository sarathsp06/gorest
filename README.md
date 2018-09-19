# GoREST REST Servive [documentation](https://documenter.getpostman.com/view/2799423/RWaLvT6K)

A Simple REST application using golang,docker,mongo  

### Technology

- Go 1.11
- Database
    - MongoDB connection details:
        - host: `mongodb`
        - port: `27017`

## Instructions
0. Clone the repo
  ```
  go get -d github.com/sarathsp06/gorest
  cd $GOAPTH/src/github.com/sarathsp06/gorest
  ```
- Run `bash setup.sh` to setup the environment,once
    - creates a directory under current diretory to attach as log volume
- Run `bash start.sh` to run the  service in docker
- Vist [localhost:8080](http://localhost:8080) to see the details of the running instance
- Refer to the  [documentation](https://documenter.getpostman.com/view/2799423/RWaLvT6K) for detail on API
- Logs are available in "$PWD/logs"
    - access logs in `$PWD/logs/access.log`
    - application specific logs in `$PWD/logs/app.log`


**NOTE**: 
* Run `make` to see all the possible make commands
* Run `su -c "setenforce 0"` , if logs  directory is empty
* If page,pageNum invalid the service would assume default (20,1)
* I have added google key in config file the account is not paid, change the config in `config.json.prod` with a valid key to run properly

## Available Resources

| Name   | Method      | URL                    | Protected |
| ---    | ---         | ---                    | ---       |
| About  | `GET`       | `/`                    | ✘         |
| List   | `GET`       | `/orders`              | ✘         |
| Create | `POST`      | `/order`              | ✓         |
| Update | `PUT/PATCH` | `/order/{id}`         | ✓         |

### TODO
* [X] (API Doc)[https://documenter.getpostman.com/view/2799423/RWaLvT6K]
* [X] Pagination
* [X] Dockerize
* [X] Add delay metrics
* [ ] Mongo db optimization ,adding indexes etc
* [ ] **Re**-enable the basic auth

## References
* https://github.com/roblaszczak/go-cleanarch/tree/master/examples/valid-simple
