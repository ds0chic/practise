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
| 2.3 | `reduce.go` | `Reduce[T any, R any](values []T, reducer func(acc R, value T, i int) R, initial R) R` | 聚合为单个值 |
| 2.4 | `uniqby.go` | `UniqBy[T any, K comparable](values []T, key func(value T) K) []T` | 按 key 去重 |
| 2.5 | `groupby.go` | `GroupBy[T any, K comparable](values []T, key func(value T) K) map[K][]T` | 按 key 分组 |

### 3. Slice / Map：分块 / 切分 / 转换

| 编号 | 文件 | 函数签名 | 语义 |
| --- | --- | --- | --- |
| 3.1 | `chunk.go` | `Chunk[T any](values []T, size int) [][]T` | 按 size 切分为多块，size<=0 时 panic |
| 3.2 | `flatten.go` | `Flatten[T any](values [][]T) []T` | 拍平一层 |
| 3.3 | `take.go` | `Take[T any](values []T, count int) []T` | 取前 count 个 |
| 3.4 | `keys.go` | `Keys[K comparable, V any](m map[K]V) []K` | 取 map 所有 key |
| 3.5 | `keyby.go` | `KeyBy[T any, K comparable](values []T, key func(value T) K) map[K]T` | 按 key 转为 map |

### 4. encoding/json

| 编号 | 文件 | 函数签名 | 语义 |
| --- | --- | --- | --- |
| 4.1.1 | `Marshal.go` | `MarshalDemo()` | JSON 序列化 demo |
| 4.1.2 | `MarshalTag.go` | `MarshalTagDemo()` | struct tag 改字段名 demo |
| 4.1.3 | `MarshalOmitEmpty.go` | `MarshalOmitEmptyDemo()` | omitempty 省略零值 demo |
| 4.2.1 | `Unmarshal.go` | `UnmarshalDemo()` | JSON 反序列化 demo |
| 4.2.2 | `UnmarshalSlice.go` | `UnmarshalSliceDemo()` | []User 反序列化 demo |
| 4.2.3 | `UnmarshalMap.go` | `UnmarshalMapDemo()` | []map 反序列化 demo |
| 4.2.4 | `UnmarshalNested.go` | `UnmarshalNestedDemo()` | 嵌套结构反序列化 demo |
| 4.3.1 | `Encoder.go` | `EncoderDemo()` | NewEncoder 流式编码 demo |
| 4.3.2 | `Decoder.go` | `DecoderDemo()` | NewDecoder 流式解码 demo |

### 5. context

| 编号 | 文件 | 函数签名 | 语义 |
| --- | --- | --- | --- |
| 5.1.1 | `Background.go` | `BackgroundDemo()` | Background 空 context demo |
| 5.1.2 | `Cancel.go` | `CancelDemo()` | WithCancel 取消 goroutine demo |
| 5.1.3 | `Timeout.go` | `TimeoutDemo()` | WithTimeout 超时中止 demo |
| 5.1.4 | `Deadline.go` | `DeadlineDemo()` | WithDeadline 到期中止 demo |
| 5.2.1 | `Value.go` | `ValueDemo()` | WithValue 传值 demo |
| 5.2.2 | `ParentChild.go` | `ParentChildDemo()` | 父取消传播到子 demo |
| 5.2.3 | `TaskChain.go` | `TaskChainDemo()` | TaskA 调 TaskB 取消传递 demo |

### 6. net/http

| 编号 | 文件 | 函数签名 | 语义 |
| --- | --- | --- | --- |
| 6.1.1 | `Server.go` | `ServerDemo()` | GET / 返回 Hello World demo |
| 6.2.1 | `User.go` | `UserDemo()` | 打印 Method/Path/UA demo，:8081/user |
| 6.3.1 | `UserStatus.go` | `UserStatusDemo()` | 404/200 状态码 demo，:8082/user |
| 6.4.1 | `UserMethod.go` | `UserMethodDemo()` | GET/POST/405 方法分支 demo，:8083/user |
| 6.5.1 | `UserAuth.go` | `UserAuthDemo()` | 打印 UA/Authorization + 回 JSON demo，:8084/user |
| 6.6.1 | `UserBody.go` | `UserBodyDemo()` | 读 Body 打印 + 回收到 demo，:8085/user |
| 6.7.1 | `UserJSON.go` | `UserJSONDemo()` | POST JSON 解析回显 demo，:8086/user |
| 6.8.1 | `ClientGet.go` | `ClientGetDemo()` | GET 拉页面 + 打印 Status/Body demo |
| 6.8.2 | `ClientTimeout.go` | `ClientTimeoutDemo()` | 带 3s context 超时 GET demo |

