---
tags: TD, Go
---

# [IA04] TD2 goroutines et synchronisation en Go

|  Information |   Valeur              |
| :------------ | :------------- |
| **Auteurs** | Sylvain Lagrue ([sylvain.lagrue@utc.fr](mailto:sylvain.lagrue@utc.fr))|
| | Khaled Belahcene (khaled.belahcene@utc.fr) |
| **Licence** | Creative Common [CC BY-SA 3.0](https://creativecommons.org/licenses/by-sa/3.0) |
| **Version document** | 0.5.1 |


## Exercice 1 : Premières goroutines

1. Créer une fonction `compte(n int)` qui affiche les nombres de 0 à n (exclu). Lancer cette fonction normalement, puis à partir d'une goroutine.
2. Créer une fonction `compteMsg(n int, msg string)` qui fait la même chose mais affiche `msg` avant chaque nombre de 0 à n. Lancer la fonction dans 2 goroutines différentes. Que se passe-t-il ?
3. Créer une fonction `compteMsgFromTo(start int, end int, msg string)` qui compte de `start` inclus à `end` exclu (tout en affichant le message avant chaque nombre).
4. Lancer 10 goroutines comptant de 0 à 100, chacune s'occupant d'une dizaine.

## Exercice 2 : Data race
Soit le programme suivant :


```golang
package main

import (
    "fmt"
)

var n = 0

func f() {
    n++
}

func main() {
    for i := 0; i < 10000; i++ {
        go f()
    }

    fmt.Println("Appuyez sur entrée")
    fmt.Scanln()
    fmt.Println("n:", n)
}
```

1. Qu'affiche le programme (essayer d'imaginer le résultat avant de le lancer) ? Pourquoi ? Que se passe-t-il si l'on enlève le `Scanln` ? Pourquoi ?
2. Corriger le problème.


## Exercice 3 : bonne année !

Afficher un compte à rebours : 5, 4, 3, 2, 1... Bonne année ! Donner 3 versions (utilisant respectivement `Sleep`, `After` et `Tick`)


## Exercice 4 : problème gestion concurrente de tableaux

Créer 2 versions des fonctions suivantes (une normale et une parallèle distribuée) sur des tableaux d'entier.

Tester les fonctions et chronométrer le temps d'éxécution de chacune pour des tableaux de plusieurs centaines de Mo (idéalement générés aléatoirement, à l'aide de `foreach` par exemple). Quelles versions sont les plus rapides ? Augmenter le nombre de goroutines. Que se passe-t-il ? Vérifier que les tableaux sont bien identiques (version avec et sans goroutine) à l'aide de la fonction `Equal` sans goroutine.



### Quelques fonctions

```golang
// remplit tab avec la valeur v
func Fill(tab []int, v int)
```

```golang
// applique f sur chaque élément de tab et remplace la valeurÒ
func ForEach(tab []int, f func (int) int)
```

```golang
// copy le tableau src dans dest
func Copy(src []int, dest []int)
```

```golang
// vérifie que tab1 et tab2 sont identiques
func Equal(tab1 []int, tab2 []int)
```

### Pour aller plus loin...

```golang
// trouve la valeur et le premier élément de tab vérifiant f (-1, 0 si rien n'est trouvé)
func Find(tab []int, f func (int) bool) (index int, val int)
```

```golang
// crée un nouveau tableau à partit de tab tel que les éléments du nouveau tableau sont composés des éléments du précédent tableau auxquels on a appliqué la fonction f
func Map(tab []int, f func (int) int) []int
```

```golang
// reduit le tableau à un entier (agrégation de tous les éléments avec la fonction f)
func Reduce(tab []int, init int, func f(int, int) int) int
```


```golang
// vérifie que la condition f est vérifiée sur tous les éléments du tableau tab
func Every(tab []int, f func (int) bool)
```

```golang
// vérifie que la condition f est vérifiée sur au moins un élément du tableau tab
func Any(tab []int, f func (int) bool)
```

```golang
// tri le tableau tab
func Sort(tab []int, func comp(int, int) bool)
```


## Exercice 5 : Ping Pong

1. Créer 2 agents qui se répondent ping/pong via un channel. On utilisera les structures et interfaces suivantes : 
```golang
type Agent interface {
	Start()
}

type PingAgent struct {
	ID string
	c  chan string
}

type PongAgent struct {
	ID string
	c  chan string
}
``` 
2. Améliorer le programme précédent afin d'avoir un seul ponger et de multiples pinger.


