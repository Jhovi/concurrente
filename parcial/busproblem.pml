#define signal(s) s++
#define wait(s) atomic {s>0 -> s--}

byte riders = 0
byte mutex = 1
byte multiplex = 50
byte bus = 0
byte allAboard = 0

active proctype Bus(){
    do
    ::
        wait(mutex)
        if
        ::(riders > 0)-> {
            signal(bus)
            wait(allAboard)
        }
        ::else->
        fi
        signal(mutex)

        depart()
    od
}

active proctype Riders(){
    do
    ::
        wait(multiplex)
        wait(mutex)
        riders++
        signal(mutex)

        wait(bus)
        signal(multiplex)

        boardBus()

        riders--
        if
        ::(riders == 0) -> signal(allAboard)
        ::else-> signal(bus)
        fi
    od
}