### 7. Gin

| 编号 | 文件 | 函数签名 | 语义 |
| --- | --- | --- | --- |
| 7.1.1 | `GinUser.go` | `GinUserDemo()` | GET /user 回 JSON demo，:8087/user |
| 7.1.2 | `GinQuery.go` | `GinQueryDemo()` | c.Query 取参回显 demo，:8088/user |
| 7.1.3 | `GinParam.go` | `GinParamDemo()` | c.Param 路径参数 demo，:8089/user/:id |
| 7.1.4 | `GinBind.go` | `GinBindDemo()` | ShouldBindJSON 绑定 demo，:8090/user |
| 7.2.1 | `GinAuth.go` | `GinAuthDemo()` + `AuthMiddleware` | token 中间件 demo，:8091/user |
| 7.2.2 | `GinGroup.go` | `GinGroupDemo()` | Group 分组 + 中间件复用 demo，:8092/api/... |
| 7.2.3 | `GinUserID.go` | `GinUserIDDemo()` | c.Set/c.Get 传 userID demo，:8093/user/info |

### 8. database/sql

| 编号 | 文件 | 函数签名 | 语义 |
| --- | --- | --- | --- |
| 8.1.1 | `DBPing.go` | `DBPingDemo()` | MySQL Open + Ping demo，需本地库 |
| 8.1.2 | `DBQueryRow.go` | `DBQueryRowDemo()` | QueryRow 单行查询 + Scan demo，需 users 表 |
| 8.1.3 | `DBQuery.go` | `DBQueryDemo()` | Query 全表查询 + Next/Scan 逐行 demo，需 users 表 |
| 8.2.1 | `DBExec.go` | `DBExecDemo()` | 建表 + 新增 + 回 ID/影响行数 demo |

约定：遍历 / 组合类回调统一为 `func(value T, i int)`，`i` 为元素下标；`UniqBy` / `GroupBy` 按设计为 `func(value T) K`（只需按值取 key，无下标）。

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

practise.Reduce([]int{1, 2, 3, 4}, func(acc int, v int, _ int) int {
    return acc + v
}, 0)
// 10

practise.UniqBy([]string{"a", "aa", "aaa", "a", "bb"}, func(s string) int {
    return len(s)
})
// []string{"a", "aa", "aaa"}

practise.GroupBy([]int{1, 2, 3, 4}, func(v int) int {
    return v % 2
})
// map[int][]int{1: {1, 3}, 0: {2, 4}}
```

### 3. Slice / Map：分块 / 切分 / 转换

```go
practise.Chunk([]int{1, 2, 3, 4, 5}, 2)
// [][]int{{1, 2}, {3, 4}, {5}}

practise.Flatten([][]int{{1, 2}, {3, 4}, {5}})
// []int{1, 2, 3, 4, 5}

practise.Take([]int{1, 2, 3, 4}, 2)
// []int{1, 2}

practise.Keys(map[string]int{"a": 1, "b": 2})
// []string{"a", "b"}（顺序随机）

practise.KeyBy([]string{"a", "aa", "aaa"}, func(s string) int {
    return len(s)
})
// map[int]string{1: "a", 2: "aa", 3: "aaa"}
```

### 4. encoding/json

```go
// Marshal.go MarshalDemo() JSON 序列化 demo
// 输出 {"Name":"Tom","Age":22}

// MarshalTag.go MarshalTagDemo() struct tag 改字段名 demo
// 输出 {"your name":"Tom","your age":22}

// MarshalOmitEmpty.go MarshalOmitEmptyDemo() omitempty 省略零值 demo
// 输出 {"your name":"Tom"}

