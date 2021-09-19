using System;
using System.Collections.Generic;

namespace HW_1_Ancient_Crypto
{
    public static class MyExtensions
    {
        public static IEnumerable<int> GetUnicodeCodePoints(this string s)
        {
            Console.WriteLine(s + " Lenght: " + s.Length);
            Console.WriteLine(s + " utf32 Length: " + s.GetUtf32Length());
            for (int i = 0; i < s.Length; i++)
            {
                var unicodeCodePoint = 0;
                try
                {
                    unicodeCodePoint = char.ConvertToUtf32(s, i);
                }
                catch (ArgumentException ex)
                {
                    Console.WriteLine("Illegal input");
                    Console.WriteLine(ex.ToString());
                }

                if (unicodeCodePoint > 0xffff)
                {
                    i++;
                }
                yield return unicodeCodePoint;
            }
        }

        public static string Utf32Substring(this string s, int startIndex, int length)
        {
            List<int> unicodePoints = new List<int>();
            for (int i = 0; i < s.Length; i++)
            {
                var unicodeCodePoint = 0;
                try
                {
                    unicodeCodePoint = char.ConvertToUtf32(s, i);
                }
                catch (ArgumentException ex)
                {
                    Console.WriteLine("Illegal input");
                    Console.WriteLine(ex.ToString());
                }

                if (unicodeCodePoint > 0xffff)
                {
                    i++;
                }
                unicodePoints.Add(unicodeCodePoint);
            }
            var retString = "";
            for (var i = startIndex; i < startIndex + length; i++)
            {
                retString += char.ConvertFromUtf32(unicodePoints[i]);
            }
            return retString;
        }

        public static int GetUtf32Length(this string s)
        {
            var length = 0;
            for (var i = 0; i < s.Length; i++)
            {
                var unicodeCodePoint = 0;
                try
                {
                    unicodeCodePoint = char.ConvertToUtf32(s, i);
                }
                catch (ArgumentException ex)
                {
                    Console.WriteLine("Illegal input");
                    Console.WriteLine(ex.ToString());
                }

                if (unicodeCodePoint > 0xffff)
                {
                    i++;
                }

                length++;
            }

            return length;
        }
    }
}