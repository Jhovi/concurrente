#define wait(s) atomic {s > 0 -> s--}
#define wait(fork[N + 1])
#define signal(s) s++

#define N 5
byte filo[N]
byte fork[N+1]
byte pos 0

active proctype filosofo(){
    do
    ::
    println("Pensando")
    wait(pos)
    wait(pos +1)


    println("Comiendo")
    signal(pos)
    signal(pos + 1)

    od
}