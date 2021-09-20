using System;
using System.Linq;

namespace HW_1_Ancient_Crypto
{
    class Program
    {
        private const int MaxUtf32 = 0x10F800; //actually 0x10FFFF but I took UTF8's ceiling
        private enum StringType
        {
            PlainText,
            CipherText,
            KeyText
        }

        private static int RotateCharacterValue(int charValue, int keyValue, StringType type)
        {
            if(type == StringType.KeyText) throw new Exception("Illegal argument value" + type);
            for (int i = 0; i < keyValue; i++)
            {
                charValue = type == StringType.CipherText ? charValue - 1 : charValue + 1;
                charValue = charValue < 0 ? MaxUtf32 - charValue : charValue;
                if (charValue is >= 0xd800 and <= 0xdfff)
                {
                    charValue = type == StringType.CipherText ? 0xd800 - 1 : 0xdfff + 1 ;
                }
                while (char.IsControl(char.ConvertFromUtf32(charValue), 0))
                {
                    charValue = type == StringType.CipherText ? charValue - 1 : charValue + 1;
                    charValue = charValue < 0 ? MaxUtf32 - charValue : charValue;
                    if (charValue is >= 0xd800 and <= 0xdfff)
                    {
                        charValue = type == StringType.CipherText ? 0xd800 - 1 : 0xdfff + 1;
                    }
                }
            }

            return charValue;
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
                if (shiftString?.Trim().ToUpper() == "C") return;
                if (!int.TryParse(shiftString, out var shiftInt))
                {
                    Console.WriteLine("Not a valid shift!");
                    continue;
                }
                inputIsValid = true;
                shiftInt = shiftInt < 0 ? MaxUtf32 - shiftInt : shiftInt;
                Console.WriteLine("Shift: " + shiftInt + " (it's only positive numbers)")
                var ciphertext = "";
                foreach (var unicodeCodePoint in shiftable.GetUnicodeCodePoints())
                {
                    var unicodeValue = RotateCharacterValue(unicodeCodePoint, shiftInt, type);
                    ciphertext += char.ConvertFromUtf32(unicodeValue);
                }
                Console.WriteLine("Result: \"" + ciphertext + "\"");
            } while (!inputIsValid);
        }

        static void DoVigenere(string inputText, StringType type)
        {
            var keyString = "";
            var inputValid = false;
            while (!inputValid)
            {
                keyString = GetInputString(StringType.KeyText);
                if (keyString?.Trim().ToUpper() == "C")
                {
                    return;
                }

                if (keyString == null)
                {
                    Console.WriteLine("Invalid input!");
                }
                else
                {
                    inputValid = true;
                }
                
            }

            if (keyString.GetUtf32Length() != inputText.GetUtf32Length())
            {
                while (keyString.GetUtf32Length() < inputText.GetUtf32Length())
                {
                    keyString += keyString;
                }
                keyString = keyString.Utf32Substring(0, inputText.GetUtf32Length());
            }

            var cipherText = "";
            foreach (var codePoints in inputText.GetUnicodeCodePoints().Zip(keyString.GetUnicodeCodePoints()))
            {
                var unicodeValue = RotateCharacterValue(codePoints.First, codePoints.Second, type);
                cipherText += char.ConvertFromUtf32(unicodeValue);
            }
            Console.WriteLine("Result: \"" + cipherText + "\"");
        }

        static void Main()
        {
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
                    var type = ed == "E" ? StringType.PlainText 
                        : ed == "D" ? StringType.CipherText
                        : StringType.KeyText;
                    if (type == StringType.KeyText)
                    {
                        Console.WriteLine("Please enter one of the correct options.");
                        continue;
                    }

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