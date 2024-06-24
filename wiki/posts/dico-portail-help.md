---
title: "Dictionnaire - Portail"
date: 2021-03-20
draft: false
categories:
tags:
- dictionnaire
cover: "/media/editeur.png"
style: bee-doc
#menu:
#  page:
#    weight: 30
---
<!--more-->
### Touches fonctions

{{< texte rouge >}}**ctrl+s**{{< /texte >}} : sauvegarde du document  
{{< texte rouge >}}**ctrl+/**{{< /texte >}} : mise en commentaire de la méta-données  
{{< texte rouge >}}**ctrl+espace**{{< /texte >}} : liste et complétion des mots-clés et rubriques  

```
title: "Beedule"
info: "Framework de développement WEB en Yaml"
icon-file: "/bee/static/img/beedule.png"
# Paramètres globaux
parameters:
  __amount_min: 1200 # mise minimum
  __cost: 0.0047 # coût en % d'une transaction
  __optimum: 0.035 # seuil minimum à atteindre en % pour vendre

# Liste des applications gérées
applications:
  admin:
    title: "Gestion des Utilisateurs"
    image: "/bee/static/img/tools.png"
    group: admin
    tables-views:
    - table-name: users
      view-name: vall
    - table-name: groups
      view-name: vall
  repas:
    title: "Repas Éleveurs"
    group: repas
    image: "/bee/data/repas/repas.jpg"
    tables-views: 
    - table-name: repas
      view-name: vall
```