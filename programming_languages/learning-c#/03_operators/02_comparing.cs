using System;

class Program 
{
    static void Main() 
    {
        var n1 = 7;
        var n2 = 5;

        Console.WriteLine(n1 > n2);
        Console.WriteLine(n1 < n2);
        Console.WriteLine(n1 >= n2);
        Console.WriteLine(n1 <= n2);
        Console.WriteLine(n1 == n2);
        Console.WriteLine(n1 != n2);
        Console.WriteLine((n1 > n2).GetType());
    }
}
