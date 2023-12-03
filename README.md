# general_learning
general learning

## C++ 
###  install <gtest/gtest.h> liblary on mac m1
#### To install and use Google Test (gtest) on macOS (including M1), you can follow these steps:
* Open Terminal and run the following command to install Homebrew:
```/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"``` 

* Install CMake (if not already installed):
Homebrew is the recommended way to install CMake:
```brew install cmake```

* Install Google Test:
You can use Homebrew to install Google Test:
```brew install googletest```




brew install cmake

brew install googletest

gtest-config --version

g++ -std=c++14 -o add_test add.cpp add_test.cpp -lgtest -lgtest_main -pthread                                                                          

