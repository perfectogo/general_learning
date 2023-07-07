#include <iostream>

using namespace std;

int main () {

    char a[] = "abcdeff";

    for (int i = 0; i <= sizeof(a) - 1; i++) {

        int count = 0;
        
        for (int j = 0; j <= sizeof(a) - 1; j++) {
            if (a[i] == a[j]) {
                count++;
            }
        }

        cout << char(a[i]) << count << "; ";
    } 


    return 0;
} 



// find the number of each array unique elements
// determination array