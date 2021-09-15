using System.Collections.Generic;

namespace HW_1_Ancient_Crypto
{
    public static class MyExtensions
    {
        public static IEnumerable<int> GetUnicodeCodePoints(this string s)
        {
            for (int i = 0; i < s.Length; i++)
            {
                int unicodeCodePoint = char.ConvertToUtf32(s, i);
                if (unicodeCodePoint > 0xffff)
                {
                    i++;
                }
                yield return unicodeCodePoint;
            }
        }
    }
}