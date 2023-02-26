package calculator

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

var knowledgeLibrary = *ast.NewKnowledgeLibrary()

type MortgageCalculatorService struct {
}

func NewRuleEngineSvc() *MortgageCalculatorService {
	// you could add your cloud provider here instead of keeping rule file in your code.
	buildRuleEngine()
	return &MortgageCalculatorService{}
}

func buildRuleEngine() {
	ruleBuilder := builder.NewRuleBuilder(&knowledgeLibrary)

	// Read rule from file and build rules
	ruleFile := pkg.NewFileResource("rules.grl")
	err := ruleBuilder.BuildRuleFromResource("Rules", "0.0.1", ruleFile)
	if err != nil {
		panic(err)
	}

}

func (svc *MortgageCalculatorService) Execute(mortgageCalc MortgageCalc) error {
	dataCtx := ast.NewDataContext()
	err := dataCtx.Add("MC", mortgageCalc)
	if err != nil {
		panic(err)
	}

	knowledgeBase := knowledgeLibrary.NewKnowledgeBaseInstance("MortgageRules", "0.0.1")

	engine := engine.NewGruleEngine()
	err = engine.Execute(dataCtx, knowledgeBase)
	if err != nil {
		panic(err)
	}
	return nil
}
