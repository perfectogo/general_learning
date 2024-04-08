using System;

public class MainClass
{
    static void Main()
    {
        OddsOneToHundred();
    }

    private static void oneToTen() {
        for (int i = 1; i < 10; i++) 
        {
            Console.WriteLine(i);
        }
    }

    private static void TenToOne() {
        for (int i = 10; i > 0; i--) 
        {
            Console.WriteLine(i);
        }
    }

    private static void OddsOneToHundred() {
        for (int i = 1; i < 100; i+=2) 
        {
            Console.WriteLine(i);
        }
    }
}
