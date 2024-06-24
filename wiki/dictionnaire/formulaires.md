## Les formulaires

```
forms:
  fadd:
    title: "Fiche Compte"
    groupe: admin
    elements:
      user_name:
        order: 10
      user_email:
        order: 20
  fedit:
    title: "Fiche compte"
    group: admin
    elements:
      user_name:
        order: 10
      user_email:
        order: 20
  fview:
    title: "Fiche Compte"
    group: admin
    elements:
      user_name:
        order: 10
      user_email:
        order: 20
      _SECTION_MDP:
        order: 30
      user_is_admin:
        order: 50
      user_groupes:
        order: 60
      _pwd_change:
        order: 70
  fprofil:
    title: "Mon profil"
    group: owner
    elements:
      user_name:
        order: 10
      user_email:
        order: 20
      _SECTION_MDP:
        order: 30
      user_is_admin:
        order: 50
      user_groupes:
        order: 60
      _pwd_change:
        order: 70
  fedit_owner:
    title: "Fiche compte"
    group: owner
    elements:
      user_name:
        order: 10
        read-only: true
      user_email:
        order: 20
  fpwd:
    title: "Changer le mot de passe"
    group: owner
    elements:
      user_name:
        order: 10
        read-only: true
      _new_password:
        order: 20
      _confirm_password:
        order: 30
        # protected: true
    check-sqlite:
      - "select 'les mots de passe ne sont pas identiques' where '{_new_password}' <> '{_confirm_password}' "
    post-sql:
      - "update users set user_password = '{_new_password}' where user_name = '{user_name}'"
  fmdp:
    title: "Sécurité"
    group: admin
    elements:
      user_name:
        order: 10
        read-only: true
      user_is_admin:
        order: 30
      user_groupes:
        order: 40
```
## Exemple : formulaire de visualisation "fview"

![](../images/form-ex-view.png)

## Exemple : formulaire de visualisation "fpwd"

![](../images/form-ex-edit.png)
