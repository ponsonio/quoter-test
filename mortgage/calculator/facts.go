package calculator

type AmortizationPeriod int64

const (
	Accelerated AmortizationPeriod = 0
	Biweekly    AmortizationPeriod = 1
	Monthly     AmortizationPeriod = 2
)

type MortgageCalc struct {
	TotalPropertyValue float64
	AmortizationPeriod AmortizationPeriod
	DownPayment        float64
	Period             int32
	Valid              bool
	RequiresCMHC       bool
	CMHCCost           float32
	PaymentAmount      float64
	Errors             []string
}

func (mc *MortgageCalc) AddError(err string) {
	mc.Errors = append(mc.Errors, err)
}
