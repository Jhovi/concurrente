
byte turn = 1

proctype P(){
    byte n = 0
    do
    :: n = n + 1
    :: turn = turn + 1
    od
}

active proctype Q(){
    do
    ::
     printf("NCS 1\n")
     printf("NCS 2\n")
     printf("NCS 3\n")
    (turn == 1) ->

    printf("Critical 1\n")
    printf("Critical 2\n")
    printf("Critical 3\n")

    turn = 2
    od
}

active proctype R(){
    do
    ::
     printf("NCS 1\n")
     printf("NCS 2\n")
     printf("NCS 3\n")
    (turn == 2) ->

    printf("Critical 1\n")
    printf("Critical 2\n")
    printf("Critical 3\n")

    turn = 1
    od
}