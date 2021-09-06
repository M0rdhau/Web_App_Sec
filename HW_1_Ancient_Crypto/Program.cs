using System;
using System.Text;

namespace HW_1_Ancient_Crypto
{
    class Program
    {
        static void doCaesar()
        {
            var inputIsValid = false;
            var shiftString = "";
            var shiftInt = 0;
            Console.WriteLine("caeeesar");
            do
            {
                Console.WriteLine("Please input the shift, or 'C' to cancel: ");
                shiftString = Console.ReadLine();
                if (!int.TryParse(shiftString, out shiftInt) && shiftString?.Trim().ToUpper() != "C")
                {
                    Console.WriteLine("Not a valid string!");
                    continue;
                }
                Console.WriteLine($"Caesar shift: {shiftInt}");
                
                Console.WriteLine("Please input the plaintext, or 'C' to cancel: ");
                var plaintext = Console.ReadLine() ?? "";
                var utf8 = new UTF8Encoding();
                var utf32 = new UTF32Encoding();
                var utf8bytes = utf8.GetBytes(plaintext);
                var utf32bytes = utf32.GetBytes(plaintext);
                // for (int i = 0; i < utf8bytes.Length; i++)
                // {
                //     Console.WriteLine(utf8bytes[i]);
                // }
                // for (int i = 0; i < utf32bytes.Length; i++)
                // {
                //     Console.WriteLine(utf32bytes[i]);
                // }
                
                var base64Str = System.Convert.ToBase64String(utf8bytes);
                byte[] bytes = Convert.FromBase64String(base64Str);
                for (int i = 0; i < bytes.Length; i++)
                {
                    Console.WriteLine(bytes[i]);
                }

            } while (!inputIsValid);
        }

        static void Main(string[] args)
        {
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
                    Console.WriteLine("Please enter one of the correct options.");
                    Console.WriteLine("C for Caesar");
                    Console.WriteLine("V for Vigenere");
                    Console.WriteLine("E for Exit");
                    buff = Console.ReadLine();
                    buff = buff?.Trim().ToUpper();
                }
            } while (buff != "E");

            Console.WriteLine(buff);
        }
    }
}