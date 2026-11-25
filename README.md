# Vadik
> Simple AI agent implemented without AI :3

## Goals
- Get the purest agent, no bloating
- Learn how AI agents work

## TODO
- Rate limiting
- Static builds

## Security levels
There are two modes now:
- `0` (agent does everything without confirmation)
- `1` (agent asks before execute every command/writing file)

## How to run?
Just fill in `config.json` and execute:
```bash
go run main.go -config your_config.json
```
You can also run `vadik` in docker container:
```bash
sudo docker run --rm -it -v ./config.json:/app/config.json:ro vadik
```

## How to build?
Via docker:
```bash
sudo docker build -t vadik -f docker/Dockerfile.build .
```
