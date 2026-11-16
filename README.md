# Groupie Tracker

Petit site web en Go qui affiche la discographie de l'artiste JUL en exploitant l'API
Spotify. Les données sont récupérées côté serveur, et les filtres (année, type, feat) sont calculés en Go.

## Prérequis

- Go 1.26 ou supérieur
- Un compte Spotify (premium)
- Une application créée sur le [dashboard développeur Spotify](https://developer.spotify.com/dashboard/) pour obtenir un Client ID et un Client Secret

## Configuration des variables d'environnement

1. Copier `.env.example` en `.env` à la racine du projet.
2. Renseigner les deux valeurs récupérées sur le dashboard Spotify :

```
ClientID = "ton-client-id"
Secret = "ton-client-secret"
```

## Lancer le projet

```
go run .
```

Le serveur écoute sur [http://localhost:8080](http://localhost:8080). Au démarrage,
il récupère un token Spotify et l'intégralité des titres de JUL (albums + singles)
avant de servir les pages.

## Routes

| Méthode | Route | Description |
|---|---|---|
| GET | `/` | Liste tous les titres de JUL. Filtrable par `?year=`, `?type=single\|album` et `?feat=true\|false`|
| GET | `/tracks/{id}` | Détail d'un titre : nom, cover, album, artiste(s), date de sortie, lien vers Spotify |
| GET | `/notfound` | Page 404 personnalisée |
| GET | `/spotifyError` | Page d'erreur affichée quand Spotify est injoignable |

## Divers

- La collection Postman utilisée pour tester les endpoints Spotify se trouve à la
  racine : `API - Spotify.postman_collection.json`.
