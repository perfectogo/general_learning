#include<stdio.h>

void print_matrix(int mat[3][3]) {
    int i, j;

    for (i = 0; i < 3; i++) {
        for (j = 0; j < 3; j++) {
            printf("%d ", mat[i][j]);
        }
        printf("\n");
    }
}

int main() {

  // Checkpoint 1 code goes here.
  int arr[9][9];

  // Checkpoint 2 code goes here.  
  int matrix[3][3][3] = {{2, 9, 11},{3, 1, 4},{8, 3, 1}};

  print_matrix(matrix);


}

