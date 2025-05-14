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

To run miggrate or create migration file do with makefile

```
make create name=create_users_table 
make migrate
make rollback
```

### Running
Run docker compose up then

```
go run main.go serve
```