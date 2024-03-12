package main

import (
	"fmt"
)

type Rule struct {
	Condition func() bool
	Action   func()
}

type RuleEngine struct {
	Rules []Rule
}

func (re *RuleEngine) AddRule(rule Rule) {
	re.Rules = append(re.Rules, rule)
}

func (re *RuleEngine) Run() {
	for _, rule := range re.Rules {
		if rule.Condition() {
			rule.Action()
		}
	}
}

func main() {
	re := RuleEngine{}

	// Add a rule that prints "Hello, world!" if the current time is after 12pm.
	re.AddRule(Rule{
		Condition: func() bool {
			return time.Now().Hour() > 12
		},
		Action: func() {
			fmt.Println("Hello, world!")
		},
	})

	// Run the rule engine.
	re.Run()
}
