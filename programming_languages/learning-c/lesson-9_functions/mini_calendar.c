#include <stdio.h>
#include <stdbool.h>

// Function to check if a year is a leap year
bool is_leap_year(int year) {
  if (year % 4 != 0) {
    return false;
  } else if (year % 100 == 0) {
    if (year % 400 != 0) {
      return false;
    }
  }
  return true;
}

// Function to add days to a given date
void add_days_to_date(int* dd, int* mm, int* yy, int days_left_to_add) {
  int days_in_month[] = {0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
  int days_left_in_month;

  while (days_left_to_add > 0) {
    days_left_in_month = days_in_month[*mm] - *dd + 1;

    if (*mm == 2 && is_leap_year(*yy)) {
      days_left_in_month += 1;
    }

    if (days_left_to_add > days_left_in_month) {
      days_left_to_add -= days_left_in_month;
      days_left_to_add -= 1;
      *dd = 1;
      if (*mm == 12) {
        *mm = 1;
        *yy += 1;
      } else {
        *mm += 1;
      }
    } else {
      *dd += days_left_to_add;
      days_left_to_add = 0;
    }
  }
}

int main() {
  int dd, mm, yy, days_left_to_add;

  printf("Enter a year between 1800 and 10000: ");
  scanf("%d", &yy);

  if (is_leap_year(yy)) {
    printf("Leap Year\n");
  } else {
    printf("Not Leap Year\n");
  }

  printf("Enter a date (mm dd yy) between 1800 and 10000: ");
  scanf("%d %d %d", &mm, &dd, &yy);

  printf("Enter the number of days to add: ");
  scanf("%d", &days_left_to_add);

  add_days_to_date(&dd, &mm, &yy, days_left_to_add);

  printf("New date after adding %d days: %d %d %d\n", days_left_to_add, mm, dd, yy);

  return 0;
}
