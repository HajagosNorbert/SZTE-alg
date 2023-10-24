8 bites állománynál 256 karaktert tudunk reprezentálni. Amennyiben véletlen eloszlással dolgozunk, feltételezhetjük, hogy szinte minden karakterből ugyan annyi van a tömörítendő szövegben. Először is, mi történik, ha pontosan ugyan annyi van?

Legyen n a bemenet hossza.
Legyen egy ciklus a huffman kódolásban, amikor a prioritási sorban a felére csökken az adatok mennyisége.

A Huffman algoritmus egy teljesen kiegyenlített bináris fát fog létrehozni, 256 levéllel. Ez azt jelenti, hogy log 256 mély a fa, tehát 8. 8 hosszú minden karakter elkódolása, pont mint azelőtt, ergo nem tömörítettünk semmit. Ennek a kódolásnak a mérete 8*n, 

Ha az inputban van különbség a karakterek mennyiségét illetően, az befolyásolhatja a huffam algoritmus generálta fát. Ez esetben, annak érdekében, hogy a leggyakoribb karaktereknek kisebb legyen a kódolása, a fa nem lesz teljesen kiegyensúlyozva.
A mi esetünkben ha elég nagy a karakterek előfordulásának száma, ez nem fog előfordulni.
Ciklusonként az lesz a jellemző, hogy 2 ugyan olyan mély, teljesen kiegyenlített fát összekötünk egy új fa bal és jobb gyerekeként. 
Ahhoz, hogy ez a minta megszűnjön, az kell, hogy a kis előfordulási különbségek összeadódjanak a ciklusokon át annyira, hogy egyszer legyen egy olyan m levélszámú fa, melynek előfordulási összege nagyobb, mint egy m*2 levelet tartalmazó fának az lőfordulási összege.

EDDIG JÓ----------


Az, hogy mennyire fog eltérni a teljesen kiegyenlített fától az eredmény az inputban szereplő karakterek mennyiségétől függ.
Ha például minden i. karkater hossza 0 -tól 256-ig 2^i + 1, akkor egy láncolt listát fogunk kapni, azzal a különbséggel, hogy az utolsó levélnek lesz egy testvére.
A mi esetünkben ez nem fog megörténni az alapfeltételek miatt.
Viszont ez és a teljesen kiegyensúlyozott fa között lesz a generált fánk, azt tekintve, hogy mekkora a jobb és bal fa mélyége közötti különbség.

Az lesz a jellemző, hogy 2 ugyan olyan mély, teljesen kiegyenlített fát összekötünk egy új fa bal és jobb gyerekeként. Ahogy a kis előfordulási különbségek összeadódnak a ciklusokon át,előfordulhat, hogy az egyik ciklusban a bal oldali fának még csak fele annyi gyereke van, mint a jobb oldalinak. 
