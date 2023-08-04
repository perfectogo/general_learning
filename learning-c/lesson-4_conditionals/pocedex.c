/*
    1. Here we have a switch statement!
    Let’s add 3 more cases right before default:

        case 7 that outputs “Squirtle”
        case 8 that outputs “Wartortle”
        case 9 that outputs “Blastoise”

    2. Remove break; from case 7.
    What do you think will happen when you run it?
*/

#include <stdio.h>

int main() {

  int number = 7;

  switch(number) {
    case 1:
      printf("Bulbasaur\n");
      break;
    case 2:
      printf("Ivysaur\n");
      break;
    case 3:
      printf("Venusaur\n");
      break;
    case 4:
      printf("Charmander\n");
      break;
    case 5:
      printf("Charmeleon\n");
      break;
    case 6:
      printf("Charizard\n");
      break;
    case 7:
      printf("Squirtle\n");
      // break;
    case 8:
      printf("Wartortle\n");
      break;
    case 9:
      printf("Blastoise\n");
      break;
    default:
      printf("Unknown\n");
      break;
  }
}