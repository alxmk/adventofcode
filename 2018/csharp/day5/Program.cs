using System;

namespace day5
{
    class Program
    {
        static int diff = 'a' - 'A';
        static void Main(string[] args)
        {
            string input = System.IO.File.ReadAllText(@"input.txt");

            System.Console.WriteLine("Part one: {0}", react(input).Length);

            int min = int.MaxValue;

            for (char c = 'a'; c <= 'z'; c++)
            {
                int len = react(input.Replace(c.ToString(), string.Empty).Replace(char.ToUpper(c).ToString(), string.Empty)).Length;
                if (len < min)
                    min = len;
            }

            System.Console.WriteLine("Part two: {0}", min);
        }

        static bool match(char a, char b)
        {
            return ((a - b) == diff) || ((b - a) == diff);
        }

        static string react(string input)
        {
            int index = 0;

            while (true)
            {
                if (index + 1 >= input.Length)
                    break;

                if (match(input[index], input[index + 1]))
                {
                    input = input.Remove(index, 2);
                    if (index > 0)
                        index--;
                }
                else
                    index++;
            }

            return input;
        }
    }
}
