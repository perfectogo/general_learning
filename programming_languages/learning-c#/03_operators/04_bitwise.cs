using System;

// & - AND, | - OR, ^ - XOR, ~ - NOT , << - Left Shift , >> - Right Shift
class Program
{
    static void Main () {

        var num1 = 5; // 101 
        var num2 = 7; // 111
       
        // OR
        /*
            1 | 1 = 1
            0 | 1 = 1
            1 | 0 = 1
            0 | 0 = 0
        */
        var or = num1 | num2;
       //  010 = 101 XOR 111 

        Console.WriteLine(or);

        // XOR
        /*
            1 ^ 1 = 0
            0 ^ 1 = 1
            1 ^ 0 = 1
            0 ^ 0 = 0
        */
        var xor = num1 ^ num2;
       //  010 = 101 XOR 111 

        Console.WriteLine(xor);

        // AND
        /*
            1 & 1 = 1
            0 & 1 = 0
            1 & 0 = 0
            0 & 0 = 0
        */
        var and = num1 & num2;
       //  010 = 101 XOR 111 

        Console.WriteLine(and);

        // Left Shift
        /*
        */
        var num3 = 7; // 111
       
        num3 = num3<<3;
              //111000

        Console.WriteLine(num3);

        // Rigt Shift
        /*
        */
        var num4 = 7; // 10100
       
        num4 = num4>>3;
              //

        Console.WriteLine(num4);
    }
}