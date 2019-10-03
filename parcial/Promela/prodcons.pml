#define wait(s) atomic {s > 0 -> s--}
#define signal(s) s++

#define N 5
byte buffer [N] 
byte pos = 0

byte notEmpty = 0
byte notFull = N
byte sbuff = 1

active[4] proctype Producer(){
    byte d 
    do
    ::
        d++
        wait(notFull)
        wait(sbuff)
        buffer[pos] = d
        pos++
        signal(sbuff)
        signal(notEmpty)
    od
}

active[4] proctype Consumen(){
    byte d 
    do
    ::
    wait(notEmpty)
    wait(sbuff)
    pos--
    d = buffer[pos]
    signal(sbuff)
    signal(notFull)
    printf("Consumiendo: %d\n",d)
    od
}