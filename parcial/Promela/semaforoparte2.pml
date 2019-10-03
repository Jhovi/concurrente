#define wait(s) atomic{ s > 0 -> s--}
#define signal(s) s++

byte s = 1
byte c = 0

active[2] proctype P() {
    do
    ::
    
    wait(s)
    c++
    assert(c < 2)
    c--


    signal(s)
    od
}