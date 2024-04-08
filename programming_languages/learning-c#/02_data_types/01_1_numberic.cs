using System;

class Program
{
    static void Main()
    {
        sbyte mySByte = 127;
        byte myByte = 255;
        short myShort = 32767;
        ushort myUShort = 65535;
        int myInt = 2147483647;
        uint myUInt = 4294967295;
        long myLong = 9223372036854775807;
        ulong myULong = 18446744073709551615;

        Console.WriteLine("sbyte: " + mySByte);
        Console.WriteLine("byte: " + myByte);
        Console.WriteLine("short: " + myShort);
        Console.WriteLine("ushort: " + myUShort);
        Console.WriteLine("int: " + myInt);
        Console.WriteLine("uint: " + myUInt);
        Console.WriteLine("long: " + myLong);
        Console.WriteLine("ulong: " + myULong);
    }
}
