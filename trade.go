package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

const (
	itemMANA = iota
	itemFOOD
)

const (
	manaPotionName = "Mana Potion"
	manaPotionDesc = "Insta-recover 5 mana"
	foodName       = "20 lb package of food"
	foodDesc       = "You eat it... to live... you know..."
)

type TradePrices struct {
	manaBuy  int
	manaSell int
	foodBuy  int
	foodSell int
}

var marketPrices = map[*Location]TradePrices{
	riga:          TradePrices{5, 3, 20, 16},
	warsaw:        TradePrices{10, 6, 21, 16},
	berlin:        TradePrices{12, 10, 21, 15},
	hanover:       TradePrices{12, 8, 22, 16},
	cologne:       TradePrices{13, 7, 22, 17},
	brussels:      TradePrices{15, 9, 28, 19},
	paris:         TradePrices{18, 10, 25, 19},
	bordeaux:      TradePrices{16, 11, 26, 22},
	san_sebastian: TradePrices{19, 3, 23, 19},
	madrid:        TradePrices{15, 2, 44, 28},
}

func openTradeMenu(g *gocui.Gui, v *gocui.View) error {
	if currentCity == nil {
		WorldState.logLock.Lock()
		WorldState.Log = append(WorldState.Log, LogEntry{WorldState.Date, "Not in a city, no one to trade with..."})
		WorldState.logLock.Unlock()
		logView, err := g.View("control")
		if err != nil {
			return err
		}
		WorldState.logLock.Lock()
		printLogsToView(WorldState.Log, logView)
		WorldState.logLock.Unlock()
		return nil
	}
	currentTrade = struct {
		manaBuy  int
		manaSell int
		foodBuy  int
		foodSell int
	}{
		0,
		0,
		0,
		0,
	}
	g.SetLayout(tradeLayout)
	return nil
}

func tradeLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("tradeBackground", 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Trading"
		p := WorldState.Player
		fmt.Fprintf(v, "You currently have:\n\tMoney: %d\n\tMana: %d\n\tFood: %d", p.Money, p.Mana, p.Food)
	}
	if v, err := g.SetView("tradeInfo", 0, maxY*3/4, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		printTradeInfo(0, v)
	}
	if v, err := g.SetView("tradingPrices", maxX/2+4, maxY/2-3, maxX/2+7, maxY/2+3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		pricelist := marketPrices[currentCity]
		fmt.Fprintln(v, pricelist.manaBuy)
		fmt.Fprintln(v, pricelist.manaSell)
		fmt.Fprintln(v, pricelist.foodBuy)
		fmt.Fprintln(v, pricelist.foodSell)
	}
	if v, err := g.SetView("tradingamount", maxX/2+7, maxY/2-3, maxX/2+10, maxY/2+3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, 0)
		fmt.Fprintln(v, 0)
		fmt.Fprintln(v, 0)
		fmt.Fprintln(v, 0)
	}
	if v, err := g.SetView("tradingtotals", maxX/2+10, maxY/2-3, maxX/2+16, maxY/2+3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, 0)
		fmt.Fprintln(v, 0)
		fmt.Fprintln(v, 0)
		fmt.Fprintln(v, 0)
		fmt.Fprintln(v, 0)
	}
	if v, err := g.SetView("tradingView", maxX/2-16, maxY/2-3, maxX/2+4, maxY/2+3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = fmt.Sprintf("Trading in %s", currentCity.name)
		v.Highlight = true
		fmt.Fprintln(v, "Buy Mana Potions")
		fmt.Fprintln(v, "Sell Mana Potions")
		fmt.Fprintln(v, "Buy Food Package")
		fmt.Fprintln(v, "Sell Food Package")
		fmt.Fprintln(v, "Total")
		err := g.SetCurrentView("tradingView")
		if err != nil {
			return err
		}
	}
	return nil
}

var currentTrade = struct {
	manaBuy  int
	manaSell int
	foodBuy  int
	foodSell int
}{
	0,
	0,
	0,
	0,
}

func currentTradeCost() int {
	pricelist := marketPrices[currentCity]
	return (currentTrade.manaBuy*pricelist.manaBuy + currentTrade.foodBuy*pricelist.foodBuy) - (currentTrade.manaSell*pricelist.manaSell + currentTrade.foodSell*pricelist.foodSell)
}

func updateTradeViews(g *gocui.Gui) error {
	v, err := g.View("tradingamount")
	if err != nil {
		return err
	}
	v.Clear()
	v.SetOrigin(0, 0)
	fmt.Fprintln(v, currentTrade.manaBuy)
	fmt.Fprintln(v, currentTrade.manaSell)
	fmt.Fprintln(v, currentTrade.foodBuy)
	fmt.Fprintln(v, currentTrade.foodSell)
	pricelist := marketPrices[currentCity]
	v, err = g.View("tradingtotals")
	if err != nil {
		return err
	}
	pricelist = marketPrices[currentCity]
	v.Clear()
	v.SetOrigin(0, 0)
	fmt.Fprintln(v, currentTrade.manaBuy*pricelist.manaBuy)
	fmt.Fprintln(v, currentTrade.manaSell*pricelist.manaSell)
	fmt.Fprintln(v, currentTrade.foodBuy*pricelist.foodBuy)
	fmt.Fprintln(v, currentTrade.foodSell*pricelist.foodSell)
	fmt.Fprintln(v, currentTradeCost())
	return nil
}

