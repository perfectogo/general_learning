#include<bits/stdc++.h>

using namespace std;

class fileLogger {
    private:
        /* data */
    public:
        void handle(string error) {
            cout << "error: " << error << endl;
        }
        fileLogger(/* args */);
        ~fileLogger();
};

fileLogger::fileLogger(/* args */) {}
fileLogger::~fileLogger() {}

class customer {
    private: fileLogger logger =  fileLogger();
    public: 
        void add() {
            try {
                cout << "ma'lumotni bazaga saqlash kodi\n";
            }
            catch(const exception& e) {
                logger.handle(e.what());
            }
        }

        customer();
        ~customer();
};

customer::customer() {}
customer::~customer() {}

int main () {
    customer s;

    s.add();

    return 0;
}