func fizzbuzz(n int) {
	for i := 0; i <= n; i++ {
		var say string
		if i%3 == 0 {
			say += "Fizz"
		}
		if i%5 == 0 {
			say += "Buzz"
		}
		if say != "" {
			fmt.Println(say)
		} else {
			fmt.Println(fmt.Sprintf("%d", i))
		}
	}
}
