# practise

Go 泛型小练习：对照 [samber/lo](https://github.com/samber/lo) 复刻常用集合函数。

## 已实现

### 1. Slice：遍历 / 回调处理类

| 编号 | 文件 | 函数签名 | 语义 |
| --- | --- | --- | --- |
| 1.1 | `contain.go` | `Contains[T comparable](list []T, item T) bool` | 是否包含某个值 |
| 1.2 | `filter.go` | `Filter[T any](filter func(value T, i int) bool, values []T) []T` | 按谓词过滤 |
| 1.3 | `find.go` | `Find[T any](find func(value T, i int) bool, values []T) (T, bool)` | 找第一个满足谓词的元素 |
| 1.4 | `map.go` | `Map[T any, R any](values []T, mapper func(value T, i int) R) []R` | `T -> R` 映射 |
| 1.5 | `foreach.go` | `Foreach[T any](values []T, each func(value T, i int))` | 遍历执行副作用 |
| 1.6 | `count.go` | `Count[T comparable](values []T, target T) int` | 统计某值出现次数 |
| 1.7 | `countby.go` | `CountBy[T any](values []T, predicate func(value T, i int) bool) int` | 按谓词计数 |
| 1.8 | `everyby.go` | `EveryBy[T any](values []T, predicate func(value T, i int) bool) bool` | 是否全部满足谓词 |

### 2. 泛型组合

| 编号 | 文件 | 函数签名 | 语义 |
| --- | --- | --- | --- |
| 2.1 | `filtermap.go` | `FilterMap[T any, R any](values []T, mapper func(value T, i int) (R, bool)) []R` | 过滤+映射一次完成 |
| 2.2 | `flatmap.go` | `Flatmap[T any, R any](values []T, mapper func(value T, i int) []R) []R` | 映射后拍平一层 |

约定：所有带下标的回调统一为 `func(value T, i int)`，`i` 为元素下标。

## 示例

### 1. Slice：遍历 / 回调处理类

```go
practise.Contains([]int{1, 2, 3}, 2)
// true

practise.Filter(func(v int, _ int) bool {
    return v%2 == 0
}, []int{1, 2, 3, 4})
// []int{2, 4}

v, ok := practise.Find(func(v int, _ int) bool {
    return v > 2
}, []int{1, 2, 3})
// 3, true

practise.Map([]int{1, 2, 3}, func(v int, _ int) string {
    return strconv.Itoa(v)
})
// []string{"1", "2", "3"}

practise.Count([]int{1, 2, 2, 3}, 2)
// 2

practise.CountBy([]int{1, 2, 3, 4}, func(v int, _ int) bool {
    return v%2 == 0
})
// 2

practise.EveryBy([]int{2, 4, 6}, func(v int, _ int) bool {
    return v%2 == 0
})
// true
```

### 2. 泛型组合

```go
practise.FilterMap([]int{1, 2, 3, 4}, func(v int, _ int) (string, bool) {
    if v%2 != 0 {
        return "", false
    }
    return strconv.Itoa(v), true
})
// []string{"2", "4"}

practise.Flatmap([]int{1, 2, 3}, func(v int, _ int) []string {
    return []string{strconv.Itoa(v), strconv.Itoa(v)}
})
// []string{"1", "1", "2", "2", "3", "3"}
```

## 跑法

```bash
go vet ./...
go test ./...
```

目前无单测，`go vet` 用于保证泛型签名可编译。

## 笔记

- [编写时的笔记（飞书）](https://iqeubg8au73.feishu.cn/wiki/Arkewkep5ibzvpkX6JzcxZVonvb?from=from_copylink)
