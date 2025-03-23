# TriangleTree

📐 `TriangleTree` is a lightweight Go module that calculates the **maximum path sum** in a triangle of integers. It supports input from either:

- A 2D array (`[][]int`)
- A URL pointing to a JSON file containing a 2D array

---

## 🚀 Features

- ✅ Supports both array and URL input
- ✅ Validates triangle structure
- ✅ Uses bottom-up dynamic programming
- ✅ Suitable for large triangle structures
- ✅ Unit-tested with mocked URLs

---

## 📦 Installation

```bash
go get github.com/andreaharris-go/triangletree
```

---

## 🧠 Usage

### 1. Use with [][]int
```go
package main

import (
	"fmt"
	"github.com/andreaharris-go/triangletree"
)

func main() {
	input := [][]int{
		{59},
		{73, 41},
		{52, 40, 53},
		{26, 53, 6, 34},
	}
	result, err := triangletree.TreeSum(input)
	if err != nil {
		panic(err)
	}
	fmt.Println("Max path sum:", result) // Output: 237
}

```
---
### 2. Use with JSON URL
```go
url := "https://raw.githubusercontent.com/andreaharris-go/triangletree/mock.json"
result, err := triangletree.TreeSum(url)
if err != nil {
	panic(err)
}
fmt.Println("Max path sum:", result) // Output: 7273

```
---
## 📑 JSON Format Example
```json
[
  [59],
  [73, 41],
  [52, 40, 53],
  [26, 53, 6, 34]
]

```
Make sure each row `i` has exactly `i+1` elements to be considered a valid triangle.
---
## 🧪 Running Unit Tests
```shell
go test ./...
```
Includes:
- Valid array input
- Invalid triangle structure
- Valid and invalid URL input
- Mocked HTTP test server
---
## 📜 License

MIT License

---
## 🙋‍♂️ Author

Developed by [Duckdev84](https://linktr.ee/duckdev84)
