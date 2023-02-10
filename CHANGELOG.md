# CHANGELOG

Historique des modifications

### À venir :
- eddy portail.yaml ne pas présenter les autres fichiers 

- bug eddy si nom rubrique commence par le début d'une autre
- bug refresh cotation
- eddy: ctrl+h

- `fixed` 
- `added` 
- `deleted` 
- `changed` 

5.1.0 10 février 2023
- `changed` attention le nom d'une rubrique jointure ne doit pas être préfixé par un _

5.0.3 17 mars 2022
- `changed` dialogue de confirmation plus petite
- `changed` readme qui pointe sur "apprendred beedule"

5.0.2 14 mars 2022
- `changed` nouvell base chinook.sqlite
- `fixed` debug=true pris en compte

5.0.1 14 mars 2022
- `changed` items qui n'affichent plus la clef de l'item en saisie
- `changed` ergonomie page portail
- `changed` couleur bouton eddy
- `fixed` suppression d'un article qui ne se faisait plus

5.0.0 13 mars 2022
- `added` element: compute-sqlite class-sqlite default-sqlite format-sqlite hide-sqlite 
- `added` formulaire: check-sqlite
- `added` action: hide-sqlite
- `changed` couleur bouton supprimer enregistrer

4.4.0 10 mars 2022
- `fixed` correction 4.3.1 xsrf nom implémenté sur l'éditeur eddy
- `added` clavier numérique sur amount counter number float percent

4.3.1 10 mars 2022
- `fixed` xsrf nom implémenté sur l'éditeur eddy

4.3.0 10 mars 2022
- `fixed` annulation recherche perte de la loupe
- `fixed` correction action non prise en compte
- `fixed` nowrap pour percent float number counter date float time
- `changed` bouton + en haut dans la carte
- `changed` appui prolongé sur entête colonne pour remettre le tri initial

4.2.2 10 mars 2022
- `added` clic droit pour enlever un tri de colonne

4.2.1 9 mars 2022
- `fixed` input xsrf

4.2.0 9 mars 2022
- `added` recherche possible dans les vues des cartes
- `added` tri possible dans les vues des cartes

4.1.2 4 mars 2022
- `changed` cast(rub as text) ajouté dans la requête du moteur de recherche
- `changed` pas de log de la datasourec

4.1.1 3 mars 2022
- `added` Ajout lib postgres

4.1.0 3 mars 2022
- `added` Ajout driver postgres

4.0.0 2 mars 2022
- `changed` Désormais les paramètres et dictionnaire d'une application sont définis dans un seul répertoire
- `added` paramètre portail=chemin/portail.yaml dans app.conf ou custom.conf
- `added` changement majeur dans dictionnaire application.yaml dans le répertoire de l'application
- `fixed` correction enregistrement de portail.yaml dans l'application courante
- `changed` beedule.log seulement accessible dans eddy

3.9.2 27 février 2022
- `changed` log load dictionnaire
- `added` <Enter> envoie un submit du formulaire

3.9.1 24 février 2022
- `fixed` icon-name dans element et non plus dans element.params

3.9.0 23 février 2022
- `added` ajout fichier logs/beedule.log tournant sur 10 jours
- `added` visualiseur du log si développeur

3.8.5 23 février 2022
- `fixed` correction select d'un record avec une clé issue d'une jointure

3.8.4 21 février 2022
- `changed` réglage ergonomie

3.8.3 21 février 2022
- `fixed` arrangement titre des cards
- `changed` Libellé du partage d'application

3.8.2 13 février 2022
- `fixed` erreur appel eddy

3.8.1 13 février 2022
- `fixed` Menu Profil caché si anonymous

3.8.0 12 février 2022
- `added` Application partageable si shareable: true

3.7.3 9 février 2022
- `fixed` tag wrap dans card
- `added` bouton supprimer dans view
- `added` bouton edit dans view 
- `added` bouton partager l'application 

3.7.2 30 janvier 2022
- `fixed` vue dashbord pas totalement visible si infooter

