package main

import "fmt"

// func Add(i *int) {
// 	*i = *i * 5

// }

// ポインタQ4.関数を用いて、変数numに5をかけてください。

func Add(i *int) {
	*i = *i * 5
}

// func Add2(i2 *string) {
// 	*i2 = *i2 + "Good"
// }

func main() {

	// Q1.値に100を持つ変数nを定義して、nのアドレスを表示してください。
	var n int = 100
	fmt.Println(&n)

	// Q2.値をnのアドレスとしたポインタ型変数pを定義し、pのアドレスとpの値を表示してください。
	var p *int = &n
	fmt.Println(p)  // アドレス
	fmt.Println(*p) // 値

	// Q3.関数を用いずに、変数nとpの値を同時に200へ変更し、nとpの値を表示してください。
	n = 200
	fmt.Println(n)
	fmt.Println(*p)

	// Q4.関数を用いて、値に100を持つ変数numに5をかけてください。
	var num int = 100
	Add(&num)
	fmt.Println(num)

	/*
		// Q1.値に100を持つ変数nを定義して、nのアドレスを表示してください。
		// var n int = 100
		// fmt.Println(&n)

		// var s string = "HEllo"
		// fmt.Println(&s)

		// Q2.値をnのアドレスとしたポインタ型変数pを定義し、pのアドレスとpの値を表示してください。
		/*
			var p *int = &n
			fmt.Println(p)  // pのアドレス
			fmt.Println(*p) // pの値

		var pt *string = &s
		fmt.Println(pt)
		fmt.Println(*pt)
	*/

	// Q3.関数を用いずに、変数nとpの値を同時に200へ変更し、nとpの値を表示してください。
	/*
			n = 200
			fmt.Println(n)
			fmt.Println(*p)


		s = "Go"
		fmt.Println(s)
		fmt.Println(*pt)
	*/

	// Q4.関数を用いて、変数numに5をかけてください。
	/*
		var num int = 100
		Add(&num)
		fmt.Println(num)
	*/

	/*
		Add2(&s)
		fmt.Println(s)

		Add2(pt)
		fmt.Println(*pt)
	*/

}
