package main

import (
	"testing"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

func BenchmarkRankFind5From64_2w(b *testing.B) {
	sliceLen := 20000
	stringInSliceLen := 64
	srcLen := 5
	targets := make([]string, 0, sliceLen)
	//st := time.Now()
	for i := 0; i < sliceLen; i++ {
		targets = append(targets, RandStringBytesMaskImprSrc(stringInSliceLen))
	}
	//fmt.Println("len of slice:", len(targets))
	//fmt.Printf("len of slice:%d, len of each string:%d, len of src:%d\n", sliceLen, stringInSliceLen, srcLen)
	//fmt.Println("gen slice time cost:", time.Now().Sub(st))
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		srcStr := RandStringBytesMaskImprSrc(srcLen)
		fuzzy.RankFind(srcStr, targets)
	}
}

func BenchmarkRankFind5From64_2k(b *testing.B) {
	sliceLen := 2000
	stringInSliceLen := 64
	srcLen := 5
	targets := make([]string, 0, sliceLen)
	//st := time.Now()
	for i := 0; i < sliceLen; i++ {
		targets = append(targets, RandStringBytesMaskImprSrc(stringInSliceLen))
	}
	//fmt.Printf("len of slice:%d, len of each string:%d, len of src:%d\n", sliceLen, stringInSliceLen, srcLen)
	//fmt.Println("gen slice time cost:", time.Now().Sub(st))
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		srcStr := RandStringBytesMaskImprSrc(srcLen)
		fuzzy.RankFind(srcStr, targets)
	}
}

func BenchmarkRankFind5From64_200(b *testing.B) {
	sliceLen := 200
	stringInSliceLen := 64
	srcLen := 5
	targets := make([]string, 0, sliceLen)
	//st := time.Now()
	for i := 0; i < sliceLen; i++ {
		targets = append(targets, RandStringBytesMaskImprSrc(stringInSliceLen))
	}
	//fmt.Printf("len of slice:%d, len of each string:%d, len of src:%d\n", sliceLen, stringInSliceLen, srcLen)
	//fmt.Println("gen slice time cost:", time.Now().Sub(st))
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		srcStr := RandStringBytesMaskImprSrc(srcLen)
		fuzzy.RankFind(srcStr, targets)
	}
}

func BenchmarkRankFin16From64_2w(b *testing.B) {
	sliceLen := 20000
	stringInSliceLen := 64
	srcLen := 16
	targets := make([]string, 0, sliceLen)
	//st := time.Now()
	for i := 0; i < sliceLen; i++ {
		targets = append(targets, RandStringBytesMaskImprSrc(stringInSliceLen))
	}
	//fmt.Printf("len of slice:%d, len of each string:%d, len of src:%d\n", sliceLen, stringInSliceLen, srcLen)
	//fmt.Println("gen slice time cost:", time.Now().Sub(st))
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		srcStr := RandStringBytesMaskImprSrc(srcLen)
		fuzzy.RankFind(srcStr, targets)
	}
}

func BenchmarkRankFind16From64_2k(b *testing.B) {
	sliceLen := 2000
	stringInSliceLen := 64
	srcLen := 16
	targets := make([]string, 0, sliceLen)
	//st := time.Now()
	for i := 0; i < sliceLen; i++ {
		targets = append(targets, RandStringBytesMaskImprSrc(stringInSliceLen))
	}
	//fmt.Printf("len of slice:%d, len of each string:%d, len of src:%d\n", sliceLen, stringInSliceLen, srcLen)
	//fmt.Println("gen slice time cost:", time.Now().Sub(st))
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		srcStr := RandStringBytesMaskImprSrc(srcLen)
		fuzzy.RankFind(srcStr, targets)
	}
}

func BenchmarkRankFind16From64_200(b *testing.B) {
	sliceLen := 200
	stringInSliceLen := 64
	srcLen := 16
	targets := make([]string, 0, sliceLen)
	//st := time.Now()
	for i := 0; i < sliceLen; i++ {
		targets = append(targets, RandStringBytesMaskImprSrc(stringInSliceLen))
	}
	//fmt.Printf("len of slice:%d, len of each string:%d, len of src:%d\n", sliceLen, stringInSliceLen, srcLen)
	//fmt.Println("gen slice time cost:", time.Now().Sub(st))
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		srcStr := RandStringBytesMaskImprSrc(srcLen)
		fuzzy.RankFind(srcStr, targets)
	}
}
