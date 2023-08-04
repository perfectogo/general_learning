/*
    1. Add an else statement that prints “Fail\n”.

    2. Add a second if / else statement with the condition grade2 > 60 that prints “Pass\n” if true and “Fail\n” if otherwise. Run the program again to see how the else clause gets skipped when the condition isn’t false!
*/

#include <stdio.h>

int main() {

  int grade1 = 59;
  int grade2 = 90;

  if (grade1 > 60) {
    printf("Pass\n");
  } else {
    printf("Fail\n");
  }

  if (grade2 > 60) {
    printf("Pass\n");
  } else {
    printf("Fail\n");
  }

}