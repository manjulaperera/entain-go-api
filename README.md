## Entain BE Technical Test

This test has been designed to demonstrate your ability and understanding of technologies commonly used at Entain. 

Please treat the services provided as if they would live in a real-world environment.

### Directory Structure

- `api`: A basic REST gateway, forwarding requests onto service(s).
- `racing`: A very bare-bones racing service.

```
entain/
├─ api/
│  ├─ proto/
│  ├─ main.go
├─ racing/
│  ├─ db/
│  ├─ proto/
│  ├─ service/
│  ├─ main.go
├─ README.md
```

### Development Environment Setup For Windows Users

1. Setup Linux dev environment in Docker

```powershell
docker run --name ubuntu-dev -e HOST_IP=10.0.0.8 --expose 3000 -p 3000:3000 -p 9000:9000 -p 8000:8000 -v //E/Manjula/Tests/entain-go-api:/src -t -i tecnickcom/alldev /bin/bash
```

NOTE: You must replace the `HOST_IP` and the `Manjula/Tests/entain-go-api` path with the correct host ip and path to your source folder respectively

[See this YouTube video if you need detailed instructions](https://www.youtube.com/watch?v=CGCn0b4FOfs)

2. Once the docker container is setup try to attach to it using a new CLI window (in docker) or in a Terminal window in VS code.

List available docker containers and pick the correct container id to attach

```powershell
docker ps
```

Then attach the correct docker container 

```powershell
docker attach b7b9b0000d42
```

3. Install HomeBrew (Since it's not available in the `tecnickcom/alldev` docker image)

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
```

4. Run these two commands in your terminal to add Homebrew to your PATH

```bash
echo 'eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"' >> /root/.profile
eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"
```
	
5. Install the Homebrew dependencies if you have sudo access

```bash
sudo apt-get install build-essential
```

[See here for more information](https://docs.brew.sh/linux)
	
6. Install GCC. For this test we need a specific version of it

```bash
brew install gcc@5
```

### Getting Started

1. Install Go (latest).

```bash
brew install go
```

... or [see here](https://golang.org/doc/install).

2. Install `protoc`

```
brew install protobuf
```

... or [see here](https://grpc.io/docs/protoc-installation/).

2. In a terminal window, start our racing service...

```bash
cd ./racing

go build && ./racing
➜ INFO[0000] gRPC server listening on: localhost:9000
```

3. In another terminal window, start our api service...

```bash
cd ./api

go build && ./api
➜ INFO[0000] API server listening on: localhost:8000
```

4. Make a request for races... 

```bash
curl -X "POST" "http://localhost:8000/v1/list-races" \
     -H 'Content-Type: application/json' \
     -d $'{
  "filter": {}
}'
```

### Changes/Updates Required

Ideally, we'd like to see you push this repository up to Github/Gitlab/Bitbucket and lodge a Pull/Merge Request for each of the following tasks. This means, we'd end up with 5x PR's in total. Each PR should target the previous so they build on one-another. This will allow us to review your changes as best as we possibly can.

... and now to the test! Please complete the following tasks.

1. Add another filter to the existing RPC, so we can call `ListRaces` asking for races that are visible only.
   > We'd like to continue to be able to fetch all races regardless of their visibility, so try naming your filter as logically as possible. https://cloud.google.com/apis/design/standard_methods#list
2. We'd like to see the races returned, ordered by their `advertised_start_time`
   > Bonus points if you allow the consumer to specify an ORDER/SORT-BY they might be after. 
3. Our races require a new `status` field that is derived based on their `advertised_start_time`'s. The status is simply, `OPEN` or `CLOSED`. All races that have an `advertised_start_time` in the past should reflect `CLOSED`. 
   > There's a number of ways this could be implemented. Just have a go!
4. Introduce a new RPC, that allows us to fetch a single race by its ID.
   > This link here might help you on your way: https://cloud.google.com/apis/design/standard_methods#get
5. Create a `sports` service that for sake of simplicity, implements a similar API to racing. This sports API can be called `ListEvents`. We'll leave it up to you to determine what you might think a sports event is made up off, but it should at minimum have an `id`, a `name` and an `advertised_start_time`.

> Note: this should be a separate service, not bolted onto the existing racing service. At an extremely high-level, the diagram below attempts to provide a visual representation showing the separation of services needed and flow of requests.
> 
> ![](example.png)


**Don't forget:**

> Document and comment! Please make sure your work is appropriately documented/commented, so fellow developers know whats going on.

**Note:**

To aid in proto generation following any changes, you can run `go generate ./...` from `api` and `racing` directories.

Before you do so, please ensure you have the following installed. You can simply run the following command below in each of `api` and `racing` directories.

```
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 google.golang.org/genproto/googleapis/api google.golang.org/grpc/cmd/protoc-gen-go-grpc google.golang.org/protobuf/cmd/protoc-gen-go
```

### Good Reading

- [Protocol Buffers](https://developers.google.com/protocol-buffers)
- [Google API Design](https://cloud.google.com/apis/design)
- [Go Modules](https://golang.org/ref/mod)
- [Ubers Go Style Guide](https://github.com/uber-go/guide/blob/2910ce2e11d0e0cba2cece2c60ae45e3a984ffe5/style.md)

### Test Cases
Use following to test the changes mentioned above.

1. Get all race meets

```bash
curl -X "POST" "http://localhost:8000/v1/list-races" \
     -H 'Content-Type: application/json' \
     -d '{
  "filter": {}
}'
```

2. Get all race meets with meeting ids 2, 5 and 7

```bash
curl -X "POST" "http://localhost:8000/v1/list-races" \
     -H 'Content-Type: application/json' \
     -d '{
  "filter": {
	"meetingIds": [2, 5, 7]
  }
}'
```

3. Get all visible racing meets

```bash
curl -X "POST" "http://localhost:8000/v1/list-races" \
     -H 'Content-Type: application/json' \
     -d '{
  "filter": {
	"meetingVisibility": true
  }
}'
```

4. Get all invisible racing meets

```bash
curl -X "POST" "http://localhost:8000/v1/list-races" \
     -H 'Content-Type: application/json' \
     -d '{
  "filter": {
	"meetingVisibility": false
  }
}'
```

5. Get all invisible racing meets with meeting ids 5 and 7

```bash
curl -X "POST" "http://localhost:8000/v1/list-races" \
     -H 'Content-Type: application/json' \
     -d '{
  "filter": {
	"meetingIds": [5, 7],
	"meetingVisibility": false
  }
}'
```

6. Get all invisible racing meets with meeting ids 5 and 7 ordered by advertised start time in descending order

```bash
curl -X "POST" "http://localhost:8000/v1/list-races" \
     -H 'Content-Type: application/json' \
     -d '{
  "filter": {
	"meetingIds": [5, 7],
	"meetingVisibility": false
  },
  "order_by": {
	  "order_by_fields": [{
		"field": "advertised_start_time",
		"direction": 1
	  }]
  }
}'
```

7. Get all invisible racing meets with meeting ids 5 and 7 ordered by multiple fields. In this case by meet name and then by advertised start time in descending order

```bash
curl -X "POST" "http://localhost:8000/v1/list-races" \
     -H 'Content-Type: application/json' \
     -d '{
  "filter": {
	"meetingIds": [7, 9],
	"meetingVisibility": false
  },
  "order_by": {
	  "order_by_fields": [
	  {
		"field": "name",
		"direction": 1
	  },
	  {
		"field": "advertised_start_time",
		"direction": 1
	  }]
  }
}'
```

8. Return an error if an invalid column name is given for the order by expression

```bash
curl -X "POST" "http://localhost:8000/v1/list-races" \
     -H 'Content-Type: application/json' \
     -d '{
  "filter": {
	"meetingIds": [7, 9],
	"meetingVisibility": false
  },
  "order_by": {
	  "order_by_fields": [{
		"field": "bla_bla",
		"direction": 1
	  }]
  }
}'
```

9. Order by clause is optional

```bash
curl -X "POST" "http://localhost:8000/v1/list-races" \
     -H 'Content-Type: application/json' \
     -d '{
  "filter": {
	"meetingIds": [7, 9],
	"meetingVisibility": false
  }
}'
```

8. Specifying the order by direction is also optional. If it's not specified the order direction for that field will be ascending by default.

```bash
curl -X "POST" "http://localhost:8000/v1/list-races" \
     -H 'Content-Type: application/json' \
     -d '{
  "filter": {
	"meetingIds": [7, 9],
	"meetingVisibility": false
  },
  "order_by": {
	  "order_by_fields": [{
		"field": "name"
	  }]
  }
}'
```