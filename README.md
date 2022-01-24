# cowabunga

Handy slice 🍕 functions using go generics

## Usage

```go
import "github.com/iancanderson/cowabunga"

ints := []int{1, 2, 3}

allAreBig := All(ints, func(n int) bool { return n > 2 })
// false

anyAreBig := Any(ints, func(n int) bool { return n > 2 })
// true

numBig := Count(ints, func(n int) bool { return n > 2 })
// 1

lastTwo := Drop(ints, 1)
// [2, 3]

lastOne := DropWhile(ints, func(n int) bool { return n < 3 })
// [3]

EachCons(ints, 2, func(nums []int) {
	fmt.Println(nums)
})
// [1, 2]
// [2, 3]

EachSlice(ints, 2, func(nums []int) {
	fmt.Println(nums)
})
// [1, 2]
// [3]

sum := 0
sum = EachWith(ints, &sum, func(num int, result *int) {
	*result += n
})
// 6

bigNumbers := Filter(ints, func(n int) bool { return n > 2 })
// [3]

bigNumbersIncremented := FilterMap(ints, func(n int) *string {
	if n > 2 {
		str := strconv.Itoa(n + 1)
		return &str
	} else {
		return nil
	}
})
// ["4"]

firstGreaterThanOne := Find(ints, func(n int) bool { return n > 1 })
// [&2]

firstNum := First(ints)
// 1

hasATwo := IsMember(ints, 2)
// true

lastNum := Last(ints)
// 3

strings := Map(ints, func(n int) string { return strconv.Itoa(n + 1) })
// ["2", "3", "4"]

firstTwo := Take(ints, 2)
// [1, 2]
```
