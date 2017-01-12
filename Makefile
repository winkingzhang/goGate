default: build

build:
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dist/goGate main.go

buildstatic:
    CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o dist/goGate main.go

builddocker: build
    docker build -t winkingzhang/gogate -f Dockerfile .

builddockeralpine: build
    docker build -t winkingzhang/gogate:alpine -f Dockerfile .

builddockerstatic: buildstatic
    docker build -t winkingzhang/gogate:static -f Dockerfile.static .