using System;
using System.Collections.Generic;
using System.Text;

namespace HW_1_Ancient_Crypto
{
    class Program
    {
        private const int MAX_UTF32 = 0x10FFFF;
        private enum StringType
        {
            
        }

        static string GetInputString( bool  )
        {
            bool isValid = false;
            var plaintext = "";
            while (!isValid)
            {
                Console.WriteLine("Please input the plaintext, or 'C' to cancel: ");
                plaintext = Console.ReadLine() ?? "";
                if (plaintext == "")
                {
                    Console.WriteLine("Not a valid string!");
                    continue;
                }
                isValid = true;
            }
            return plaintext;
        }

        static void doCaesar()
        {
            var inputIsValid = false;
            var shiftString = "";
            var shiftInt = 0;
            do
            {
                Console.WriteLine("Please input the shift, or 'C' to cancel: ");
                shiftString = Console.ReadLine();
                if (!int.TryParse(shiftString, out shiftInt) && shiftString?.Trim().ToUpper() != "C")
                {
                    Console.WriteLine("Not a valid shift!");
                    continue;
                }
                Console.WriteLine($"Caesar shift: {shiftInt}");
                
                
                inputIsValid = true;
                var enc = new UTF32Encoding();
                
                var ciphertext = "";
                foreach (int unicodeCodePoint in plaintext.GetUnicodeCodePoints())
                {
                    ciphertext += char.ConvertFromUtf32((unicodeCodePoint + shiftInt) % MAX_UTF32);
                }
                Console.WriteLine("Ciphertext is: " + ciphertext);
            } while (!inputIsValid);
        }

        static void Main(string[] args)
        {
            string s = "a🌀c🏯";
            foreach(int unicodeCodePoint in s.GetUnicodeCodePoints())
            {
                Console.WriteLine(unicodeCodePoint);
                Console.WriteLine(char.ConvertFromUtf32(unicodeCodePoint));
                Console.WriteLine(char.ConvertFromUtf32(unicodeCodePoint + 10));
            }
            Console.WriteLine("EZ Encryption EZE");
            
            // bool isOk = true;
            // for (int i = 0; i < buff.Length; i++)
            // {
            //     Console.WriteLine();
            //     isOk = (buff[i] - '0' > 64 && buff[i] - '0' < 91) || (buff[i] - '0' > 96 && buff[i] - '0' < 123);
            // }
            String buff;
            do
            {
                Console.WriteLine("Please Select the encryption:");
                Console.WriteLine("C for Caesar");
                Console.WriteLine("V for Vigenere");
                Console.WriteLine("E for Exit");
                buff = Console.ReadLine();
                buff = buff?.Trim().ToUpper();
                if (buff != "E")
                {
                    Console.WriteLine("Enter 'E' for Encryption, 'D' for Decryption");
                    String ed  = Console.ReadLine();
                    ed = ed?.Trim().ToUpper();
                    // Console.WriteLine("Enter plaintext:");
                    // String plaintext = Console.ReadLine();
                    if (buff == "C")
                    {
                        doCaesar();
                    }
                    else
                    {
                        Console.WriteLine("Please input the Cipher:");
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