using System;

class Program
{
    static void Main()
    {
       int[] numbers = new int[]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10};

       PrintArray1(numbers);

    }

    static void declarations()
    {
        int[] array0;
        int[] array1 = new int[]{1, 2, 3, 4, 5, 6};
        var array2 = new int[]{1, 2, 3, 4, 5, 6};
        var array3 = new int[5];

        array0 = array1;
        array0[0] = 100;

        Console.WriteLine(array0);
        Console.WriteLine(array1[0]); // array is reference type
        Console.WriteLine(array2[0]);
        Console.WriteLine(array3[0]);

        PrintArray(array1);
    }

    static void PrintArray(int[] array) 
    {
        for (int i=0; i < array.Length; i++)
        {
            Console.WriteLine(array[i]);
        }
    }

    static void PrintArray1(int[] array)
    {
        foreach (var number in array) 
        {
            Console.WriteLine(number);
        }
    }

    static void MostCommonMethods() 
    {
        
    }
}