, funkcionális programozás: nincs mutáció, minden kifejezés
Syntax
 - fn típus deklaráció
 - fn definíció
 - fn hívás
 - tail pozíció
```elm
Num a -> Num a
sum = \n ->
    if n == 1 then
        n
    else
        n + sum (n - 1)
```
Transzformáció!

```elm
Num a, Num a -> Num a
sum = \n, acc ->
    if n == 0 then
        acc
    else
        sum (n - 1) (acc+n)
```

- TCO után. 
- ";" kiemelése, Mono fázis

```elm
sum = \n, acc ->
    mark j
    if n == 0 then
        acc
    else
        n = n-1;
        acc = (acc+n);
        goto j;
```

Láncolt lista, ami "a" típusú elemeket tartalmaz (ConsList):

```elm
ConsList a : [ Nil, Cons a (ConsList a) ]

strList = Cons "egyelemű" Nil

numList = Cons 1 (Cons 2 (Cons 3 Nil))
```

- Mi a map fn
- Típus deklaráció
- Pattern Matching
- ez nem tail rekurzív, de Tail Rekurzív Moduló Konstruktor
- rajz

```elm
map : ConsList a, (a -> b) -> ConsList b
map = \list, f ->
    when list is
        Cons head tail -> Cons (f head) (map tail f)
        Nil -> Nil

double = \num -> num*2
doubledList = map numList double
```

- Cons-ból a `(map tail f)` -t megjelöljük Hole-nak (inicializálatlan memóriaterület). Mindig csak egy Hole lehet. Minden mást kiszámolunk
- Destination Passing Style - mint az acc
- Oppurtinistic In Place Mutation

```elm
map : ConsList a, (a -> b) -> ConsList b
map = \l, f ->
    when l is
        Nil -> Nil
        Cons head tail -> 
            y = f head
            dst = Cons y Hole
            mapDps tail dst f;
            dst


mapDps : ConsList a, (a -> b) -> void
mapDps = \l, dst, f->
    when l is
        Nil -> setTail dst Nil;
        Cons head tail ->
            y = f head
            newDst = Cons head Hole
            setTail dst newDst;
            mapDps tail newDst f 
```

