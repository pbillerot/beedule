
## Le dictionnaire Chinook

![](./images/chinook-schema.png)

## application.yaml

```
# application CHINOOK
app-id: chinook
title: "Chinook, Magasin de multimédias"
image: "/bee/data/chinook/chinook.jpg"
group: chinook
menu: 
- table-id: Artist
  view-id: vall
  in-footer: true
- table-id: Album
  view-id: vall
  in-footer: true
- table-id: Track
  view-id: vall
  in-footer: true
- table-id: PlayList
  view-id: vall
- table-id: Customer
  view-id: vall
  in-footer: true
- table-id: Employee
  view-id: vall
- table-id: Invoice
  view-id: vall
- table-id: MediaType
  view-id: vall
- table-id: Genre
  view-id: vall
```

## Album.yaml

```
# Table des Albums
setting:
  alias-db: chinook
  key: AlbumId
  col-display: Title
  icon-name: "circle"

elements:
  AlbumId:
    type: counter
    label-long: "Id"
    label-short: "Id"
  Title:
    type: text
    label-long: "Album"
    label-short: "Album"
  ArtistId:
    type: list
    label-long: "Artiste"
    label-short: "Artiste"
    items-sql: "SELECT ArtistId as 'key', Name as 'label' From Artist order by Name"
  _image:
    type: "image"
    label-long: "Pochette"
    label-short: "Pochette"
    width: s
    params: 
      src: "/bee/data/chinook/chinook.jpg"
      url: "/bee/data/chinook/chinook.jpg"

views:
  vall:
    form-add: fedit
    form-edit: fedit
    form-view: fview
    deletable: false
    type: table
    group: chinook
    title: "Albums"
    icon-name: "circle"
    order-by: Title
    with-line-number: true
    card:
      header:
        - AlbumId
      meta:
      description:
        - Title
      extra:
    elements:
      AlbumId:
        order: 10
        hide: true
      Title:
        order: 20
      ArtistId:
        order: 30
        hide: false
  
  vartist:
    form-add: fedit
    form-edit: fedit
    form-view: fview
    deletable: true
    type: table
    group: chinook
    title: "Albums"
    icon-name: "circle"
    order-by: Title
    with-line-number: true
    elements:
      AlbumId:
        order: 10
      Title:
        order: 20

forms:
  fedit:
    title: "Album"
    groupe: chinook
    elements:
      AlbumId:
        order: 10
      Title:
        order: 20
      ArtistId:
        order: 30
  fview:
    title: "Album"
    groupe: chinook
    elements:
      AlbumId:
        order: 10
        hide: true
      Title:
        order: 20
      ArtistId:
        order: 30
      _card_pochette:
        type: card
        icon-name: "circle"
        label-long: "Pochette"
        order: 40
      _image:
        order: 45
      _vTrack:
        type: card
        order: 100
        label-long: Morceaux
        width: max
        icon-name: "file audio"
        args:
          AlbumId: "{AlbumId}"
        params:
          table: Track
          view: valbum
          where: "Track.AlbumId = '{AlbumId}'"
```

## Artist.yaml

```
# Table des Artists
setting:
  alias-db: chinook
  key: ArtistId
  col-display: Name
  icon-name: "user secret"

elements:
  ArtistId:
    type: number
    label-long: "Id"
    label-short: "Id"
  Name:
    type: text
    label-long: "Artiste"
    label-short: "Artiste"

views:
  vall:
    form-add: fedit
    form-edit: fedit
    form-view: fview
    deletable: false
    type: table
    group: chinook
    title: "Artistes"
    icon-name: "user secret"
    order-by: Name
    card:
      header:
        - ArtistId
      meta:
      description:
        - Name
      extra:
    elements:
      ArtistId:
        order: 10
      Name:
        order: 20

forms:
  fedit:
    title: "Artiste"
    groupe: chinook
    elements:
      ArtistId:
        order: 10
      Name:
        order: 20

  fview:
    title: "Artiste"
    groupe: chinook
    elements:
      ArtistId:
        order: 10
      Name:
        order: 20
      _image:
        order: 30
      _card_albums:
        type: card
        order: 100
        label-long: "Albums de {Name}"
        width: l
        args:
          ArtistId: "{ArtistId}"
        params:
          table: Album
          view: vartist
          where: "ArtistId = '{ArtistId}'"
```

