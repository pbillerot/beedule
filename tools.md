# BEEDULE DEVELOPPEMENT

## Installation de GO sur ma Debian Buster
```console
cd ~/Téléchargements
wget https://golang.org/dl/go1.14.6.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz
```

## Ajout de modules 
### markdown
https://github.com/gomarkdown/markdown
https://github.com/russross/blackfriday
- go get -u github.com/gomarkdown/markdown

### Scheduler
- https://github.com/MichaelS11/go-scheduler
Pour traduire l'expression 
- https://github.com/bradymholt/cRonstrue

### dans .profil
```console
export PATH=$PATH:/usr/local/go/bin
export GOPATH=~/go
```

### Mémo de commandes (pour le débutant que je suis)
- go mod init beedule
- **go mod tidy** pour mettre à jour go.mod
- go get ./...
- go get -u

 - git tag v1.0.0
 - git push --tags origin

