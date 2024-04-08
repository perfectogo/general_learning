using System;

class Program 
{
    static void oneToTen() 
    {
        var i = 0;

        while (i < 10)
        {
            Console.WriteLine(i);

            i++;
        }
    }

    static void TenToOne() 
    {
        var i = 10;

        while (i > 0)
        {
            Console.WriteLine(i);

            i--;
        }
    }
    static void Main()
    {
        TenToOne();
    }
}