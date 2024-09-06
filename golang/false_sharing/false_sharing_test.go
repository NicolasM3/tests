package main

// import (
// 	"sync"
// 	"sync/atomic"
// 	"testing"

// 	"golang.org/x/sys/cpu"
// )

// type AtomicBatchOperations interface {
// 	DummyOperation()
// }

// type NoPad struct {
// 	a uint64
// 	b uint64
// 	c uint64
// }

// func (myatomic *NoPad) DummyOperation() {
// 	atomic.AddUint64(&myatomic.b, 1)
// 	atomic.AddUint64(&myatomic.c, 1)
// 	atomic.AddUint64(&myatomic.a, 1)
// }

// type Pad struct {
// 	a uint64
// 	_ cpu.CacheLinePad
// 	b uint64
// 	_ cpu.CacheLinePad
// 	c uint64
// 	_ cpu.CacheLinePad
// }

// func (myatomic *Pad) DummyOperation() {
// 	atomic.AddUint64(&myatomic.a, 1)
// 	atomic.AddUint64(&myatomic.b, 1)
// 	atomic.AddUint64(&myatomic.c, 1)
// }

// type HybridPad struct {
// 	a uint64
// 	_ [4]uint64
// 	b uint64
// 	c uint64
// 	_ [4]uint64
// }

// func (myatomic *HybridPad) DummyOperation() {
// 	atomic.AddUint64(&myatomic.a, 1)
// 	atomic.AddUint64(&myatomic.b, 1)
// 	atomic.AddUint64(&myatomic.c, 1)
// }

// func testAtomicIncrease(myatomic AtomicBatchOperations) {
// 	paraNum := 50000
// 	addTimes := 100
// 	var wg sync.WaitGroup
// 	wg.Add(paraNum)
// 	for i := 0; i < paraNum; i++ {
// 		go func() {
// 			for j := 0; j < addTimes; j++ {
// 				myatomic.DummyOperation()
// 			}
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()

// }
// func BenchmarkNoPad(b *testing.B) {
// 	myatomic := &NoPad{}
// 	b.ResetTimer()
// 	testAtomicIncrease(myatomic)
// }

// func BenchmarkHybridPad(b *testing.B) {
// 	myatomic := &HybridPad{}
// 	b.ResetTimer()
// 	testAtomicIncrease(myatomic)
// }

// func BenchmarkPad(b *testing.B) {
// 	myatomic := &Pad{}
// 	b.ResetTimer()
// 	testAtomicIncrease(myatomic)
// }
