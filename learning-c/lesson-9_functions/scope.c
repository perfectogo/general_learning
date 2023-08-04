#include <stdio.h>

char* myMessage = "This is my global message!";

void myFunc() {
  char* myMessage = "This is my local message!";
  printf("%s\n", myMessage);
}

int main() {
  int myNumber = 39; // You can change this value to test different outputs

  if (myNumber <= 50) {
    printf("%d\n", myNumber);
    printf("%s\n", myMessage);
  } else {
    int myNumber = 500;
    printf("%d\n", myNumber);
    myFunc();
  }

  return 0;
}
