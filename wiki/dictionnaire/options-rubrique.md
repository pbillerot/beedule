## Options des rubriques

## args
```
# Pour passer des arguments au formulaire d'ajout
  _card_ecriture:
    type: card
    order: 300
    label-long: "Écritures du compte"
    icon-name: "list"
    width: xxl
    params:
      without-frame: true
      table: journal
      view: vcard
      where: "journal.compte = '{id}' and (journal.rapproche < 2 or journal.rapproche is null)" 
    args:
      compte: "{id}"
      _ecart: "{ecart}"
```
## class-sqlite
```
# Class CSS calculé en SQL utilisée dans les vues
class-sqlite: "select case when {ptf_gain} > 0 then 'green' when {ptf_gain} < 0 then 'red' end"
```
## col-align
```
# Pour aligné le contenu dans la cellule left | center | right
col-align: "left"
```
## col-nowrap
```
# Pour éviter la césure du contenu dans la cellule d'une vue
col-nowrap: true
```
## compute-sqlite
```
# Pour calculer la valeur de la rubrique dans un formulaire en édition
# La rubrique sera protégée
compute-sqlite: "select '{orders_buy}' * '{orders_quantity}' + '{orders_buy}' * '{orders_quantity}' * '{__cost}'"
```
## dataset
```
# Pour passer des paramètres nommés à un plugin javascript
dataset: 
  classjquery: "select 'bee-chart-quotes'"
  title: "select 'Cours deptf_id}'"
  quotes: "select open as matin, close as soir from quotes where id = '{ptf_id}' order by date"
  ...
```
## default
```
# Valeur par défaut
default: "buy"
```
## default-sqlite
```
# Valeur par défaut calculée en SQL
default-sqlite: "select datetime('now', 'localtime')"
```
## format-sqlite
```
# Mise en forme d'une valeur
format-sqlite: "select printf('%3.2f Mo', {Bytes}.00/1000000, 'unixepoch')"
...
format-sqlite: "select strftime('%M:%S', {Milliseconds}/1000, 'unixepoch')"
```
## grid
```
# Class pour donner la largeur du champ dans le formulaire "four wide field" 16 colonnes possibles
grid: "height wide field" # 2 colonnes de champs par ligne
```
## group
```
# Groupe autorisé à accéder à cette rubrique - Si "owner" c'est l'enregistreement qui ne sera pas visible
group: "admin"
```
## help
```
# Texte d'aide sur la rubrique
help: "Durée en millisecondes"
```
## hide
```
# Pour cacher le champ ou la colonne
hide: true
```
## hide-on-mobile
```
# La colonne sera cachée sur les mobiles (largeur écran <768px
hide-on-mobile: true
```
## hide-sqlite
```
# Pour cacher le champ ou la colonne via une requête SQL
hide-sqlite: "select 'ok' where ..."
```
## items
```
# Liste de clé/valeur pour un champ de type "list" ou "radio"
items:
  - key: "cheque"
    label: "Chèque"
  - key: "espece"
    label: "Espèce"
  - key: "virement"
    label: "Virement"
```
## items-sql
```
# Liste de clé/valeur pour un champ de type "list" ou "radio" via un requête SQL
items-sql: "select group_id as key, group_id as label from groups order by group_id"
```
## jointure
```
# colonne issue d'une jointure avec une ou d'autres tables
# ne pas oublier de préfixer la colonne par le nom de la table si des noms de colonne sont identiques dans les tables jointes
jointure:
  join: "left outer join ptf on ptf_id = id"
  column: "ptf.ptf_top"
...
jointure:
  # join: on ne renseignera pas le join si la jointure a déjà été faite par ailleurs
  column: "ptf.ptf_rem"
...
jointure:
  join: ""
  column: "ptf.date || '-' || ptf.id"
...
jointure:
  join: ""
  column: "((quotes.adjclose-quotes.close1)/quotes.close1)*100"
```
## label-long
```
# Label du champ dans un formulaire
label-long: "Montant net"
```
## label-short
```
# Label de la colonne dans une vue
label-short: "Net"
```
## max
```
# Valeur maximum du champ numérique
max: 1199.99
```
## max-length
```
# Longueur maximum du texte saisie dans le champ
max-length: 9
```
## min
```
# Valeur minimum du champ numérique
min: 2
```
## min-length
```
# Longueur minimum du texte saisie dans le champ
min-length: 2
```
## order
```
# N° d'ordre d'affichage du champ ou colonne dans un formulaire ou vue
order: 110
```
## params
```
# Ensemble de paramètres nommés complémentaires en fonction du type de rubrique
_action_sell:
  type: "button"
  group: "trader"
  label-long: "Vendre cette valeur..."
  params:
    url: "/bee/edit/picsou/orders/vachat/feditsell/{orders_id}?orders_order=sell&orders_sell={orders_quote}"
...
_image_day:
  type: image
  label-long: "Graph du jour"
  label-short: "Graph J"
  width: max
  icon-name: "emblem-photos"
  params:
    src: "/bee/data/picsou/png/day/{orders_ptf_id}.png"
    url: "/bee/data/picsou/png/day/{orders_ptf_id}.png"
```
## pattern
https://www.w3schools.com/tags/att_input_pattern.asp
```
# Expression régulière pour contrôler la saisie dans un champ
user_name:
  type: text
  label-long: "Nom ou pseudo"
  label-short: "Nom ou pseudo"
  pattern: "[a-zA-Z0-9]+"
  place-holder: "Pseudo constitué avec seulement des lettres et chiffres"
  required: true
user_password:
  type: password
  label-long: "Mot de passe"
  pattern: "[a-zA-Z0-9_-]+"
  required: true
  min-length: 3
```
## place-holder
```
# Texte indicatif qui s'efface dès que l'on active le champ du formulaire
place-holder: "Pseudo constitué avec seulement des lettres et chiffres"
```
## protected
```
# Champ généralement calculé
# Le champ ne pourra pas être mis à jour directement par l'utilisateur
protected: true
```
## read-only
```
# Champ en lecture seule.
read-only: true
```
## required
```
# Champ obligatoire en saisie (si bien sûr le champ est en mise jour)
required: true
```
## type
[Pour en savoir plus voir la page "type de rubriques"](/dictionnaire/types-rubrique/)
```
# Type de rubrique
type: amount button checkbox counter date datetime duration email float image list markdown month number pdf percent plugin section tag tel text time radio url week
```
## width
```
# Largeur du champ dans un formulaire : s m l xl max 
# 150px 360px 450px 600px 100%
width: s m l xl max
```