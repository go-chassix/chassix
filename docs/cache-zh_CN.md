# 缓存
chassis-go 支持memory、redis两种开箱即用的缓存。

cache接口定义如下：
```go
package cache


type Store interface {
	Get(key string) (val interface{}, ok bool)
	Set(key string, val interface{}) (ok bool)
	Delete(key string)
	Contains(key string) bool
}

```
### 基于golang-lru的memory cache使用注意事项
- Set(key,val)val的类型应该与创建缓存时注册的类型一致

```go
package main
import (
    "fmt"
    
    "c6x.io/chassis/cache"
)
func main(){
    var numberType int
    if cs,err:=cache.NewMemoryCacheStore("test",numberType,100);err == nil{
        cs.Set("foo",5)
        if val,ok:= cs.Get("foo");ok {
        	fmt.Printf("val :%d\n",val.(int))
        }
    }
}

```

- NewMemoryCacheStore(name,type,size) size不能<=0时

### Redis cache store使用注意事项

redis cache store 采用set类型存储键值对, 在NewCacheStore的时可为其元素指定过期时间，redis存储时以指定的name为前缀。例如：

```go
package main
import (
    "fmt"

    "c6x.io/chassis/cache"
)
type testT struct {
	A string
	B int
	Inner
}
type Inner struct {
    C string
}
func main (){
    if cs, err :=cache.NewRedisCacheStore("test", &testT{}, 0);err!=nil {
        t := &testT{
            A: "test",
            B: 10,
            Inner: Inner{C: "s"},
        }
        cs.Set("foo",t)
        if val,ok:=cs.Get("foo");ok {
            fmt.Printf("val: %+v\n",val.(*testT))
        }
    }       
}
```
以上代码在redis存储的key为 test:foo
