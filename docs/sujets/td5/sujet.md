---
tags: TP, TD, Go
---

# [IA04] - TD/TP 5 Bureau de vote multi-agent

## Question

1. Récupérer sur Moodle le package `restagentdemo` puis l'exécuter.

```shell
go get gitlab.utc.fr/lagruesy/ia04/demos/restagentdemo
```

2. Inspecter le code. Pourquoi le serveur implémente-t-il une exclusion mutuelle ? 
3. À partir de cet exemple, programmer un bureau de vote simple (sans liste d'émargement) dont on pourra programmer la règle de vote. En particulier, réaliser un constructeur pour les agents votants (avec éventuellement une initialisation aléatoire de leurs préférences), adapter le type de requête et de réponse, etc.

Deux extensions de navigateurs pour faire des requêtes REST : 

- pour chrome : <https://chrome.google.com/webstore/detail/rested/eelcnbccaccipfolokglfhhmapdchbfg?hl=fr&>
- pour firefox : <https://addons.mozilla.org/fr/firefox/addon/rested/>

## Si vous avez fini... Plus petit/plus grand

> Il s'agit d'un jeu (très simple) à information cachée : les joueurs conviennent d'une plage de nombre, l'un des joueurs en choisit un (noté $x$) et le note, et l'autre doit deviner sa valeur en faisant des propositions successives $`n_1, ... n_k`$. Le joueur qui a choisi le nombre répond "gagné" si $`x = n_i`$, "+ haut" si $`n_i < x`$ ou "plus bas" si $`n_i > x`$.

Programmer le jeu du "+ petit / + grand" à l'aide d'une architecture client-serveur. Un client pourra lancer une partie en faisant une requête au serveur qui stipulera les bornes de la plage de valeur. Le serveur choisira aléatoirement un nombre dans cette plage (différent pour chaque partie), et comptera le nombre de coups joués.

On pourra considérer les extensions suivantes :

1. **Joueur artificiel** : programmer un agent artificiel cherchant à deviner le nombre. On pourra considérer une stratégie aléatoire, puis une stratégie optimale sachant que le nombre est choisi avec une probabilité uniforme.

2. **Evil + petit / + grand** : le joueur qui choisit le nombre triche et peut modifier sa valeur, mais ne doit pas être pris en flagrant délit.
