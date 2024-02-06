# Desafio Técnico: openlibrary_api

API written in Go using Negroni framework that searches strings in Open Library API with a client written in Python that makes the requests through the Go API and prints the result into the console. Both apps run inside a Docker container.

## API

This API receives requests from Runner at the endpoint `/search-by-title` and respond to it with a list of titles of documents that matches with the search string in any field of the document. It has a limit of 20000 results (see to-do list below).

### Build

`sudo docker build -t api:alexandre-kupac ./api`

### Usage

`sudo docker run -ti --rm -p 3000:3000 api:alexandre-kupac`

## Runner

A Python script to make requests to Go API, format the responses, and display them. It receives the search string as an argument.

### Build

`docker build -t runner:alexandre-kupac ./runner`

### Usage

`sudo docker run -ti --rm --network="host" runner:alexandre-kupac "Os Sertões"`

## TODO

- [ ] create pagination logic
