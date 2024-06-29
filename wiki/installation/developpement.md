## Développement du framework Beedule

Beedule est développé avec le langage Go en utilisant le framework Beego.

<a href="https://github.com/pbillerot/beedule" target="_blank">Beedule sur Github</a>

**Beego** est un framework écrit en Go (diminutif de Golang). Il permet de coder des sites Internet (tout comme Django sur Python, Ruby On Rails sur Ruby, etc…). 

#### Architecture MVC de l'application

- `conf` : le fichier de configuration
- `controllers` : les contrôleurs
- `models` : les modèles
- `routers`: le fichier de routage
- `static` : pour les fichiers de type css, images et javascript dans leur dossier respectif
- `test` : les fichiers de test
- `views` : les vues

![](../images/beego-mvc.jpg)

## Installation de golang

depuis la debian 12 c'est tout simple

```bash
sudo apt install golang
echo "export PATH=~/go/bin:/usr/local/go/bin:${PATH}" | sudo tee -a $HOME/.profile source
echo export GOPATH=~/go | sudo tee -a $HOME/.profile source
source $HOME/.profile
go version
```

## Installation de beego

- <a href="https://beego.wiki/docs/install/install/" target="_blank">Beego wiki</a>

## Installation de l'outil bee

- <a href="https://beego.wiki/docs/install/bee/" target="_blank">Beego bee</a>