## Customer.yaml
```
# Table des Clients
setting:
  alias-db: chinook
  key: CustomerId
  col-display: FirstName
  icon-name: "user"

elements:
  CustomerId:
    type: number
    label-long: "Id"
    label-short: "Id"
  FirstName:
    type: text
    label-long: "Nom"
    label-short: "Nom"
  LastName:
    type: text
    label-long: "Prénom"
    label-short: "Prénom"
  Company:
    type: text
    label-long: "Entreprise"
    label-short: "Entreprise"
  Address:
    type: text
    label-long: "Adresse"
    label-short: "Adresse"
  City:
    type: text
    label-long: "Ville"
    label-short: "Ville"
  PostalCode:
    type: text
    label-long: "Code postal"
    label-short: "Code postal"
  Phone:
    type: text
    label-long: "Téléphone"
    label-short: "Téléphone"
  Fax:
    type: text
    label-long: "Fax"
    label-short: "Fax"
  Email:
    type: email
    label-long: "Email"
    label-short: "Email"
  SupportRepId:
    type: list
    label-long: "Suivi par"
    label-short: "Suivi par"
    items-sql: "SELECT EmployeeId as 'key', LastName as 'label' From Employee order by LastName"

views:
  vall:
    form-add: fedit
    form-edit: fedit
    form-view: fview
    deletable: false
    type: card
    group: chinook
    title: "Clients"
    icon-name: "user"
    order-by: CustomerId
    card:
      header:
        - CustomerId
        - FirstName
        - LastName
      meta:
        - Company
      description:
        - Address
        - PostalCode
        - City
      extra:
        - Phone
        - Fax
        - Email
      footer:
        - SupportRepId
    elements:
      CustomerId:
      FirstName:
      LastName:
      Company:
      Address:
      City:
      PostalCode:
      Phone:
      Fax:
      Email:
      SupportRepId:

  vwallet:
    form-add: fedit
    form-edit: fedit
    form-view: fview
    deletable: false
    type: card
    group: chinook
    title: "Clients"
    icon-name: "user"
    order-by: CustomerId
    elements:
      CustomerId:
        order: 10
      FirstName:
        order: 20
      LastName:
        order: 30
      Company:
        order: 40

forms:
  fedit:
    title: "Client"
    groupe: chinook
    elements:
      CustomerId:
        order: 10
      FirstName:
        order: 20
      LastName:
        order: 30
      Company:
        order: 40
      Address:
        order: 50
      City:
        order: 60
      PostalCode:
        order: 70
      Phone:
        order: 80
      Fax:
        order: 90
      Email:
        order: 100
      SupportRepId:
        order: 110

  fview:
    title: "Client"
    groupe: chinook
    elements:
      CustomerId:
        order: 10
        hide: true
      FirstName:
        order: 20
      LastName:
        order: 30
      Company:
        order: 40
      Address:
        order: 50
      City:
        order: 60
      PostalCode:
        order: 70
      Phone:
        order: 80
      Fax:
        order: 90
      Email:
        order: 100
      SupportRepId:
        order: 110
      _factures:
        type: card
        order: 300
        label-long: "Factures"
        icon-name: "file invoice"
        args:
          CustomerId: "{CustomerId}"
        params:
          table: Invoice
          view: vclient
          where: "CustomerId = '{CustomerId}'"
```

