- authentication via [JWT](https://jwt.io/)
- routes mapping via [echo](https://github.com/labstack/echo)
- Database migration [goose](https://github.com/pressly/goose)
- File Watchers [air](https://github.com/cosmtrek/air)
- PosgreSQL Database [go-pg](https://github.com/go-pg/pg)

## Table of Contents

- [Install & Use](#install-and-use)
- [Folder Structure](#folder-structure)
- [ API Docs ](#api-docs)

## Install and Use

Start by cloning this repository

```sh
# HTTPS
$ git clone https://github.com/wdhafin/eFishery-test.git
```

install goose and run migration

```sh
# install goose
go get -u github.com/pressly/goose/cmd/goose
# run migration
cd db/migration
goose postgres "user=<user> password=<password> dbname=<dbname> sslmode=disable" up
```

install air using

```sh
go get -u github.com/cosmtrek/air
```

then config the database on config/main.yml

```sh
pg:
  host: "127.0.0.1"
  port: "5432"
  dbname: ""
  user: ""
  password: ""
  debug: false
```

then


```sh
# cd into project root
$ go mod vendor
# start the api
$ air
```

## Folder Structure

### `/internal`

Private application and library code. This is the code you don't want others importing in their applications or libraries. Note that this layout pattern is enforced by the Go compiler itself. See the Go 1.4 [`release notes`](https://golang.org/doc/go1.4#internalpackages) for more details. Note that you are not limited to the top level `internal` directory. You can have more than one `internal` directory at any level of your project tree.

You can optionally add a bit of extra structure to your internal packages to separate your shared and non-shared internal code. It's not required (especially for smaller projects), but it's nice to have visual clues showing the intended package use. Your actual application code can go in the `/internal/app` directory (e.g., `/internal/app/myapp`) and the code shared by those apps in the `/internal/pkg` directory (e.g., `/internal/pkg/myprivlib`).

### `/pkg`

Library code that's ok to use by external applications (e.g., `/pkg/mypubliclib`). Other projects will import these libraries expecting them to work, so think twice before you put something here :-) Note that the `internal` directory is a better way to ensure your private packages are not importable because it's enforced by Go. The `/pkg` directory is still a good way to explicitly communicate that the code in that directory is safe for use by others. The [`I'll take pkg over internal`](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/) blog post by Travis Jeffery provides a good overview of the `pkg` and `internal` directories and when it might make sense to use them.

It's also a way to group Go code in one place when your root directory contains lots of non-Go components and directories making it easier to run various Go tools (as mentioned in these talks: [`Best Practices for Industrial Programming`](https://www.youtube.com/watch?v=PTE4VJIdHPg) from GopherCon EU 2018, [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0) and [GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go](https://www.youtube.com/watch?v=3gQa1LWwuzk)).

See the [`/pkg`](pkg/README.md) directory if you want to see which popular Go repos use this project layout pattern. This is a common layout pattern, but it's not universally accepted and some in the Go community don't recommend it. 

It's ok not to use it if your app project is really small and where an extra level of nesting doesn't add much value (unless you really want to :-)). Think about it when it's getting big enough and your root directory gets pretty busy (especially if you have a lot of non-Go app components).

### `/configs`

Configuration file templates or default configs.

Put your `confd` or `consul-template` template files here.

### `/init`

System init (systemd, upstart, sysv) and process manager/supervisor (runit, supervisord) configs.

### `/entity`

Represent Database table or you can call this folder is Model

### `/module`

Main process of application
 - handler
 - usecase
 - Store
 - Interface

## API Docs
[API Docs](https://documenter.getpostman.com/view/10259308/TVK76LEq)