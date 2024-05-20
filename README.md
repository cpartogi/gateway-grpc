# Application Documentation

## Requirement
1. Go version 1.20 above
2. Docker for unit test
3. Swag, for api documentation. Installation command : go install github.com/swaggo/swag/cmd/swag@latest


## Setup
1. Git clone this repository
2. Install application with command : 
```bash 
make init-app 
```
3. Copy appplication config with command : 
```bash 
cp config-example.toml config.toml
```
4. Edit `config.toml` file based on your own configuration.
5. Create protobuf folder with command : 
```bash 
mkdir pb
```
6. Generate Proto with command : 
```bash 
make proto-gen
```


## Run Application
1. Run with command : 
```bash 
make run
```
2. API url example : [POST] localhost:10001/user/register 
3. Open browser for api documentation : localhost:10001/swagger/