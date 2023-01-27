#include<bits/stdc++.h>

using namespace std;

class srp {
    public: 
        void add() {
            try {
                cout << "ma'lumotni bazaga saqlash kodi\n";
            }
            catch(const std::exception& e) {
                std::cerr << e.what() << '\n';
            }
        }

        srp();
        ~srp();
};

srp::srp() {}
srp::~srp() {}

int main () {
    srp s;

    s.add();

    return 0;
}