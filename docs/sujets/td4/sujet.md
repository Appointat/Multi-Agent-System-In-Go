---
tags: TD, Comsoc
---

# [IA04] TD4 : théorie du choix social, méthodes et propriétés 

|  Information |   Valeur              |
| :------------ | :------------- |
| **Auteurs** | Sylvain Lagrue ([sylvain.lagrue@hds.utc.fr](mailto:hds.sylvain.lagrue@utc.fr))|
|             | Khaled Belahcene (khaled.belahcene@hds.utc.fr) |
|             | Jérome Gaigne (jerome.gaigne@hds.utc.fr) |
|             | Hénoïk Willot (henoik.willot@hds.utc.fr) |
| **Licence** | Creative Common [CC BY-SA 3.0](https://creativecommons.org/licenses/by-sa/3.0) |
| **Version document** | 1.1.1 |

## Exercice 1 : méthodes simples de vote

### Calcul de résultats

On considère le profil suivant.

| 2 | 1 | 1 | 1 |
|---|---|---|---|
| 3 | 3 | 4 | 1 |
| 1 | 4 | 3 | 4 |
| 2 | 2 | 2 | 3 |
| 4 | 1 | 1 | 2 |

Calculer le résultat des méthodes de vote suivantes :

- Méthode de Borda ;
- Méthode par *scoring* en utilisant le vecteur suivant de scores $`⟨20, 21, 22, 23⟩`$ ;
- STV (*Simple Transferable Vote*).

On pourra utiliser la fonction de départage suivant : le plus le numéro d'une alternative est petit, le plus elle est préférée.

### Quelques propriétés (final A21)

1. Montrer (directement, avec un contre-exemple) que le vote majoritaire ne respecte pas le critère de Condorcet.
2. Montrer (directement, avec un contre-exemple) que le vote de Borda ne respecte pas IIA (*independence of irrelevant alternatives*, en français indépendance des alternatives non disponibles).

## Exercice 2 : méthodes basées sur les graphes

On considère le profil suivant.

| 5 | 4 | 3 |
|---|---|---|
| a | b | d |
| b | c | c |
| c | d | a |
| d | a | b |

### Question

1. Calculer le graphe de majorité.

### Graphe de majorité et méthode de Copeland 

2. Calculer le résultat de la méhode de Copeland.
3. Quel score pourrait obtenir un gagnant de Condorcet ?
4. Montrer que la méthode de Copeland est Condorcet cohérente.
5. En quoi la méthode de Copeland est-elle liée au graphe de majorité ?
6. En quoi peut-on dire que Copeland est la "plus petite réparation" du graphe de majorité ?

### Autres méthodes

Appliquer au profil les méthodes suivantes : 

1. Young/Condorcet
2. Dogson
3. Kramer-Simpson

## Exercice 3 : single peaked preferences

Considérons le profil suivant.

| 4 | 2 | 3 | 1 |
|---|---|---|---|
| 1 | 2 | 4 | 3 |
| 2 | 3 | 5 | 4 |
| 3 | 4 | 3 | 2 |
| 4 | 1 | 2 | 5 |
| 5 | 5 | 1 | 1 |

### Questions 

1. Montrer que les préférences sont à pic simple.
2. Calculer le résultat de la méthode de l'électeur médian.
3. Vérifier que c'est bien un gagnant de Condorcet.
4. Trouver une manipulation est-il possible ?

## Exercice 4 : représentation des préférences des agents (final A21)

### Relations et scores

Soit $`\succsim`$ une relation transitive, réflexive et complète sur un ensemble fini $`E`$. On admet qu'une telle relation peut être *représentée* par une fonction de valeur $`V_\succsim: E \rightarrow \mathbb{N}`$ de manière à ce que $`a \succsim b \iff V_\succsim (a) \ge V_\succsim (b)`$.

1. Donner le graphe complet de la relation représentée par les scores suivants :

| a | b | c | d | e | f | g |
|---|---|---|---|---|---|---|
| 5 | 9 | 23| 4 | 9 | 12| 4 |


2. Proposer une représentation pour la relation suivante :

$`\mathcal{R} = \{(a, a), (a, d), (a, f)`$ ,$`(b, a), (b, b) ,(b, c), (b, d), (b, f)`$, $`(c, a), (c, b), (c, c), (c, d), (c, f)`$, $`(d, d), (d, f), (e, a), (e, b), (e, c), (e, d), (e, e), (f, f), (e, f) \}`$

### La relation inconnue

Soit $`≿_1`$ et $`≿_2`$ deux relations réflexives et transitives sur $`E`$.
Soit $`\mathcal{R}`$ la relation binaire sur $`E`$ définie de la manière suivante :
```math
∀a ∈ E\ ∀b ∈ E \quad a\mathcal{R} b \iff (a≿_1 b\  \text{et non}(b{≿_1} a))\  \text{ou}\  (a ≿_1 b \text{ et }  b ≿_1 a \text{ et } a ≿_2 b)
```

Montrer que :

3. La relation $`\mathcal{R}`$ est réflexive et transitive.
4. Si $`≿_1`$ ou $`≿_2`$ est complète, $`\mathcal{R}`$ l'est aussi.
5. Lorsque l'entier $`N`$ est suffisamment grand, la fonction $`U := N V_{≿_1} + V_{≿_2}`$ représente $`\mathcal{R}`$.
6. Quel type de relation représente cette relation ?
