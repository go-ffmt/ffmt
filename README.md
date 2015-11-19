# Display a friendly fmt for golang


## Test
```
go test -v github.com/wzshiming/ffmt

```
```
=== RUN   Test_fmt
[
 main.<anonym>{
  Msg:"Display a friendly fmt for golang"
  msg:<private>
  Stru:slice[
   main.<anonym>{
    Msg:"" AA:slice[
     0 0 0 0 0
     0 0 0 0 0
     0 0 0 0 0
     0 0 0 0 0
    ]
   }
   main.<anonym>{
    Msg:"Test"
    AA:slice[
     2222 3333
     0 0 0 0 0
     0 0 0 0 0
     0 0 0 0 0
     0 0 0
    ]
   }
  ]
  Floats:slice[
   2.1 3.3 0
   0 0 0 0 0
   0 0 0 0 0
   0 0 0 0 0
   0 0
  ]
  Ints:slice[
   slice[
    1 4
   ]
   slice[
    3
   ]
  ]
  Maps:map[
   "bb":"bye world"
   "aa":"hi world"
  ]
  B:true
 }
 "{
  "Msg":"Display a friendly fmt for golang"
 ,"Stru":[
   {
    "Msg":"","AA":[
     0,0,0,0,0
    ,0,0,0,0,0
    ,0,0,0,0,0
    ,0,0,0,0,0
    ]
   }
  ,{
    "Msg":"Test"
   ,"AA":[
     2222,3333
    ,0,0,0,0,0
    ,0,0,0,0,0
    ,0,0,0,0,0
    ,0,0,0
    ]
   }
  ]
 ,"Floats":[
   2.1,3.3,0
  ,0,0,0,0,0
  ,0,0,0,0,0
  ,0,0,0,0,0
  ,0,0
  ]
 ,"Ints":[
   [
    1,4
   ]
  ,[
    3
   ]
  ]
 ,"Maps":{
   "aa":"hi world"
  ,"bb":"bye world"
  }
 ,"B":true
 }"
]
--- PASS: Test_fmt (0.00s)
PASS
ok  	github.com/wzshiming/ffmt	0.033s

```




## Install

```
go get -u github.com/wzshiming/ffmt
```