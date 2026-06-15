package main

import (
	"fmt"
	"moviebooking/booking"
	"time"
)

func main() {
	svc := booking.GetInstance()

	// 1. Cities.
	nyc := svc.AddCity("New York")
	_ = svc.AddCity("Los Angeles")

	// 2. Movies.
	matrix := booking.NewMovie("The Matrix", 120*time.Minute)
	avengers := booking.NewMovie("Avengers: Endgame", 170*time.Minute)
	svc.AddMovie(matrix)
	svc.AddMovie(avengers)

	// 3. A screen with two rows (A and B) of 10 seats; the back half is premium.
	screen := booking.NewScreen()
	for i := 1; i <= 10; i++ {
		seatType := booking.Regular
		if i > 5 {
			seatType = booking.Premium
		}
		screen.AddSeat(booking.NewSeat(1, i, seatType))
		screen.AddSeat(booking.NewSeat(2, i, seatType))
	}

	// 4. Cinema.
	if _, err := svc.AddCinema("AMC Times Square", nyc.Id, []*booking.Screen{screen}); err != nil {
		fmt.Println("failed to add cinema:", err)
		return
	}

	// 5. Shows.
	svc.AddShow(matrix, screen, time.Now().Add(2*time.Hour), booking.NewWeekendPricingStrategy())
	svc.AddShow(avengers, screen, time.Now().Add(5*time.Hour), booking.NewWeekdayPricingStrategy())

	alice := svc.CreateUser("Alice", "alice@example.com")

	fmt.Println("\n--- Alice's booking flow ---")
	shows := svc.FindShows("Avengers: Endgame", "New York")
	if len(shows) == 0 {
		fmt.Println("no shows found")
		return
	}
	selected := shows[0]

	available := availableSeats(selected)
	fmt.Printf("Available seats for %q at %s: %v\n",
		selected.Movie.Title, selected.StartTime.Format(time.Kitchen), seatIDs(available))

	desired := []*booking.Seat{available[2], available[3]}
	fmt.Printf("Alice selects seats: %v\n", seatIDs(desired))

	b, err := svc.BookTickets(alice.Id, selected.Id, desired,
		booking.NewCreditCardPayment("1234-5678-9876-5432", 123))
	if err != nil {
		fmt.Println("Booking failed:", err)
	} else {
		fmt.Println("\n--- Booking successful! ---")
		fmt.Println("Booking ID:    ", b.Id)
		fmt.Println("User:          ", b.User.Name)
		fmt.Println("Movie:         ", b.Show.Movie.Title)
		fmt.Printf("Seats:          %v\n", seatIDs(b.Seats))
		fmt.Printf("Total amount:   $%.2f\n", b.TotalAmount)
		fmt.Println("Payment status:", b.Payment.Status)
	}

	fmt.Println("\nSeat status after Alice's booking:")
	for _, s := range desired {
		fmt.Printf("  seat %s: %s\n", s.Id, &s.SeatStatus)
	}

}

func availableSeats(show *booking.Show) []*booking.Seat {
	var out []*booking.Seat
	for _, s := range show.Screen.Seats {
		if s.SeatStatus == booking.Available {
			out = append(out, s)
		}
	}
	return out
}

func seatIDs(seats []*booking.Seat) []string {
	ids := make([]string, len(seats))
	for i, s := range seats {
		ids[i] = s.Id
	}
	return ids
}
