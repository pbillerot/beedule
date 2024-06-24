## Le dictionnaire

La programmation de l'application se fait à travers un dictionnaire.

Le dictionnaire est organisé sous la forme de 3 types de fichiers **yaml**.
- **portail.yaml** pour configurer la présentation du portail
- puis à raison d'un répertoire par application avec
  - **application.yaml** le fichier pour définir le logo et les menus
  - un **table.yaml** par table de l'application pour déclarer la **connexion** à la base de données, les **rubriques**, **vues** et **formulaires**
  - des fichiers divers : images scripts sql javascript
