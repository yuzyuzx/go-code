package main

import (
	"fmt"
)

/**
 * 引数番目までのフィボナッチ数列を配列で返す
 */
func fibonacciArray(n int) []int {
  result := []int{0, 1}

  for i := 2; i <= n; i++ {
    result = append(result, result[i - 2] + result[i - 1])
  }

  return result[0:n + 1]
}

/**
 * 引数番目のフィボナッチ数を返す
 */
func fibonacci(n int) int {
  if n == 0 {
    return 0;
  }

  a, b := 0, 1
  for i:= 2; i <= n; i++ {
    tmp := a
    a = b
    b += tmp
  }

  return b
}

func main() {
  number := 20

  for i := 0; i <= number; i++ {
    fmt.Printf("number:%d\n", i)

    result := fibonacciArray(i);
    fmt.Println(result)

    fibonacciResult := fibonacci(i)
    fmt.Println(fibonacciResult)

    fmt.Println()
  }
}
