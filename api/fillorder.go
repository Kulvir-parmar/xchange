package api

/*
User1 -> receives (quantity * price USDT) and pays (quantity BTC)
User2 -> receives (quantity BTC) and pays (quantity * price USDT)
*/
func flipBalance(userId1, userId2 string, price, quantity float64) {
	var user1, user2 *User

	/*
		IRL users will be fetched from db.
		Most likely userId will be indexed for faster lookup.
	*/
	for _, user := range Users {
		switch user.UserId {
		case userId1:
			user1 = &user
		case userId2:
			user2 = &user
		}
	}

	user1.Balance["USDT"] += quantity * price
	user1.Balance["BTC"] -= quantity

	user2.Balance["USDT"] -= quantity * price
	user2.Balance["BTC"] += quantity
}

func fillOrder(order Order) float64 {
	remainingQuantity := order.Quantity

	if order.Side == "BID" {
		for idx, ask := range asks {
			if ask.Price > order.Price {
				continue
			}

			if ask.Quantity > remainingQuantity {
				asks[idx].Quantity -= remainingQuantity
				flipBalance(ask.UserId, order.UserId, ask.Price, remainingQuantity)
				return 0.0

			} else {
				remainingQuantity -= ask.Quantity
				flipBalance(ask.UserId, order.UserId, ask.Price, ask.Quantity)
				asks = append(asks[:idx], asks[idx+1:]...)
			}
		}

	} else if order.Side == "ASK" {
		for idx, bid := range bids {
			if bid.Price < order.Price {
				continue
			}

			if bid.Quantity > remainingQuantity {
				bid.Quantity -= remainingQuantity
				flipBalance(order.UserId, bid.UserId, order.Price, remainingQuantity)
				return 0.0

			} else {
				remainingQuantity -= bid.Quantity
				flipBalance(order.UserId, bid.UserId, order.Price, bid.Quantity)
				bids = append(bids[:idx], bids[idx+1:]...)
			}
		}
	}

	return remainingQuantity
}
