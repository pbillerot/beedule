## Container Caddy Server

[Caddy Server](https://caddyserver.com/docs/) le frontal web, c'est l'élément le plus important. Il sera chargé  
   - de contrôler le trafic http (:80) et https (:443)
   - de renouveller le certificat lié au nom de domaine
   - de gérer les authentifications pour certaines [URI](https://fr.wikipedia.org/wiki/Uniform_Resource_Identifier)
   - de rediriger les flux vers les autres containers en fonction des URI
   - de journaliser les accès et les erreurs

## dans Docker

`/volshare/docker/caddy/docker-compose.yaml`

```
#
# CADDY SERVER
#
version: "3.3"
services:
  caddy:
    # https://hub.docker.com/_/caddy?tab=description
    image: caddy:latest
    container_name: caddy
    restart: always
    ports:
      - 80:80
      - 443:443
      # - 7890:7890
    volumes:
    - ./caddy/caddyfile.conf:/etc/caddy/Caddyfile
    - /volshare/persistent/etc:/data
    - /volshare:/volshare
    networks:
    - web

volumes:
  certs:

networks:
  web:
    driver: bridge
```

## Config Caddy

`/volshare/docker/caddy/caddyfile.conf`

```
# CADDY SERVER

# Glogal option pour demande certificat
{
  email contact@domain.com
}
# Redirection sur votre domaine
http://domain.com, https://domain.com {
  redir https://www.domain.com
}
# définition zone
www.domain.com {
  # Serveur de fichiers statiques
  rewrite /store /store/
  handle_path /store/* {
    root * /volshare/data/store
    file_server browse
  }
  # Protéger une uri avec un mot de passe
  # Pour hasher un mot de passe :
  # docker exec -it caddy caddy hash-password
  basicauth /uri/* {
    <user> <hashed_password_base64>
    # exemple
    # admin JDJhJDE0JGthTkFWZzlwZG55aUdwbVo5RFBOZE9EOVMzUjhjSTI0SXJDV1lIWTVQdU42dmswcHlhN3dl
  }
}
```
