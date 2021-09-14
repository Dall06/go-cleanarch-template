# Tempalte
No previous versions, still working on this one to reach v1. 

API Service for pssword app mobilde

based on :
    https://github.com/bxcodec/go-clean-arch/tree/9e174b8b0bbdfbab69bc293bb2905b2bb622155c

    https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate


## Installation

NOT Required. it has autodeploy to docker

To test app is necessary to create a .env file following .en.example file located in root folder

## Usage

```golang
go run main.go
```

## Changes
-- No domain layer for DDD good practices

-- cmd package that contains files that runs the service

-- domain entity (user) is now inside in its own internal procedures

-- Jwt claims now inside the JWT file in middleware

-- src folder now called pkg

-- queries now located in the repository of the user

-- Test files ordered

## License
[MIT](https://choosealicense.com/licenses/mit/)