- authentication via [JWT](https://jwt.io/)
- routes mapping via [express-routes-mapper](https://github.com/aichbauer/express-routes-mapper)

## Table of Contents

- [Install & Use](#install-and-use)
- [Folder Structure](#folder-structure)
- [ API Docs ](#api-docs)

## Install and Use

Start by cloning this repository

```sh
# HTTPS
$ git clone https://github.com/wdhafin/eFishery-test.git
# checkout to fetch-api branch
$ git checkout fetch-api
```
then config the port in config/index.js

```sh
const config = {
  #config port
  port: '9999',
  privateRoutes,
};
```
then


```sh
# cd into project root
$ npm i
# start the api
$ npm start
```

## Folder Structure

This boilerplate has 4 main directories:

- api - for controllers, models, services, etc.
- config - for routes, database, etc.

## API Docs
[API Docs](https://documenter.getpostman.com/view/10259308/TVK76LEq)
