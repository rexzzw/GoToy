package main

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
)

type CacheAPI struct {
    Redis redis.Conn
}
var (
    CacheV1 = &CacheAPI{}
)

func init() {
    c, err := redis.Dial("tcp", ":6379")
    if err != nil {
        fmt.Printf("make connection to redis failed, err:%+v\n", err)
    }
    CacheV1.Redis = c
}

func main() {
    key := "xxx"
    // reply, err := CacheV1.Redis.Do("SET", key, true)
    // if err != nil {
    //     fmt.Printf("set key to redis failed, err:%+v\n", err)
    //     return
    // }
    // fmt.Printf("set key result: %v\n", reply)

    checkResult, err := CacheV1.Redis.Do("EXISTS", key)
    if err != nil {
        fmt.Printf("check exist key failed, err:%v\n", err)
        return
    }
    if result, ok := checkResult.(int64); ok {
        fmt.Printf("check exist key result:%v\n", result==1)
    }
    return
}
