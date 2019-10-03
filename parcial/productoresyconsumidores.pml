#define wait(s) atomic {s>0 -> s--}
#define signal(s) s++
#define N 5

byte buffer[N]
byte pos = 0
byte sbuff = 1
byte notEmpty = 0
byte notFull = N

active proctype Producer(){
    byte d
    do
    ::
        d++
        wait(notFull)
        wait(sbuff)
        buffer[pos]  = d
        pos++
        signal(sbuff)
        signal(notEmpty)
    od
}


active proctype Consumer(){
    byte d
    do
    ::
        wait(notEmpty)
        signal(sbuff)
        pos--
        d = buffer[pos]
        signal(sbuff)
        signal(notFull)
        printf("COnsumiendo: %d\n",d)
    od
}