func myFunc(a []int) int {
  if a == nil || len(a) == 0 {
    return 0
  }
  sum := 0
  for i := 0; i < len(a); i++ {
    sum += a[i]
  }
  return sum
} 
