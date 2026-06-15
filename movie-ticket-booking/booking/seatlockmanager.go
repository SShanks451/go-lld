package booking

import (
	"errors"
	"fmt"
	"slices"
	"sync"
	"time"
)

type SeatLockManager struct {
	mu      sync.Mutex
	timeout time.Duration
	locks   map[string]map[string]string
	timers  map[string]*time.Timer
	closed  bool
}

func NewSeatLockManager(timeout time.Duration) *SeatLockManager {
	return &SeatLockManager{
		timeout: timeout,
		locks:   make(map[string]map[string]string),
		timers:  map[string]*time.Timer{},
	}
}

func (slm *SeatLockManager) LockSeats(show *Show, seats []*Seat, userId string) error {
	slm.mu.Lock()
	defer slm.mu.Unlock()

	if slm.closed {
		return errors.New("Locking manager is closed.")
	}

	for _, seat := range seats {
		if seat.SeatStatus != Available {
			return fmt.Errorf("Seat id: %v is locked", seat.Id)
		}
	}

	_, ok := slm.locks[show.Id]
	if !ok {
		slm.locks[show.Id] = map[string]string{}
	}

	for _, seat := range seats {
		seat.SeatStatus = Locked
		slm.locks[show.Id][seat.Id] = userId
	}

	seatsCopy := slices.Clone(seats)
	key := slm.createKey(show.Id, userId)
	slm.timers[key] = time.AfterFunc(slm.timeout, func() {
		slm.mu.Lock()
		defer slm.mu.Unlock()
		slm.releaseLocked(show.Id, seatsCopy, userId)
	})

	var ids []string
	for _, seat := range seats {
		ids = append(ids, seat.Id)
	}

	fmt.Printf("For user id [%v], following seats are locked: [I%v]", userId, ids)

	return nil
}

func (slm *SeatLockManager) UnlockSeats(show *Show, seats []*Seat, userId string) {
	slm.mu.Lock()
	defer slm.mu.Unlock()

	key := slm.createKey(show.Id, userId)
	slm.clearTimer(key)
	slm.releaseLocked(show.Id, seats, userId)
}

func (slm *SeatLockManager) ConfirmSeats(show *Show, seats []*Seat, userId string) error {
	slm.mu.Lock()
	defer slm.mu.Unlock()

	showlocks := slm.locks[show.Id]
	for _, seat := range seats {
		if owner, ok := showlocks[seat.Id]; !ok || owner != userId || seat.SeatStatus != Locked {
			return fmt.Errorf("For user id [%v], lock is not found on seat id [%v]", userId, seat.Id)
		}
	}

	for _, seat := range seats {
		seat.SeatStatus = Booked
		delete(showlocks, seat.Id)
	}

	if len(showlocks) == 0 {
		delete(slm.locks, show.Id)
	}

	slm.clearTimer(slm.createKey(show.Id, userId))

	return nil
}

func (slm *SeatLockManager) releaseLocked(showId string, seats []*Seat, userId string) {
	showlocks, ok := slm.locks[showId]
	if !ok {
		return
	}

	for _, seat := range seats {
		if owner, ok := showlocks[seat.Id]; ok && owner == userId {
			delete(showlocks, owner)
			if seat.SeatStatus == Locked {
				seat.SeatStatus = Available
			}
		}
	}

	if len(showlocks) == 0 {
		delete(slm.locks, showId)
	}
}

func (slm *SeatLockManager) createKey(showId, userId string) string {
	return showId + "_" + userId
}

func (slm *SeatLockManager) clearTimer(key string) {
	if timer, ok := slm.timers[key]; ok {
		timer.Stop()
		delete(slm.timers, key)
	}
}
