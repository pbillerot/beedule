## Les rubriques

Une rubrique représente une **colonne de la table** ou un **champ temporaire**.  
Un champ temporaire sera préfixé par un undescore `_pwd_change`  
Une rubrique peut être utilisée dans les libellés et scripts, le nom sera encadré alors par des accolades `{user_name}`  
Beedule propose aussi des macros `{$<macro>}` pour récupérer des données techniques [Toutes les macros...](macros.md)

Les propriétés des rubriques seront définies dans la section `elements:` mais aussi lors de l'utilisation de la rubrique dans une vue `views:` ou formulaire `forms:`.

Le comportement de la rubrique dans une vue ou formulaire est principalement lié à son **type** (text, password, email, checkbox...) [Tous les types...](types-rubrique.md)

En fonction du type de rubrique des options viendront compléter ses caractéristiques [Toutes les options...](options-rubrique.md)

## Exemple de rubriques

```
# la section elements: va énumérer les rubriques qui seront utilisées dans les vues et formulaires.
elements:
  user_name:
    type: text
    label-long: "Nom ou pseudo"
    label-short: "Nom ou pseudo"
    pattern: "[a-zA-Z0-9]+"
    required: true
  user_password:
    type: password
    label-long: "Mot de passe"
    pattern: "[a-zA-Z0-9_-]+"
    required: true
    min-length: 3
  user_email:
    type: email
    label-long: "Email"
    label-short: "Email"
    required: true
  user_is_admin:
    type: checkbox
    label-long: "Administrateur"
    label-short: "Administrateur"
    col-align: center
  user_groupes:
    type: tag
    label-long: "Groupes"
    label-short: "Groupes"
    col-align: center
    items-sql: "select group_id as key, group_id as label from groups order by group_id"
  _pwd_change:
    type: button
    label-long: "Changer le mot de passe..."
    params:
      url: "/bee/edit/admin/users/vall/fpwd/{user_name}"
  _SECTION_MDP:
    type: section
    label-long: "Sécurité"
    icon-name: lock
    params:
      form: fmdp
  _new_password:
    type: password
    label-long: "Nouveau mot de passe"
    pattern: "[a-zA-Z0-9_-]+"
    required: true
    min-length: 3
  _confirm_password:
    type: password
    label-long: "Confirmer le mot de passe"
    pattern: "[a-zA-Z0-9_-]+"
    required: true
    min-length: 3
```