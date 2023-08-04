/*
    1.In the accompanying code block, write code that prints the element that is in the fourth row and second column of matrix.

    2.Loop through matrix and add up all the elements. Assign your answer to the sum variable.
*/

#include<stdio.h>

int main() {
  int matrix[][4] = {{14, 10, 6, 4}, {3, 7, 18, 11}, {13, 9, 5, 17}, {19, 12, 2, 1}}; 
  int sum = 0;

  // Checkpoint 1 code goes here.
  printf("%i\n", matrix[3][1]);

  // Checkpoint 2 code goes here.
   for(int q = 0; q < sizeof(matrix)/sizeof(matrix[0]); q++){
    for(int t = 0; t < sizeof(matrix[0])/sizeof(int); t++){
      sum += matrix[q][t];
    }
  }
  
}