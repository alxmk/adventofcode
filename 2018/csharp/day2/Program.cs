using System;
using System.Collections.Generic;

namespace day2
{
    class Program
    {
        static void Main(string[] args)
        {
            string[] lines = System.IO.File.ReadAllLines(@"input.txt");

            int twoCount = 0, threeCount = 0;

            // Part one
            foreach (string line in lines)
            {
                Dictionary<char, int> charCount = new Dictionary<char, int>();

                foreach (char c in line)
                {
                    int count;
                    if (!charCount.TryGetValue(c, out count))
                    {
                        charCount.Add(c, 1);
                        continue;
                    }

                    charCount[c]++;
                }

                bool foundTwo = false, foundThree = false;
                foreach (KeyValuePair<char, int> kv in charCount)
                {
                    if (kv.Value == 2 && !foundTwo)
                    {
                        twoCount++;
                        foundTwo = true;
                    }
                    if (kv.Value == 3 && !foundThree)
                    {
                        threeCount++;
                        foundThree = true;
                    }
                    if (foundTwo && foundThree)
                    {
                        break;
                    }
                }
            }

            // Part two
            int index = 0, matching = 0;
            foreach (string line in lines)
            {
                bool found = false;
                for (int i = index; i < lines.Length; i++)
                {
                    // No need to compare a string with itself
                    if (i == index)
                        continue;

                    int pos = 0, diff = 0;
                    foreach (char c in line)
                    {
                        // System.Console.WriteLine("Comparing: {0} {1}", line, lines[i]);
                        if (c != lines[i][pos++])
                        {
                            // System.Console.WriteLine("Chars: {0} {1}", c, lines[i][pos - 1]);
                            diff++;
                        }
                        if (diff == 2)
                            break;
                    }

                    if (diff == 1)
                    {
                        found = true;
                        matching = i;
                        break;
                    }

                }
                if (found)
                    break;

                index++;
            }

            System.Console.WriteLine("Success: {0} {1} {2}", twoCount * threeCount, lines[index], lines[matching]);
        }
    }
}
