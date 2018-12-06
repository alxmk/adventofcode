using System;

namespace day4
{
    class Program
    {
        static void Main(string[] args)
        {
            string[] lines = System.IO.File.ReadAllLines(@"input.txt");

            int guard = 0, asleepAt = 0;
            int[,] sleepSchedule = new int[10000, 60];

            foreach (string line in lines)
            {
                int minute = parseMinute(line);

                if (line.Contains("Guard"))
                {
                    guard = parseGuard(line);
                }
                if (line.Contains("falls asleep"))
                {
                    asleepAt = minute;
                }
                if (line.Contains("wakes up"))
                {
                    for (int i = asleepAt; i < minute; i++)
                    {
                        sleepSchedule[guard, i]++;
                    }
                    asleepAt = -1;
                }
            }

            int sleepyGuard = 0, minutesAsleep = 0, sleepiestMinute = 0, overallSleepiestGuard = 0, overallSleepiestMinute = 0, overallSleepiestAmount = 0;
            for (int g = 0; g < sleepSchedule.GetLength(0); g++)
            {
                int thisGuardMins = 0, thisGuardSleepiest = 0, thisGuardSleepiestMinute = 0;
                for (int minute = 0; minute < sleepSchedule.GetLength(1); minute++)
                {
                    int sleepyMinutes = sleepSchedule[g, minute];
                    thisGuardMins += sleepyMinutes;
                    if (sleepyMinutes > thisGuardSleepiest)
                    {
                        thisGuardSleepiest = sleepyMinutes;
                        thisGuardSleepiestMinute = minute;
                    }
                    if (sleepyMinutes > overallSleepiestAmount)
                    {
                        overallSleepiestAmount = sleepyMinutes;
                        overallSleepiestGuard = g;
                        overallSleepiestMinute = minute;
                    }
                }

                if (thisGuardMins > minutesAsleep)
                {
                    minutesAsleep = thisGuardMins;
                    sleepyGuard = g;
                    sleepiestMinute = thisGuardSleepiestMinute;
                }
            }

            System.Console.WriteLine("Part 1: {0} {1} {2}; Part 2: {3}", sleepyGuard, sleepiestMinute, sleepyGuard * sleepiestMinute, overallSleepiestGuard * overallSleepiestMinute);
        }

        static int parseMinute(string line)
        {
            string[] parts = line.Split(new char[] { ' ', ':', ']' });

            return Int32.Parse(parts[2]);
        }

        static int parseGuard(string line)
        {
            string[] parts = line.Split(new char[] { ' ', '#' });

            return Int32.Parse(parts[4]);
        }
    }
}
