# CHANGELOG

Historique des modifications

### À venir :
- filtres avec < > !

- `fixed`
- `added`
- `deleted`
- `changed`

5.6.1 - 30 juillet 2025
- `removed` suppression fonction de partage d'application
- `removed` option Application.Target

5.6.0 - 28 juillet 2025
- `added` icone-name dans Action
- `added` doc structure de dictionnaire

5.5.5 - 25 juin 2025
- `added` icone-about sur l'accueil et apropos

5.5.4 - 25 juin 2025
- `fixed` correction syntaxe golang

5.5.3 - 25 juin 2025
- `fixed` ajout icone et titre sur la page d'accueil

5.5.2 - 20 juin 2025
- `changed` icone accueil et aprops changée

5.5.1 - 25 mai 2025
- `changed` golang 1.24 go-sqlite3 v1.14.22
- `changed` fomantic-ui@2.9.4 jquery-3.7.1.min.js
- `changed` Doc installation sous Ubuntu 25.04
- `fixed` liste multiple qui affichait plus d'items qu'attendu

5.4.3 - 25 décembre 2024
- `fixed` recherche cases insensitive dans une vue

5.4.2 - 5 juillet 2024
- `added` path / routé sur /bee

5.4.1 - 25 juin 2024
- `fixed` lien wiki dans l'éditeur

5.4.1 - 25 juin 2024
- `added` url /wiki

5.4.0 - 24 juin 2024
- `changed` guide beedule wiki intégré dans la webapp /bee/wiki/beedoc

5.3.3 - 5 janvier 2024
- `fixed` prise en compte changement d'année dans les échéances

5.3.2 - 2 janvier 2024
- `changed` codemirror5 / codemirror
- `fixed` edyy-hint.js of codemirror move to eddy

5.3.1 - 2 janvier 2024
- `fixed` codemirror link .css

5.3.0 - 2 janvier 2024
- `changed` codemirror 5.61.1 vers 5.65.16

5.21.1 - 26 juillet 2023
- `changed` margin top sur button dans card

5.21.0 - 25 juillet 2023
- `added` application.wiki pour définir le répertoire du wiki de l'application

5.20.5 - 21 juillet 2023
- `fixed` style-sqlite accepte désormais plusieurs attributs

5.20.4 - 20 juillet 2023
- `fixed` vue dans form : fusion des where

5.20.3 - 30 juin 2023
- `changed` image dans liste table sans lien

5.20.2 - 30 juin 2023
- `changed` image dans liste table

5.20.1 - 29 juin 2023
- `changed` page login en card

5.19.3 - 25 juin 2023
- `added` affichage image mini dans liste table

5.19.2 - 25 juin 2023
- `changed` action-press sans affichage

5.19.1 - 24 juin 2023
- `changed` action-press disbled sur la ligne

5.19.0 - 24 juin 2023
- `added` action-press sur appui long sur article

5.18.0 - 16 juin 2023
- `added` appui long sur article actif

5.17.3 - 14 juin 2023
- `fixed` correction filtre sur colonne numérique

5.17.2 - 7 juin 2023
- `fixed` abandon appui long sur article

5.17.1 - 4 juin 2023
- `fixed` encadrement card selected qui ne faisait plus

5.17.0 - 3 juin 2023
- `added` appui long sur un article pour le sélectionner (encadrer) - à suivre...

5.16.2 - 2 juin 2023
- `fixed` actions dans une formview traitées maintenant

5.16.1 - 22 mai 2023
- `deleted` suppression de l'option element.PostAction

5.16.0 - 22 mai 2023
- `fixed` petites corrections suite rédaction de la doc

5.15.2 - 18 mai 2023
- `fixed` filtre avec recherche dans la liste et mémorisation saisie

5.15.1 - 17 mai 2023
- `changed` manifest.pwa.lson intégré dans header html

5.15.0 - 16 mai 2023
- `added` manifest.com pour intégration WPA (Progressive Web App)

5.14.8 - 14 mai 2023
- `changed` overflow-x visible

5.14.7 - 14 mai 2023
- `fixed` le bouton recherche n'était pas caché lors de l'activation de la recherche

5.14.6 - 13 mai 2023
- `added` batman ajout cret - coloriage si erreur

5.14.5 - 12 mai 2023
- `added` batman ajout date run
- `added` format type datetime

5.14.4 - 11 mai 2023
- `added` batman erreur sql et shell dans colonne result

