using System;

namespace Diffie_Hellman
{
    class Program
    {
        static void Main(string[] args)
        {
            //sqrt of maxvalue is more than actual sqrt
            //so square(sqrt(maxvalue)) > maxvalue == overflow
            Console.WriteLine("Hello World!");
            ulong ceiling =(ulong) Math.Floor(Math.Sqrt(ulong.MaxValue));
            ulong notCeiling = (ceiling * ceiling) - 1;
            ulong otherCeiling = (ulong) Math.Floor(Math.Sqrt(ceiling));
            Console.WriteLine(Math.Sqrt(ulong.MaxValue));
            Console.WriteLine(ceiling);
            Console.WriteLine(notCeiling);
            Console.WriteLine(otherCeiling);
        }
    }
}