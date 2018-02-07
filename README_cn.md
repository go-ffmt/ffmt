# golang 友好的数据格式化

 - [English](./README.md)
 - [简体中文](./README_cn.md)

## 安装

``` shell
go get -u -v gopkg.in/ffmt.v1
```

## 用法

[API 文档](http://godoc.org/gopkg.in/ffmt.v1)

## 示例

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

## MIT许可证

版权所有©2017-2018 wzshiming <[https://github.com/wzshiming](https://github.com/wzshiming)>。

MIT是[MIT许可证](https://opensource.org/licenses/MIT)许可的开源软件。
