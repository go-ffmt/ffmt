# Display a friendly fmt for golang

## Install

``` shell
go get -u -v gopkg.in/ffmt.v1
```

## Usage

``` golang
// Friendly display
ffmt.Puts(a ...interface{}) (int, error) // The go stlye friendly display of data
ffmt.P(a ...interface{}) (int, error) // The go stlye friendly display of data and types
ffmt.Pjson(a ...interface{}) (int, error) // The json stlye friendly display data

// Mark line
ffmt.Mark(a ...interface{}) // Output prefix current line position
ffmt.MarkStack(skip int, a ...interface{}) // Output prefix stack line position
ffmt.MarkStackFull() // Output stack full

// Table 
ffmt.ToTable(t interface{}, is ...interface{}) [][]string // Data to table data
ffmt.FmtTable(b [][]string) (ss []string) // Format table data
```

## Examples

[examples](./examples/main.go)

``` golang
m := map[string]interface{}{
    "hello": "w",
    "A": []int{
        1, 2, 3, 4, 5, 6,
    },
}

ffmt.Puts(m)
/*
{
 "A": [
  1 2 3
  4 5 6
 ]
 "hello": "w"
}
*/

ffmt.P(m)
/*
map{
 string(A): slice[
  int(1) int(2) int(3)
  int(4) int(5) int(6)
 ]
 string(hello): string(w)
}
*/

ffmt.Pjson(m)
/*
{
 "A": [
  1,2,3
 ,4,5,6
 ]
,"hello": "w"
}
*/

m0 := ffmt.ToTable(m, m)
ffmt.Puts(m0)
/*
[
 [
  "A"
  "hello"
 ]
 [
  "[1 2 3 4 5 6]"
  "w"
 ]
]
*/

m1 := ffmt.FmtTable(m0)
ffmt.Puts(m1)
/*
[
 "A             hello "
 "[1 2 3 4 5 6] w     "
]
*/

ffmt.Mark("hello")
/*
main.go:25  hello
*/
```




## MIT License

Copyright © 2017-2018 wzshiming<[https://github.com/wzshiming](https://github.com/wzshiming)>.

MIT is open-sourced software licensed under the [MIT License](https://opensource.org/licenses/MIT).
