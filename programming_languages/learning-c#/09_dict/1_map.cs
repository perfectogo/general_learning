using System;
using System.Collections.Generic;

class Program
{
    static void Main()
    {
        //declarations();
        per();

    }

    static void declarations()
    {
        //Dictionary<string, string> dict0; // Declaration, not initialized (null by default)
        Dictionary<int, string> dict1 = new Dictionary<int, string> { { 1, "bir" }, { 2, "ikki" } }; // Initialized with items

        foreach (var kvp in dict1) // Corrected foreach loop syntax
        {
            Console.WriteLine($"{kvp.Key}:{kvp.Value}");
        }
    }

    static void per()
    {
        Dictionary<int, string> dict1 = new Dictionary<int, string> { { 1, "bir" }, { 2, "ikki" } }; // Initialized with items

        string value;
        if (dict1.TryGetValue(2, out value)) // Attempt to get the value for key 3
        {
            Console.WriteLine(value); // Key 3 exists, print its value
        }
        else
        {
            Console.WriteLine($"Key {3} not found in dictionary."); // Key 3 doesn't exist
        }
    }
}
