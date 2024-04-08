using System;

class Program
{
    static void Main()
    {
        var num = 5;

        if (num % 3 == 0 && num % 5 == 0)
        {
            Console.WriteLine("FIZZBUZZ");

            return;
        }
        else if (num % 3 == 0)
        {
            Console.WriteLine("FIZZ");
            
            return;
        }
        
        Console.WriteLine("BUZZ");
    }
}