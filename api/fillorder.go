package api

import "fmt"

func getUser(userId string) (User, error) {
	for _, user := range users {
		if user.UserId == userId {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("User not found")
}

func flipBalance(userId1, userId2 string, price, quantity float64) {
	user1, err := getUser(userId1)
	if err != nil {
		fmt.Println(err)
		return
	}

	user2, err := getUser(userId2)
	if err != nil {
		fmt.Println(err)
		return
	}

	user1.Balance[TICKER] += quantity
	user1.Balance["USDT"] -= price * quantity

	user2.Balance[TICKER] -= quantity
	user2.Balance["USDT"] += price * quantity
}

func fillOrder(order Order) float64 {
	var remainingQuantity float64

	if order.Side == "BID" {
		for _, askOrder := range asks {
			if askOrder.Price > order.Price {
				continue
			}

			if askOrder.Quantity > order.Quantity {
				askOrder.Quantity -= order.Quantity
				flipBalance(order.UserId, askOrder.UserId, order.Price, order.Quantity)
				return 0.0
			} else {
				remainingQuantity = order.Quantity - askOrder.Quantity
				flipBalance(order.UserId, askOrder.UserId, order.Price, askOrder.Quantity)
				asks = asks[1:]
			}
		}

	} else if order.Side == "ASK" {
		for _, bidOrder := range bids {
			if bidOrder.Price < order.Price {
				continue
			}

			if bidOrder.Quantity > order.Quantity {
				bidOrder.Quantity -= order.Quantity
				flipBalance(bidOrder.UserId, order.UserId, order.Price, order.Quantity)
				return 0.0
			} else {
				remainingQuantity = order.Quantity - bidOrder.Quantity
				flipBalance(bidOrder.UserId, order.UserId, order.Price, bidOrder.Quantity)
				bids = bids[1:]
			}
		}
	}

	return remainingQuantity
}
