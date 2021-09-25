using System;
using System.Linq;

namespace HW_1_Ancient_Crypto
{
    class Program
    {
        private const int MaxUtf32 = 0x10FFFF; //actually 0x10FFFF but I took UTF8's ceiling
        private enum StringType
        {
            PlainText,
            CipherText,
            KeyText
        }

        private enum EncType
        {
            Vigenere,
            Caesar
        }

        private static int NormalizeCharValue(int charValue, StringType type)
        {
            // Normalize the actual number we're going to shift with
            // In c#, number field springs up from number 0 and flows outwards
            // Hence, modulo operator works weirdly.
            // Assume |A| > |B| and A > 0 and B < 0
            // Modulo operator, in B % A should return ||B| - |A||
            // But what it actually returns, is -||B|%|A||
            // Next line remedies that
            charValue = charValue < 0 ? MaxUtf32 + charValue : charValue;
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

        private static int RotateCharacterValue(int charValue, int keyValue, StringType type)
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

        

        private static string DoCaesar(string shiftable, int shiftInt, StringType type )
        {
            // Calculate the actual number we're going to shift with
            // In c#, number field springs up from number 0 and flows outwards
            // Hence, modulo operator works weirdly.
            // Assume |A| > |B| and A > 0 and B < 0
            // Modulo operator, in B % A should return ||B| - |A||
            // But what it actually returns, is -||B|%|A||
            // Next line remedies that
            shiftInt = shiftInt < 0 ? MaxUtf32 + shiftInt : shiftInt;
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

        static string DoVigenere(string inputText, string keyString, StringType type)
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
        
        private static string GetInputString( StringType type)
        {
            // convert string type definitions to strings
            var placeholder = type == StringType.CipherText ? "ciphertext"
                : type == StringType.KeyText ? "key string"
                : "plaintext";
            Console.WriteLine($"Please input the {placeholder}, or nothing to return to previous menu");
            var text = Console.ReadLine() ?? "";
            if (text == "")
            {
                return "";
            }
            return text;
        }
        
        static void VigenereCaesarMenu(EncType encType, StringType stringType)
        {
            var continueLooping = true;
            while (continueLooping)
            {
                // Set names
                string encName = encType == EncType.Caesar ? "Caesar" : "Vigenere";
                string encMethod = stringType == StringType.CipherText ? "Decrypt" : "Encrypt";
                Console.WriteLine("===============================");
                Console.WriteLine($"{encName} {encMethod}ion");
                Console.WriteLine("===============================");
                var text = GetInputString(stringType);
                var resultString = "";
                //if input is blank, return to previous menu
                if (text == "")
                {
                    return;
                }
                if (encType == EncType.Caesar)
                {
                    var inputIsValid = false;
                    do
                    {
                        Console.WriteLine("Please input the shift integer, or 'C' to return to previous menu");
                        var shiftString = Console.ReadLine();
                        //if input is C, return to previous menu
                        if (shiftString?.Trim().ToUpper() == "C") return;
                        if (!int.TryParse(shiftString, out var shiftInt))
                        {
                            Console.WriteLine("Not a valid shift!");
                            continue;
                        }
                        //required to exit the loop
                        inputIsValid = true;
                        resultString = DoCaesar(text, shiftInt, stringType);
                    } while (!inputIsValid);
                }
                else
                {
                    var keyString = GetInputString(StringType.KeyText);
                    //if input is blank, return to previous menu
                    if (keyString == "")
                    {
                        return;
                    }
                    resultString = DoVigenere(text, keyString, stringType);

                }
                Console.WriteLine();
                Console.WriteLine("Result is:");
                Console.WriteLine($"###{resultString}###");
                Console.WriteLine($"Would you like to {encMethod}  with {encName} again?");
                Console.WriteLine("[Y]es");
                Console.WriteLine("Anything else to go back to the previous menu");
                var buff = Console.ReadLine();
                buff = buff?.Trim().ToUpper();
                continueLooping = buff == "Y";
            }
        }

        static bool EncryptDecryptMenu(EncType encType)
        {
            while (true)
            {
                Console.WriteLine("===============================");
                Console.WriteLine(encType == EncType.Caesar ? "Caesar" : "Vigenere");
                Console.WriteLine("===============================");
                Console.WriteLine("Would you like to [E]ncrypt or [D]ecrypt a string?");
                Console.WriteLine("Or would you like to go [B]ack to the previous menu?");
                Console.WriteLine("Or would you like to e[X]it the program?");
                //Normalize the buffer
                var buff = Console.ReadLine();
                buff = buff?.Trim().ToUpper();
                // exit the menu, don't exit the program
                if (buff == "B")
                {
                    return true;
                } 
                // exit the menu and exit the program
                if (buff == "X")
                {
                    return false;
                }
                //Start the menu over if the buffer options are not supported
                if (buff != "E" && buff != "D")
                {
                    Console.WriteLine("Invalid option!");
                    continue;
                }
                StringType stringType = buff == "E" ? StringType.PlainText : StringType.CipherText;
                VigenereCaesarMenu(encType, stringType);
            }
        }

        static void MainMenu()
        {
            var buff = "";
            while (buff != "X")
            {
                Console.WriteLine("===============================");
                Console.WriteLine("Please Select the encryption:");
                Console.WriteLine("===============================");
                Console.WriteLine("[C]aesar");
                Console.WriteLine("[V]igenere");
                Console.WriteLine("Or would you like to e[X]it?");
                //Normalize the buffer
                buff = Console.ReadLine();
                buff = buff?.Trim().ToUpper();
                //Start the menu over if the buffer options are not supported
                if (buff == "X") continue;
                if (buff != "C" && buff != "V")
                {
                    Console.WriteLine("Invalid option!");
                    continue;
                }

                EncType type = buff == "C" ? EncType.Caesar : EncType.Vigenere;
                buff = EncryptDecryptMenu(type) ? "" : "X";
            }
        }

        static void Main()
        {
            MainMenu();
        }
    }
}