## Employee.yaml
```
# Table des Clients
setting:
  alias-db: chinook
  key: EmployeeId
  col-display: EmployeeId
  icon-name: "user tie"

elements:
  EmployeeId:
    type: number
    label-long: "Id"
    label-short: "Id"
  FirstName:
    type: text
    label-long: "Prénom"
    label-short: "Prénom"
  LastName:
    type: text
    label-long: "Nom"
    label-short: "Nom"
  Name:
    type: text
    label-long: "Nom"
    label-short: "Nom"
    jointure:
      column: Employee.FirstName || ' ' || Employee.LastName 
  Title:
    type: text
    label-long: "Fonction"
    label-short: "Fonction"
  ReportsTo:
    type: list
    label-long: "Responsable"
    label-short: "Responsable"
    items-sql: "SELECT EmployeeId as 'key', FirstName || ' ' || LastName as 'label' From Employee order by FirstName"
  BirthDate:
    type: date
    label-long: "Date de naissance"
    label-short: "Né(e)"
  HireDate:
    type: date
    label-long: "Date embauche"
    label-short: "Embauché(e) le"
  Address:
    type: text
    label-long: "Adresse"
    label-short: "Adresse"
  City:
    type: text
    label-long: "Ville"
    label-short: "Ville"
  State:
    type: text
    label-long: "État"
    label-short: "État"
  Country:
    type: text
    label-long: "Région"
    label-short: "Région"
  PostalCode:
    type: text
    label-long: "Code postal"
    label-short: "Code postal"
  Phone:
    type: text
    label-long: "Téléphone"
    label-short: "Téléphone"
  Fax:
    type: text
    label-long: "Fax"
    label-short: "Fax"
  Email:
    type: text
    label-long: "Email"
    label-short: "Email"

views:
  vall:
    form-add: fedit
    form-edit: fedit
    form-view: fview
    deletable: false
    type: card
    group: chinook
    title: "Employés"
    icon-name: "user tie"
    order-by: LastName
    card:
      header:
        - EmployeeId
        - FirstName
        - LastName
      meta:
        - Title
        - ReportsTo
      description:
      extra:
        - Phone
        - Email
    elements:
      EmployeeId:
        order: 10
      FirstName:
        order: 20
      LastName:
        order: 30
      Title:
        order: 40
      ReportsTo:
        order: 50
      BirthDate:
        order: 60
      HireDate:
        order: 70
      Address:
        order: 80
      City:
        order: 90
      State:
        order: 100
      Country:
        order: 110
      PostalCode:
        order: 120
      Phone:
        order: 130
      Fax:
        order: 140
      Email:
        order: 150

forms:
  fedit:
    title: "Employé(e)"
    groupe: chinook
    elements:
      EmployeeId:
        order: 10
      FirstName:
        order: 20
      LastName:
        order: 30
      Title:
        order: 40
      ReportsTo:
        order: 50
      BirthDate:
        order: 60
      HireDate:
        order: 70
      Address:
        order: 80
      City:
        order: 90
      State:
        order: 100
      Country:
        order: 110
      PostalCode:
        order: 120
      Phone:
        order: 130
      Fax:
        order: 140
      Email:
        order: 150
  fview:
    title: "Employé(e)"
    groupe: chinook
    elements:
      EmployeeId:
        order: 10
        hide: true
      Name:
        order: 20
      Title:
        order: 40
      ReportsTo:
        order: 50
      BirthDate:
        order: 60
      HireDate:
        order: 70
      _address:
        type: card
        order: 100
        icon-name: "address card"
      Address:
        order: 110
      City:
        order: 120
      State:
        order: 130
      Country:
        order: 140
      PostalCode:
        order: 150
      Phone:
        order: 160
      Fax:
        order: 170
      Email:
        order: 180
      _portefeuille:
        type: card
        label-long: "Clients en portefeuille"
        order: 200
        width: max
        icon-name: "wallet"
        params:
          table: Customer
          view: vwallet
          where: "SupportRepId = '{EmployeeId}'"
```
## Genre.yaml
```
# Table des Genres de musique
setting:
  alias-db: chinook
  key: GenreId
  col-display: Name
  icon-name: "music"

elements:
  GenreId:
    type: counter
    label-long: "Id"
    label-short: "Id"
  Name:
    type: text
    label-long: "Genre"
    label-short: "Genre"

views:
  vall:
    form-add: fedit
    form-edit: fedit
    deletable: false
    type: table
    group: chinook
    title: "Genres de musique"
    icon-name: "music"
    order-by: GenreId
    card:
      header:
        - GenreId
      meta:
      description:
        - Name
      extra:
    elements:
      GenreId:
        order: 10
      Name:
        order: 20

forms:
  fedit:
    title: "Genre de musique"
    groupe: chinook
    elements:
      GenreId:
        order: 10
      Name:
        order: 20
```

