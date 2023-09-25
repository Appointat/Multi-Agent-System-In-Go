# Exemple de serveur Rest en Go

Un agent permet de faire de petites opérations arithmétiques (`+`, `-`, `*`) sur 2 entiers si on lui demande en REST.

Techniquement le serveur n'est pas pur REST (à vous de trouver pourquoi), mais il est fonctionnel et peut tenir la charge.

3 exécutables (indépendants) sont fournis :

* `launch-all-rest-agents` permet de lancer une démo avec un seul exécutable
* `lanch-rsagt` permet de lancer un agent de type serveur
* `lanch-rsagt` permet de lancer un agent de type client