3.7.1 30 janvier 2022
- `fixed` jointure en readonly
- `fixed` les champs readonly protected hide n'étaient pas enregistrés
- `fixed` update avec where préfixé par le nom de la table
- `changed` vue Title avec macro

3.7.0 30 janvier 2022
- `fixed` correction plusieurs vues dans un dashboard
- `fixed` correction input datetime en datetime-local
- `fixed` actions possible dans dashboard
- `added` sidebar dans footer si in-footer = true
- `changed` vue formulaire sans tri et sans recherche
- `changed` jquery 3.6.0
- `changed` fomantic-ui 2.8.8

3.6.1 9 janvier 2022
- `fixed` calcul class dans la liste card
- `changed` taille icône à large/big

3.6.0 8 janvier 2022
- `changed` icon-name rattachée à la rubrique section
- `changed` card view avec séparateur header
- `changed` type section en card

3.5.0 8 janvier 2022
- `changed` view.Mask en view.Card
- `changed` label valeur en nowrap dans les cards
- `added` view.Card.Footer
- `deleted` vue de type "image" supprimée

3.4.2 8 janvier 2022
- `added` options width dans les vues
- `added` _crud_card.html

3.4.1 7 janvier 2022
- `added` template DictCreate Get Set Unset Keys Values
- `changed` si même chart dans la vue le script ne sera chargé qu'une seule fois
- `fixed` suppression de la flèche retour sur page dashboard

3.4.0 4 janvier 2022
- `added` type epdf grace à labelmake

3.3.6 4 janvier 2022
- `fixed` width 1ère section était forcée à "s"
- `fixed` correction retour suite ouverture view à partir de dashboard
- `fixed` correction style section de type view 
- `changed` dashboard Titre de la section avec label-long
- `changed` titre sous le chart si label-long renseigné

3.3.5 3 janvier 2022
- `fixed` card dans view width non pris en en compte
- `added` type chart
- `added` vue dashboard
- `added` bouton ajouter dans la vue
- `added` app pluvio pour mettre en place la vue dashboard

3.3.4 30 décembre 2021
- `changed` image et pdf ne génèrent plus de saut de section dans form-view
- `deleted` type month week datetime

3.3.3 30 décembre 2021
- `changed` image et pdf ne génèrent plus de saut de section dans form-view
- `deleted` type markdown

3.3.2 30 décembre 2021
- `fixed` validator si en mise à jour
- `changed` dictionnaire école sous ./beedic
- `changed` tags avec le label
- `added` title image

3.3.1 29 décembre 2021
- `changed` dictionnaire école sous ./beedic
- `added` ajout validator de beego pour contrôler la saisie coté serveur

3.3.0 27 décembre 2021
- `deleted` type duration tel supprimés
- `added` édition possible des sources SQL dans l'éditeur
- `fixed` correction auncun en aucun

3.2.1 26 décembre 2021
- `fixed` correction liste des fichiers dans le menu de l'éditeur - liste globale au lieu de la liste des fichiers de l'application en cours
- `fixed` correction appel aide en ligne éditeur

3.2.0 de Noẽl 2021
- `changed` Grosses modifs de structures - le répertoire `dicodir` est à définir au niveau des applications

3.1.3 du 20 décembre 2021
----------------------
- `changed` combobox remplacé par list

3.1.2 du 18 décembre 2021
----------------------
- `changed` help not include in webapp

3.1.1 du 18 décembre 2021
----------------------
- `fixed` champ radio qui plantait l'affichage

3.1.0 du 17 décembre 2021
----------------------
- `changed` config sous volshare/data

2.2.10 du 14 juillet 2021
----------------------
- `changed` chargement de tous les fichiers dictionnaires présents dans dicodir et non plus seulement des tables déclarées dans TableView
- `added` via element.Args passage d'arguments au formulaire d'ajout - le champ sera protégé
- `added` bouton Ajout dans les vues table
- `fixed` correction retour formulaire d'ajout (utilisation du forward)
- `fixed` correction retour suppression d'un enregistrement (utilisation du forward)
- `fixed` correction order-by qui n'était pas pris en compte

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