## Invoice.yaml
```
# Table des Factures 23 mai 2023
setting:
  alias-db: chinook
  key: InvoiceId
  col-display: InvoiceId
  icon-name: "file invoice"

elements:
  InvoiceId:
    type: number
    label-long: "Ref."
    label-short: "Ref."
  CustomerId:
    type: list
    label-long: "Client"
    label-short: "Client"
    items-sql: "SELECT CustomerId as 'key', FirstName || ' ' || LastName as 'label' From Customer order by LastName"
  InvoiceDate:
    type: date
    label-long: "Date"
    label-short: "Date"
  BillingAddress:
    type: text
    label-long: "Adresse facturation"
    label-short: "Adresse facturation"
  BillingCity:
    type: text
    label-long: "Ville"
    label-short: "Ville"
  BillingState:
    type: text
    label-long: "État"
    label-short: "État"
  BillingCountry:
    type: text
    label-long: "Région"
    label-short: "Région"
  BillingPostalCode:
    type: text
    label-long: "Code postal"
    label-short: "Code postal"
    width: s
  Total:
    type: amount
    label-long: "Montant Total de la facture"
    label-short: "Total"
    read-only: true

views:
  vall:
    form-add: fedit
    form-edit: fedit
    form-view: fview
    deletable: false
    type: table
    group: chinook
    title: "Factures"
    icon-name: "file invoice"
    elements:
      InvoiceId:
        order: 10
      CustomerId:
        order: 20
      InvoiceDate:
        order: 30
      BillingAddress:
        order: 40
      BillingCity:
        order: 50
      BillingState:
        order: 60
      BillingCountry:
        order: 70
      BillingPostalCode:
        order: 80
      Total:
        order: 90

  vclient:
    form-add: fedit
    form-edit: fedit
    form-view: fview
    deletable: false
    type: table
    group: chinook
    title: "Factures"
    icon-name: "file invoice"
    with-line-number: true
    elements:
      InvoiceId:
        order: 10
      InvoiceDate:
        order: 30
      Total:
        order: 90

forms:
  fedit:
    title: "Facture"
    groupe: chinook
    elements:
      InvoiceId:
        order: 10
        hide: true
      CustomerId:
        order: 20
      InvoiceDate:
        order: 30
      BillingAddress:
        order: 40
      BillingCity:
        order: 50
      BillingState:
        order: 60
      BillingCountry:
        order: 70
      BillingPostalCode:
        order: 80
      Total:
        order: 90

  fview:
    title: "Facture"
    groupe: chinook
    elements:
      InvoiceId:
        order: 10
        hide: true
      CustomerId:
        order: 20
      InvoiceDate:
        order: 30
      BillingAddress:
        order: 40
      BillingCity:
        order: 50
      BillingState:
        order: 60
      BillingCountry:
        order: 70
      BillingPostalCode:
        order: 80
      Total:
        order: 90
      _vinvoice:
        type: card
        label-long: "Titres achetés"
        order: 200
        width: max
        icon-name: "receipt"
        args:
          InvoiceId: "{InvoiceId}"
        params:
          table: InvoiceLine
          view: vinvoice
          where: "InvoiceId = '{InvoiceId}'"
```

