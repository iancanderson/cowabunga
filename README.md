# cowabunga

Handy slice functions using go generics

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

bigNumbers := Filter(ints, func(n int) bool { return n > 2 })
// [3]

firstNum := First(ints)
// 1

hasATwo := IsMember(ints, 2)
// true

lastNum := Last(ints)
// 3

strings := Map(ints, func(n int) string { return strconv.Itoa(n + 1) })
// ["2", "3", "4"]
```
