using System;

class Program
{
    static void Main () 
    {
        var a = 15;
        var b = 17;
        var c = 19;

        var max = a;

        if (b > a && b > c)
        {
            max = b;
        }
        else if (c > a && c > b)
        { 
            max = c;
        }

        Console.WriteLine(max);
    }
}