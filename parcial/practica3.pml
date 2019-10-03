
#define signal(s) atomic {s>0 -> s++}
#define wait(s) s--

byte mutex = 1
byte count = 0

 proctype A(){

do
::
    wait(mutex)
    count = count+1
    wait(mutex)

od

}


 proctype B(){

do
::
    wait(mutex)
    count = count+1
    wait(mutex)

od

}

init{
    atomic{
        run A()
        run B()
        printf("count = %d\count",count)
    }
}