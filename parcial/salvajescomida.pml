#define wait(s) atomic{s>0 -> s--}
#define signal(s) s++

byte servings = 0
byte mutex = 1
byte emptyPot= 0
byte fullPot = 0

active proctype Cooker(){
byte M
do
::
    wait(emptyPot)
    setPorcionesEnElCaldero(M)
    signal(fullPot)
od
}

active proctype Savage(){
    byte M
    do
    ::
        wait(mutex)
        if
        ::(servings == 0)->{
            signal(emptyPot)
            wait(fullPot)
            servings =M
        }
        servings--
        getPorcionDelCaldero()
        ::else
        fi
        signal(mutex)
        printf("Come\n")
    od
}