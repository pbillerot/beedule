## Les tables

Le fichier `<table>.yaml` est découpé en 4 sections, `setting`, `elements`, `views` et `forms`.

```
# BASE DE DONNÉES
setting:
  # connecteur à la base données
  # connecteur défini dans custom.conf du serveur par l'administrateur
  alias-db: admin
  # la clé de la table
  key: user_name
  # la colonne qui sera affichée dans les vues et formulaires pour identifier l'enregistrement
  col-display: user_name
  # l'icône par défaut des vues et formulaires dans le portail
  # Indiquer le nom de l'icône choisie dans la bibilothèque fomantic
  # https://fomantic-ui.com/elements/icon.html
  icon-name: "user"

# RUBRIQUES
# liste des rubriques (les colonnes) de la table qui seront utilisées dans les vues et formulaires
elements:
  user_name:
    type: text
    ...
  user_password:
    type: password
    ...

# VUES
# liste des vues qui seront utilisées par l'application
views:
  vall:
    ...
  vprofil:
    ...

# FORMULAIRES
# liste des formulaires qui seront utilisées par l'application
forms:
  fadd:
    ...
  fedit:
    ...

```
