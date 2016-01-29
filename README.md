# 更加友好的显示数据类型的golang库
# Display a friendly fmt for golang


## Test
```
go test -v github.com/wzshiming/ffmt

```
```
=== RUN   Test_fmt
{
  "Stru": [
    {
      "Msg": ""
     ,"AA":  [
        0
       ,0
       ,0
       ,0
       ,0
      ]
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
   ,[
    ]
  ]
 ,"Maps": {
    "aa": "hi world"
   ,"bb": "bye world"
  }
 ,"B":      true
 ,"T":      "2016-01-29 11:44:10.5678027 +0800 CST"
 ,"Chan":   "chan(0x0000000000c08205c000)"
 ,"Msg":    "Display a friendly fmt for golang"
 ,"Msg2":   "你好"
 ,"Msg3":   ""
 ,"Msg4":   "hello all hello all hello all hello all hello all hello all "
 ,"Floats": [
    2.1
   ,3.3
   ,0
   ,0
   ,0
  ]
 ,"TTT": {
    "Msg4": "hello all hello all hello all hello all hello all hello all "
   ,"Stru": [
      {
        "Msg": ""
       ,"AA":  [
          0
         ,0
         ,0
         ,0
         ,0
        ]
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
   ,"B":      true
   ,"T":      "2016-01-29 11:44:10.5678027 +0800 CST"
   ,"Msg":    "Display a friendly fmt for golang"
   ,"Msg2":   "你好"
   ,"Msg3":   ""
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
     ,[
      ]
    ]
   ,"Maps": {
      "aa": "hi world"
     ,"bb": "bye world"
    }
   ,"TTT":  "ptr(0x000000000000006a2fa0)"
   ,"Chan": "chan(0x0000000000c08205c000)"
  }
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
    [
    ]
  ]
  Maps: {
    aa: hi world
    bb: bye world
  }
  B:   true
  T:   <2016-01-29 11:44:10.5678027 +0800 CST>
  TTT: &{
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
      [
      ]
    ]
    Maps: {
      aa: hi world
      bb: bye world
    }
    B:    true
    T:    <2016-01-29 11:44:10.5678027 +0800 CST>
    TTT:  <ptr(0x000000000000006a2fa0)>
    Chan: <chan(0x0000000000c08205c000)>
  }
  Chan: <chan(0x0000000000c08205c000)>
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
    [
    ]
  ]
  Maps: {
    aa: "hi world"
    bb: "bye world"
  }
  B:   true
  T:   <2016-01-29 11:44:10.5678027 +0800 CST>
  TTT: &{
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
      [
      ]
    ]
    Maps: {
      aa: "hi world"
      bb: "bye world"
    }
    B:    true
    T:    <2016-01-29 11:44:10.5678027 +0800 CST>
    TTT:  <ptr(0x000000000000006a2fa0)>
    Chan: <chan(0x0000000000c08205c000)>
  }
  Chan: <chan(0x0000000000c08205c000)>
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
    slice[
    ]
  ]
  Maps: map{
    string(aa): string(hi world)
    string(bb): string(bye world)
  }
  B:   bool(true)
  T:   <2016-01-29 11:44:10.5678027 +0800 CST>
  TTT: &struct{
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
      slice[
      ]
    ]
    Maps: map{
      string(aa): string(hi world)
      string(bb): string(bye world)
    }
    B:    bool(true)
    T:    <2016-01-29 11:44:10.5678027 +0800 CST>
    TTT:  <ptr(0x000000000000006a2fa0)>
    Chan: <chan(0x0000000000c08205c000)>
  }
  Chan: <chan(0x0000000000c08205c000)>
}
--- PASS: Test_fmt (0.00s)
PASS
ok  	github.com/wzshiming/ffmt	0.028s
```




## Install

```
go get -u github.com/wzshiming/ffmt
```
