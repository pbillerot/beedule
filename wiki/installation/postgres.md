
## POSTGRES une base de données open source

https://www.postgresqltutorial.com/postgresql-getting-started/

## dans Docker

`/volshare/docker/postgres/docker-compose.yaml`

```

# https://github.com/docker-library/docs/blob/master/postgres/README.md

services:
  postgres:
    image: postgres:latest
    container_name: postgres
    restart: unless-stopped
    user: 1000:1000
    ports:
    # - 127.0.0.1:5432:5432
    - 5432:5432
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
      PGDATA: /var/lib/postgresql/data/pgdata
      TZ: 'Europe/Paris'
      PGTZ: 'Europe/Paris'
    volumes:
    - "/volshare/data/backup:/backup"
    - "/volshare/persistent/postgres:/var/lib/postgresql/data/pgdata"
    networks:
    - docker_web

volumes:
  certs:

networks:
  docker_web:
    external: true
```

## Commande de backup

```bash
docker exec -i postgres /usr/bin/pg_dump -U user base >/volshare/data/backup/base.sql
```

## Commande de restauration

```bash
docker exec -i postrges /bin/bash -c "PGPASSWORD=password psql --username user base" < /volshare/data/backup/base.sql
```
