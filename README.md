# Desafio Técnico: openlibrary_api

## API

Build:
`sudo docker build -t api:alexandre-kupac .`

Usage:
`sudo docker run -ti --rm -p 3000:3000 api:alexandre-kupac`

## Runner

Build:
`docker build -t runner:alexandre-kupac .`

Usage:
`sudo docker run -ti --rm --network="host" runner:alexandre-kupac "Os Sertões"`
