/*
    1. Change the logical operator in the conditional statement from a > 0 && b > 0 to a > 0 || b > 0.
    What do you think will happen?

    2. Now add another if statement that checks if both a > 0 and !(b > 0) are true and prints “Positive too” if they are.
    What do you think will happen?
*/
#include <stdio.h>

int main() {

  int a = 10;
  int b = -5;

  if (a > 0 || b > 0) {
    printf("Positive\n");
  }

   if (a > 0 && !(b > 0)) {
    printf("Positive too\n");
  }
}