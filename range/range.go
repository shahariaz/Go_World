package main

func main() {
	nums := []int{6, 7, 8}
	//
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}
	for _, num := range nums {
		println(num)
	}
	m := map[string]string{"name": "golang", "age": "25"}
	for key, value := range m {
		println(key, value)
	}
	// i = starting byte of rune and c = unicode
	for i, c := range "golang" {
		println(i, string(c))
	}
}