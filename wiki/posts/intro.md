---
title: "Introduction"
date: 2021-05-05
draft: false
categories:
tags:
- intro
cover: "/media/user-manual.jpg"
---
<!--more-->
{{< diaporama >}}
{{< image image="/media/beedule-apropos.png" position="droite" taille="m" >}}

**BEEDULE** est un framework de développement d'application WEB **et** un serveur d'application WEB

C'est un [CRUD](https://fr.wikipedia.org/wiki/CRUD) pour réaliser des opérations de base sur des données :
- **C**reate : créer
- **R**ead : lire
- **U**pdate : mettre à jour
- **D**elete : supprimer

Le moteur **Beedule** va réaliser des **INSERT**, **SELECT**, **UPDATE** et **DELETE** sur des bases de données **Mysql**, **Sqlite**, **Oracle** ou **Postgres** 

Le développeur d'application Beedule ne devra connaître qu'un seul langage, le langage [SQL](https://fr.wikipedia.org/wiki/Structured_Query_Language).  
Les opérations à réaliser, les listes et formulaires seront décrits dans des fichiers textes au format [YAML](https://fr.wikipedia.org/wiki/YAML) que nous appelerons le **dictionnaire**.

L'objectif du guide est de vous présenter :

- comment installer **Beedule** sur un serveur dans un container [DOCKER](https://fr.wikipedia.org/wiki/Docker_(logiciel))
- la structure du dictionnaire
- l'éditeur du dictionnaire

