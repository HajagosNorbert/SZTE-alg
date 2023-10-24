Oldja meg a következő rekurziót: T(1) = 3, és minden n ≥ 2 háromhatványra
T(n) = 4T(n/3) + 2n − 1.

Solve the next recursion: T(1) = 3, and for each n ≥ 2 power of 3-es
T(n) = 4T(n/3) + 2n − 1.

-----

Enrolling the recursion

T(n) = 4T(n/3) + 2n − 1
T(n) = 4    [T(n/3) = 4T(n/9) + 2/3n − 1]   + 2n − 1
T(n) = 4*[4T(n/9) + 2/3n − 1] + 2n − 1 =
= 4^2*T(n/9) + 4*2/3n − 4 + 2n − 1 =
= 4^3*T(n/3^3) + 4^*2/3n − 4 + 2n − 1 =

