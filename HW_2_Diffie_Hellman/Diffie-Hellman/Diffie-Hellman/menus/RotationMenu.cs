using System;
using Diffie_Hellman.utils;
using Microsoft.VisualBasic.CompilerServices;
using static Diffie_Hellman.utils.RotationUtils;

namespace Diffie_Hellman.menus
{
    public static class RotationMenu
    {
        public enum EncType
        {
            Vigenere,
            Caesar
        }
        private static string GetInputString( RotationUtils.StringType type)
        {
            // convert string type definitions to strings
            var placeholder = type == RotationUtils.StringType.CipherText ? "ciphertext"
                : type == RotationUtils.StringType.KeyText ? "key string"
                : "plaintext";
            Console.WriteLine($"Please input the {placeholder}, or nothing to return to previous menu");
            var text = Console.ReadLine() ?? "";
            if (text == "")
            {
                return "";
            }
            return text;
        }
        
        public static void VigenereCaesarMenu(EncType encType, RotationUtils.StringType stringType)
        {
            var continueLooping = true;
            while (continueLooping)
            {
                // Set names
                string encName = encType == EncType.Caesar ? "Caesar" : "Vigenere";
                string encMethod = stringType == RotationUtils.StringType.CipherText ? "Decrypt" : "Encrypt";
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
                    var keyString = GetInputString(RotationUtils.StringType.KeyText);
                    //if input is blank, return to previous menu
                    if (keyString == "")
                    {
                        return;
                    }
                    resultString = DoVigenere(text, keyString, stringType);

                }
                Console.WriteLine();
                Console.WriteLine("Result is:");
                Console.WriteLine(resultString);
                Console.WriteLine($"Would you like to {encMethod}  with {encName} again?");
                Console.WriteLine("[Y]es");
                Console.WriteLine("Otherwise, press any button");
                var buff = Console.ReadLine();
                buff = buff?.Trim().ToUpper();
                continueLooping = buff == "Y";
            }
        }

        public static bool EncryptDecryptMenu(EncType encType)
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
                RotationUtils.StringType stringType = buff == "E" ? RotationUtils.StringType.PlainText : RotationUtils.StringType.CipherText;
                VigenereCaesarMenu(encType, stringType);
            }
        }
    }
}