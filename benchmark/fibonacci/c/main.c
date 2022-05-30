#include <stdio.h>

long long fibonacci(long long i);

//./program  56.03s user 0.11s system 98% cpu 56.863 total
int main() {
    // for (int i = 0; i < 50; i++) {
        int i = 49;
        printf("result for %d is %lld\n", i, fibonacci(i));
    // }
}

long long fibonacci(long long n) {
    if (n == 0){
      return 0;
    } else if (n == 1) {
        return 1;
    } else {
        return (fibonacci(n-1) + fibonacci(n-2));
    }
}