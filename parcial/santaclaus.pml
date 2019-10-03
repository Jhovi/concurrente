#define wait(s) atomic {s>0 -> s--}
#define signal(s) s++

byte elves = 0
byte reindeer = 0
byte santaSem = 0
byte reindeerSem = 0
byte eltex = 1
byte mutex = 1

active proctype Santa(){
    do
    ::
    wait(santaSem)
    wait(mutex)
    if
    ::(reindeer >= 9)->{
        prepareSleigh()
        signal(reindeerSem)
        reindeer -=
    }
    ::else-> if
    ::::(elves == 3)->helpElves()
    else->
    fi
    fi
    signal(mutex)
    od
}

active proctype Reindeer(){
    do
    ::
        wait(mutex)
        reindeer++
        if
        ::(reindeer == 9)-> signal(santaSem)
        ::else->
        fi
        signal(mutex)

        wait(reindeerSem)
        getHitched()
    od
}

active proctype Elves(){
    do
    ::
        wait(eltex)
        wait(mutex)
        elves++
        if
        ::(elves == 3) -> signal(santaSem)
        ::else-> signal(eltex)
        fi
        signal(mutex)

        getHelp()

        wait(mutex)
        elves--
        if
        ::(elves ==0) -> signal(eltex)
        ::else->
        fi
        signal(mutex)
    od
}