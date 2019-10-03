

byte turn = 1

active proctype p(){
    do
    ::
    printf("NCS: 1\n")
    printf("NCS: 2\n")
    printf("NCS: 3\n")

    (turn == 1)->

    printf("CS: 1\n")
    printf("CS: 2\n")
    printf("CS: 3\n")

turn = 2
    od

}

active proctype q(){
    do
    ::
    printf("NCS: 1\n")
    printf("NCS: 2\n")
    printf("NCS: 3\n")

    (turn == 2)->

    printf("CS: 1\n")
    printf("CS: 2\n")
    printf("CS: 3\n")

turn = 1
    od

}