//package main
//
//import (
//	"github.com/aws/aws-sdk-go/aws/session"
//	"github.com/aws/aws-sdk-go/service/budgets"
//)
//
//func describe() {
//	svc := budgets.New(session.New())
//	accountId := "660262896402"
//	budgetName := "MonthlyCost"
//	input := &budgets.DescribeBudgetInput{
//		AccountId: &accountId,
//		BudgetName: &budgetName,
//	}
//	svc.DescribeBudget(input)
//
//}