5.14.3 - 11 mai 2023
- `changed` batman dans echeances

5.14.2 - 10 mai 2023
- `changed` filtres avec liste et tag
- `changed` default pris en compte si <> "" lors de l'enregistrement

5.14.1 - 10 mai 2023
- `fixed` coreection pour postgres

5.14.0 - 9 mai 2023
- `fixed` édition d'une date postgres
- `added` params.WithoutFrame pour ne pas afficher le cadre d'un élément card
- `added` params.WithoutFrame avec une vue de type table

5.13.0 - 7 mai 2023
- `deleted` vue de type smart supprimée
- `added` dans un formulaire hide hide-sqlite sur une card cache les éléments à l'intérieur

5.12.5 - 4 mai 2023
- `fixed` style-sqlite accepte désormais plusieurs attributs

5.12.4 - 1er mai 2023
- `added` interface url

5.12.3 - 1er mai 2023
- `added` couleur des cumul bas de tableau
- `fixed` exec shell par scheduler

5.12.2 - 1er mai 2023
- `fixed` CodeFormat qui ne laissait plus passer les valeurs sans format

5.12.1 - 1er mai 2023
- `fixed` codemirror python

5.12.0 - 30 avril 2023
- `added` with-sum dans élément et vue

5.11.1 - 30 avril 2023
- `fixed` correction appel su shell

5.11.0 - 29 avril 2023
- `added` action shell

5.10.1 - 29 avril 2023
- `added` filtres dans une vue

5.9.2 - 25 avril 2023
- `added` help en markdown
- `changed` retour arrière wide container

5.9.1 - 24 avril 2023
- `added` codemirror python

5.9.0 - 24 avril 2023
- `added` intégration de batman

5.8.6 - 20 avril 2023
- `changed` suppression class searchable

5.8.5 - 19 avril 2023
- `fixed` correction passage d'argument à un formulaire

5.8.4 - 19 avril 2023
- `changed` couleur carte

5.8.3 - 12 avril 2023
- `fixed` correction bouton plus derrière les cartes

5.8.2 - 11 avril 2023
- `changed` card enabled et couleur des données

5.8.1 - 11 avril 2023
- `fixed` card avec crud-list-selected est mainetent encadrée

5.8.0 - 10 avril 2023
- `added` card header-right -action meta-left right extra-left -right footer-left -right -action

5.7.6 - 4 avril 2023
- `changed` format des amount avec séparateur des milliers

5.7.6 - 3 avril 2023
- `fixed` titre des vues et formulaires qui n'étaient pas affichés

5.7.5 - 3 avril 2023
- `changed` card max -> 99%

5.7.4 - 2 avril 2023
- `changed` Formulaire: champ protégés formatés dans un autre input

5.7.3 - 1 avril 2023
- `fixed` correction champ input avec le format
- `changed` vue table sans scrolling - utiliser plutot une vvue smart

5.7.2 - 1 avril 2023
- `changed` vue table scrolling si débordement

5.7.1 - 27 mars 2023
- `added` editeur coloriage sql
- `changed` message toast centré en bas

5.7.0 - 25 mars 2023
- `added` planificateur de tâches via une table tasks id, name, day, month, last_day, last_month, disabled
- `added` option ajax-sql pour récupérer des données en ajax sur l'appui d'un bouton - dataset utilisé pour valoriser les variables de la requête
- `changed` un passage d'argument à un formulaire d'ajout ne protège plus la rubrique associée

5.6.2 - 18 mars 2023
- `fixed` affichage type "tag" dans une vue "table" et "smart"

5.6.1 - 15 mars 2023
- `changed` vue type smart adapatée au smartphone

5.6.0 - 13 mars 2023
- `added` vue type smart adapatée au smartphone

5.5.0 - 10 mars 2023
- `added` bouton crud-jquery-ajax

5.4.0 8 mars 2023
- `added` element.WithScript dataset sur type button

5.3.1 7 mars 2023
- `fixed` crud_table correction calcul de la valeur de la clé

5.3.0 7 mars 2023
- `added` view style-sqlite
- `added` element.option style-sqlite
- `added` type button dans une liste table Params.SQL
- `added` element.grid pour regrouper des champs sur la même ligne
- `changed` fomantic 2.9.2 jquery 3.6.3

5.2.1 21 février 2023
- `added` paramètre "params.target" si "params.url"

5.1.1 17 février 2023
- `changed` ajout texarea dans la doc

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
