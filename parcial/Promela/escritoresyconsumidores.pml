
#define wait(s) atomic {s>0 -> s--}
#define signal(s) s++
#define N 5

byte readers = 0
byte mutex = 1
byte roomEmpty = 1

active proctype Writer(){
    do
    ::
    wait(roomEmpty)
    printf("Writer writing")
    signal(roomEmpty)
    od
}

active proctype Reader(){
    do
    ::
        wait(mutex)
        readers++
        if
        ::(readers == 1)-> wait(roomEmpty)
        ::else->
        fi

        signal(mutex)
        printf("Reader reading")
        wait(mutex)

        readers--
        if
        ::(readers == 0)-> signal(roomEmpty)
        ::else->
        fi

        signal(mutex)
    od
}