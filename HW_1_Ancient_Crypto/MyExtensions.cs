using System;
using System.Collections.Generic;

namespace HW_1_Ancient_Crypto
{
    public static class MyExtensions
    {
        private static List<int> UnicodeCodePoints(string s)
        {
            List<int> unicodePoints = new List<int>();
            for (int i = 0; i < s.Length; i++)
            {
                var unicodeCodePoint = 0;
                try
                {
                    unicodeCodePoint = char.ConvertToUtf32(s, i);
                }
                catch
                {
                    Console.WriteLine("Illegal input - text contains invalid characters!");
                    Console.WriteLine("Output will be slightly wrong");
                }

                if (unicodeCodePoint > 0xffff)
                {
                    i++;
                }
                unicodePoints.Add(unicodeCodePoint);
            }

            return unicodePoints;
        }

        public static IEnumerable<int> GetUnicodeCodePoints(this string s)
        {
            return UnicodeCodePoints(s);
        }

        public static string Utf32Substring(this string s, int startIndex, int length)
        {
            List<int> unicodePoints = UnicodeCodePoints(s);
            var retString = "";
            for (var i = startIndex; i < startIndex + length; i++)
            {
                retString += char.ConvertFromUtf32(unicodePoints[i]);
            }
            return retString;
        }

        public static int GetUtf32Length(this string s)
        {
            return UnicodeCodePoints(s).Count;
        }
    }
}