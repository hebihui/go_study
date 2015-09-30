package main

// #include <stdio.h>
// int f()
// {
//     int n,i,j;
//     scanf("%d",&n);
//     for(i = 2, j = n/2; i <= j;)
//     {
//         if(n % i == 0)
//         {
//             printf("%d ", i);
//             n /= i;
//         }
//         else{
//           ++i;
//         }
//     }
//     if(n != 1)
//         printf("%d\n", n);
// }
import (
	"fmt"
)

func main() {
	n := 821
	i := 2
	for j := n / 2; i <= j; {
		if n%i == 0 {
			fmt.Println(i)
			n /= i
		} else {
			i += 1
		}
	}
	if n != 1 {
		fmt.Println(n)
	}
}
