package main

import (
	"container/list"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	typeSuccess int = 1
	typeFail    int = 2
)

//指标
type metrics struct {
	success int64
	fail    int64
}

//滑动窗口
type SlidingWindow struct {
	bucket   int                //桶数
	curKey   int64              //当前key
	m        map[int64]*metrics //统计
	dataList *list.List
	sync.RWMutex
}

//创建滑动窗口
func NewSlidingWindow(bucket int) *SlidingWindow {
	sw := &SlidingWindow{}
	sw.bucket = bucket
	sw.dataList = list.New()
	return sw
}

//统计成功
func (sw *SlidingWindow) AddSuccess() {
	sw.incr(typeSuccess)
}

//统计失败
func (sw *SlidingWindow) AddFail() {
	sw.incr(typeFail)
}

//自增操作
func (sw *SlidingWindow) incr(t int) {
	sw.Lock()
	defer sw.Unlock()
	nowTime := time.Now().Unix()
	if _, ok := sw.m[nowTime]; !ok {
		sw.m = make(map[int64]*metrics)
		sw.m[nowTime] = &metrics{}
	}
	if sw.curKey == 0 {
		sw.curKey = nowTime
	}
	//一秒一个bucket
	if sw.curKey != nowTime {
		sw.dataList.PushBack(sw.m[nowTime])
		delete(sw.m, sw.curKey)
		sw.curKey = nowTime

		if sw.dataList.Len() > sw.bucket {
			for i := 0; i <= sw.dataList.Len()-sw.bucket; i++ {
				sw.dataList.Remove(sw.dataList.Front())
			}
		}
	}

	switch t {
	case typeSuccess:
		sw.m[nowTime].success++
	case typeFail:
		sw.m[nowTime].fail++
	default:
		log.Fatal("err type")
	}
}

//获取数据长度
func (sw *SlidingWindow) Len() int {
	return sw.dataList.Len()
}

//获取数据(space 如：5、10秒)
func (sw *SlidingWindow) Data(space int) []*metrics {
	sw.RLock()
	defer sw.RUnlock()
	var data []*metrics
	var num = 0
	var m = &metrics{}
	for i := sw.dataList.Front(); i != nil; i = i.Next() {
		one := i.Value.(*metrics)
		m.success += one.success
		m.fail += one.fail
		if num%space == 0 {
			data = append(data, m)
			m = &metrics{} //重置m
		}
		num++
	}
	return data
}

func main1() {
	rand.Seed(time.Now().UnixNano())
	sw := NewSlidingWindow(100)
	var r int
	for i := 0; i < 2000; i++ {
		r = rand.Intn(3)
		if r == 1 {
			sw.AddSuccess()
		} else {
			sw.AddFail()
		}
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
	}

	fmt.Println("sw dataList len:", sw.Len())
	for _, item := range sw.Data(3) {
		log.Printf("success: %v, fail:%v\n", item.success, item.fail)
	}
}

func main() {

	//fmt.Println(time.Now().UnixNano())
	//now := time.Now().UnixNano() / 1e6
	//fmt.Println(now)

	t := time.Now()
	fmt.Println("t:", t)
	t1 := t.Add(-time.Second * 5).Add(time.Second * 20)
	fmt.Println("t:", t1)
}
