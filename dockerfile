# http://golang.io/fr/tutoriels/deployer-un-serveur-go-avec-docker/
# https://www.cloudreach.com/en/resources/blog/cts-build-golang-dockerfiles/

# Utilisation de golang comme image de base
# Le GOPATH par défaut de cette image est /go.
FROM golang:alpine

ENV APPNAME beedule

# Installation de GCC et GIT
RUN apk add -U --no-cache build-base git

# Copie des sources de notre projet
WORKDIR /src
COPY . .

# Fetch dependencies.
# Using go get.
RUN go get -d -v

# Installation de Hugo
WORKDIR /src
RUN git clone https://github.com/gohugoio/hugo.git
WORKDIR /src/hugo
RUN go install 


# Construction du binaire toujours à partir de /src
WORKDIR /src
RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/${APPNAME}

# Pont d'entrée
ENTRYPOINT /go/bin/${APPNAME}

# Le port sur lequel notre serveur écoute
EXPOSE 3945

# docker image build -t beedule:latest .
# docker container run --name beedule -p 3945:3945 -d beedule:latest
# http://localhost:3945
# docker container ps
# docker container stop beedule
# docker container start beedule
