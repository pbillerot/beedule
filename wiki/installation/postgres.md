
## POSTGRES une base de données open source

https://www.postgresqltutorial.com/postgresql-getting-started/

## dans Docker

`/volshare/docker/postgres/docker-compose.yaml`

```
#
# PORTAINER
#
version: "3.3"
services:
  postgres:
    image: postgres:latest
    container_name: postgres
    restart: unless-stopped
    user: 1000:1000
    # ports:
    #   - 5432:5432
    environment:
      POSTGRES_USER: beedule
      # export POSTGRES_PASSWORD=password
      POSTGRES_PASSWORD: $POSTGRES_PASSWORD
      TZ: 'Europe/Paris'
    volumes:
    - "/volshare/data/backup:/backup"
    - "/volshare/persistent/postgres:/var/lib/postgresql/data"
    networks:
    - web

volumes:
  certs:

networks:
  web:
    external: true
```

## Commande de backup

```bash
docker exec -i postgres /usr/bin/pg_dump -U beedule beedule >/volshare/data/backup/beedule.sql
```

## Commande de restauration

```bash
docker exec -i postrges /bin/bash -c "PGPASSWORD=${POSTGRES_PASSWORD} psql --username beedule beedule" < /volshare/data/backup/beedule.sql
```
