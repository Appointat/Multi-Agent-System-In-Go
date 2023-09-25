---
tags: TD, Go
---

# [IA04] TD1 : montée en compétences en Go

|  Information |   Valeur              |
| :------------ | :------------- |
| **Auteurs** | Sylvain Lagrue ([sylvain.lagrue@utc.fr](mailto:sylvain.lagrue@utc.fr))|
| | Khaled Belahcene (khaled.belahcene@utc.fr) |
| **Licence** | Creative Common [CC BY-SA 3.0](https://creativecommons.org/licenses/by-sa/3.0) |
| **Version document** | 1.1.0 |

## Préliminaires

Créer l’arborescence suivante sur votre compte:

```
- go
   - bin
   - pkg
```

Le répertoire `bin` contiendra les binaires créés à l'aide de `go install`, tandis que `pkg` contiendra les packages et modules importés à l'aide de la commande `go get`.

Mettre les variables `PATH` et `GOPATH` à jour pour que tout fonctionne.

## Exercice 1 : premiers pas en Go

### Création de l'architecure et du module du premier tp

Tous les exercices seront réalisés dans un répertoire `tp1`.

```shell
> mkdir tp1
> cd tp1
> go mod init tp1 # ou votre_chemin_unique/tp1
> mkdir cmd
> mkdir cmd/hello
```

Pourquoi avoir ajouté un répertoire `cmd` ?  

### Hello world

Écrire un programme Hello World et tester les commandes suivantes (avec les bons arguments).
```
go version
go build
go run
go install
gofmt / go fmt
go doc
```

Que font-elles ? Quelle est leur utilité ?

### Pair/impair

Écrire un programme qui affiche les nombres pairs entre 1 et 1000.


## Exercice 2 : slices et tableaux

1. Écrire une fonction `Fill(sl []int)` qui prend en entrée un slice d'entiers et qui le remplit avec des nombres aléatoires.
1. Écrire une fonction `Moyenne(sl []int)` qui prend en entrée un slice d’entiers et qui renvoie la moyenne de ses éléments.
2. Écrire une fonction `ValeursCentrales(sl []int) []int` qui prend en entrée un slice d’entiers et qui renvoie la ou les valeurs centrales du slice. 
3. Écrire une fonction `Plus1(sl []int)` qui prend en entrée un slice d’entiers et qui ajoute un à chaque élément. Comment utiliser cette fonction pour tester les trois précédentes ?
4. Écrire une fonction `Compte(tab []int)` qui affiche les nombres contenus dans tab.

## Exercice 3 : les sprites

L'objectif de cet exercice est de créer différentes structures pour gérer des sprites. Selon [Wikipédia](https://fr.wikipedia.org/wiki/Sprite_(jeu_vid%C3%A9o)) :

> Dans le domaine de l’infographie et des jeux vidéo, désigne une image en deux dimensions qui peut être déplacée indépendamment du fond (ou décor) de l’affichage et est théoriquement superposé par le processeur vidéo, au moment d’envoyer le signal à l’écran.

1. Écrire une structure `Point2D` et un constructeur `NewPoint2D(x, y float64) (*Point2D)` les getters/setters `(Point2D) X() float32`, `(Point2D) Y() float32`,   `(Point2D) Clone() (Point2D)`, `Module() float64` qui calcule sa distance à l'origine.
2. Écrire la structure `Rectangle` qui contient 2 points : le sommet "en haut à gauche" et celui "en bas à droite". Créer un constructeur, des getters/setters.
4. Écrire la structure `Sprite` qui contient une position, une *hitbox*, un facteur de zoom et le nom du fichier bitmap. Écrire un constructeur, des setters/getters, des méthodes `Move` qui déplace le sprite à une autre position et `Collision` qui calcule le rectangle collision entre 2 sprites.


## Problème : les anagrammes et les palindromes

### Préliminaires

On souhaite créer un programme en Go permettant d’obtenir l’ensemble des palindromes et des anagrammes de la langue française.

Pour les différents tests, on pourra utiliser les mots suivants 

```go
dict := [...]string{"AGENT", "CHIEN", "COLOC", "ETANG", "ELLE", "GEANT", "NICHE", "RADAR"}
```

### Palindromes

1. Écrire la fonction `IsPalindrome(word string) bool` qui détermine si une chaine de caractère est un palindrome.

```go
IsPalindrome("RADAR") // true
IsPalindrome("AGENT") // false
```

2. Écrire `Palindromes(words []string) (l []string)` qui renvoie un slice sur le sous-ensemble des mots de `words` qui sont des palindromes
 
```go
Palindromes(dict) // [ "COLOC", "ELLE", "RADAR" ]
```

### Anagrammes

1. Écrire la fonction`Footprint(s string) (footprint string)` qui calcule l'empreinte de la chaine s, c'est-à-dire la chaîne avec ses lettres par ordre alphabétique. On utilisera pour cela le package `sort` de la bibliothèque standard.
```go
FootPrint("AGENT") // renvoie "AEGNT"
```

2. Écrire la fonction `Anagrams(words []string) (anagrams map[string][]string)` qui renvoie une `map` associant à chaque empreinte trouvée la liste de mots qui possèdent cette empreinte.

```go
Anagrams(dict)
/* { AEGNT: [AGENT ETANG GEANT]
     CEHIN: [CHIEN NICHE]
     CCLOO: [COLOC]
     EELL:  [ELLE]
     AADRR: [RADAR] } */
```

### Application à la langue française

1. Écrire une fonction `DictFromFile(filename string) (dict []string)` qui lit un fichier contenant un mot par ligne et qui renvoie la liste des mots contenus dans le ficher. 
2. Répondre aux questions suivantes :
    1. Quel est le(s) plus long(s) palindromes de la langue française ?
    2. Quels sont les anagrammes de agents ?
    3. Quel(s) mot(s) de la langue française contien(nen)t le plus d’anagrammes ?
    4. Existe-t-il un palindrome qui possède des anagrammes ?

## S’avancer pour la semaine prochaine…

1. Reprendre la question 2 de l'exercice 1 mais un lançant vos affichages avec le mot-clef `go`. Que se passe-t-il ? Pourquoi ? 
2. Reprendre l'exercice 2 en utilisant des goroutines.
