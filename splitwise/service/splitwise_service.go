package service

import (
	"fmt"
	"splitwise/entities"
	"sync"
)

type SplitwiseService struct {
	users  map[string]*entities.User
	groups map[string]*entities.Group
	mu     sync.Mutex
}

var (
	instance *SplitwiseService
	once     sync.Once
)

func GetInstance() *SplitwiseService {
	once.Do(func() {
		instance = &SplitwiseService{
			users:  make(map[string]*entities.User),
			groups: make(map[string]*entities.Group),
		}
	})

	return instance
}

func (s *SplitwiseService) AddUser(id, name, email string) *entities.User {
	user := entities.NewUser(id, name, email)
	s.mu.Lock()
	defer s.mu.Unlock()
	s.users[user.GetId()] = user
	return user
}

func (s *SplitwiseService) AddGroup(id, name string, members []*entities.User) *entities.Group {
	group := entities.NewGroup(id, name, members)
	s.mu.Lock()
	defer s.mu.Unlock()
	s.groups[group.GetId()] = group
	return group
}

func (s *SplitwiseService) GetUser(id string) (*entities.User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, ok := s.users[id]
	return user, ok
}

func (s *SplitwiseService) GetGroup(id string) *entities.Group {
	s.mu.Lock()
	defer s.mu.Unlock()
	group, _ := s.groups[id]
	return group
}

func (s *SplitwiseService) CreateExpense(b *entities.ExpenseBuilder) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	expense, err := b.Build()
	if err != nil {
		return err
	}

	paidBy := expense.GetPaidBy()

	splits := expense.GetSplits()
	for _, s := range splits {
		participant := s.GetUser()
		amount := s.GetAmount()
		if paidBy != participant {
			paidBy.GetBalanceSheet().AdjustBalance(participant, amount)
			participant.GetBalanceSheet().AdjustBalance(paidBy, -amount)
		}
	}

	fmt.Printf("Expense Created with amout:  %v, and description: %v\n", expense.GetAmount(), expense.GetDescription())

	return nil
}

func (s *SplitwiseService) SettleUp(payerId, payeeId string, amount float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	payer, ok := s.users[payerId]
	if !ok {
		return fmt.Errorf("Payer not found!!! No user found with id: %v", payerId)
	}

	payee, ok := s.users[payeeId]
	if !ok {
		return fmt.Errorf("Payee not found!!! No user found with id: %v", payeeId)
	}

	fmt.Printf("User: %v is settling with user: %v and the amount is %v\n", payer.GetName(), payee.GetName(), amount)

	payer.GetBalanceSheet().AdjustBalance(payee, amount)
	payee.GetBalanceSheet().AdjustBalance(payer, -amount)

	return nil
}

func (s *SplitwiseService) ShowBalanceSheet(userId string) {
	s.mu.Lock()
	user, ok := s.users[userId]
	s.mu.Unlock()
	if !ok {
		fmt.Printf("No user found with id: %v", user)
		return
	}

	user.GetBalanceSheet().ShowBalances()
}

// will code later
// func (s *SplitwiseService) SimplifyGroubDebts(groupId string) ([]*entities.Transaction, error) {
// 	group, ok := s.groups[groupId]
// 	if !ok {
// 		return nil, errors.New("No group found with this id")
// 	}

// }
