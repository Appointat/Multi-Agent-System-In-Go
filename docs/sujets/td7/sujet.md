---
tags: TD, dynamique des croyances, fusion
---

# [IA04] Sujet TD7

## Exercice : fusion et vaccins (examen final A21)

4 experts médicaux discutent de l'efficacité de 3 vaccins. Considérons les variables propositionnelles $`\{A,B,C\}`$ signifiant « le vaccin A (resp. B, C) est efficace. »

- L'expert n°1 et l'expert n°2 pensent que les vaccins A et B sont efficaces (mais ne se prononce pas pour le dernier). 
- L'expert n°3 estime que le vaccin B est efficace, mais pas le vaccin A. 
- L'expert n°4 estime que si le vaccin B est efficace, alors le vaccin C l'est également (N.B. si le vaccin B n'est pas efficace, il n'a aucune idée de l'efficacité du vaccin C).

### Questions

1. Représenter les croyances de chaque agent $`j`$ à l'aide d'une formule propositionnelle $`\varphi_j`$. Quels sont les modèles de chacune de ces formules (cf. tableau ci-après) ?
3. Qu'est-ce la distance de Hamming entre une interprétation et une formule ? Comment la calculer ?
4. Remplir le tableau suivant, sachant que $`d_H^j(\varphi_j, \omega_i)`$ représente la distance de Hamming entre l'interprétation $`ω_i`$ et $`\varphi_j`$.
5. Quel est le résultat de la fusion de ces 4 croyances en prenant comme opérateur d'agrégation la somme, la somme des carrés, le $`\max`$ ou le $`\text{GMax}`$ ?
6. Après une série massive de tests, on est maintenant sûr que le vaccin C n'est pas efficace. Comment modéliser cette contrainte d'intégrité, notée $`\mu`$ à l'aide d'une formule propositionnelle ?
7. Que devient le résultat de la fusion si l'on prend comme contrainte d'intégrité $`\mu`$ ?

| $`\Omega`$ | $`A`$ | $`B`$ | $`C`$ | $`d_H^1(\varphi_1, \omega_i)`$  | $`d_H^2(\varphi_2, \omega_i)`$ | $`d_H^3(\varphi_3, \omega_i)`$  | $`d_H^4(\varphi_4, \omega_i)`$  | $`\displaystyle\sum_{j=1}^4 d_H^j(\varphi_j, \omega_i)`$  | $`\displaystyle\sum_{j=1}^4 d_H^j(\varphi_j, \omega_i)^2`$ | $`\displaystyle\max_{j=1..4} d_H^j(\varphi_j, \omega_i)`$ | $`\displaystyle\underset{j=1..4}{\text{GMax}}\ d_H^j(\varphi_j, \omega_i)`$ |
|---|---|---|---|---|---|---|---|---|---|---|---|
| $`ω_0`$ | 0 | 0 | 0 |   |   |   |   |   |   |   |   |
| $`ω_1`$ | 0 | 0 | 1 |   |   |   |   |   |   |   |   |
| $`ω_2`$ | 0 | 1 | 0 |   |   |   |   |   |   |   |   |
| $`ω_3`$ | 0 | 1 | 1 |   |   |   |   |   |   |   |   |
| $`ω_4`$ | 1 | 0 | 0 |   |   |   |   |   |   |   |   |
| $`ω_5`$ | 1 | 0 | 1 |   |   |   |   |   |   |   |   |
| $`ω_6`$ | 1 | 1 | 0 |   |   |   |   |   |   |   |   |
| $`ω_7`$ | 1 | 1 | 1 |   |   |   |   |   |   |   |   |

