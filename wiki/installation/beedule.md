
## Container Beedule

## Script de construction de l'image

`/volshare/docker/beedule/dockerfile`

```
# http://golang.io/fr/tutoriels/deployer-un-serveur-go-avec-docker/
# https://www.cloudreach.com/en/resources/blog/cts-build-golang-dockerfiles/

# ETAPE COMPILATION
# Utilisation de golang pour compiler le projet beedule
# Le GOPATH par défaut de cette image est /go.
FROM golang:1.17-alpine as goalpine

# Installation de GCC et GIT
RUN apk add build-base git

# Installation de beedule
WORKDIR /src
RUN git clone https://github.com/pbillerot/beedule.git
WORKDIR /src/beedule

# Build avec CGO
ENV CGO_CFLAGS="-g -O2 -Wno-return-local-addr"
RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /src/beedule/beedule

# ETAPE GENERATION D'UNE IMAGE avec seulement le module compilé
FROM alpine
RUN mkdir -p /src/beedule
copy --from=goalpine /src/beedule /src/beedule
RUN apk add --update nnn

# Pont d'entrée
WORKDIR /src/beedule
ENTRYPOINT ./beedule

# Le port sur lequel notre serveur écoute
EXPOSE 3945
```

## Construction du container

`/volshare/docker/beedule/docker-compose.yaml`

```
version: "3.3"
services:
  beedule:
    build:
      context: .
    image: beedule
    container_name: beedule
    restart: unless-stopped
    user: 1000:1000
    # ports:
    #   - "7601:3945"
    volumes:
    - /volshare:/volshare
    - ./custom.conf:/src/beedule/conf/custom.conf
    networks:
    - web

volumes:
  certs:

networks:
  web:
    external: true
```

## Configuration de Beedule

`/volshare/docker/beedule/custom.conf`

```
#
# custom.conf - PERSONNALISATION SUR LE SITE
#
#debug = false # pour avoir les traces des requêtes sql
EnableXSRF = true
runmode = "prod"
debug = true

# chemin de portail.yaml
portail = "/volshare/data/beedule/portail.yaml"

# chemin de beedule.log
logpath = "/volshare/logs/beedule/beedule.log"

# SECTION déclaration des accès aux base de données et des répertoires des fichiers de données

[admin]
datasource = "user=beedule password=<mdp> host=postgres dbname=beedule sslmode=disable"
drivername = "postgres"
datadir = "/volshare/data/beedule/admin/files" 

[pluvio]
datasource = "user=beedule password=<mdp> host=postgres dbname=beedule sslmode=disable"
drivername = "postgres"
datadir = "/volshare/data/beedule/pluvio/files"

[chinook]
datasource = "/volshare/data/beedule/chinook/db/chinook.sqlite"
drivername = "sqlite3"
datadir = "/volshare/data/beedule/chinook/files"

[jecompte]
datasource = "user=beedule password=<mdp> host=postgres dbname=beedule sslmode=disable"
drivername = "postgres"
datadir = "/volshare/data/beedule/jecompte/files"
```

## Commande de construction du container

```bash
docker-compose up -d --build 
# pour nettoyer toutes les images intermédiaires
docker image prune -f
```

## Déclaration de l'URL à servir par Caddy Server

`/volshare/docker/caddy/caddyfile.yaml`

```
# BEEDULE avec dans la zone principale
www.domain.com {
  ...
  rewrite /bee /bee/
  handle_path /bee/* {
    reverse_proxy BEEDULE:3945
  }
  ...
}
```

```
# BEEDULE avec une zone spéciale
beedule.domain.com {
  rewrite / /bee/
  reverse_proxy beedule:3945
  log {
    format console
    output file /volshare/logs/caddy/beedule.log
  }
}
```
