using System;
using System.Collections.Generic;

namespace day1
{
    class Program
    {
        static void Main(string[] args)
        {
            string[] lines = System.IO.File.ReadAllLines(@"input.txt");

            bool done1 = false, done2 = false;
            int i = 0, freq = 0, part1 = 0, part2 = 0;

            Dictionary<int, bool> seen = new Dictionary<int,bool>();

            while (!(done1 && done2)) {
                int num = 0;
                if (!Int32.TryParse(lines[i%lines.Length], out num)) {
                    System.Console.WriteLine("{0} not parseable", lines[i%lines.Length]);
                    return;
                }

                freq += num;

                if (seen.ContainsKey(freq)) {
                    part2 = freq;
                    done2 = true;
                } else {
                    seen.Add(freq, true);
                }

                if (++i == lines.Length) {
                    part1 = freq;
                    done1 = true;
                }
            }
            System.Console.WriteLine("Success: {0} {1}", part1, part2);
        }
    }
}