func printTradeInfo(index int, v *gocui.View) {
	v.Clear()
	v.SetOrigin(0, 0)
	switch index {
	case 0, 1:
		v.Title = manaPotionName
		fmt.Fprintln(v, manaPotionDesc)
	case 2, 3:
		v.Title = foodName
		fmt.Fprintln(v, foodDesc)
	}
	fmt.Fprint(v, "\nPress left to decrease count, right to increase count,\nand enter to finalize the deal\n")
	fmt.Fprint(v, "The 1st column is the price per unit, the 2nd column is the amount to buy/sell,\nand the final column is the amount to pay/earn.")
}

func tradeUp(g *gocui.Gui, v *gocui.View) error {
	_, index := v.Cursor()
	if index > 0 {
		index--
		err := v.SetCursor(0, index)
		if err != nil {
			return err
		}
		info, err := g.View("tradeInfo")
		if err != nil {
			return err
		}
		printTradeInfo(index, info)
	}
	return nil
}

func tradeDown(g *gocui.Gui, v *gocui.View) error {
	_, index := v.Cursor()
	if index < 3 {
		index++
		err := v.SetCursor(0, index)
		if err != nil {
			return err
		}
		info, err := g.View("tradeInfo")
		if err != nil {
			return err
		}
		printTradeInfo(index, info)
	}
	return nil
}

func tradeLeft(g *gocui.Gui, v *gocui.View) error {
	_, index := v.Cursor()
	updated := false
	p := WorldState.Player
	t := &currentTrade
	pricelist := marketPrices[currentCity]
	currentCost := currentTradeCost()
	switch index {
	case 0:
		if t.manaBuy > 0 && int(p.Mana) >= 5*(t.manaSell-t.manaBuy+1) {
			t.manaBuy--
			updated = true
		}
	case 1:
		if t.manaSell > 0 && p.Money > currentCost+pricelist.manaSell && int(p.Attributes.MaxMana) >= 5*(t.manaBuy-t.manaSell+1)+int(p.Mana) {
			t.manaSell--
			updated = true
		}
	case 2:
		if t.foodBuy > 0 && int(p.Food) >= 20*(t.foodSell-t.foodBuy+1) {
			t.foodBuy--
			updated = true
		}
	case 3:
		if t.foodSell > 0 && p.Money > currentCost+pricelist.foodSell {
			t.foodSell--
			updated = true
		}
	}
	if updated {
		updateTradeViews(g)
	}
	return nil
}

func tradeRight(g *gocui.Gui, v *gocui.View) error {
	_, index := v.Cursor()
	updated := false
	p := WorldState.Player
	t := &currentTrade
	pricelist := marketPrices[currentCity]
	currentCost := currentTradeCost()
	switch index {
	case 0:
		if p.Money > currentCost+pricelist.manaBuy && int(p.Attributes.MaxMana) >= int(p.Mana)+5*(t.manaBuy+1-t.manaSell) && t.manaBuy < 99 {
			t.manaBuy++
			updated = true
		}
	case 1:
		if int(p.Mana) > 5*(t.manaBuy-t.manaSell-1) && t.manaSell < 99 {
			t.manaSell++
			updated = true
		}
	case 2:
		if p.Money > currentCost+pricelist.foodBuy && t.foodBuy < 99 {
			t.foodBuy++
			updated = true
		}
	case 3:
		if int(p.Food) > 20*(t.foodSell+1-t.foodBuy) && t.foodSell < 99 {
			t.foodSell++
			updated = true
		}
	}
	if updated {
		updateTradeViews(g)
	}
	return nil
}

func tradeEnter(g *gocui.Gui, v *gocui.View) error {
	cost := currentTradeCost()
	p := WorldState.Player
	t := &currentTrade
	// Check Trade is valid
	if cost > p.Money || int(p.Mana) < 5*(t.manaSell-t.manaBuy) || p.Attributes.MaxMana < uint16(int(p.Mana)+5*(t.manaBuy-t.manaSell)) || int(p.Food) < 20*(t.foodSell-t.foodBuy) {
		return nil
	}
	p.Money -= cost
	if t.manaBuy > t.manaSell {
		p.Mana += 5 * uint16(t.manaBuy-t.manaSell)
	} else {
		p.Mana -= 5 * uint16(t.manaSell-t.manaBuy)
	}
	if t.foodBuy > t.foodSell {
		p.Food += 20 * uint16(t.foodBuy-t.foodSell)
	} else {
		p.Food -= 20 * uint16(t.foodSell-t.foodBuy)
	}
	g.DeleteView("tradeBackground")
	g.DeleteView("tradeInfo")
	g.DeleteView("tradingPrices")
	g.DeleteView("tradingamount")
	g.DeleteView("tradingtotals")
	g.SetLayout(layout)
	return nil
}
