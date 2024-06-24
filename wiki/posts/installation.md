---
title: "Installation de Beedule"
date: 2021-01-30
draft: false
#categories:
tags:
- technique
cover: "/media/parametrage.jpg"
style: bee-doc
#menu:
#  page:
#    parent: technique
#    weight: 10
---
*Guide pour l'administrateur technique*
<!--more-->
{{< toc >}}
{{< diaporama >}}

Je vous propose d'installer une plateforme complète pour héberger notre application **Beedule**.

Nous utiliserons une VM (Machine Virtuelle) [DEBIAN 10](https://fr.wikipedia.org/wiki/Debian) avec le gestionnaire de conteneur [Docker](https://fr.wikipedia.org/wiki/Docker_(logiciel)) installé.

Ce document ne décrit pas l'installation d'une VM Debian et de Docker.  
Pour ma part je loue une **VPS Debian 10 Docker** chez l'hébergeur [OVH](https://www.ovhcloud.com/fr/vps/) (1 vCore 2 Go 20 Go 1 domaine.eu pour 46.16 €/an)

## Prérequis du système hôte

Système : `Debian Buster`
```shell
>uname -a
Linux vps-7d2d773f 4.19.0-16-cloud-amd64 #1 SMP Debian 4.19.181-1 (2021-03-19) x86_64 GNU/Linux
```
Gestionnaire de conteneur Docker :
```shell
>docker version
Client: Docker Engine - Community
 Version:           20.10.5
 API version:       1.41
 Go version:        go1.13.15
 ...
 
Server: Docker Engine - Community
 Engine:
  Version:          20.10.5
  API version:      1.41 (minimum version 1.12)
  Go version:       go1.13.15
  ...
```

```shell
>docker-compose version
docker-compose version 1.21.0, build unknown
docker-py version: 3.4.1
CPython version: 3.7.3
OpenSSL version: OpenSSL 1.1.1d  10 Sep 2019
```


## La plateforme Docker

{{< image image="/media/docker.png" >}}

Notre plateforme sera composée de 4 containers :

- [Caddy Server](https://caddyserver.com/docs/) le frontal web, c'est l'élément le plus important. Il sera chargé :  
   - de contrôler le trafic http (:80) et https (:443)
   - de renouveller le certificat lié au nom de domaine
   - de gérer les authentifications pour certaines [URI](https://fr.wikipedia.org/wiki/Uniform_Resource_Identifier)
   - de rediriger les flux vers les autres containers en fonction des URI
   - de journaliser les accès et les erreurs
- **Beedule** le container qui va servir l'application **Beedule**

Pour plus de confort, j'utilise  
- [Portainer](https://korben.info/portainer-io-un-outil-graphique-pour-gerer-vos-environnements-docker-en-toute-securite.html) pour gérer graphiquement l'environnement Docker  
- [Filebrowser](https://filebrowser.org/features) pour manipuler les fichiers du répertoire partagé (volshare)

Les 4 containers ont accès à la même ressource de fichiers `volshare` et les échanges entre **Caddy Server** et les autres containers se font à travers le réseau privé `web`. Ces containers ne sont pas  accessibles de l'extérieur.

La configuration de **Docker** se fait à travers le fichier `/volshare/docker/docker-compose.yaml`, 
**Caddy Server** via `/volshare/docker/caddy/caddyfile.conf`

Nous allons détailler tout cela ci-aprés.

## Volume partagé /volshare

`/volshare` est le répertoire partagé entre tous les containers.

Il aura la structure suivante :
```
/volshare
  /logs
    access.log access.0.log ... access.9.log
  /etc
    (les certificats du domaine)
  /bivouac
    (le site web Hugo administré par Victor)
  /filebrowser
    database.db
  /data (le répertoire des données à sauvegarder)
    /store
      (le répertoire des fichiers statiques servi par Caddy)
    /beedic
      le répertoire du dictionnaire de l'application beedule
  /docker (les fichiers de configuration des containers)
    docker-compose.yaml
    /caddy
      caddyfile.conf
    /filebrowser
      filebrowser.conf
    /beedule
      docker-compose.yaml
      dockerfile
      custom.conf
      build.sh
```

## Container Filebrowser

{{< image image="/media/filebrowser.gif" taille="m" position="droite" >}}

[Filebrowser](https://filebrowser.org/features) est un explorateur de fichiers en ligne.

Nous allons le régler pour explorer le répertoire `/volshare` du système hôte via l'url `/fb`

### /volshare/docker/docker-compose.yaml

```yaml
  filebrowser:
    image: filebrowser/filebrowser:latest
    container_name: filebrowser
    restart: unless-stopped
    volumes:
    - /volshare:/srv
    - /volshare/filebrowser/database.db:/database.db
    - ./filebrowser/filebrowser.json:/.filebrowser.json    
    networks:
    - web
```

### /volshare/docker/filebrowser/filebrowser.json

```json
{
  "port": 80,
  "baseURL": "/fb",
  "address": "",
  "log": "stdout",
  "database": "/database.db",
  "root": "/srv",
}
```

### /volshare/docker/caddy/caddyfile.conf

```shell
# filebrowser /fb
redir /fb /fb/
reverse_proxy /fb/* filebrowser:80
```

## Container Portainer

{{< image image="/media/portainer.png" taille="m" position="droite" >}}

[Portainer](https://korben.info/portainer-io-un-outil-graphique-pour-gerer-vos-environnements-docker-en-toute-securite.html) via son interface web permet de visualiser les ressources de la plateforme Docker, les containers, les images, les volumes, le réseau.

### /volshare/docker/docker-compose.yaml

```yaml
  portainer:
    image: portainer/portainer-ce
    container_name: portainer
    command: -H unix:///var/run/docker.sock
    restart: unless-stopped
    volumes:
    - /var/run/docker.sock:/var/run/docker.sock
    networks:
    - web
```

### /volshare/docker/caddy/caddyfile.conf

```shell
# portainer 
# on supprime le préfix /portainer après le routage
redir /portainer /portainer/
route /portainer/* {
    uri strip_prefix /portainer
    reverse_proxy portainer:9000
}
```

## Container Beedule

{{< image image="/media/beedule-apropos.png" position="droite" taille="m" >}}

Ce container va héberger l'application **Beedule** libre d'utilisation sur [Github](https://github.com/pbillerot/beedule).

**Beedule** a été écrit dans le langage [Golang](https://fr.wikipedia.org/wiki/Go_(langage)).

_À noter que les 4 containers et Docker lui-même ont été développés en Golang_

### /volshare/docker/beedule/docker-compose.yaml

```yaml
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
    - docker_web

volumes:
  certs:

networks:
  docker_web:
    external: true
```

### /volshare/docker/beedule/custom.conf

```ini
# db default ou Tables.dataSource Tables.driverName
# [alias]
# drivertype 1:Mysql 2:Sqlite 3:Oracle 4:Postgres

#debug = false # pour avoir les traces des requêtes sql
EnableXSRF = true
runmode = "production"

# répertoire du dictionnaire
dicodir = "/volshare/beedic/config"

# !!! Attention
# la section représente 
# - l'application et en même temps l'alias du connecteur à la base de données
# - la définition du connecteur à la base de données
# - le chemin du répertoire des fichiers

[admin]
drivertype = 2
datasource = "/volshare/data/beedule/beedule.sqlite"
drivername = "sqlite3"
datadir = "/volshare/data/beedule/admin" 

[picsou]
drivertype = 2
# datasource = ."/database/picsou.sqlite"
datasource = "/volshare/data/picsou/picsou.sqlite"
drivername = "sqlite3"
datadir = "/volshare/data/picsou"
```

### /volshare/docker/caddy/caddyfile.conf

```shell

# beedule
redir /bee /bee/
reverse_proxy /bee/* beedule:3945

```

## Image Beedule

Ci-après le script qui permet de construire l'image Beedule.

_À noter que ce container intègre le moteur **Hugo** et des utilitaires **Git** **SSH** et **Ftp** pour déployer les sites sur des serveurs externes_

### /volshare/docker/victor/dockerfile

```dockerfile
# ETAPE COMPILATION
# Utilisation de golang pour compiler le projet beedule
# Le GOPATH par défaut de cette image est /go.
FROM golang:1.15-alpine as goalpine

# Installation de GCC et GIT
RUN apk add build-base git

# Installation de beedule
WORKDIR /src
RUN git clone https://github.com/pbillerot/beedule.git
WORKDIR /src/beedule

# Build avec CGO
RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /src/beedule/beedule

# ETAPE GENERATION D'UNE IMAGE avec seulement le module compilé
FROM alpine
RUN mkdir -p /src/beedule
copy --from=goalpine /src/beedule /src/beedule
RUN apk add --update nano

# Pont d'entrée
WORKDIR /src/beedule
ENTRYPOINT ./beedule

# Le port sur lequel notre serveur écoute
EXPOSE 3945
```

### /volshare/docker/victor/build.sh

```shell
# Construction de l'image
docker-compose up -d --build 
docker image prune -f
```

## Container Caddy

### /volshare/docker/docker-compose.yaml

Version complète

_À seul ce container présente des ports accessible de l'extérieur (ports 80 et 443)_

```yaml
version: "3.3"
services:

  caddy:
    # https://hub.docker.com/_/caddy?tab=description
    image: caddy:latest
    container_name: caddy
    restart: unless-stopped
    ports:
    - '80:80'
    - '443:443'
    volumes:
    - './caddy/caddyfile.conf:/etc/caddy/Caddyfile'
    - '/volshare/etc:/data'
    - '/volshare/data/store:/srv'
    - '/volshare:/volshare'
    networks:
    - web
    
    filebrowser:
    ...
    
    portainer:
    ...

volumes:
  certs:

networks:
  web:
    driver: bridge

```

### /volshare/docker/caddy/caddyfile.conf

Version complète

Dans ce fichier on trouvera :
- l'`email` qui sera utilisé pour demander le certificat SSL du domaine
- le nom du domaine du site
- éventuellement la liste de d'adresses IP indésirables
- les répertoires servis par l'explorateur web intégré à Caddy Server
- les fichiers logs 
- et enfin les différents `reverse_proxy` redirigés vers nos containers (qui ont été détaillés plus haut)

```shell
# Configuration du serveur Caddy
# https://caddyserver.com/docs/

# GLOBAL option
# https://www.ssllabs.com/ssltest/analyze.html?d=mon.domaine.com
{
    email mon.email@gmail.com
}

# HOST
mon.domaine.com

# blacklist - sites indésirables
@blaklist {
    remote_ip 94.130.212.180 134.119.20.10
}
handle @blaklist {
    respond "Refused!" 403
}

# Serveur de fichiers statiques
redir /store /store/
handle_path /store/* {
    root * /volshare/data/store
    file_server browse
}

# Log du trafic (rotation automatique tous les 100 Mo (10 logs) 
log {
    output file /volshare/log/access.log
    format single_field common_log
}

# filebrowser
# ...
# portainer
# ...
# beedule
# ...

```

## Procédure d'installation

```shell
cd /volshare/docker
# creation/mise à jour des containers sans reconstruction des images
docker-compose up -d
# nettoyage des images intermédiaires
docker image prune -f
```

## Démarrage de Beedule

{{< image image="/media/beedule-login.png" position="droite" taille="m" >}}

Dans votre navigateur préféré taper l'URL :

[https://mon.domaine.fr/bee](https://mon.domaine.fr/bee)

puis renseigner le code utilisateur et son mot de passe que vous avez configuré

et vous devriez avoir l'écran d'accueil de Beedule :

