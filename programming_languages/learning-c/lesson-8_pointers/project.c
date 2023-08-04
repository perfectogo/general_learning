#include <stdio.h>
#include <string.h>

void copy(char* dst, char* src) {
  while (*src != '\0') {
    *dst = *src; // Copy the character pointed to by src into the address stored in dst
    src++;      // Move src pointer to the next character in the source string
    dst++;      // Move dst pointer to the next available space in the destination string
  }
  *dst = '\0';  // Add the null terminator at the end of the copied string
}

int main() {
  char srcString[] = "We promptly judged antique ivory buckles for the next prize!";
  int srcLength = strlen(srcString);

  char dstString[srcLength + 1]; // Allocate enough space for dstString (including null terminator)

  // Using the copy function to copy srcString into dstString
  copy(dstString, srcString);

  printf("Source String: %s\n", srcString);
  printf("Copied String: %s\n", dstString);

  return 0;
}
