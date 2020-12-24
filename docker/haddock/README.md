# CONFIG DU SERVEUR HÉBERGEMENT

## Construction de l'image

  docker build -t haddock .

## Volume volshare

```bash
mkdir /volshare
mkdir /volshare/src
cd /volshare/src
git clone git@github.com:pbillerot/beedule.git
```


## Rotation des logs de TRAEFIK

Dans le fichier

    /etc/logrotate/traefik

recopier

```bash
/volshare/traefik/access.log {
  su root root
  daily
  rotate 7
  nocompress
  missingok
  create 644 root root
  postrotate
  docker ps -a | grep traefik | awk '{print $1}' | xargs docker restart
  endscript
}
```

## Git de Beedule

```bash
mkdir src
cd src
git clone git@github.com:pbillerot/beedule.git
```

## Customize bash shell

    cat ~/.bashrc
```bash
export HISTTIMEFORMAT="%d/%m/%y %T "
export PS1='\u@\h:\W \$ '
alias l='ls -CF'
alias la='ls -A'
alias ll='ls -alF'
alias ls='ls --color=auto'
source /etc/profile.d/bash_completion.sh
export PS1="\[\e[31m\][\[\e[m\]\[\e[38;5;172m\]\u\[\e[m\]@\[\e[38;5;153m\]\h\[\e[m\] \[\e[38;5;214m\]\W\[\e[m\]\[\e[31m\]]\[\e[m\]\\$ "```

