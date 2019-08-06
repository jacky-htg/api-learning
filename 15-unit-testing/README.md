# Rebel Services

Unit testing

Tasks:
- create docker-compose.yml to start and stop mysql container.
- create services/database/databasetest/docker.go to handle command start and stop mysql container using docker-compose.
- create file tests/tests.go to provide main services.
- models/models_test/user_test.go to handle unit test of user models.
- go test models/models_test/user_test.go 

## File Changes :


## New File :
- docker-compose.yml
- services/database/databasetest/docker.go
- tests/tests.go
- models/models_test/user_test.go

## Adding Dependency :
- github.com/google/go-cmp/cmp