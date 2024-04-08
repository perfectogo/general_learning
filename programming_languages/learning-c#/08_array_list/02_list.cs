using System;
using System.Collections.Generic; // added using directive for List<T>
using System.Linq;

class Program
{
    static void Main()
    {
        // List<int> numbers = new List<int> { 1, 2, 3, 4, 5, 6, 7, 8 }; // Corrected list initialization

        // // Call the declarations method to see different list declarations
        // declarations();

        //addItemToList();
        appendItemToList();
    }

    static void declarations()
    {
       // List<int> numbers0; // Declaration, but not initialized (null by default)
        List<int> numbers1 = new List<int>(); // Initialized as empty
        List<int> numbers2 = new List<int> { 1, 2, 3, 4, 5, 6 }; // Initialized with items
        var numbers3 = new List<int> { 1, 2, 3, 4, 5, 6 }; // Initialized with items using var

        // Printing the lists. Note that you need to iterate over the list or convert it to string representation
        Console.WriteLine("numbers1:");
        foreach (var num in numbers1)
        {
            Console.WriteLine(num);
        }

        Console.WriteLine("numbers2:");
        foreach (var num in numbers2)
        {
            Console.WriteLine(num);
        }

        Console.WriteLine("numbers3:");
        foreach (var num in numbers3)
        {
            Console.WriteLine(num);
        }
    }

    static void addItemToList()
    {
        var numbers1 = new List<int>{1, 2, 3, 4, 5, 6, 7};

        numbers1.Add(12);

        foreach (var num in numbers1)
        {
            Console.WriteLine(num);
        }
    }

    static void appendItemToList()
    {
        var numbers1 = new List<int>{1, 2, 3, 4, 5, 6, 7};

        var numbers2 = numbers1.Append(12);

        foreach (var num in numbers2)
        {
            Console.WriteLine(num);
        }
    }
}
