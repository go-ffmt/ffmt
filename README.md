# Display a friendly fmt for golang


## Test
```
go test -v github.com/wzshiming/ffmt

```
```
=== RUN Test_fmt
{
 "Msg":"Display a friendly fmt for golang"
,"Stru":{
  "Msg":"Test"
 ,"AA":[
   2222
  ,3333
  ]
 }
,"Floats":[
  2.1
 ,3.3
 ]
,"Ints":[
  [
   1
  ,4
  ]
 ,[
   3
  ]
 ]
,"Maps":{
  "aa":"hi world"
 ,"bb":"bye world"
 }
}

{
 Display
 a
 friendly
 fmt
 for
 golang
 {
  Test
  [
   2222
   3333
  ]
 }
 [
  2.1
  3.3
 ]
 [
  [
   1
   4
  ]
  [
   3
  ]
 ]
 map[
  aa:hi
  world
  bb:bye
  world
 ]
}

--- PASS: Test_fmt (0.00s)
PASS
ok      github.com/wzshiming/ffmt       0.325s

```




## Install

```
go get -u github.com/wzshiming/ffmt
```