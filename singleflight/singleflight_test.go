package singleflight

import (
	"golang.org/x/sync/singleflight"
	"strconv"
	"sync"
	"testing"
	"time"
)

type user struct {
	id       int
	name     string
	password string
	email    string
	token    string
}

func getUserByID(id int) user {
	// 模拟数据库查询耗时
	time.Sleep(time.Millisecond)
	return user{}
}

func BenchmarkBufferWithPool(b *testing.B) {
	var wg sync.WaitGroup

	for n := 0; n < b.N; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = getUserByID(1024)
		}()
	}

	wg.Wait()
}

// 使用 singleflight 原语进行并发限制
func getUserByID2(sg *singleflight.Group, id int) user {
	// 使用 id 作为 key
	v, _, _ := sg.Do(strconv.Itoa(id), func() (interface{}, error) {
		// 模拟数据库查询耗时
		time.Sleep(time.Millisecond)
		return user{}, nil
	})
	return v.(user)
}

func BenchmarkBufferWithPool2(b *testing.B) {
	var wg sync.WaitGroup
	var sg singleflight.Group

	for n := 0; n < b.N; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = getUserByID2(&sg, 1024)
		}()
	}

	wg.Wait()
}
