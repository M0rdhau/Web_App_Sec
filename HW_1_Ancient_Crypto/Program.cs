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
            // If result is negative, make sure to loop back on the number field from maximum point
            charValue = charValue < 0 ? MaxUtf32 + charValue : charValue;
            // loop back on the number field up from 0
            charValue %= MaxUtf32;
            // If character value lands in the surrogate number range, get out of there
            if (charValue is >= 0xd800 and <= 0xdfff)
            {
                charValue = type == StringType.CipherText ? 0xd800 - 1 : 0xdfff + 1 ;
            }

            return charValue;
        }

        private static int RotateCharacterValue(int charValue, int keyValue, StringType type)
        {
            if(type == StringType.KeyText) throw new Exception("Illegal argument value: " + type);
            // Rotate character value. If encrypting, add key value, otherwise subtract
            charValue = type == StringType.CipherText ? charValue - keyValue : charValue + keyValue;
            charValue = NormalizeCharValue(charValue, type);
            while (char.IsControl(char.ConvertFromUtf32(charValue), 0))
            {
                // Rotate character value out of the undesirable number range.
                // If encrypting, add key value, otherwise subtract
                charValue = type == StringType.CipherText ? charValue - 1 : charValue + 1;
                charValue = NormalizeCharValue(charValue, type);
            }

            return charValue;
        }

        

        private static string DoCaesar(string shiftable, int shiftInt, StringType type )
        {
            shiftInt = shiftInt < 0 ? MaxUtf32 + shiftInt : shiftInt;
            shiftInt = shiftInt % MaxUtf32;
            Console.WriteLine("Shift: " + shiftInt);
            var ciphertext = "";
            foreach (var unicodeCodePoint in shiftable.GetUnicodeCodePoints())
            {
                var unicodeValue = RotateCharacterValue(unicodeCodePoint, shiftInt, type);
                ciphertext += char.ConvertFromUtf32(unicodeValue);
            }

            return ciphertext;
        }

        static string DoVigenere(string inputText, string keyString, StringType type)
        {
            // var inputValid = false;
            // while (!inputValid)
            // {
            //     keyString = GetInputString(StringType.KeyText, out var something);
            //     if (keyString?.Trim().ToUpper() == "C")
            //     {
            //         return;
            //     }
            //
            //     if (keyString == null)
            //     {
            //         Console.WriteLine("Invalid input!");
            //     }
            //     else
            //     {
            //         inputValid = true;
            //     }
            //     
            // }

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

            return cipherText;
        }

        
        private static string GetInputString( StringType type, out bool cancelMenu )
        {
            cancelMenu = false;
            var text = "";
            var placeholder = type == StringType.CipherText ? "ciphertext"
                : type == StringType.KeyText ? "key string"
                : "plaintext";
            Console.WriteLine("Please input the " + placeholder + ", or nothing to return to previous menu");
            text = Console.ReadLine() ?? "";
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
                var text = GetInputString(stringType, out var isCancelled);
                var resultString = "";
                //if input is blank, return to previous menu
                if (isCancelled)
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
                        inputIsValid = true;
                        resultString = DoCaesar(text, shiftInt, stringType);
                    } while (!inputIsValid);
                }
                else
                {
                    
                }
                Console.WriteLine("Result is: " + resultString);
                Console.WriteLine("Would you like to " + encMethod + " with " + encName + " again?");
                Console.WriteLine("[Y]es");
                Console.Write("Anything else to go back to the previous menu");
                var buff = Console.ReadLine();
                buff = buff?.Trim().ToUpper();
                continueLooping = buff == "Y";
            }
            return;
        }

        static bool EncryptDecryptMenu(EncType encType)
        {
            while (true)
            {
                Console.WriteLine("Would you like to [E]ncrypt or [D]ecrypt a string?");
                Console.WriteLine("Or would you like to go [B]ack to the previous menu?");
                Console.Write("Or would you like to e[X]it the program?");
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
                if (buff != "C" && buff != "V")
                {
                    Console.Write("Invalid option!");
                    continue;
                }
                StringType stringType = buff == "E" ? StringType.PlainText : StringType.CipherText;
                VigenereCaesarMenu(encType, stringType);
            }
            // done encrypting, exit the menu, don't exit the program
            return true;
        }

        static void MainMenu()
        {
            var buff = "";
            while (buff != "X")
            {
                Console.WriteLine("Please Select the encryption:");
                Console.WriteLine("[C]aesar");
                Console.WriteLine("[V]igenere");
                Console.Write("Or would you like to e[X]it?");
                //Normalize the buffer
                buff = Console.ReadLine();
                buff = buff?.Trim().ToUpper();
                //Start the menu over if the buffer options are not supported
                if (buff != "C" && buff != "V" && buff != "X")
                {
                    Console.Write("Invalid option!");
                    continue;
                }

                EncType type = buff == "C" ? EncType.Caesar : EncType.Vigenere;
                buff = EncryptDecryptMenu(type) ? "" : "X";
            }
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


                    var text = GetInputString(type, out var isCancelled);
                    if(isCancelled) continue;
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