package main

// スライスQ8.slに要素[6 7]を追加してください。

import "fmt"

func main() {
	// Q1.実行結果が[1 2 3 4 5]となるスライスslを作成し、表示してください。
	sl := []int{1, 2, 3, 4, 5}
	fmt.Println(sl)

	// Q2.slの最初の要素を取得してください。
	fmt.Println(sl[0])

	// Q3.slの2番目から3番目の要素を取得してください。
	fmt.Println(sl[1:3])
	// Q4.最後の要素を取得してください。
	fmt.Println(sl[len(sl)-1])
	// Q5.slの最初と最後の要素以外すべてを取得してください。
	fmt.Println(sl[1 : len(sl)-1])

	// Q6.slの要素数を調べてください。
	fmt.Println(len(sl))
	// Q7.slのキャパシティ(容量)を調べてください。
	fmt.Println(cap(sl))
	// Q8.slに要素[6 7]を追加してください。
	fmt.Println(append(sl, 6, 7))

	/*
		// Q1.実行結果が[1 2 3 4 5]となるスライスslを作成し、表示してください。
		sl := []int{1, 2, 3, 4, 5}
		fmt.Println(sl)
		// 実行結果が[1 2 3 4 5]となるスライスslがあります。
		// Q2.slの最初の要素を取得してください。
		fmt.Println(sl[0])
		// Q3.slの2番目から3番目の要素を取得してください。
		fmt.Println(sl[1:3])
		// Q4.最後の要素を取得してください。
		fmt.Println(sl[len(sl)-1])
		// Q5.slの最初と最後の要素以外すべてを取得してください。
		fmt.Println(sl[1 : len(sl)-1])

		// Q6.slの要素数を調べてください。
		fmt.Println(len(sl))
		// Q7.slのキャパシティ(容量)を調べてください。
		fmt.Println(cap(sl))

		// Q8.slに要素[6 7]を追加してください。
		fmt.Println(append(sl, 6, 7))

		// Q1.実行結果が[1 2 3 4 5]となるスライスslを作成してください。
		sl2 := []string{"A", "B", "C"}

		// Q2.slの最初の要素を取得してください。
		fmt.Println(sl2[0])

		// Q3.slの2番目から3番目の要素を取得してください。
		fmt.Println(sl2[1:3])

		// Q4.最後の要素を取得してください。
		fmt.Println(sl2[len(sl2)-1])

		// Q5.slの最初と最後の要素以外すべてを取得してください。
		fmt.Println(sl2[1 : len(sl2)-1])

		// Q6.slの要素数を調べてください。
		fmt.Println(len(sl2))
		// Q7.slのキャパシティ(容量)を調べてください。
		fmt.Println(cap(sl2))
		// Q8.slに要素[6 7]を追加してください。
		fmt.Println(append(sl2, "D", "E"))

	*/

}
