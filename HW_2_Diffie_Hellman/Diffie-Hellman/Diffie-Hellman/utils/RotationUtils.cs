using System;
using System.Linq;

namespace Diffie_Hellman.utils
{
    public static class RotationUtils
    {
        public const int MaxUtf32 = 0x10FFFF;
        public enum StringType
        {
            PlainText,
            CipherText,
            KeyText
        }
        public static int NormalizeCharValue(int charValue, StringType type)
        {
            // Normalize the actual number we're going to shift with
            // In c#, number field springs up from number 0 and flows outwards
            // Hence, modulo operator works weirdly.
            // Assume |A| > |B| and A > 0 and B < 0
            // Modulo operator, in B % A should return ||B| - |A||
            // But what it actually returns, is -||B|%|A||
            // Next line remedies that
            while (charValue < 0)
            {
                charValue = MaxUtf32 + charValue;
            }
            // loop back on the number field up from 0
            charValue %= MaxUtf32;
            // If character value lands in the surrogate number range, get out of there
            // If encrypting, make it 0xe000, otherwise 0xd799 (right before and after the range)
            if (charValue is >= 0xd800 and <= 0xdfff)
            {
                charValue = type == StringType.CipherText ? 0xd800 - 1 : 0xdfff + 1 ;
            }

            return charValue;
        }

        public static int RotateCharacterValue(int charValue, int keyValue, StringType type)
        {
            if(type == StringType.KeyText) throw new Exception($"Illegal argument value: {type}");
            // Rotate character value. If encrypting, add key value, otherwise subtract
            charValue = type == StringType.CipherText ? charValue - keyValue : charValue + keyValue;
            charValue = NormalizeCharValue(charValue, type);
            while (char.IsControl(char.ConvertFromUtf32(charValue), 0))
            {
                // Rotate character value out of the undesirable (control character) number range.
                // If encrypting, add key value, otherwise subtract
                charValue = type == StringType.CipherText ? charValue - 1 : charValue + 1;
                charValue = NormalizeCharValue(charValue, type);
            }

            return charValue;
        }

        

        public static string DoCaesar(string shiftable, int shiftInt, StringType type )
        {
            // Calculate the actual number we're going to shift with
            // In c#, number field springs up from number 0 and flows outwards
            // Hence, modulo operator works weirdly.
            // Assume |A| > |B| and A > 0 and B < 0
            // Modulo operator, in B % A should return ||B| - |A||
            // But what it actually returns, is -||B|%|A||
            // Next line remedies that
            while (shiftInt < 0)
            {
                shiftInt = MaxUtf32 + shiftInt;
            }
            // regular modulo operation
            shiftInt = shiftInt % MaxUtf32;
            Console.WriteLine($"Shift: {shiftInt}");
            var ciphertext = "";
            // rotate each character by the given shift
            foreach (var unicodeCodePoint in shiftable.GetUnicodeCodePoints())
            {
                var unicodeValue = RotateCharacterValue(unicodeCodePoint, shiftInt, type);
                ciphertext += char.ConvertFromUtf32(unicodeValue);
            }

            return ciphertext;
        }

        public static string DoVigenere(string inputText, string keyString, StringType type)
        {
            // Normalizing the key string
            if (keyString.GetUtf32Length() != inputText.GetUtf32Length())
            {
                //Make sure that key string is at least the size of cipher string, or more
                while (keyString.GetUtf32Length() < inputText.GetUtf32Length())
                {
                    keyString += keyString;
                }
                //Make sure they're the same length by truncating the possibly longer key string
                keyString = keyString.Utf32Substring(0, inputText.GetUtf32Length());
            }
            var cipherText = "";
            // rotate the value of ith character in plaintext by the UTF32 value of the ith character in 
            // key string
            foreach (var codePoints in inputText.GetUnicodeCodePoints().Zip(keyString.GetUnicodeCodePoints()))
            {
                var unicodeValue = RotateCharacterValue(codePoints.First, codePoints.Second, type);
                cipherText += char.ConvertFromUtf32(unicodeValue);
            }

            return cipherText;
        }
    }
}