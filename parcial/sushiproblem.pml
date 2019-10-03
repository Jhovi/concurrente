#define wait(s) atomic{s>0 -> s--}
#define signal(s) s++

byte eating = 0
byte waiting = 0
byte mutex = 1
byte block = 0
bool wait = false

active proctype Sushi(){
    do
    ::
        wait(mutex)
        if
        ::(wait) -> {waiting++
        signal(mutex)
        wait(block)
        }
        ::else->{
            eating++
            eating = 5
            wait = eating
            signal(mutex)
        }
        fi

        printf("Comiendo sushi\n")

        wait(mutex)
        eating--
        if
        ::(eating == 0)->{
            n = min(5,waiting)
            waiting -=n
            eating = 5
            wait = eating
            signal(block)
        }
        ::else->
        fi
        signal(mutex)
    od
}