// Unmarshal.go UnmarshalDemo() JSON 反序列化 demo
// 输出 name: Tom / age: 22

// UnmarshalSlice.go UnmarshalSliceDemo() []User 反序列化 demo
// 输出 [{Tom 22} {Bob 25}] / 25

// UnmarshalMap.go UnmarshalMapDemo() []map 反序列化 demo
// 输出 [map[age:22 name:Tom] map[age:25 name:Bob]] / Bob

// UnmarshalNested.go UnmarshalNestedDemo() 嵌套结构反序列化 demo
// 输出 {Tom 22 {Shanghai Nanjing Road}} / Shanghai

// Encoder.go EncoderDemo() NewEncoder 流式编码 demo
// 输出 {"name":"Tom","age":22}

// Decoder.go DecoderDemo() NewDecoder 流式解码 demo
// 输出 {Tom 22} / Tom / 22
```

### 5. context

```go
// Background.go BackgroundDemo() Background 空 context demo
// 输出 context.Background

// Cancel.go CancelDemo() WithCancel 取消 goroutine demo
// 输出 任务运行 x3 / 任务结束

// Timeout.go TimeoutDemo() WithTimeout 超时中止 demo
// 输出 任务运行中 x3 / context deadline exceeded / 任务中止

// Deadline.go DeadlineDemo() WithDeadline 到期中止 demo
// 输出 任务运行中 x3 / context deadline exceeded / 任务中止

// Value.go ValueDemo() WithValue 传值 demo
// 输出 13145

// ParentChild.go ParentChildDemo() 父取消传播到子 demo
// 输出 child canceled / context canceled

// TaskChain.go TaskChainDemo() TaskA 调 TaskB 取消传递 demo
// 输出 TaskA 开始 / TaskB 正在运行 x3 / TaskB 收到取消信号 / context canceled
```

### 6. net/http

```go
// Server.go ServerDemo() GET / 返回 Hello World demo
// 浏览器访问 http://localhost:8080/ 输出 Hello World

// User.go UserDemo() 打印 Method/Path/UA demo
// 浏览器访问 http://localhost:8081/user，控制台输出 Method/Path/UA

// UserStatus.go UserStatusDemo() 404/200 状态码 demo
// 浏览器访问 http://localhost:8082/user 输出 user not found（404）

// UserMethod.go UserMethodDemo() GET/POST/405 方法分支 demo
// 浏览器 GET http://localhost:8083/user 输出 获取 user 信息

// UserAuth.go UserAuthDemo() 打印 UA/Authorization + 回 JSON demo
// 浏览器访问 http://localhost:8084/user 输出 {"status":"ok"}，控制台输出 UA/token

// UserBody.go UserBodyDemo() 读 Body 打印 + 回收到 demo
// POST body 到 http://localhost:8085/user，控制台打印 body，页面输出 收到数据

// UserJSON.go UserJSONDemo() POST JSON 解析回显 demo
// POST {"name":"Tom","age":22} 到 http://localhost:8086/user，原样回显 JSON，错 JSON 回 400

// ClientGet.go ClientGetDemo() GET 拉页面 + 打印 Status/Body demo
// go run ./cmd/demo clientget，控制台输出 Status: 200 OK + example.com 页面

// ClientTimeout.go ClientTimeoutDemo() 带 3s context 超时 GET demo
// go run ./cmd/demo clienttimeout，正常输出页面，超 3s 则打印 context deadline exceeded
```

### 7. Gin

```go
// GinUser.go GinUserDemo() GET /user 回 JSON demo
// 浏览器访问 http://localhost:8087/user 输出 {"age":22,"name":"tom"}
// 需 gin 依赖：go get github.com/gin-gonic/gin（已进 go.mod/go.sum）

// GinQuery.go GinQueryDemo() c.Query 取参回显 demo
// 浏览器访问 http://localhost:8088/user?name=tom&age=22 输出 {"age":"22","name":"tom"}

// GinParam.go GinParamDemo() c.Param 路径参数 demo
// 浏览器访问 http://localhost:8089/user/1 输出 {"user_id":"1"}

