# Forum (B1) / date.go

<img alt="Go" src="https://img.shields.io/badge/go-%2300ADD8.svg?&style=for-the-badge&logo=go&logoColor=white"/><img alt="JavaScript" src="https://img.shields.io/badge/javascript-%23323330.svg?&style=for-the-badge&logo=javascript&logoColor=%23F7DF1E"/>
<img alt="HTML5" src="https://img.shields.io/badge/html5-%23E34F26.svg?&style=for-the-badge&logo=html5&logoColor=white"/>
<img alt="CSS3" src="https://img.shields.io/badge/css3-%231572B6.svg?&style=for-the-badge&logo=css3&logoColor=white"/>

## Lancement du serveur :

- Telecharger le dossier .
- Aller a la racine du dossier .
- Lancer la commande ``go run date.go`` .
- Remplir les informations de MySQL pour le premier lancement :

| Demandé | Exemple de réponse attendu |
| ----------- | ----------- |
| Host| localhost ou adresse IP  ( ex : 192.168.1.1 ) |
| Port | Port de MySQL  ( ex : 3306 ) |
| DbName | Nom de la base de donnée à utiliser  ( ex : forum ) |
| UserName | Nom d'utilisateur MySQL pouvant crééer et gérer les tables  ( ex : root ) |
| Password | Mot de passe de l'utilisateur  ( ex : root1234 ) |
- Se rendre sur le site avec l'adresse ``http://localhost:8080`` ou ``http://(adresse IP):8080`` affiché dans le terminal pour les pc du même réseau.

## Contenu du projet :

| Dossier | Contenu |
| ----------- | ----------- |
| db| Initialisation et intéraction avec la BDD|
| handlers| Gestionnaire des différentes routes du site|
| db| Initialisation et intéraction avec la BDD|
| lib| Fonctions et types utilisé par le site et le terminal|
| static| Dossier contenant les styles et contenu de la page (CSS, JS, images...)|
| templates| Fichiers HTML de la page|
| date.go| Gestion des routes et lancement du serveur|

# Problèmes :

- La fonctionnalité de like n'existe pas (juste récupéré mais pas possible d'en ajouter) .
- La page de settings n'existe pas, ce qui renvoie une ``erreur 404`` .
- La gestion des erreurs existe seulement pour ``erreur 404``, ``erreur lors de l'inscription ou connexion`` ( les autres renvoie une page blanche avec l'erreur dessus ) .
