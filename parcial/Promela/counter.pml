
byte n = 0

proctype P(){
    byte temp
    byte i

    for (i : 1..10){
        temp = n
        n = temp + 1
    }


}

init {
    atomic { //lanza cosas ininterrupidamente
        run P()
        run P()
    }

    (_nr_pr == 1)-> printf("n = %d\n",n)
    assert(n > 2)
}