// GinBind.go GinBindDemo() ShouldBindJSON 绑定 demo
// POST {"name":"tom","age":"22"} 到 http://localhost:8090/user 回显，错格式回 400 {"error":"格式错误"}

// GinAuth.go GinAuthDemo() token 中间件 demo
// 不带/错 token 访问 http://localhost:8091/user 回 401 {"error":"unauthorized"}，带 Authorization: 123456 回 Tom/22

// GinGroup.go GinGroupDemo() Group 分组 + 中间件复用 demo
// POST http://localhost:8092/api/public/login 回 login success；GET /api/user/info（带 token）回 Tom/22，不带回 401

// GinUserID.go GinUserIDDemo() c.Set/c.Get 传 userID demo
// GET http://localhost:8093/user/info（带 Authorization: 123456）回 {"name":"Tom","user_id":1001}，不带回 401
```

### 8. database/sql

```go
// DBPing.go DBPingDemo() MySQL Open + Ping demo
// go run ./cmd/demo dbping，库通了输出 数据库连接成功，否则 数据库连接失败 + err
// 需 mysql 驱动：go get github.com/go-sql-driver/mysql（已进 go.mod/go.sum）

// DBQueryRow.go DBQueryRowDemo() QueryRow 单行查询 + Scan demo
// go run ./cmd/demo dbqueryrow，需 users 表有 id=1，输出 id name age（如 1 Tom 22)

// DBQuery.go DBQueryDemo() Query 全表查询 + Next/Scan 逐行 demo
// go run ./cmd/demo dbquery，逐行输出全表 id name age

// DBExec.go DBExecDemo() 建表 + 新增 + 回 ID/影响行数 demo
// go run ./cmd/demo dbexec，输出 新增用户ID + 影响行数（无 users 表会自动建）
```
```

## 跑法

```bash
go vet ./...
```

`go vet` 保证签名可编译（练习记录，不写单测）。

### 本地看 http 网页实际输出

```bash
go run ./cmd/demo server      # :8080 / -> Hello World
go run ./cmd/demo user        # :8081 /user -> request received
go run ./cmd/demo userstatus  # :8082 /user -> user not found（404）
go run ./cmd/demo usermethod  # :8083 /user -> 获取 user 信息（GET）
go run ./cmd/demo userauth    # :8084 /user -> {"status":"ok"}
go run ./cmd/demo userbody    # :8085 /user -> 收到数据（POST 带 body）
go run ./cmd/demo userjson    # :8086 /user -> 回显 POST 的 JSON
go run ./cmd/demo clientget   # 控制台输出 Status + example.com 页面（client 类不用浏览器）
go run ./cmd/demo clienttimeout  # 同上，但请求带 3s context 超时
go run ./cmd/demo ginuser     # :8087 /user -> {"age":22,"name":"tom"}
go run ./cmd/demo ginquery    # :8088 /user?name=tom&age=22 -> 回显 query 参数
go run ./cmd/demo ginparam    # :8089 /user/1 -> {"user_id":"1"}
go run ./cmd/demo ginbind     # POST JSON 到 :8090 /user -> 绑定回显，错格式 400
go run ./cmd/demo ginauth     # :8091 /user -> 带正确 token 回 Tom/22，否则 401
go run ./cmd/demo gingroup    # :8092 /api/public/login + /api/user/info|order
go run ./cmd/demo ginuserid   # :8093 /user/info 带 token -> user_id 1001，否则 401
go run ./cmd/demo dbping      # 控制台输出 数据库连接成功/失败（需本地 MySQL 3306 有 test 库）
go run ./cmd/demo dbqueryrow  # 控制台输出 id name age（需 users 表有 id=1）
go run ./cmd/demo dbquery     # 控制台逐行输出全表 id name age
go run ./cmd/demo dbexec      # 控制台输出 新增用户ID + 影响行数
```

浏览器打开对应地址即可，每加一个 server Demo 就在 `cmd/demo/main.go` 加一行 case。

## 笔记

- [编写时的笔记（飞书）](https://iqeubg8au73.feishu.cn/wiki/Arkewkep5ibzvpkX6JzcxZVonvb?from=from_copylink)
