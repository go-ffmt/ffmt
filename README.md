# 更加友好的显示数据类型的golang库
# Display a friendly fmt for golang


## Test
```
go test -v github.com/wzshiming/ffmt

```
```
=== RUN   Test_fmt
{
  "Msg2":   "你好"
 ,"Floats": [
    2.1
   ,3.3
   ,0
   ,0
   ,0
  ]
 ,"Msg4": "hello all hello all hello all hello all hello all hello all "
 ,"Stru": [
    {
      "AA": [
        0
       ,0
       ,0
       ,0
       ,0
      ]
     ,"Msg": ""
    }
   ,{
      "Msg": "Test"
     ,"AA":  [
        2222
       ,3333
       ,0
       ,0
       ,0
      ]
    }
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
 ,"T": "2016-01-28 14:52:51.2765128 +0800 CST"
 ,"Msg": "Display a friendly fmt for golang"
 ,"Msg3": ""
}
{
  Msg:  Display a friendly fmt for golang
  Msg2: 你好
  Msg3:
  Msg4: hello all hello all hello all hello all hello all hello all
  msg:  <private>
  Stru: [
    {
      Msg:
      AA:  [
        0
        0
        0
        0
        0
      ]
    }
    {
      Msg: Test
      AA:  [
        2222
        3333
        0
        0
        0
      ]
    }
  ]
  Floats: [
    2.1
    3.3
    0
    0
    0
  ]
  Ints: [
    [
      1
      4
    ]
    [
      3
    ]
  ]
  Maps: {
    bb: bye world
    aa: hi world
  }
  B: true
  T: <2016-01-28 14:52:51.2765128 +0800 CST>
}
{
  Msg:  "Display a friendly fmt for golang"
  Msg2: "你好"
  Msg3: ""
  Msg4: "hello all hello all hello all hello all hello all hello all "
  msg:  <private>
  Stru: [
    {
      Msg: ""
      AA:  [
        0
        0
        0
        0
        0
      ]
    }
    {
      Msg: "Test"
      AA:  [
        2222
        3333
        0
        0
        0
      ]
    }
  ]
  Floats: [
    2.1
    3.3
    0
    0
    0
  ]
  Ints: [
    [
      1
      4
    ]
    [
      3
    ]
  ]
  Maps: {
    "aa": "hi world"
    "bb": "bye world"
  }
  B: true
  T: <2016-01-28 14:52:51.2765128 +0800 CST>
}
struct{
  Msg:  string(Display a friendly fmt for golang)
  Msg2: string(你好)
  Msg3: string()
  Msg4: string(hello all hello all hello all hello all hello all hello all )
  msg:  <private>
  Stru: slice[
    struct{
      Msg: string()
      AA:  array[
        int(0)
        int(0)
        int(0)
        int(0)
        int(0)
      ]
    }
    struct{
      Msg: string(Test)
      AA:  array[
        int(2222)
        int(3333)
        int(0)
        int(0)
        int(0)
      ]
    }
  ]
  Floats: array[
    float32(2.1)
    float32(3.3)
    float32(0)
    float32(0)
    float32(0)
  ]
  Ints: slice[
    slice[
      int(1)
      int(4)
    ]
    slice[
      int(3)
    ]
  ]
  Maps: map{
    string(aa): string(hi world)
    string(bb): string(bye world)
  }
  B: bool(true)
  T: <2016-01-28 14:52:51.2765128 +0800 CST>
}
--- PASS: Test_fmt (0.00s)
PASS
ok  	github.com/wzshiming/ffmt	0.026s
```




## Install

```
go get -u github.com/wzshiming/ffmt
```
