func myFunc(a []int) int {
  sum := 0
  for i := 0; i < len(a); i++ {
    sum += a[i]
  }
  return sum
}