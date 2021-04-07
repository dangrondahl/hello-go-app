# Hello world in Go

Super simple web server app made for demo purposes.
Heavy inspired by (https://blog.gopheracademy.com/advent-2017/kubernetes-ready-service/)[https://blog.gopheracademy.com/advent-2017/kubernetes-ready-service/]

## Usage

### Build

```bash
docker build . -t hello-go-app
```

### Run

```bash
docker run -p <some_available_port>:8080 hello-go-app
```
