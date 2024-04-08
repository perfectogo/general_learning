using System;

class Program
{
    static void Main() {
        // in c#, every type is a class

        // int classes
        int number1 = 32;
        Int16 number2 = 16;
        Int32 number3 = 32;
        Int64 number4 = 64;

        Console.WriteLine(number1.GetType());
        Console.WriteLine(number2.GetType());
        Console.WriteLine(number3.GetType());
        Console.WriteLine(number4.GetType());

        // float class
        float fNum1 = number2;
        float fNum2 = 32.0f;

        Console.WriteLine(fNum1);
        Console.WriteLine(fNum2);
        Console.WriteLine(fNum1.GetType());
        Console.WriteLine(fNum2.GetType());

        // double class
        double dNum1  = 64.0;
        Console.WriteLine(dNum1);
        Console.WriteLine(dNum1.GetType());

        // decimal class
        decimal deNum = 34;
        Console.WriteLine(deNum);
        Console.WriteLine(deNum.GetType());
    }
}