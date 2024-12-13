# Go: Out-of-bounds array access in function myFunc

This repository demonstrates a potential out-of-bounds read in Go when accessing an array.  The `myFunc` function calculates the sum of elements in an integer array. However, it doesn't handle cases where the array is `nil` or empty, leading to a potential panic.

**Bug:** The code fails to check for `nil` or empty array before iterating, resulting in a runtime panic if the input `a` is `nil` or empty.

**Solution:** The solution adds a check for `nil` or empty arrays, preventing the out-of-bounds access.  If the array is `nil` or empty, it gracefully returns 0.

**How to run:**
1. Clone this repository.
2. Navigate to the repository's directory.
3. Run `go run bug.go` to observe the panic.
4. Run `go run bugSolution.go` to see the corrected, panic-free version.