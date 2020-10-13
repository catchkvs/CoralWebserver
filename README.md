# CoralWebserver
Coral Webserver provides simple way to host static website with basic login and contact us features


## Getting started
1. Clone the repo: `git clone https://github.com/catchkvs/CoralWebserver.git`
2. Go to the root : `cd CoralWebserver`
3. Run the build: `go build -o server pkg/main/server.go`
4. Run `./server`
### Using docker container
1. Build the docker image: `docker build -t coral-webserver:0.1 .`
2. Run the docker container: `docker run -p 3040:3040 coral-webserver:0.1`

