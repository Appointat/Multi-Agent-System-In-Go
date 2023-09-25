# Serveur de vote

|  Information |   Valeur              |
| :------------ | :------------- |
| **Auteurs** | Sylvain Lagrue ([sylvain.lagrue@utc.fr](mailto:sylvain.lagrue@utc.fr))|
| **Licence** | Creative Common [CC BY-SA 3.0](https://creativecommons.org/licenses/by-sa/3.0) |
| **Version document** | 1.2.1 |

## Travail demandé

1. Un service de vote ("serveur REST") implémentant différentes méthodes de vote
2. Des agents clients pouvant voter sur ce serveur
3. Un petit fichier `readme.md` succinct expliquant comment lancer votre travail (idéalement via `go install`, cf. infra) et résumant ce que vous avez réellement implémenté
4. Idéalement, votre application doit être installable via un `go install` et le gitlab de l'UTC
5. Vous pouvez bien entendu enrichir les fonctionnalités à votre gré (à renseigner dans le `readme.md`)

**Date de rendu Moodle :** le 07 novembre (23h59)

## API demandée du web service

On décrit ici l'API demandée afin de pouvoir lancer automatiquement une batterie de tests.

### Commande `/new_ballot`

- Requête : `POST`
- Objet `JSON` envoyé

| propriété  | type        | exemple de valeurs possibles                                      |
|------------|-------------|-------------------------------------------------------------------|
| `rule`      | `string`       | `"majority"`,`"borda"`, `"approval"`, `"stv"`, `"kemeny"`,... |
| `deadline`  | `string`       | `"Tue Nov 10 23:00:00 UTC 2009"`                               |
| `voter-ids` | `[string,...]` | `["ag_id1", "ag_id2", "ag_id3"]`                                       |
| `#alts`     | `int`          | `12` |   

*Remarques :* la deadline représente la date de fin de vote. Pour celle-ci, utiliser la bibliothèque standard de `Go`, en particulier le package `time`. La propriété `#alts` représente le nombre d'alternatives, numérotées de 1 à `#alts`.

- Code retour

| Code retour | Signification |
|-------------|---------------|
| `201`       | vote créé     |
| `400`       | bad request   |
| `501` 	  | not implemented |

- Objet `JSON` renvoyé (si `201`)

| propriété  | type | exemple de valeurs possibles                                  |
|------------|-------------|-----------------------------------------------------|
| `ballot-id`    | `string` | `"vote12"` |

### Commande `/vote`

- Requête : `POST`
- Objet `JSON` envoyé

| propriété   | type | exemple de valeurs possibles |
|------------|-------------|------------------------|
| `agent-id` | `string` | `"ag_id1"` |
| `vote-id`  | `string` | `"vote12"` |
| `prefs`    | `[int,...]` | `[1, 2, 4, 3]` |
| `options`  | `[int,...]` | `[3]` |

*Remarque :*`options` est facultatif et permet de passer des renseignements supplémentaires (par exemple le seuil d'acceptation en approval)

- code retour

| Code retour | Signification |
|-------------|---------------|
| `200`       | vote pris en compte  |
| `400`       | bad request          |
| `403`       |	vote déjà effectué   |
| `501` 	  | Not Implemented      |
| `503`       | la deadline est dépassée |

### Commande `/result`

- Requête : `POST`
- Objet `JSON` envoyé

| propriété  | type | exemple de valeurs possibles                                  |
|------------|-------------|-----------------------------------------------------|
| `ballot-id`    | `string` | `"vote12"` |


- code retour

| Code retour | Signification   |
|-------------|-----------------|
| `200`       | OK              |
| `425`       | Too early       |
| `404`       |	Not Found       |

- Objet `JSON` renvoyé (si `200`)

| propriété   | type | exemple de valeurs possibles |
|------------|-------------|------------------------|
| `winner`   | `int`       | `4`                    |
| `ranking`  | `[int,...]` | `[2, 1, 4, 3]`         |

*Remarque :* la propriété `ranking` est facultative.

