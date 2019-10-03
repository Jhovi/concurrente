
bool wantp = false
bool wantq = false
byte cc = 0

active proctype P(){
    do
    ::
        (wantq == false)
        wantp = true 
        cc++
      
        wantp = false
        cc--

    od

}

active proctype Q(){
    do
    ::
        (wantp == false)
        wantq = true
        cc++
        wantq = false
        cc--
    od
}