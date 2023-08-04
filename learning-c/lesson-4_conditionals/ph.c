/*
    1. In chemistry, pH is a scale used to specify the acidity or basicity of an aqueous solution.
    Write an if statement that checks if ph is greater than 7, then print “Basic” if it is.
    Checkpoint 2 Passed

    2. Change double ph = 7.9 to double ph = 4.6.
    Now attach an else if statement to that if statement and have it check if ph is less than 7, then print “Acidic” if it does.

    3. Change double ph = 4.6 to double ph = 7.
    Add an else statement that prints “Neutral” if all conditions fail to pass.
*/
#include <stdio.h>
 
int main() {
 
  double ph = 7;

  // Start the if, else if, else here:
  if (ph > 7) {
    printf("Basic\n");
  } else if (ph < 7) {
    printf("Acidic\n");
  } else {
    printf("Neutral\n");
  }
}