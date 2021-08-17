package main

import "fmt"

// 设想有一个戏剧演出团，演员们经常要去各种场合表演喜剧。通常，客户会指定几出剧目
// 而剧团则根据观众人数及剧目类型来向观众收费。该团目前出演两种戏剧：喜剧以及悲剧
// 给客户发出账单时，剧团还会根据到场观众的数量给出观众量积分优惠，下次客户再请
// 剧团表演时可以使用积分获得折扣

type Invoice struct {
	Customer     string
	Performances []Performance
}

type Performance struct {
	PlayID  string
	Audience int
}

type Play struct {
	Name string
	Type string
}

func main() {

}


// version one
func statement(invoice Invoice, plays map[string]Play)string {
	var totalAmount, volumeCredits, thisAmount int
	result := fmt.Sprintf("Statement for %s \n", invoice.Customer)
	for _, performance := range invoice.Performances {
		play := plays[performance.PlayID]
		switch play.Type {
		case "tragedy":
			thisAmount=40000
			if performance.Audience>30{
				thisAmount+=1000*(performance.Audience-30)
			}
		case "comedy":
			thisAmount=30000
			if performance.Audience>20{
				thisAmount+=10000+500*(performance.Audience-20)
			}
			thisAmount+=300*performance.Audience
		default:
		}
		if performance.Audience>30{
			volumeCredits+=performance.Audience
		}
		if "comedy"==play.Type{
			volumeCredits += performance.Audience/5
		}
		result+=fmt.Sprintf("%s: %d %d seats", play.Name, thisAmount/100, performance.Audience)
		totalAmount+=thisAmount
	}
	result += fmt.Sprintf("Amount owed is %d", totalAmount / 100)
	result += fmt.Sprintf("you earn %d credits", volumeCredits)
	return result
}


// version two: 增加适量的缩进: 可读性上升少许
func statementV2(invoice Invoice, plays map[string]Play)string {
	var totalAmount, volumeCredits, thisAmount int
	result := fmt.Sprintf("Statement for %s \n", invoice.Customer)

	for _, performance := range invoice.Performances {
		play := plays[performance.PlayID]
		switch play.Type {
		case "tragedy":
			thisAmount = 40000
			if performance.Audience > 30{
				thisAmount += 1000 *(performance.Audience - 30)
			}
		case "comedy":
			thisAmount = 30000
			if performance.Audience > 20{
				thisAmount += 10000+500 * (performance.Audience - 20)
			}
			thisAmount += 300 * performance.Audience
		default:
			panic("panic")
		}

		if performance.Audience > 30{
			volumeCredits += performance.Audience
		}

		if "comedy" == play.Type{
			volumeCredits += performance.Audience/5
		}

		result += fmt.Sprintf("%s: %d %d seats", play.Name, thisAmount/100, performance.Audience)
		totalAmount += thisAmount
	}
	result += fmt.Sprintf("Amount owed is %d", totalAmount / 100)
	result += fmt.Sprintf("you earn %d credits", volumeCredits)
	return result
}


// version three: 增加适量代码
func statementV3(invoice Invoice, plays map[string]Play)string {
	var totalAmount, volumeCredits, thisAmount int
	result := fmt.Sprintf("Statement for %s \n", invoice.Customer)

	for _, performance := range invoice.Performances {
		// 获取剧的属性
		play := plays[performance.PlayID]
		// 计算总金额
		switch play.Type {
		case "tragedy":
			thisAmount = 40000
			if performance.Audience > 30{
				thisAmount += 1000 *(performance.Audience - 30)
			}
		case "comedy":
			thisAmount = 30000
			if performance.Audience > 20{
				thisAmount += 10000+500 * (performance.Audience - 20)
			}
			thisAmount += 300 * performance.Audience
		default:
			panic("panic")
		}

		//增加 Credit
		if performance.Audience > 30{
			volumeCredits += performance.Audience
		}
		// 每5个额外观众，增加积分
		if "comedy" == play.Type{
			volumeCredits += performance.Audience/5
		}

		// 打印相应的参数
		result += fmt.Sprintf("%s: %d %d seats", play.Name, thisAmount/100, performance.Audience)
		totalAmount += thisAmount
	}
	result += fmt.Sprintf("Amount owed is %d", totalAmount / 100)
	result += fmt.Sprintf("you earn %d credits", volumeCredits)
	return result
}


// version five：每一段长代码中的所有注释，基本上都可以抽象出一个函数
func statementV3(invoice Invoice, plays map[string]Play)string {
	result := fmt.Sprintf("Statement for %s \n", invoice.Customer)
	for _, performance := range invoice.Performances {
		result += fmt.Sprintf("%s %d %d seats", playFor(performance, plays),
			amountFor(performance, playFor(performance, plays)), performance.Audience)
	}
	result += fmt.Sprintf("Amount owed is %d", totalAmount(invoice.Performances, plays))
	result += fmt.Sprintf("You earned: %d credits", totalCredits(invoice.Performances, plays))
	return result
}

func totalAmount(performances []Performance, plays map[string]Play) int {
	var result int
	for _, performance := range performances {
		result += amountFor(performance, playFor(performance, plays))
	}
	return result
}

func totalCredits(performances []Performance, plays map[string]Play) int {
	var result int
	for _, performance := range performances {
		result += volumeCreditsFor(performance, playFor(performance, plays))
	}
	return result
}

func amountFor(performance Performance, play Play) int {
	var thisAmount int
	switch play.Type {
	case "tragedy":
		thisAmount = 40000
		if performance.Audience > 30{
			thisAmount += 1000 *(performance.Audience - 30)
		}
	case "comedy":
		thisAmount = 30000
		if performance.Audience > 20{
			thisAmount += 10000+500 * (performance.Audience - 20)
		}
		thisAmount += 300 * performance.Audience
	default:
		panic("panic")
	}
	return thisAmount
}

func playFor(performance Performance, plays map[string]Play) Play {
	return plays[performance.PlayID]
}

func volumeCreditsFor(performance Performance, play Play) int {
	volumeCredits := 0

	if performance.Audience > 30 {
		volumeCredits += performance.Audience
	}
	if play.Type == "comedy" {
		volumeCredits += performance.Audience / 5
	}
	return volumeCredits
}