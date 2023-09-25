---
tags: TD, Go
---

# [IA04] TD6 - Vote : Résultats axiomatiques, recherche de consensus et manipulation

Prévoir environ 1h pour chacun des thèmes.

On pourra utiliser les utilitaires fournis pour le calcul des permutations accessibles ici : <https://gitlab.utc.fr/lagruesy/ia04/>

## I. Résultats axiomatiques

On rappelle le théorème de l'électeur médian : 

> **Théorème :** Si un nombre impair de votants a des *préférences à pic simple*, alors le candidat préféré par l'électeur médian est un vainqueur de Condorcet.

1. En déduire une démonstration du théorème suivant :


> **Théorème :** Si un nombre impair de votants ont des *préférences à pic simple*, alors aucune règle de vote respectant le *critère de Condorcet* n'est *manipulable*.

2. **Pour aller plus loin :** 
On rappelle les définitions suivantes, pour une SCF décisive $``v``$ :

> * pour un profil $`\mathbf(P)`$ et deux alternatives $`x \ne y`$, on note $`N_{x\succ y}^\mathbf{P}`$ l'ensemble des votants préférant $`x`$ à $`y`$ dans $`\mathbf{P}`$
> * *Unanimité :* pour tout $`x, y, \mathbf{P}`$ $`\quad`$ on a 
$`N_{x\succ y}^\mathbf{P} = N \Rightarrow y \ne v(\mathbf P)`$
> * *Indépendance :* pour tout $`x, y, \mathbf P, \mathbf{P'}`$ 
$`v(\mathbf P) = x \ \text{et}\  N_{x\succ y}^\mathbf{P} = N_{x\succ y}^\mathbf{P'} \Rightarrow v(\mathbf{P'})\ne y`$
> * *Monotonie (forte)* : pour tout $`x, \mathbf P, \mathbf{P'}`$,
$`v(\mathbf P) = x\ \text{et}\ \forall y\in A\setminus\{x\}\  N_{x\succ y}^\mathbf{P} \subseteq N_{x\succ y}^\mathbf{P'} \Rightarrow v(\mathbf{P'}) = x`$


démontrer que Surjectivité + Monotonie $`\Rightarrow`$ Unanimité dans le cas de SCF décisives.

## II. Implémentation de la règle de Kemeny

> **Définition :** la distance d'édition entre deux rangements $`\succ`$ et $`\succ'`$ d'un ensemble $`A`$ est égale au nombre de paires d'éléments de $`A`$ pour lesquelles les rangements sont en désaccord. On la  note fréquemment $`\tau(\succ, \succ')`$ (fonction Tau de Kendall <https://en.wikipedia.org/wiki/Kendall_tau_distance>).

1. Implémenter une fonction calculant la distance d'édition entre deux rangements.
2. Implémenter une fonction calculant la distance entre un rangement et un profil, définie comme la somme des distances d'édition entre ce rangement et tous les rangements du profil.
3. Implémenter la règle de Kemeny à l'aide d'un algorithme brutal, consistant à :
    * énumérer tous les rangements possibles,
    * calculer leur distance au profil de vote, et
    * sélectionner le plus proche.
4. Tester l'implémentation, par exemple à l'aide de l'un des exemples traités sur la [page Wikipédia de la règle de Kemeny](https://en.wikipedia.org/wiki/Kemeny%E2%80%93Young_method).

## III. Calcul de manipulations

A l'aide de l'implémentation déjà effectuée des différentes règles de votes, on veut répondre à la question posée dans le cours, à savoir l'existence de manipulation pour le profil de vote suivant dans le cadre de chacune de ces règles.

| Votants | 1 | 2 | 3 | 4 | 5 |
|--------|---|---|---|---|----|
| Rang 1 | a | b | e | e | ? |
| Rang 2 | b | a | c | c | ?  |
| Rang 3 | c | d | b | b | ?  |
| Rang 4 | d | e | a | a | ?  |
| Rang 5 | e | c | d | d | ?  |


> **Définition :** en présence d'information incomplète sur les préférences (par exemple, lorsque l'on connait les votes de tous les électeurs sauf un), on dit qu'un candidat est :
> * *vainqueur nécessaire* s'il gagne l'élection pour toutes les manières possibles de compléter les préférences,
> * *vainqueur possible* s'il gagne l'élection pour au moins une manière possible de compléter les préférences,
> * *vainqueur impossible* s'il ne gagne l'élection dans aucune configuration possible.

1. En faisant appel aux fonctions déjà réalisées qui implémentent différentes règles de vote, programmer une fonction `possibleWinners(P Profile) map[candidate] []candidate` qui, étant donné le profil de vote de tous les électeurs sauf un, détermine les vainqueurs possibles de l'élection, ainsi qu'une des complétions correspondantes. On procèdera de manière *brutale*, en énumérant tous les profils de vote pour l'électeur manquant.
2. A l'aide de cette fonction, programmer `isPossibleWinner(P Profile, c candidate) bool` et `isNecessaryWinner(P Profile, c candidate) bool` qui renvoient respectivement les vainqueurs nécessaires et impossibles.
3. A l'aide de `possibleWinners`, programmer la fonction `bestResponse(P profile, pref []candidate) (winner candidate, ballot [] candidate)` qui, étant donné un profil incomplet P et les préférences de l'électeur manquant, renvoie la meilleure stratégie de vote de ce candidat (en cas d'ex aequo, on renverra un des bulletins les plus proche des préférences sincères du candidat au sens de la distance d'édition).
