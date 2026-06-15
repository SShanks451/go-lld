package booking

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

const defaultLockTimeout = 5 * time.Minute

type MovieBookingService struct {
	mu             sync.RWMutex
	users          map[string]*User
	cities         map[string]*City
	cinemas        map[string]*Cinema
	movies         map[string]*Movie
	shows          map[string]*Show
	lockManager    *SeatLockManager
	bookingManager *BookingManager
}

var (
	instance *MovieBookingService
	once     sync.Once
)

func GetInstance() *MovieBookingService {
	once.Do(func() {
		lockManager := NewSeatLockManager(defaultLockTimeout)
		instance = &MovieBookingService{
			users:          make(map[string]*User),
			cities:         make(map[string]*City),
			cinemas:        make(map[string]*Cinema),
			movies:         make(map[string]*Movie),
			shows:          make(map[string]*Show),
			lockManager:    lockManager,
			bookingManager: NewBookingManager(lockManager),
		}
	})

	return instance
}

func (mbs *MovieBookingService) AddCity(name string) *City {
	city := NewCity(name)
	mbs.mu.Lock()
	mbs.cities[city.Id] = city
	mbs.mu.Unlock()

	return city
}

func (mbs *MovieBookingService) AddCinema(name, cityId string, screens []*Screen) (*Cinema, error) {
	mbs.mu.Lock()
	defer mbs.mu.Unlock()

	city, ok := mbs.cities[cityId]
	if !ok {
		return nil, errors.New("City not found with this id")
	}

	cinema := NewCinema(name, city, screens)
	mbs.cinemas[cinema.Id] = cinema

	return cinema, nil
}

func (mbs *MovieBookingService) AddMovie(movie *Movie) {
	mbs.mu.Lock()
	defer mbs.mu.Unlock()

	mbs.movies[movie.Id] = movie
}

func (mbs *MovieBookingService) AddShow(movie *Movie, screen *Screen, startTime time.Time, pricingStrategy PricingStartegy) *Show {
	show := NewShow(movie, screen, startTime, pricingStrategy)

	mbs.mu.Lock()
	mbs.shows[show.Id] = show
	mbs.mu.Unlock()

	return show
}

func (mbs *MovieBookingService) CreateUser(name, email string) *User {
	user := NewUser(name, email)

	mbs.mu.Lock()
	mbs.users[user.Id] = user
	mbs.mu.Unlock()

	return user
}

func (mbs *MovieBookingService) BookTickets(userId, showId string, seats []*Seat, paymentStrategy PaymentStrategy) (*Booking, error) {
	mbs.mu.Lock()
	user, userOk := mbs.users[userId]
	show, showOk := mbs.shows[showId]
	mbs.mu.Unlock()

	if !userOk {
		return nil, errors.New("User not found")
	}

	if !showOk {
		return nil, errors.New("Show not found")
	}

	booking, err := mbs.bookingManager.CreateBooking(user, show, seats, paymentStrategy)
	if err != nil {
		return nil, fmt.Errorf("Error while booking: %v", err)
	}

	return booking, nil
}

func (mbs *MovieBookingService) FindShows(movieTitle, cityName string) []*Show {
	mbs.mu.Lock()
	defer mbs.mu.Unlock()

	var result []*Show
	for _, show := range mbs.shows {
		if strings.EqualFold(show.Movie.Title, movieTitle) {
			for _, cinema := range mbs.cinemas {
				for _, screen := range cinema.Screens {
					if screen == show.Screen && strings.EqualFold(cinema.City.Name, cityName) {
						result = append(result, show)
					}
				}
			}
		}
	}

	return result
}
