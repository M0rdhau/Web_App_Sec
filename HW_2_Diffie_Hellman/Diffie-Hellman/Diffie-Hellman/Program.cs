using System;
using static Diffie_Hellman.menus.RotationMenu;
using static Diffie_Hellman.utils.CryptoUtils;

namespace Diffie_Hellman
{
    class Program
    {
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
            // MainMenu();
        }
    }
}