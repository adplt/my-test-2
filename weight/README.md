# Weight Application

This is a bare-bones example of a Weight application providing a REST API.

The entire application is contained within the `app` folder.

`main.go` is an index of the project.

`.env.example` is the property for `.env` file. All configuration using `.env`. In this file, there is 3 major env in one file (dev, test, and prod). For set which env do you use, run this one for the first time: ``

## Mac

    export ENV=dev

## Windows

    Set-Content -Path ENV -Value 'dev'

Note:
1. For windows, do this command on the `Windows PowerShell` application. Or you can setting this environtment on Windows Environtment.
2. For the env value:
    - Development: `dev`
    - Test: `test`
    - Production: `prod`

## Install / Sync Depencies

    go mod tidy

## Run the app

    go run main.go

then you can go to open web application by: `http://localhost:[set up port]/`

## Run the tests

    go test ./... -v

## Build the app

    go build




# REST API

The REST API to the example app is described below. For this one, there is a swagger url that explain all the API for this microservices. The swagger path is: `[url]/swagger/index.html`. To open swagger, make sure that .env property properly (sync with current env do you use):

## Development

    DEV__SWAGGER__ALLOW=TRUE
    DEV__SWAGGER__URL=[back end url]

## Test
    TEST__SWAGGER__ALLOW=TRUE
    TEST__SWAGGER__URL=[back end url]

## Production
    PROD__SWAGGER__ALLOW=TRUE
    PROD__SWAGGER__URL=[back end url]
