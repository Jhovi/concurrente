#define wait(s) atomic{s>0 -> s--}
#define signal(s)s++
#define N 4

byte customers = 0
byte mutex = 1
byte customer = 0
byte barber = 0
byte customerDone = 0
byte barberDone = 0

active proctype Customer(){
    do
    ::
    wait(mutex)
    if
    ::(customers == N)->{
        signal(mutex)
        balk()
    }
    else->
    fi
    customers++
    signal(mutex)

    signal(customer)
    wait(barber)

    printf("Cortandome el cabello")
    signal(customerDone)

    wait(mutex)
    customers--
    signal(mutex)

    od
}

active proctype Barber(){
    do
    ::
        wait(customer)
        signal(barber)

        printf("Cortando")

        wait(customerDone)
        signal(barberDone)
    od
}