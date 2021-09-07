**Fastest Mirror Finder**
==============
**Helps you to find the fastest mirror among [Debian mirrors](https://www.debian.org/mirror/list) and present it to you using REST API.**

I got the idea from [Hands-On RESTful Web Services with Go](https://www.packtpub.com/product/hands-on-restful-web-services-with-go-second-edition/9781838643577) book, but I totally changed the implementation, added [gorilla/mux](github.com/gorilla/mux) as router, added tests, added Redis as cache, added Swagger, Dockerized the whole project, Continuous Integration using Github Actions, ...

**Production Environment (Using Docker)**
------------
If you have docker and docker-compose installed, then simply you can run:

(Make sure you are in root directory of this project which is the folder that contains `main.go`)
```bash
docker-compose up
```
And that's it :) now you can serve [Swagger Doc](https://github.com/farbodahm/lets-go/blob/main/fastestMirrorFinder/docs/swagger.yaml) to see available APIs.

Tests (Inside Docker container)
------------
To run tests inside Docker container, run:
```bash
docker-compose run app go test -v ./...
```

APIs Document (Using Swagger)
------------
1) *to generate spec: (optional, needed after changing things in handlers or router)*
Make sure you have [go-swagger](https://goswagger.io/install.html) installed, then simply you can run:
```bash
swagger generate spec -o ./docs/swagger.yaml --scan-models . 
```
2) *to serve spec: (In other word, see available APIs)*
```bash
swagger serve ./docs/swagger.yaml
```

**Development Environment (Without Docker)**
------------
1) If you want to run the project without Docker, first thing first, you need to install [Redis](https://redis.io):
```bash
sudo apt install redis-server 
```

2) Install project's dependencies:
```bash
go mod download
```

3) Change `REDIS_ADDR`'s value to `localhost` in [./internal/api/database/redis.go](https://github.com/farbodahm/lets-go/blob/main/fastestMirrorFinder/internal/api/database/redis.go)

4) Build the project:
```bash
go build
```

5) Run the project:
```bash
./fastestMirrorFinder
```

Tests
------------
To run tests without Docker, run:
```bash
go test -v ./...
```

**Note:** If you don't want to run tests inside Docker container, you can simply comment `RUN apk add go gcc` in [Dockerfile](https://github.com/farbodahm/lets-go/blob/main/fastestMirrorFinder/Dockerfile#L22), to make the build inside Docker even faster.

Project's Structure
------------
Because I wrote this project on the first week of my journey on learning Golang, I tried to follow a famous and big project's structure as far as I can, so I chose [CockroachDB](https://github.com/cockroachdb/cockroach) and tried to dig deep into it; which was really helpful and formative; I *highly* recommend it to any one else who is currently reading this :)

How does it work?
------------
Really simple :)

1) For the first request, it will scrap [Debian's mirrors](https://www.debian.org/mirror/list), then extract related urls using Regex, then save it to Redis to use it as cache. ([mirrors package](https://github.com/farbodahm/lets-go/blob/main/fastestMirrorFinder/pkg/mirrors/fetch.go))
(Instead of using Regex for parsing scraped page, we could use golang.org/x/net/html)

2) Send a GET request to each URL and compute the latency while saving the one which has the lowest latency. ([fastest_mirror package](https://github.com/farbodahm/lets-go/blob/main/fastestMirrorFinder/pkg/fastest_mirror/find.go))

3) Present the result using REST APIs. ([api package](https://github.com/farbodahm/lets-go/tree/main/fastestMirrorFinder/internal/api))

Useful Links
------------
- https://www.ardanlabs.com/blog/2017/02/design-philosophy-on-packaging.html

- https://github.com/cockroachdb/cockroach

- https://blog.questionable.services/article/testing-http-handlers-go

- https://quii.gitbook.io/learn-go-with-tests/questions-and-answers/http-handlers-revisited

- https://schier.co/blog/a-simple-web-scraper-in-go

- https://gobyexample.com/regular-expressions

- https://github.com/go-redis/redis

- https://tutorialedge.net/golang/go-redis-tutorial/