## InvoiceLine.yaml
```
# Table des Factures
setting:
  alias-db: chinook
  key: InvoiceLineId
  col-display: InvoiceLineId
  icon-name: "receipt"

elements:
  InvoiceLineId:
    type: counter
    label-long: "Id"
    label-short: "Id"
  InvoiceId:
    type: number
    label-long: "Facture"
    label-short: "Facture"
  TrackId:
    type: number
    label-long: "Morceau"
    label-short: "Morceau"
  UnitPrice:
    type: amount
    label-long: "P.U."
    label-short: "P.U."
  Quantity:
    type: number
    label-long: "Quantité"
    label-short: "Quantité"
  TrackName:
    type: text
    label-short: "Titre"
    jointure: 
      join: "left outer join Track on Track.TrackId = InvoiceLine.TrackId left outer join Album on Album.AlbumId = Track.AlbumId left outer join Artist on Artist.ArtistId = Album.ArtistId"
      column: Track.Name
  AlbumTitle:
    type: text
    label-short: "Album"
    jointure: 
      column: Album.Title
  ArtistName:
    type: text
    label-short: "Artiste"
    jointure: 
      column: Artist.Name    

views:
  vall:
    form-add: fedit
    form-edit: fedit
    deletable: false
    type: table
    group: chinook
    title: "Produits"
    icon-name: "receipt"
    elements:
      InvoiceLineId:
        order: 10
      InvoiceId:
        order: 10
      TrackId:
        order: 20
      UnitPrice:
        order: 30
      Quantity:
        order: 40

  vinvoice:
    form-add: fedit
    form-edit: fedit
    deletable: false
    type: table
    group: chinook
    title: "Produits"
    icon-name: "receipt"
    with-line-number: true
    elements:
      InvoiceLineId:
        order: 10
        hide: true
      InvoiceId:
        order: 20
        hide: true
      TrackId:
        order: 30
        hide: true
      ArtistName:
        order: 40
      AlbumTitle:
        order: 50
      TrackName:
        order: 60
      Quantity:
        order: 70
      UnitPrice:
        order: 80

forms:
  fedit:
    title: "Produit"
    groupe: chinook
    elements:
      InvoiceLineId:
        order: 10
      InvoiceId:
        order: 20
        protected: true
      TrackId:
        order: 30
      UnitPrice:
        order: 40
      Quantity:
        order: 50
```

## MediaType.yaml
```
# Table des Types de Média
setting:
  alias-db: chinook
  key: MediaTypeId
  col-display: Name
  icon-name: "audio description"

elements:
  MediaTypeId:
    type: number
    label-long: "Id"
    label-short: "Id"
  Name:
    type: text
    label-long: "Type de média"
    label-short: "Type de média"

views:
  vall:
    form-add: fedit
    form-edit: fedit
    deletable: false
    type: table
    group: chinook
    title: "Types de Media"
    icon-name: "audio description"
    order-by: MediaTypeId
    card:
      header:
        - MediaTypeId
      meta:
      description:
        - Name
      extra:
    elements:
      MediaTypeId:
        order: 10
      Name:
        order: 20

forms:
  fedit:
    title: "Type de Média"
    groupe: chinook
    elements:
      MediaTypeId:
        order: 10
      Name:
        order: 20
```

## PlayList.yaml
```
# Table des PlayLists
setting:
  alias-db: chinook
  key: PlaylistId
  col-display: Name
  icon-name: "align justify"

elements:
  PlaylistId:
    type: number
    label-long: "Id"
    label-short: "Id"
  Name:
    type: text
    label-long: "Playliste"
    label-short: "Playliste"
  TrackId:
    type: number
    label-long: "Morceaux"
    label-short: "Morceaux"

views:
  vall:
    form-add: fedit
    form-edit: fedit
    form-view: fview
    deletable: false
    type: table
    group: chinook
    title: "Playliste"
    icon-name: "align justify"
    order-by: Name
    card:
      header:
        - PlaylistId
      meta:
      description:
        - Name
      extra:
    elements:
      PlaylistId:
        order: 10
      Name:
        order: 20

forms:
  fedit:
    title: "PlayListe de musique"
    groupe: chinook
    elements:
      PlaylistId:
        order: 10
      Name:
        order: 20
  fview:
    title: "PlayListe"
    groupe: chinook
    elements:
      PlaylistId:
        order: 10
        hide: false
      Name:
        order: 20
        hide: false
      _tracks:
        type: card
        order : 100
        label-long: "Titres de la playliste : {Name}"
        width: max
        icon-name: "file audio"
        params:
          without-frame: false
          table: Track
          view: vplaylist
          where: "TrackId in (select TrackId from PlaylistTrack where PlaylistTrack.PlaylistId = '{PlaylistId}')"
```

