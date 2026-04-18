package main

type PricingStrategy interface {
	Calculate(ticket *Ticket) float64
}

type FixedPricingStrategy struct{}
type HourlyPricingStrategy struct{}

func (fpx *FixedPricingStrategy) Calculate(ticket *Ticket) float64 {
	return 100.00
}

func (fpx *HourlyPricingStrategy) Calculate(Ticket *Ticket) float64 {
	// some logic
	return 50.5
}

type CostComputation struct {
	Strategy PricingStrategy
}

func NewCostComputation(strategy PricingStrategy) *CostComputation {
	return &CostComputation{
		Strategy: strategy,
	}
}

func (cc *CostComputation) Compute(ticket *Ticket) float64 {
	return cc.Strategy.Calculate(ticket)
}
