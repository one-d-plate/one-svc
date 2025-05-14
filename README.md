## Go Boilerplate

### How to install
Install require depedencies
```
go mod tidy
```

### Migration
Installation goose

```
go install github.com/pressly/goose/v3/cmd/goose@latest
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
source ~/.bashrc
```

### Running
Run docker compose up then

```
go run main.go serve
```