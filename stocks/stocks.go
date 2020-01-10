package stocks

type StockEnv struct {
	steps      int     //how many steps have been taken
	reward     int     //what reward was received for step
	shares     int     //how many shares are owned
	funds      float64 //available to buy things with
	totalFunds float64 //total portfolio
	done       bool
}
