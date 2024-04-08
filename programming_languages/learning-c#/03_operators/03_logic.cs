using System;

class Program 
{
    static void Main() 
    {
        var n1 = true;
        var n2 = false;

        Console.WriteLine(n1 || n2);
        Console.WriteLine(n1 && n2);
        Console.WriteLine(!true);
        Console.WriteLine((n1 || n2).GetType());
    }
}
