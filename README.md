# practise

Go 泛型小练习：对照 [samber/lo](https://github.com/samber/lo) 复刻常用集合函数，熟悉 `any / comparable` 约束和 `func(value T, index int)` 谓词写法。

## 已实现

| 文件 | 函数签名 | 语义 |
| --- | --- | --- |
| `contain.go` | `Contains[T comparable](list []T, item T) bool` | 是否包含某个值 |
| `filter.go` | `Filter[T any](filter func(value T, i int) bool, values []T) []T` | 按谓词过滤 |
| `find.go` | `Find[T any](find func(value T, i int) bool, values []T) (T, bool)` | 找第一个满足谓词的元素 |
| `map.go` | `Map[T any, R any](values []T, mapper func(value T, i int) R) []R` | `T -> R` 映射 |
| `foreach.go` | `Foreach[T any](values []T, each func(value T, i int))` | 遍历执行副作用 |
| `count.go` | `Count[T comparable](values []T, target T) int` | 统计某值出现次数 |
| `countby.go` | `CountBy[T any](values []T, predicate func(value T, i int) bool) int` | 按谓词计数 |
| `everyby.go` | `EveryBy[T any](values []T, predicate func(value T, i int) bool) bool` | 是否全部满足谓词 |

约定：所有带下标的回调统一为 `func(value T, i int)`，`i` 为元素下标。

## 示例

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

## 跑法

```bash
go vet ./...
go test ./...
```

目前无单测，`go vet` 用于保证泛型签名可编译。

## 路线图

- [x] `Contains` / `Filter`
- [x] `Find` / `Map` / `Foreach`
- [x] `Count` / `CountBy` / `EveryBy`
- [ ] `Reduce` / `Reject` / `Uniq`
- [ ] `GroupBy` / `Keys` / `Values`

下一步建议先写 `Reduce`，它是聚合类函数的通用底座。

## 笔记

- [编写时的笔记（飞书）](https://iqeubg8au73.feishu.cn/wiki/Arkewkep5ibzvpkX6JzcxZVonvb?from=from_copylink)
