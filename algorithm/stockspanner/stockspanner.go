package stockspanner

type day struct {
	price int
	span  int
}

type StockSpanner struct {
	days []*day
}

func Constructor() StockSpanner {
	return StockSpanner{
		days: make([]*day, 0, 10000),
	}
}

func (this *StockSpanner) Next(price int) int {
	span := 1
	for len(this.days) > 0 {
		lastDay := this.days[len(this.days)-1]
		if lastDay.price > price {
			break
		} else {
			// pop lastDay
			this.days = this.days[:len(this.days)-1]
			span += lastDay.span
		}
	}
	this.days = append(this.days, &day{
		price: price,
		span:  span,
	})
	return span
}
