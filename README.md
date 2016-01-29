# 更加友好的显示数据类型的golang库
# Display a friendly fmt for golang


## Test
```
go test -v github.com/wzshiming/ffmt

```
```
=== RUN   Test_fmt
{
  "B":      true
 ,"Chan":   "chan(0x0000000000c082060000)"
 ,"Floats": [
   2.1,3.3,0
  ,0  ,0  ,0
  ]
 ,"Ints": [
   [
    1,4    ,5
   ,1,4    ,5
   ,6,11999,0
   ]
  ,[
    3
   ]
  ,[
   ]
  ]
 ,"Maps": {
   "aa": "hi world"
  ,"bb": "bye world"
  }
 ,"Msg":  "Display a friendly fmt for golang"
 ,"Msg2": "你好"
 ,"Msg3": "hello all hello all hello all hello all hello all hello all "
 ,"Msgs": [
   "hello","world"
  ,"bey"  ,"bey"
  ]
 ,"Stru": [
   {
    "AA": [
     0,0,0,0
    ,0,0,0,0
    ]
   ,"Msg": ""
   }
  ,{
    "AA": [
     2222,3333,0,0
    ,0   ,0   ,0,0
    ]
   ,"Msg": "Test"
   }
  ]
 ,"T":   "2016-01-29 16:35:36.0838884 +0800 CST"
 ,"TTT": {
   "B":      true
  ,"Chan":   "chan(0x0000000000c082060000)"
  ,"Floats": [
    2.1,3.3,0
   ,0  ,0  ,0
   ]
  ,"Ints": [
    [
     1,4    ,5
    ,1,4    ,5
    ,6,11999,0
    ]
   ,[
     3
    ]
   ,[
    ]
   ]
  ,"Maps": {
    "aa": "hi world"
   ,"bb": "bye world"
   }
  ,"Msg":  "Display a friendly fmt for golang"
  ,"Msg2": "你好"
  ,"Msg3": "hello all hello all hello all hello all hello all hello all "
  ,"Msgs": [
    "hello","world"
   ,"bey"  ,"bey"
   ]
  ,"Stru": [
    {
     "AA": [
      0,0,0,0
     ,0,0,0,0
     ]
    ,"Msg": ""
    }
   ,{
     "AA": [
      2222,3333,0,0
     ,0   ,0   ,0,0
     ]
    ,"Msg": "Test"
    }
   ]
  ,"T":   "2016-01-29 16:35:36.0838884 +0800 CST"
  ,"TTT": "ptr(0x000000000000006a6040)"
  }
}
{
  Msg:  Display a friendly fmt for golang
  Msg2: 你好
  Msg3: hello all hello all hello all hello all hello all hello all
  msg:  <private>
  Msgs: [
   hello world
   bey   bey
  ]
  Stru: [
   {
    Msg:
    AA:  [
     0 0 0 0
     0 0 0 0
    ]
   }
   {
    Msg: Test
    AA:  [
     2222 3333 0 0
     0    0    0 0
    ]
   }
  ]
  Floats: [
   2.1 3.3 0
   0   0   0
  ]
  Ints: [
   [
    1 4     5
    1 4     5
    6 11999 0
   ]
   [
    3
   ]
   [
   ]
  ]
  Maps: {
   aa: hi world
   bb: bye world
  }
  B:   true
  T:   <2016-01-29 16:35:36.0838884 +0800 CST>
  TTT: &{
   Msg:  Display a friendly fmt for golang
   Msg2: 你好
   Msg3: hello all hello all hello all hello all hello all hello all
   msg:  <private>
   Msgs: [
    hello world
    bey   bey
   ]
   Stru: [
    {
     Msg:
     AA:  [
      0 0 0 0
      0 0 0 0
     ]
    }
    {
     Msg: Test
     AA:  [
      2222 3333 0 0
      0    0    0 0
     ]
    }
   ]
   Floats: [
    2.1 3.3 0
    0   0   0
   ]
   Ints: [
    [
     1 4     5
     1 4     5
     6 11999 0
    ]
    [
     3
    ]
    [
    ]
   ]
   Maps: {
    aa: hi world
    bb: bye world
   }
   B:    true
   T:    <2016-01-29 16:35:36.0838884 +0800 CST>
   TTT:  <ptr(0x000000000000006a6040)>
   Chan: <chan(0x0000000000c082060000)>
  }
  Chan: <chan(0x0000000000c082060000)>
}
{
  Msg:  "Display a friendly fmt for golang"
  Msg2: "你好"
  Msg3: "hello all hello all hello all hello all hello all hello all "
  msg:  <private>
  Msgs: [
   "hello" "world"
   "bey"   "bey"
  ]
  Stru: [
   {
    Msg: ""
    AA:  [
     0 0 0 0
     0 0 0 0
    ]
   }
   {
    Msg: "Test"
    AA:  [
     2222 3333 0 0
     0    0    0 0
    ]
   }
  ]
  Floats: [
   2.1 3.3 0
   0   0   0
  ]
  Ints: [
   [
    1 4     5
    1 4     5
    6 11999 0
   ]
   [
    3
   ]
   [
   ]
  ]
  Maps: {
   aa: "hi world"
   bb: "bye world"
  }
  B:   true
  T:   <2016-01-29 16:35:36.0838884 +0800 CST>
  TTT: &{
   Msg:  "Display a friendly fmt for golang"
   Msg2: "你好"
   Msg3: "hello all hello all hello all hello all hello all hello all "
   msg:  <private>
   Msgs: [
    "hello" "world"
    "bey"   "bey"
   ]
   Stru: [
    {
     Msg: ""
     AA:  [
      0 0 0 0
      0 0 0 0
     ]
    }
    {
     Msg: "Test"
     AA:  [
      2222 3333 0 0
      0    0    0 0
     ]
    }
   ]
   Floats: [
    2.1 3.3 0
    0   0   0
   ]
   Ints: [
    [
     1 4     5
     1 4     5
     6 11999 0
    ]
    [
     3
    ]
    [
    ]
   ]
   Maps: {
    aa: "hi world"
    bb: "bye world"
   }
   B:    true
   T:    <2016-01-29 16:35:36.0838884 +0800 CST>
   TTT:  <ptr(0x000000000000006a6040)>
   Chan: <chan(0x0000000000c082060000)>
  }
  Chan: <chan(0x0000000000c082060000)>
}
struct{
  Msg:  string(Display a friendly fmt for golang)
  Msg2: string(你好)
  Msg3: string(hello all hello all hello all hello all hello all hello all )
  msg:  <private>
  Msgs: slice[
   string(hello) string(world)
   string(bey)   string(bey)
  ]
  Stru: slice[
   struct{
    Msg: string()
    AA:  array[
     int(0) int(0) int(0) int(0)
     int(0) int(0) int(0) int(0)
    ]
   }
   struct{
    Msg: string(Test)
    AA:  array[
     int(2222) int(3333) int(0) int(0)
     int(0)    int(0)    int(0) int(0)
    ]
   }
  ]
  Floats: array[
   float32(2.1) float32(3.3) float32(0)
   float32(0)   float32(0)   float32(0)
  ]
  Ints: slice[
   slice[
    int(1) int(4)     int(5)
    int(1) int(4)     int(5)
    int(6) int(11999) int(0)
   ]
   slice[
    int(3)
   ]
   slice[
   ]
  ]
  Maps: map{
   string(aa): string(hi world)
   string(bb): string(bye world)
  }
  B:   bool(true)
  T:   <2016-01-29 16:35:36.0838884 +0800 CST>
  TTT: &struct{
   Msg:  string(Display a friendly fmt for golang)
   Msg2: string(你好)
   Msg3: string(hello all hello all hello all hello all hello all hello all )
   msg:  <private>
   Msgs: slice[
    string(hello) string(world)
    string(bey)   string(bey)
   ]
   Stru: slice[
    struct{
     Msg: string()
     AA:  array[
      int(0) int(0) int(0) int(0)
      int(0) int(0) int(0) int(0)
     ]
    }
    struct{
     Msg: string(Test)
     AA:  array[
      int(2222) int(3333) int(0) int(0)
      int(0)    int(0)    int(0) int(0)
     ]
    }
   ]
   Floats: array[
    float32(2.1) float32(3.3) float32(0)
    float32(0)   float32(0)   float32(0)
   ]
   Ints: slice[
    slice[
     int(1) int(4)     int(5)
     int(1) int(4)     int(5)
     int(6) int(11999) int(0)
    ]
    slice[
     int(3)
    ]
    slice[
    ]
   ]
   Maps: map{
    string(aa): string(hi world)
    string(bb): string(bye world)
   }
   B:    bool(true)
   T:    <2016-01-29 16:35:36.0838884 +0800 CST>
   TTT:  <ptr(0x000000000000006a6040)>
   Chan: <chan(0x0000000000c082060000)>
  }
  Chan: <chan(0x0000000000c082060000)>
}
--- PASS: Test_fmt (0.00s)
PASS
ok  	github.com/wzshiming/ffmt	0.026s
```




## Install

```
go get -u github.com/wzshiming/ffmt
```
