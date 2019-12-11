package quickperm

import "fmt"

func Permutate(a []int, c chan []int) {
   N := len(a)
   p := make([]int, N+1)
   for i := 0; i < N+1; i++ {
      p[i] = i
   }
   i := 0
   for i < N {
      p[i]--
      j := i % 2 * p[i]
      a[i], a[j] = a[j], a[i]
      perm := make([]int, len(a))
      copy(perm, a)
      c <- perm
      i = 1
      for p[i] == 0 {
         p[i] = i
         i++
      }
   }
   close(c)
}

func main() {
   var a = []int{1, 2, 3, 4, 5}
   c := make(chan []int)
   go Permutate(a, c)
   for perm := range c {
      fmt.Println(perm)
   }
}
