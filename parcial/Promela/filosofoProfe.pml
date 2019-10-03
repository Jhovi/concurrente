#define wait(s) atomic { s > 0 -> s--}
#define signal(s) s++

#define N 5
byte fork[N] = { 1, 1, 1, 1, 1}

active[N-1] proctype P1(){
    do
    ::
    printf("Pensando")
    wait(fork[_pid])
    wait(fork[_pid+1])
    printf("Comiendo")
    signal(fork[_pid])
    signal(fork[_pid+1])
    od
}

active proctype P2(){
    do
    ::
    printf("L Pensando\n")
    wait(fork[0])
    wait(fork[N-1])
    printf("L Comiendo")
    signal(fork[0])
    signal(fork[N-1])
    od
}