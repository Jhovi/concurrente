
byte turn = 1
byte cc = 0

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
     cc++
    (turn == 1) ->
   
    assert(cc < 2)
    cc--
    turn = 2
    od
}

active proctype R(){
    do
    ::
    cc++
    (turn == 2) ->
   
    assert(cc < 2)
    cc--
    turn = 1
    od
}