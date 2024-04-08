using System;

class Program 
{
    static void Main () {
        string numberString = "123";
        int parsedInt;

        // Using int.Parse method
        parsedInt = int.Parse(numberString);
        Console.WriteLine("Parsed Int value: " + parsedInt); // Output: 123

        // Using int.TryParse method
        if (int.TryParse(numberString, out parsedInt))
        {
            Console.WriteLine("Parsed Int value: " + parsedInt); // Output: 123
        }
        else
        {
            Console.WriteLine("Failed to parse the string.");
        }
    }
}