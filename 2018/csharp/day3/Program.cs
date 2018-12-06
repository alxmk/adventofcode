using System;

namespace day3
{
    class Program
    {
        static void Main(string[] args)
        {
            string[] lines = System.IO.File.ReadAllLines(@"input.txt");

            int[,] sheet = new int[1000, 1000];

            foreach (string line in lines)
            {
                int[] coords = Split(line);

                // System.Console.WriteLine("{0},{1} {2}x{3}", coords[0], coords[1], coords[2], coords[3]);

                for (int x = coords[0]; x < coords[0] + coords[2]; x++)
                {
                    for (int y = coords[1]; y < coords[1] + coords[3]; y++)
                    {
                        // System.Console.WriteLine("{0},{1}++", x, y);
                        sheet[x, y]++;
                    }
                }
            }

            int count = 0;

            foreach (int value in sheet)
            {
                if (value > 1)
                {
                    count++;
                }
            }

            System.Console.WriteLine("Part 1: {0}", count);

            foreach (string line in lines)
            {
                int[] coords = Split(line);

                bool found = true;

                for (int x = coords[0]; x < coords[0] + coords[2]; x++)
                {
                    for (int y = coords[1]; y < coords[1] + coords[3]; y++)
                    {
                        if (sheet[x, y] != 1)
                        {
                            found = false;
                            break;
                        }
                    }
                }

                if (found)
                {
                    System.Console.WriteLine(line);
                    break;
                }
            }
        }

        static int[] Split(string line)
        {
            string[] parts = line.Split(new char[] { ' ', ',', 'x', ':' });

            return new int[] { Int32.Parse(parts[2]), Int32.Parse(parts[3]), Int32.Parse(parts[5]), Int32.Parse(parts[6]) };
        }
    }
}
