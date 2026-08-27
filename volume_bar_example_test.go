package fmtx

import "fmt"

func ExampleVolumeBar() {
	orig := EnableColor
	EnableColor = false // keep the docs readable; color is on by default
	defer func() { EnableColor = orig }()

	fmt.Println(VolumeBar(5, 10, 10))     // int: half full
	fmt.Println(VolumeBar(0, 10, 10))     // int: empty
	fmt.Println(VolumeBar(10, 10, 10))    // int: full
	fmt.Println(VolumeBar(12, 10, 10))    // int: clamped to the width
	fmt.Println(VolumeBar(0.25, 1.0, 10)) // float: a quarter
	fmt.Println(VolumeBar(5, 0, 10))      // invalid max -> empty
	// Output:
	// █████░░░░░
	// ░░░░░░░░░░
	// ██████████
	// ██████████
	// ██░░░░░░░░
	//
}

func ExampleVolumeBar2() {
	orig := EnableColor
	EnableColor = false
	defer func() { EnableColor = orig }()

	fmt.Println(VolumeBar2(3, 3, 10, 10))   // 3 + 3 out of 10
	fmt.Println(VolumeBar2(0, 5, 10, 10))   // only the second segment
	fmt.Println(VolumeBar2(5, 0, 10, 10))   // only the first segment
	fmt.Println(VolumeBar2(10, 10, 10, 10)) // full
	fmt.Println(VolumeBar2(6, 6, 10, 10))   // clamps at the width
	// Output:
	// ███▒▒▒░░░░
	// ▒▒▒▒▒░░░░░
	// █████░░░░░
	// ██████████
	// ██████▒▒▒▒
}

func ExampleRangeVolume() {
	fmt.Println(RangeVolume(20, 60, 0, 100, 10)) // interval 20..60
	fmt.Println(RangeVolume(0, 100, 0, 100, 10)) // whole range
	fmt.Println(RangeVolume(80, 20, 0, 100, 10)) // swapped from/to is fine
	fmt.Println(RangeVolume(50, 50, 0, 100, 10)) // single point -> nothing to show
	// Output:
	// ░░████░░░░
	// ██████████
	// ░░██████░░
	// ░░░░░░░░░░
}

func ExampleRangeVolume2() {
	fmt.Println(RangeVolume2(10, 40, 60, 90, 0, 100, 10)) // two separate ranges
	fmt.Println(RangeVolume2(10, 40, 30, 90, 0, 100, 10)) // overlapping: second wins
	fmt.Println(RangeVolume2(40, 10, 30, 90, 0, 100, 10)) // swapped ranges still work
	// Output:
	// ░███░░▒▒▒░
	// ░██▒▒▒▒▒▒░
	// ░██▒▒▒▒▒▒░
}
