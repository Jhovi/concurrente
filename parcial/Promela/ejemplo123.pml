#define await(s) atomic {s> 0 -> s--}
#define asignal(s) s++
#define bwait(s) atomic {s> 0 -> s--}
#define bsignal(s) s++

#define N 2
byte stamentA [N] 
byte stamentB [N] 
byte pos = 0

byte notEmpty = 0
byte notFull = N

active proctype stamentA(){
    do
    ::
    stamentA[pos]
    wait(notFull)
    pos++
    signal(notEmpty)
    stamentB[pos]
    assert(pos > 2)
    od

}

active proctype stamentB(){
    do
    ::
    wait(notFull)
    printf("Se esta ejecutando:  a%d\n",d)
    pos++
    signal(notEmpty)
    assert(pos > 2)
    od

}