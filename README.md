# BEEDULE

Framework de développement d'application WEB en Yaml

- Utilisation du framework MVC https://beego.me/
- Utilisation du CSS https://fomantic-ui.com/

Beedule est un CRUD dont les spécifications sont décrites dans un fichier Yaml

```yaml
# Table groups
setting:
  alias-db: admin
  key: group_id
  col-display: group_id
  icon-name: "users"

elements:
  group_id:
    type: text
    label-long: "Groupes"
    label-short: "Groupes"
  group_note:
    type: textarea
    label-long: "Note"
    label-short: "Note"

views:
  vall:
    form-add: fedit
    form-edit: fedit
    deletable: true
    group: admin
    title: "Groupes"
    icon-name: users
    order-by: group_id
    mask:
      header:
        - group_id
      meta:
      description:
        - group_note
      extra:
    elements:
      group_id:
      group_note:

forms:
  fedit:
    title: "Fiche Groupe"
    groupe: admin
    elements:
      group_id:
        order: 10
      group_note:
        order: 20
```