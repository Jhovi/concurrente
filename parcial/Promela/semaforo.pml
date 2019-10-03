#define wait(s) atomic{ s > 0 -> s--}
#define signal(s) s++

byte s = 1

active[2] proctype P() {
    do
    ::
    printf("NCS 1\n")
    printf("NCS 2\n")
    printf("NCS 3\n")
    wait(s)

    printf("CS 1\n")
    printf("CS 2\n")
    printf("CS 3\n")

    signal(s)
    od
}