## Track.yaml
```
# Table des Morceaux
setting:
  alias-db: chinook
  key: TrackId
  col-display: Name
  icon-name: "file audio"

elements:
  TrackId:
    type: counter
    label-long: "Id"
    label-short: "Id"
  Name:
    type: text
    label-long: "Titre"
    label-short: "Titre"
  AlbumId:
    type: list
    label-long: "Album"
    label-short: "Album"
    items-sql: "SELECT AlbumId as 'key', Title as 'label' From Album order by Title"
  MediaTypeId:
    type: list
    label-long: Type média"
    label-short: "Type média"
    width: m
    items-sql: "SELECT MediaTypeId as 'key', Name as 'label' From MediaType order by Name"
  GenreId:
    type: list
    label-long: "Genre de musique"
    label-short: "Genre"
    width: m
    items-sql: "SELECT GenreId as 'key', Name as 'label' From Genre order by Name"
  Composer:
    type: text
    label-long: "Compositeur"
    label-short: "Compositeur"
  Milliseconds:
    type: number
    label-long: "Durée en mn"
    label-short: "Durée"
    help: "Durée en millisecondes"
    format-sqlite: "select strftime('%M:%S', {Milliseconds}/1000, 'unixepoch')"
  Bytes:
    type: number
    label-long: "Taille"
    label-short: "Taille"
    help: "Taille en octets"
    format-sqlite: "select printf('%3.2f Mo', {Bytes}.00/1000000, 'unixepoch')"
    col-align: right
    col-nowrap: true
  UnitPrice:
    type: amount
    label-long: "P.U."
    label-short: "P.U."
    help: "Montant en €"
    col-nowrap: true
  ArtistName:
    type: text
    label-long: "Artiste"
    label-short: "Artiste"
    jointure:
      join: "left outer join Album on Album.AlbumId = Track.AlbumId left outer join Artist on Artist.ArtistId = Album.ArtistId"
      column: Artist.Name

views:
  vall:
    form-add: fedit
    form-edit: fedit
    form-view: fview
    deletable: true
    type: table
    group: chinook
    title: "Morceaux"
    icon-name: "file audio"
    order-by: Artist.Name, Track.AlbumId
    limit: 50
    elements:
      ArtistName:
        order: 1 
      AlbumId:
        order: 3
      TrackId:
        order: 10
        hide: true
      Name:
        order: 20
      MediaTypeId:
        order: 40
      GenreId:
        order: 50
      Composer:
        order: 60
      Composer:
        order: 60
      Milliseconds:
        order: 70
      Bytes:
        order: 80
      UnitPrice:
        order: 90

  vplaylist:
    type: table
    group: chinook
    title: "Morceaux"
    icon-name: "file audio"
    with-line-number: true
    elements:
      ArtistName:
        order: 1 
      AlbumId:
        order: 3
      TrackId:
        order: 10
        hide: true
      Name:
        order: 20
      GenreId:
        order: 50
      Milliseconds:
        order: 70

  valbum:
    form-add: fedit
    form-edit: fedit
    deletable: true
    type: table
    group: chinook
    title: "Morceaux"
    icon-name: "file audio"
    order-by: Name
    limit: 50
    with-line-number: true
    elements:
      TrackId:
        order: 10
        hide: true
      Name:
        order: 20
      MediaTypeId:
        order: 40
      GenreId:
        order: 50
      Composer:
        order: 60
      Composer:
        order: 60
      Milliseconds:
        order: 70
      Bytes:
        order: 80
      UnitPrice:
        order: 90

forms:
  fedit:
    title: "Morceau"
    groupe: chinook
    elements:
      TrackId:
        order: 10
        hide: true
      AlbumId:
        order: 20
        protected: true
      _group_title:
        type: card
        order: 100
      Name:
        order: 110
      Composer:
        order: 120
      _group_media:
        type: card
        order: 200
      MediaTypeId:
        order: 210
      GenreId:
        order: 220
      _group_data:
        type: card
        order: 300
      Milliseconds:
        order: 310
      Bytes:
        order: 320
      _group_last:
        type: card
        order: 400
      UnitPrice:
        order: 410

  fview:
    title: "Morceau"
    groupe: chinook
    elements:
      TrackId:
        order: 10
        hide: true
      Name:
        order: 20
      MediaTypeId:
        order: 40
      GenreId:
        order: 50
      Composer:
        order: 60
      Milliseconds:
        order: 70
      Bytes:
        order: 80
      UnitPrice:
        order: 90
      _card_Album:
        type: card
        order: 100
        label-long: Album
        icon-name: circle 
      AlbumId:
        order: 110
      _card_Artist:
        type: card
        order: 200
        label-long: Artiste
        icon-name: user secret 
      ArtistName:
        order: 210
```
