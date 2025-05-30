package sort

import (
	"context"
	"fmt"
	rand2 "math/rand"
	"reflect"
	"runtime"
	"slices"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

func TestSort(t *testing.T) {
	arrSize := 10000000
	arr := make([]int, arrSize)
	wg := &sync.WaitGroup{}
	gcCount := runtime.NumCPU()
	fmt.Println("gcCount:", gcCount)
	for i := 0; i < gcCount; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			r := rand2.New(rand2.NewSource(gtime.Now().UnixNano() + int64(index)))
			for j := index * arrSize / gcCount; j < (index+1)*arrSize/gcCount; j++ {
				arr[j] = r.Intn(1000000)
			}
		}(i)
	}
	wg.Wait()

	//fmt.Println("arr", arr)

	ctx, cancel := context.WithTimeout(context.Background(), 60*gtime.S)
	defer cancel()
	//
	arr0 := slices.Clone(arr)
	//fmt.Println("before sort:", arr0)
	wg.Add(1)
	go MeasureExecutionTime(ctx, func(ctx context.Context, ints []int, b bool) {
		sort.Ints(ints)
	}, arr0, true, wg)

	arr1 := slices.Clone(arr)
	//fmt.Println("before sort:", arr1)
	wg.Add(1)
	go MeasureExecutionTime(ctx, InsertSort, arr1, true, wg)

	arr2 := slices.Clone(arr)
	//fmt.Println("before sort:", arr2)
	wg.Add(1)
	go MeasureExecutionTime(ctx, SelectSort, arr2, true, wg)

	//arr22 := slices.Clone(arr)
	////fmt.Println("before sort:", arr2)
	//wg.Add(1)
	//go MeasureExecutionTime(SelectSort, arr22, false, wg)

	arr3 := slices.Clone(arr)
	//fmt.Println("before sort:", arr3)
	wg.Add(1)
	go MeasureExecutionTime(ctx, BubbleSort, arr3, true, wg)

	arr4 := slices.Clone(arr)
	//fmt.Println("before sort:", arr4)
	wg.Add(1)
	go MeasureExecutionTime(ctx, BubbleSort2, arr4, true, wg)

	arr5 := slices.Clone(arr)
	//fmt.Println("before sort:", arr5)
	wg.Add(1)
	go MeasureExecutionTime(ctx, ShellSort, arr5, true, wg)

	arr6 := slices.Clone(arr)
	//fmt.Println("before sort:", arr6)
	wg.Add(1)
	go MeasureExecutionTime(ctx, QuickSort, arr6, false, wg)

	arr7 := slices.Clone(arr)
	wg.Add(1)
	go MeasureExecutionTime(ctx, QuickSort2, arr7, false, wg)

	arr8 := slices.Clone(arr)
	wg.Add(1)
	go MeasureExecutionTime(ctx, QuickSort3, arr8, false, wg)

	done := make(chan struct{})
	go func() {
		wg.Wait() // 等待所有协程完成
		done <- struct{}{}
	}()

	startTime := time.Now()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	quitSelect := false
	for !quitSelect {
		select {
		case <-ticker.C:
			fmt.Println("process running..., elapsed time:", time.Since(startTime).String())
		case <-done:
			quitSelect = true
		}
	}

	arr9 := slices.Clone(arr)
	wg.Add(1)
	ctx2, cancel2 := context.WithTimeout(context.Background(), 60*gtime.S)
	defer cancel2()
	go MeasureExecutionTime(ctx2, ParallelSort, arr9, false, wg)
	startTime = time.Now()
	go func() {
		wg.Wait() // 等待所有协程完成
		done <- struct{}{}
	}()
	var once sync.Once
	for {
		select {
		case <-ticker.C:
			fmt.Println("parallelSort process running..., elapsed time:", time.Since(startTime).String())
		case <-done:
			once.Do(func() {
				close(done)
			})
			return
		case <-ctx2.Done():
			once.Do(func() {
				close(done)
			})
			return
		}
	}

	//fmt.Println("arr2:", arr2)
}

// MeasureExecutionTime 用于测量排序函数的执行时间
func MeasureExecutionTime(ctx context.Context, sortFunc func(context.Context, []int, bool), arr []int, asc bool, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		arr = nil
		runtime.GC()
	}()

	//	获取函数名
	fullFuncName := runtime.FuncForPC(reflect.ValueOf(sortFunc).Pointer()).Name()
	funcName := fullFuncName[strings.LastIndex(fullFuncName, ".")+1:] // 提取最后的函数名

	startTime, endTime := &gtime.Time{}, &gtime.Time{}

	done := make(chan struct{})

	go func() {
		defer close(done)
		startTime = gtime.Now()
		sortFunc(ctx, arr, asc)
		endTime = gtime.Now()
	}()
	select {
	case <-ctx.Done():
		fmt.Printf("sort func: %s,context done\n", funcName)
	case <-done:
		fmt.Printf("sort func: %s,Execution time: %s\n", funcName, endTime.Sub(startTime).String())
	}

	//fmt.Printf("sort func: %s,asc: %v,result:%v\n", funcName, asc, arr)
}
