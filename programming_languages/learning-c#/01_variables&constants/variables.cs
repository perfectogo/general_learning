using System;

class Program 
{
    static void Main() 
    {
        // declaring variable with var key word
        var num = 45;
        Console.WriteLine(num);

        // declaring variable with type
        int number1 = 7;
        string word1 = "this is a string value";
        
        string a = $"{number1} {word1}";

        Console.WriteLine(a);
    }
}
