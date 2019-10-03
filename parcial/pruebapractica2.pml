

bool wantp = false
bool wantq = false
byte cc = 0
active proctype p(){

    do
    ::
 cc++
        (wantq == false)->
         
        wantp = true
      
       
        cc--
        wantp = false
 assert(cc<2)
    od
}



active proctype q(){

    do
    ::
        
 cc++
        (wantp == false)->
        
        wantq = true
       
       
        cc--
        

        wantq = false
 assert(cc<2)
    od
}