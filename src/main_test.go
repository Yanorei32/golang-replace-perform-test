package main

import (
	"regexp"
	"strings"
	"testing"
)

func BenchmarkAppend_Trim(b *testing.B) {
	base_str := "xx	"
	bench_str := ""

	for i := 0; i < 128; i++ {
		bench_str += base_str
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = strings.Trim(bench_str, "	")
	}
}

func BenchmarkAppend_ReplaceAll(b *testing.B) {
	base_str := "xx	"
	bench_str := ""

	for i := 0; i < 128; i++ {
		bench_str += base_str
	}

	b.ResetTimer()


	for i := 0; i < b.N; i++ {
		_ = strings.ReplaceAll(bench_str, "	", "")
	}
}

func BenchmarkAppend_Regex(b *testing.B) {
	base_str := "xx	"
	bench_str := ""

	for i := 0; i < 128; i++ {
		bench_str += base_str
	}

	regObj := regexp.MustCompile("\t")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = regObj.ReplaceAllString(bench_str, "")
	}
}

func BenchmarkAppend_Replacer(b *testing.B) {
	base_str := "xx	"
	bench_str := ""

	for i := 0; i < 128; i++ {
		bench_str += base_str
	}

	replacer := strings.NewReplacer("\t", "")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = replacer.Replace(bench_str)
	}
}

