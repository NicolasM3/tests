package main

import (
	"sync"
	"testing"

	"golang.org/x/sys/cpu"
)

type Operation interface {
	DummyOperation(wg *sync.WaitGroup)
}

type Result struct {
	sumA int64
	sumB int64
}

type PaddedResult struct {
	sumA int64
	_    cpu.CacheLinePad
	sumB int64
}

func genericCalc(op Operation) {
	var wg sync.WaitGroup
	wg.Add(2)
	op.DummyOperation(&wg)
	wg.Wait()
}

func (r *Result) DummyOperation(wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 1000000; i++ {
			r.sumA += 1 // Calcula sumA
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000000; i++ {
			r.sumB += 1 // Computes sumB
		}
		wg.Done()
	}()
}

func (r *PaddedResult) DummyOperation(wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 1000000; i++ {
			r.sumA += 1 // Computes sumA
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000000; i++ {
			r.sumB += 1 // Computes sumB
		}
		wg.Done()
	}()
}

func BenchmarkNoPad(b *testing.B) {
	b.ResetTimer()
	result := Result{}
	genericCalc(&result)
}

func BenchmarkPad(b *testing.B) {
	b.ResetTimer()
	result := PaddedResult{}
	genericCalc(&result)
}
