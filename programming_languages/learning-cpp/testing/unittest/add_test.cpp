// add_test.cpp
#include <gtest/gtest.h>
#include "add.h"

TEST(AddFunctionTest, PositiveNumbers) {
    EXPECT_EQ(add(2, 3), 5);
}

TEST(AddFunctionTest, NegativeNumbers) {
    EXPECT_EQ(add(-2, -3), -5);
}

TEST(AddFunctionTest, MixedNumbers) {
    EXPECT_EQ(add(-2, 3), 1);
    EXPECT_EQ(add(2, -3), -1);
}

int main(int argc, char** argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}