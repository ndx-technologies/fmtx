package fmtx

import "fmt"

func ExampleSpark() {
	fmt.Println(Spark(map[int]int{0: 1, 1: 3, 2: 6, 3: 8, 4: 4, 5: 2}, 6))
	fmt.Println(Spark(map[int]int{0: 1, 1: 1, 2: 1, 3: 8, 4: 1, 5: 1}, 6)) // one peak
	// Output:
	// ▁▃▆█▄▂
	// ▁▁▁█▁▁
}

func ExampleSparkLine() {
	fmt.Println(SparkLine([]int{1, 2, 3, 4, 5, 6, 7, 8}, 8, 8)) // one cell per level
	fmt.Println(SparkLine([]int{2, 8}, 8, 5))                   // two values stretched
	fmt.Println(SparkLine([]int{1, 3, 5, 7, 9}, 9, 5))          // lowest maps to a space
	// Output:
	// ▁▂▃▄▅▆▇█
	// ▂▂▂██
	//  ▂▄▆█
}
