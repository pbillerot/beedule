# CHANGELOG

Historique des modifications

### À venir :
- aide sur l'application Application.Help url
- eddy: ctrl+h
- utiliser localStorage à la place des cookies

2.2.9 du 9 juin 2021
----------------------
- `removed` raz historique git
- `changed` chinook : toutes les tables visualisables avec jointures

2.2.8 du 6 juin 2021
----------------------
- `changed` Width maw-with: 100% pour mobile
- `changed` format appliqué dans le formulaire edit pour les rubriques qui ne sont pas en saisie
- `changed` help non affiché pour les champs disabled

2.2.7 du 5 juin 2021
----------------------
- `changed` Params.Width dans Element.Width
- `added` width par default par type de rubrique

2.2.6 du 4 juin 2021
----------------------
- `added` Params.Width list image
- `fixed` correction reprise édition suite sauvegarde

2.2.5 du 4 juin 2021
----------------------
- `added` Params.Width s m l xl max pour les sections image pdf et list
- `added` template function dict
- `fixed` correction reprise saisie après ctrl+s
- `changed` codemirror 5.58.1 -> 5.61.1

2.2.4 du 4 juin 2021
----------------------
- `fixed` datadir par aliasdb

2.2.3 du 4 juin 2021
----------------------
- `fixed` attribut jointure pour indiquer que la rubrique est le résulat d'une jointure

2.2.2 du 3 juin 2021
----------------------
- `added` col-nowrap
- `added` format-sql
- `added` help sur rubrique intégrée 
- `added` view: with-line-number
- `added` eddy: menu vertical avec liste des fichiers dicodir pour les ouvrir

2.2.1 du 2 juin 2021
----------------------
- `removed` jointure.join jointure.column 
- `added` Jointure.Join Jointure.Column 
- `changed` historique navigation pour gérer le retour arrière

2.2.0 du 31 mai 2021
----------------------
- `added` dictionnaire exemple avec base chinook
- `added` élement de type section avec params view table where
- `changed`vue table en sous-module pour réutilisation dans formView

2.1.0 du 24 mai 2021
----------------------
- `added` /bee/help -> aide en ligne
- `added` lien doc dans readme
- `changed` config chargé une seule fois dans main

2.0.13 du 23 mai 2021
----------------------
- `added` template crud_list*.html renommés 
- `fixed` correction edit de portail.yaml

2.0.11 du 23 mai 2021
----------------------
- `fixed` pointer actif sur liste de type table si formulaire view ou edit

2.0.10 du 23 mai 2021
----------------------
- `fixed` xsrfdata dans eddy.html

2.0.9 du 23 mai 2021
----------------------
- `added` fond d'écran

2.0.8 du 21 mai 2021
----------------------
- `removed` dico : params.path supprimé car en doublon avec params.src

2.0.7 du 21 mai 2021
----------------------
- `removed` suppression de config car externe à la webapp

2.0.6 du 20 mai 2021
----------------------
- `changed` répertoire du dictionnaire extérieure à la webapp
- `changed` répertoire du dictionnaire défini dans custom.conf

2.0.5 du 20 mai 2021
----------------------
- `changed` répertoire statique "datadir" défini dans custom.conf avec path: /bee/data/aliasdb
- `changed` ok url image sous /bee/data/aliasdb
- `added` dico: ajout gestion ptf
- `added` dico: ajout gestion orders

2.0.4 du 19 mai 2021
----------------------
- `changed` eddy : taille titre en h3 sur fond inversé en gris
- `added` eddy: la position du curseur sur le ctrl+s est mémorisée
- `added` eddy: ajout de la liste des rubriques _r dans l'auto-compléteur
- `added` eddy: coloriage des rubriques
- `fixed` correction suppression d'un article

2.0.3 du 18 mai 2021
----------------------
- `added` eddy : éditeur des fichiers de /config en ligne
- `added` eddy : auto-complétion ctrl+space dans eddy.js
- `added` eddy : erreurs rechargement du dictionnaire affichées dans la fenêtre

2.0.2 du 17 mai 2021
----------------------
- `added` eddy : éditeur des fichiers de /config en ligne

2.0.1 du 17 mai 2021
----------------------
- `fixed` retour profil sur view non autorisée
- `changed` icônes des view en bleu  dans le menu
- `fixed` bouton retour enlevé sur view vprofil

2.0.0 du 15 mai 2021
----------------------
- `changed` dictionnaire dans fichier Yaml sous /config
- `removed` suppression de la gestion de batch
- `removed` suppression de la gestion de tâches
- `removed` suppression de la gestion de paramètres en base - maintenant dans config/portail.yaml

1.7.4 du 15 mai 2021
----------------------
- `changed` fin version 1
- `changed` début déclaration dico dans fichier.yaml

1.7.3 du 12 mai 2021
----------------------
- `changed` go mod initand tidy
- `changed` set runmode="production" si docker

1.7.2 du 12 mai 2021
----------------------
- `changed` logs v2 - dataset non redéfini

1.7.0 du 11 mai 2021
----------------------
- `changed` passer en beego v2
- `added` ajout du changelod
- `removed` suppression gestion répertoire Hugo

1.6.6 du 6 avril 2021
----------------------
- `changed` graphique dans une vue

x.x.x 
----------------------
- historique non reportée dans ce document

###### Types de changements:
`added` *pour les nouvelles fonctionnalités.*  
`changed` *pour les changements aux fonctionnalités préexistantes.*  
`deprecated` *pour les fonctionnalités qui seront bientôt supprimées*.  
`removed` *pour les fonctionnalités désormais supprimées.*  
`fixed` *pour les corrections de bugs.*  
`security` *en cas de vulnérabilités.*  
