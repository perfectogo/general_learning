using System;

class Program 
{
    static void Main() {
        int intValue = 10;
        double doubleValue;

        // Explicitly cast int to double
        doubleValue = (double)intValue;
        Console.WriteLine("Double value: " + doubleValue); // Output: 10.0

        Console.WriteLine((string)intValue);
    }
}