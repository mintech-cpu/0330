package main

import (
	"fmt"
)

// Q10.引数にinterface型のaをとる関数interを作成し、aのデータ型が数値型であれば「int」文字列型であれば「string」それ以外には「?」を表示するswitch文を作成してください。
func inter(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("?")
	}
}

func main() {
	// Q1.forループで0~9を表示してください。
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// Q2.forループ0~9で、実行結果を「0 1 2 4 5」にしてください。
	for i2 := 0; i2 < 10; i2++ {
		if i2 == 3 {
			continue
		} else if i2 == 6 {
			break
		} else {
			fmt.Println(i2)
		}
	}
	// Q3.1~5の値を持つ配列をforループで回し、1~5を表示してください。
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(i)
	}
	// Q4.「A B C」の値を持つ配列arrを作り、インデックス番号とともに表示してください。
	arr2 := [...]string{"A", "B", "C"}
	for i, v := range arr2 {
		fmt.Println(i, v)
	}
	// Q5.「A B C」の値を持つスライスslを作り、インデックス番号とともに表示してください。
	sl := []string{"A", "B", "C"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	// Q6.「apple 100」の値を持つマップmを作り、実行結果を「apple 100」となるようにしてください。
	m := map[string]int{"apple": 100}
	for v, k := range m {
		fmt.Println(v, k)
	}

	// Q7.0~19才をyoung、20~65才未満をadult、65才以上をseniorとするswitch文を作成し、30才はどれに当てはまるか表示してください。

	age := 16
	switch {
	case 0 <= age && age < 20:
		fmt.Println("young")
	case 20 <= age && age < 65:
		fmt.Println("adult")
	default:
		fmt.Println("senior")
	}

	// Q8.interface型の変数xを作成して、値を100としてください。
	var x interface{} = 100
	fmt.Println(x)

	// Q9. Q8の変数xのデータ型が、数値型であれば「int」文字列型であれば「string」それ以外には「?」を表示するswitch文を作成してください。
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("?")
	}
	fmt.Println(x)

	// Q11. Q10の関数をmain関数で実行してください。引数には100を入れます。
	inter(100)
	inter("hello")
	inter(3.14)
}
