package weather

import "fmt"

// https://pkg.go.dev/testing#hdr-Examples
func ExampleToS() {
	w := new(Weather)
	w.Area = "Tokyo"
	w.HeadLine = "Good weather"
	w.Body = "You will be able to dry your laundry!"

	fmt.Println(w.ToS())
	// Output:
	// Tokyoの天気です。
	// Good weather
	// You will be able to dry your laundry!
}
