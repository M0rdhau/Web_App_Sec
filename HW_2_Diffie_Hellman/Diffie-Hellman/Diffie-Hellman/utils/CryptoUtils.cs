using System;
using System.Collections.Generic;

namespace Diffie_Hellman.utils
{
    public class CryptoUtils
    {
        public enum DHTypes
        {
            Prime,
            Base,
            
        }

        public static ulong ExtendedEuclid(ulong e, ulong d)
        {
            ulong[] x = {1, 0, d};
            ulong[] y = {0, 1, e};
            ulong[] t = {0, 0, 0};
            var i = 1;
            do
            {
                ulong q = 0;
                if (i == 1)
                {
                    q = x[2] / y[2];
                    for (var j = 0; j < 3; j++)
                    {
                        t[j] = x[j] - (q * y[j]);
                    }
                }
                else
                {
                    for (var j = 0; j < 3; j++)
                    {
                        x[j] = y[j];
                        y[j] = t[j];
                    }
                    q = x[2] / y[2];
                    for (int j = 0; j < 3; j++)
                    {
                        t[j] = x[j] - (q * y[j]);
                    }
                }
                i++;
                if (y[2] == 0)
                {
                    return 0;
                }
            } while (y[2] != 1);
            return y[1];
        }

        public static ulong ModPow(ulong mod, ulong pow, ulong number)
        {
            // number^pow%mod
            if (pow is 1) return number % mod;
            if (mod is 1 or 0) return number;
            if (number is 1 or 0) return number;

            ulong carryOver = number;
            while (true)
            {
                if (pow == 1) return carryOver % mod;
                carryOver = (carryOver * number) % mod;
                pow--;
            }
        }

        public static bool IsDivisible(ulong number, List<ulong> divisors)
        {
            var isDivisible = false;
            foreach(var divisor in divisors)
            {
                if (divisor > Math.Sqrt(number)) break;
                if (number % divisor == 0)
                {
                    isDivisible = true;
                    break;
                }
            }
            return isDivisible;
        }
        
        

    }
}