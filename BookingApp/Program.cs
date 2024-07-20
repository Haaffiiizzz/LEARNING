using System;

namespace BookingApp
{
    class Program
    {
        static void Main(string[] args)
        {
            string eventName = "Go Conference";
            const int totalTickets = 50;
            int ticketsLeft = totalTickets;

            Console.WriteLine($"Welcome to the {eventName} booking application");
            Console.WriteLine($"There are a total of {totalTickets} tickets. There are {ticketsLeft} left.");

            while (ticketsLeft > 0)
            {
                Console.Write("Enter your name: ");
                string userName = Console.ReadLine();

                Console.Write("Enter the number of tickets you want to book: ");
                bool isValidNumber = int.TryParse(Console.ReadLine(), out int userTickets);

                if (isValidNumber && userTickets > 0 && userTickets <= ticketsLeft)
                {
                    ticketsLeft -= userTickets;
                    Console.WriteLine($"Thank you {userName}! You have successfully booked {userTickets} tickets.");
                    Console.WriteLine($"{ticketsLeft} tickets are left.\n");
                }
                else
                {
                    Console.WriteLine("Invalid number of tickets. Please try again.\n");
                }
            }

            Console.WriteLine("All tickets are booked. Thank you!");
        }
    }
}
