using System;

class Program
{
    static void Main()
    {
        var ys = (a: -9, 0, 67, 100);

        // Accessing the second element by index
        var secondElement = ys.Item2;

        // Checking the type of the second element
        Console.WriteLine("Type of the second element: " + secondElement.GetType());
    }
}

