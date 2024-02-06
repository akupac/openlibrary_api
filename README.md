# openlibrary_api

## API

Build:
`docker build -t api:alexandre-kupac .`

Usage:
`docker run -ti --rm -p 3000:3000 api:alexandre-kupac`

## Runner

Build:
`docker build -t api:alexandre-kupac .`

Usage:
`docker run -ti --rm -e API_PORT=3000 runner:alexandre-kupac "Lord of the Rings"`
