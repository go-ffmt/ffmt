# Display a friendly fmt for golang


## Test
```
go test -v github.com/wzshiming/ffmt

```
```
=== RUN   Test_fmt
[
 struct {
  Msg: "Display a friendly fmt for golang"
  Msg2: "hello all"
  Msg3: "hello all2"
  Msg4: "hello all hello all hello all hello all hello all hello all "
  msg: < private >
  Stru: slice [
   struct {
    Msg: ""
    AA: array [
     int ( 0 )
     int ( 0 )
     int ( 0 )
     int ( 0 )
     int ( 0 )
    ]
   }
   struct {
    Msg: "Test"
    AA: array [
     int ( 2222 )
     int ( 3333 )
     int ( 0 )
     int ( 0 )
     int ( 0 )
    ]
   }
  ]
  Floats: array [
   float32 ( 2.1 )
   float32 ( 3.3 )
   float32 ( 0 )
   float32 ( 0 )
   float32 ( 0 )
  ]
  Ints: slice [
   slice [
    int ( 1 )
    int ( 4 )
   ]
   slice [
    int ( 3 )
   ]
  ]
  Maps: map [
   "aa": "hi world"
   "bb": "bye world"
  ]
  B: bool ( true )
 }
 "{
  "Msg": "Display a friendly fmt for golang"
 ,"Msg2": "hello all"
 ,"Msg3": "hello all2"
 ,"Msg4": "hello all hello all hello all hello all hello all hello all "
 ,"Stru": [
   {
    "Msg": ""
   ,"AA": [
     0
    ,0
    ,0
    ,0
    ,0
    ]
   }
  ,{
    "Msg": "Test"
   ,"AA": [
     2222
    ,3333
    ,0
    ,0
    ,0
    ]
   }
  ]
 ,"Floats": [
   2.1
  ,3.3
  ,0
  ,0
  ,0
  ]
 ,"Ints": [
   [
    1
   ,4
   ]
  ,[
    3
   ]
  ]
 ,"Maps": {
   "aa": "hi world"
  ,"bb": "bye world"
  }
 ,"B": true
 }
 "
]
 --- PASS: Test_fmt (0.00s)
=== RUN   Test_Now
time.Time < 2016-01-05 23: 35: 07.3397014 +0800 CST >
 
< invalid >
 --- PASS: Test_Now (0.00s)
PASS
ok  	github.com/wzshiming/ffmt	0.035s
```




## Install

```
go get -u github.com/wzshiming/ffmt
```