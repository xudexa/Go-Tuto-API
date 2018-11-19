# Création d’un Web Service simpliste en GO
**Description :**
Web Service de gestion d’une Vidéothèque avec gestion des prêts.
Utilisation au maximum des package de base du Go à quelques exceptions.

**Le besoin :**
- Lister les films de la vidéothèque (avec filtre et tri)
- CRUD d’une vidéo (Create, Read, Update, Delete)
- Lister nos contacts
- CRUD d’un contact
- Gestion des prêts :
	- Prêt et Retour d’une vidéo
	- Connaitre le stock de notre vidéothèque

Le projet va s’orienter autour de ce besoin en se découpant en plusieurs exercices simples qui donneront forme à notre Web Service.
 
### 1.	Création d’un serveur web et de ses routes associées (utilisation du package gorilla/mux).

Pour rappel, d’un point vu purement conventionnel les méthodes suivantes, pour le CRUD, sont associées :
- PUT : Create
- POST : Update
- GET : Read
- DELETE : no coment
Avant de commencer il faut définir nos ressources (routes) et le fonctionnement de notre Web Service :
- Routes pour gérer les vidéos
- Routes pour gérer les contacts
- Routes pour gérer les prêts

#### Exercice 1 :
- Création d’un serveur web sur un port d’écoute en l’interrogeant, le programme doit nous retourner un message « Hello Ernesto ».

**Aides :**

> Le routeur est géré avec le routeur du package gorilla/mux.
> A l’instanciation du routeur StrictSlash() lorsqu’il est à true redirige les routes du types /contacts/ vers /contacts.

Pour définir une route nous utilisons plusieurs fonctions: 
- Methods() définit la méthode gérer par notre route(GET, POST, PUT, …). 
- Path() correspond à la route à laquelle on veut faire correspondre notre action. 
- Name() est le nom que l’on souhaite donner à notre route. 
- HandlerFunc() spécifie la fonction à laquelle cette route est liée.

#### Exercice 2 :
- Prérequis exercice précédent
- Récupération du/des paramètre(s) d’une demande à notre web service.
- Des routes un peu plus concrètes pour les 5 méthodes du CRUD qui retournerons un simple message différent avec le/les paramètre(s) passé(s) et le statut approprié.

**Aides :**
> La fonction mux.Vars(*http.Request) retourne un tableau associatif de chaine des paramètres passé en paramètre.

Les statuts valides du package net/http GO sont :
- GET : http.StatusOK
- POST : http.StatusAccepted ou http.StatusOKt
- PUT : http.StatusCreated
- DELETE : http.StatusOKt

#### Exercice 3 :
- Standardiser la réponse et la renvoyer au format JSON comme suit :
	- ReponseAPI : { ‘‘Meta’’ : {}, ‘’Data’’ : ‘’ ‘’ }
	- Meta { ‘’Objet’’ : ‘’…’’,’’Total’’ :0,’’Offset’’ :0,’’Count’’ :0}
	- Data : la chaine réponse

**Aide : **

En renvoyant un json en réponse, il faut penser à prévenir le navigateur à l’aide d’un header que la réponse sera formater Json :
> w.Header().Set("Content-Type", "application/json; charset=UTF-8")

#### Exercice 4 : 
- Prérequis exercice précédent
- Définir l’ensemble des routes du web service (pour le CRUD contacts et vidéos) avec une fonction distincte par route qui pour le moment renvoi juste un message standardisé.

### 2.	Interroger une Base de donnés

Dans cette étape nous allons créer un accès à un serveur SQL, pour pouvoir interroger une base de données.
Modéliser les contacts en BDD : ID, Nom, Prénom

#### Exercice 1 :
- Réaliser un une connexion sans package autre que celui de base et le driver MSSQL

**Aides :**

Driver utilisé pour se connecter au MSSQL : https://github.com/denisenkom/go-mssqldb
> Chaine de connexion :
> sqlserver://user:password@instance:port?database=base&encrypt=disable&parseTime=true

*Nb : liste des drivers DB gérés : https://github.com/golang/go/wiki/SQLDrivers*

#### Exercice 2 :
- Réaliser une structure pour inscrire les informations de connexion
- Réaliser une fonction de connexion
- Réaliser une fonction pour lire la liste des enregistrements
- Réaliser une fonction pour lire un enregistrement

**Aides :**
```
type DatabaseSetting struct {
    SettingName  string 
    Type         string 
    DatabaseName string 
    HostName     string 
    Port         int    
    Driver       string 
    User         string 
    Password     string 
}
```

#### Exercice 3 :
- Ajouter / modifier / Supprimer des enregistrements
#### Exercice 4 :
- Prérequis exercice précédent.
- Modéliser la vidéothèque 
- Définir un ensemble de fonctions répondant au CRUD pour la vidéothèque
 
### 3.	Réaliser un package

Cette étape a pour but de comprendre le fonctionnement des packages, leur utilisation et la portée des variables, interface et fonctions.

#### Exercice 1 :
- Réalisation d’un package settings :
	- Définir une structure devant supporter le paramétrage.
- Réaliser une fonction qui alimente une variable locale au package.
- Réaliser une fonction qui renverra la variable et l’afficher.

**Aides :**
```
type ServeurSetting struct {
    Environnement struct {
        Name string `json:"name"`
        HTTP struct {
            ListenPort int    `json:"listenPort"`
            Path       string `json:"path"`
        } `json:"http"`
        ModeDebug     bool             `json:"modeDebug"`
        LevelLogger   int              `json:"levelLogger"`
    } `json:"environnement"`
}
```

#### Exercice 2 :
- Réaliser un fichier *.json qui reprendra les valeurs renseigné en dur. 
- Réaliser une fonction qui permet : 
	- D’ouvrir et récupérer le contenu du fichier json.
	- Sérialiser et alimenter le fichier json dans la variable

**Aides :**
	
	Le fichier settings.json sera à placer au niveau du main.go.	
```	
{
    "environnement":{
        "name" : "Dev",
        "http" : {
            "listenPort":8882,
            "path" : "/dev"
        },
        "modeDebug" : true,
        "levelLogger" : 5
    }
}
```

> Attention aux déclarations de l’import des packages.

#### Exercice 3 :
- Reprendre l’exercice 2 en ajoutant :
	- Plusieurs configurations d’environnement
	- En modifiant la fonction de chargement en passant en paramètre d’entrée le nom de la configuration à charger
 
### 4.	Packaging et assemblage des étapes précédentes pour réaliser notre web service

#### Exercice 1 :
- Prérequis :
	- Exercice 4 de l’étape 1
	- Exercice 4 de l’étape 3
- Packager l’ensembles des fonctions appelé par le router dans un package « controlers »
- Package la gestion des routes dans un package « routes »
- Refaire le lancement du web service sur le port d’écoute en paramètre

#### Exercice 2 :
- Prérequis exercice précédent.
- Packager les structures dans un package : « models »
- Packager les requêtes dans un package : « dbacces »
- Réaliser la connexion de la base de données avant le lancement du Web service avec les informations de connexion inscrites dans le paramétrage.
#### Exercice 3 :
- Prérequis exercice précédent.
- Alimenter vos fonctions CRUD du package « controllers » avant de renvoyer la réponse formatée.
#### Exercice 4 :
- Mettre en place les routes et la gestion des prêts / retour de vos vidéos.
