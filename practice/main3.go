package main

// プラスα. 「Hello Go」を「Hello go」に変更
// プラスα. sの文字列に「Hello」が含まれているか？

import (
	"fmt"
	"strings"
)

func main() {
	// Q1. 3 + 1 の答えを表示
	// Q2. 3 - 1 の答えを表示
	// Q3. 3 × 1 の答えを表示
	// Q4. 3 / 1 の答えを表示
	// Q5. 30 / 4 の余りを表示
	fmt.Println(3 + 1)
	fmt.Println(3 - 1)
	fmt.Println(3 * 1)
	fmt.Println(3 / 1)
	fmt.Println(3 % 1)

	// Q6. 実行結果に「3 + 1 = 4」を表示
	// Q7. 実行結果に「3 - 1 = 2」を表示
	// Q8. 実行結果に「3 × 1 = 3」を表示
	// Q9. 実行結果に「3 / 1 = 3」を表示
	fmt.Println("3 + 1 =", 4)
	fmt.Println("3 - 1 =", 2)
	fmt.Println("3 × 1 =", 3)
	fmt.Println("3 / 1 =", 3)

	// Q10.「Hello Go」を表示
	fmt.Println("Hello Go")

	// プラスα. 「Hello Go」2番目の文字「e」を表示
	var n string = "Hello Go"
	fmt.Println(string(n[1]))
	// プラスα. 「Hello Go」を「Hello go」に変更
	fmt.Println(strings.Replace(n, "G", "g", 1))
	// プラスα. sの文字列に「Hello」が含まれているか？
	fmt.Println(strings.Contains(n, "Hello"))

}
