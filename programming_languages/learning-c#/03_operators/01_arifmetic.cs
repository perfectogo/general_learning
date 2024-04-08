using System;

class Program 
{
    static void Main() {
        var n1 = 7;
        var n2 = 5;

        Console.WriteLine("Arifmetic");
        Console.WriteLine(n1 + n2);
        Console.WriteLine(n1 - n2);
        Console.WriteLine(n1 * n2);
        Console.WriteLine(n1 / n2);
        Console.WriteLine(n1 % n2);

        // SHORT
         Console.WriteLine("Short");
        var n = 1;

        n+=1;
        Console.WriteLine(n);

        n-=1;
        Console.WriteLine(n);

        n*=1;
        Console.WriteLine(n);

        n/=1;
        Console.WriteLine(n);

        n%=1;
        Console.WriteLine(n);

        Console.WriteLine("Increment");
        //INCREMENT
        //int n3 = 0;
        var n11 = 7;
        var n12 = 7;
        var n21 = 5;

        //postfix
        var n3 = n21 + n11++;
        Console.WriteLine(n3);

        // postfix
        var n4 = n21 + ++n12;
        Console.WriteLine(n4);

        //DECREMENT
          //postfix
        var n5 = n21 + n11--;
        Console.WriteLine(n5);

        // postfix
        var n6 = n21 + --n12;
        Console.WriteLine(n6);
    }
}