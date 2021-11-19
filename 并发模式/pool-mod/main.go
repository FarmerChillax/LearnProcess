package main

import (
	"io"
	"log"
	"math/rand"
	"pool/pool"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines   = 25
	pooledResources = 2
)

type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

var idCounter int32

func CreateConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)
	return &dbConnection{ID: id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// 创建用来管理连接的池
	p, err := pool.New(CreateConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	// 使用资源池中的链接来完成查询
	for query := 0; query < maxGoroutines; query++ {
		//
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()

}

// 用来测试连接的资源池
func performQueries(query int, p *pool.Pool) {
	// 从资源池获取一个连接资源
	conn, err := p.Acquire()

	if err != nil {
		log.Println(err)
		return
	}
	// 查询结束将资源放回资源池
	defer p.Release(conn)

	// 模拟查询响应
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Second)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
