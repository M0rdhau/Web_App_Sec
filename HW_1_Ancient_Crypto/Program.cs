using System;
using System.Linq;

namespace HW_1_Ancient_Crypto
{
    class Program
    {
        private const int MaxUtf32 = 0x10FFFF;
        private enum StringType
        {
            PlainText,
            CipherText,
            KeyText
        }

        private static string GetInputString( StringType type )
        {
            var isValid = false;
            var text = "";
            var placeholder = type == StringType.CipherText ? "ciphertext"
                : type == StringType.KeyText ? "key string"
                : "plaintext";
            while (!isValid)
            {
                Console.WriteLine("Please input the " + placeholder + ", or 'C' to cancel: ");
                text = Console.ReadLine() ?? "";
                if (text == "")
                {
                    Console.WriteLine("Not a valid string!");
                    continue;
                }
                isValid = true;
            }
            return text;
        }

        private static void DoCaesar(string shiftable, StringType type )
        {
            var inputIsValid = false;
            do
            {
                Console.WriteLine("Please input the shift, or 'C' to cancel: ");
                var shiftString = Console.ReadLine();
                if (!int.TryParse(shiftString, out var shiftInt) && shiftString?.Trim().ToUpper() != "C")
                {
                    Console.WriteLine("Not a valid shift!");
                    continue;
                }
                inputIsValid = true;
                shiftInt = Math.Abs(shiftInt);
                var ciphertext = "";
                foreach (int unicodeCodePoint in shiftable.GetUnicodeCodePoints())
                {
                    int unicodeValue;
                    if (type == StringType.PlainText)
                    {
                        unicodeValue = (unicodeCodePoint + shiftInt) % (MaxUtf32 + 1);
                        if (unicodeValue < 0x20) unicodeValue += 0x20;
                    }else if (type == StringType.CipherText)
                    {
                        unicodeValue = (unicodeCodePoint - shiftInt) % (MaxUtf32 + 1);
                        if (unicodeValue < 0x20) unicodeValue = MaxUtf32 - (0x20 - unicodeValue);
                    }
                    else
                    {
                        throw new Exception("Illegal argument value" + type);
                    }
                    Console.WriteLine(unicodeValue + char.ConvertFromUtf32(unicodeValue));
                    ciphertext += char.ConvertFromUtf32(unicodeValue);
                }
                Console.WriteLine("Result: \"" + ciphertext + "\"");
            } while (!inputIsValid);
        }

        static void DoVigenere(string inputText, StringType type)
        {
            var inputIsValid = false;
            do
            {
                Console.WriteLine("Please enter key string, or 'C' to cancel:");
                var keyString = GetInputString(StringType.KeyText);
                if (keyString.Length > inputText.Length)
                {
                    Console.WriteLine("Key string and Input string should be of the same length - " + inputText.Length);
                    continue;
                }

                inputIsValid = true;
                var cipherText = "";
                foreach (var codePoints in inputText.GetUnicodeCodePoints().Zip(keyString.GetUnicodeCodePoints()))
                {
                    var unicodeValue = 0;
                    if (type == StringType.PlainText)
                    {
                        unicodeValue = (codePoints.First + codePoints.Second) % (MaxUtf32 + 1);
                        if (unicodeValue < 0x20) unicodeValue += 0x20;
                    }else if (type == StringType.CipherText)
                    {
                        unicodeValue = (codePoints.First - codePoints.Second) % (MaxUtf32 + 1);
                        if (unicodeValue < 0x20) unicodeValue = MaxUtf32 - (0x20 - unicodeValue);
                    }
                    Console.WriteLine(unicodeValue + char.ConvertFromUtf32(unicodeValue));
                    cipherText += char.ConvertFromUtf32(unicodeValue);
                }
                Console.WriteLine("Result: \"" + cipherText + "\"");
            } while (!inputIsValid);
        }

        static void Main()
        {
            Console.WriteLine("EZ Encryption: EZE ");
            
            String buff;
            do
            {
                Console.WriteLine("Please Select the encryption:");
                Console.WriteLine("C for Caesar");
                Console.WriteLine("V for Vigenere");
                Console.WriteLine("E for Exit");
                buff = Console.ReadLine();
                buff = buff?.Trim().ToUpper();
                if (buff == "V" || buff == "C")
                {
                    Console.WriteLine("Enter 'E' for Encryption, 'D' for Decryption");
                    var ed  = Console.ReadLine();
                    ed = ed?.Trim().ToUpper();
                    var type = ed == "E" ? StringType.PlainText : StringType.CipherText;
                    var text = GetInputString(type);
                    if (buff == "C")
                    {
                        DoCaesar(text, type);
                    }
                    else
                    {
                        DoVigenere(text, type);
                    }
                }
                else
                {
                    if(buff == "E") continue;
                    Console.WriteLine("Please enter one of the correct options.");
                    Console.WriteLine("C for Caesar");
                    Console.WriteLine("V for Vigenere");
                    Console.WriteLine("E for Exit");
                    buff = Console.ReadLine();
                    buff = buff?.Trim().ToUpper();
                }
            } while (buff != "E");
        }
    }
}