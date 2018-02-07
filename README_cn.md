# golang 友好的数据格式化

 - [English](./README.md)
 - [简体中文](./README_cn.md)

## 安装

``` shell
go get -u -v gopkg.in/ffmt.v1
```

## 用法

``` golang
// 友好的显示
func Puts(a ...interface{}) (int, error) // 以golang风格友好格式化显示数据
func P(a ...interface{}) (int, error) // 以golang风格友好格式化显示数据并显示类型
func Pjson(a ...interface{}) (int, error) // 以json风格友好格式化数据

// 标记行
func Mark(a ...interface{}) // 输出当前行号
func MarkStack(skip int, a ...interface{}) // 输出栈位置的行号
func MarkStackFull() // 输出完整的栈

// 表格 
func ToTable(t interface{}, is ...interface{}) [][]string // 数据转表格
func FmtTable(b [][]string) (ss []string) // 表格格式化
```

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
