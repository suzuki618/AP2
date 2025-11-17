package main

import (
	"fmt"
	"math/rand"
	"time"
)

// randomInts は 0〜99 の乱数を n 個生成して返します。
// main 側で作った *rand.Rand を渡すことで再現性とシードの扱いを分離します。
func randomInts(n int, rnd *rand.Rand) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = rnd.Intn(100)
	}
	return s
}

// countOccurrences はスライス内の値ごとの出現回数を返します。
func countOccurrences(a []int) map[int]int {
	m := make(map[int]int)
	for _, v := range a {
		m[v]++
	}
	return m
}

// containsAll は randoms が input に含まれる各値を必要な回数満たしているかをチェックします。
// input に同じ値が複数ある場合、その回数分含まれていることを要求します。
func containsAll(randoms, input []int) bool {
	need := countOccurrences(input)
	have := countOccurrences(randoms)
	for v, req := range need {
		if have[v] < req {
			return false
		}
	}
	return true
}

func main() {
	// ローカル乱数生成器を作成（Seed を明示的に渡す）
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 問題2: 関数を使って randoms と input を生成
	randoms := randomInts(10, rnd)
	input := randomInts(3, rnd)

	// 初期の配列を表示
	fmt.Println("initial randoms:", randoms)
	fmt.Println("input:", input)

	// 問題1相当: input の各値が randoms に何個含まれているかを出力
	counts := countOccurrences(randoms)
	totalMatches := 0
	for _, v := range input {
		c := counts[v]
		fmt.Printf("value %d appears %d time(s) in randoms\n", v, c)
		totalMatches += c
	}
	fmt.Printf("total matches: %d\n", totalMatches)

	// 問題3: input の数字が全部含まれるまで繰り返す
	// 繰り返しでは、input に使われている値を必要な回数分保持し、
	// その他の要素をランダムに置き換えていく。
	maxIterations := 100000
	iterations := 0
	for !containsAll(randoms, input) {
		iterations++
		if iterations > maxIterations {
			fmt.Printf("stopped after %d iterations (gave up)\n", iterations)
			break
		}

		need := countOccurrences(input)
		// 保持済みカウント
		kept := make(map[int]int)

		// 各インデックスについて、まだ必要な値なら保持し、そうでなければ新しい乱数を入れる
		for i, v := range randoms {
			if need[v] > 0 && kept[v] < need[v] {
				// この値は input で必要なので保持
				kept[v]++
				continue
			}
			// 置換
			randoms[i] = rnd.Intn(100)
		}
	}

	fmt.Println("final randoms:", randoms)
	fmt.Printf("iterations until contains all input values: %d\n", iterations)

	// 最後に最終状態での一致数を出力
	finalCounts := countOccurrences(randoms)
	finalTotal := 0
	for _, v := range input {
		c := finalCounts[v]
		fmt.Printf("(final) value %d appears %d time(s) in randoms\n", v, c)
		finalTotal += c
	}
	fmt.Printf("(final) total matches: %d\n", finalTotal)
}

