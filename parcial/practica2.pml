
bool wantp = false
bool wantq = false

active proctype p(){

    do
    ::
        printf("NCS: 1\n")
        printf("NCS: 2\n")
        printf("NCS: 3\n")

        (wantq == false)->
        wantp = true

        printf("CRITICAL: 1\n")
        printf("CRITICAL: 2\n")
        printf("CRITICAL: 3\n")

        wantp = false

    od
}



active proctype q(){

    do
    ::
        printf("NCS: 1\n")
        printf("NCS: 2\n")
        printf("NCS: 3\n")

        (wantp == false)->
        wantq = true

        printf("CRITICAL: 1\n")
        printf("CRITICAL: 2\n")
        printf("CRITICAL: 3\n")

        wantq = false

    od
}