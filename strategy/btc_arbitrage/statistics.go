package btc_arbitrage

import (
	"fmt"
	"os"
	"sync"
)

var bboMgr *BBOMgr

func Init() {
	bboMgr = &BBOMgr{}
	bboMgr.Init()
}

type BBO struct {
	Ask int
	Bid int
}

type BBOMgr struct {
	sync.RWMutex
	HUOBIBBO             BBO
	OKCoinBBO            BBO
	HUOBIAskDifOKCoinBid []int
	OKCoinAskDifHUOBIBid []int

	fileHUOBIToOKCoin *os.File
	fileOKCoinToHUOBI *os.File
}

const (
	reservedCnt = 50
	toInt       = 100000
)

func (mgr *BBOMgr) Init() {
	mgr.HUOBIAskDifOKCoinBid = make([]int, 0, reservedCnt)
	mgr.OKCoinAskDifHUOBIBid = make([]int, 0, reservedCnt)

	mgr.fileHUOBIToOKCoin, _ = os.OpenFile("huobi-ask-okcoin-bid.txt", os.O_APPEND|os.O_RDWR, 0666)
	mgr.fileOKCoinToHUOBI, _ = os.OpenFile("okcoin-ask-huobi-bid.txt", os.O_APPEND|os.O_RDWR, 0666)
}

func (mgr *BBOMgr) UpdateHUOBI(ask, bid int) {
	mgr.update(ask, bid, &mgr.HUOBIBBO)
}

func (mgr *BBOMgr) UpdateOKCoin(ask, bid int) {
	mgr.update(ask, bid, &mgr.OKCoinBBO)
}

func (mgr *BBOMgr) update(ask, bid int, bbo *BBO) {
	mgr.Lock()

	bbo.Ask, bbo.Bid = ask, bid
	mgr.HUOBIAskDifOKCoinBid = append(mgr.HUOBIAskDifOKCoinBid, mgr.HUOBIBBO.Ask-mgr.OKCoinBBO.Bid)
	mgr.OKCoinAskDifHUOBIBid = append(mgr.OKCoinAskDifHUOBIBid, mgr.OKCoinBBO.Ask-mgr.HUOBIBBO.Bid)

	huobiDifOKCoin := []int{}
	if reservedCnt == len(mgr.HUOBIAskDifOKCoinBid) {
		huobiDifOKCoin = make([]int, reservedCnt)
		copy(huobiDifOKCoin, mgr.HUOBIAskDifOKCoinBid)
		mgr.HUOBIAskDifOKCoinBid = []int{}
	}

	okcoinDifHUOBI := []int{}
	if reservedCnt == len(mgr.OKCoinAskDifHUOBIBid) {
		okcoinDifHUOBI = make([]int, reservedCnt)
		copy(okcoinDifHUOBI, mgr.OKCoinAskDifHUOBIBid)
		mgr.OKCoinAskDifHUOBIBid = []int{}
	}

	mgr.Unlock()

	if len(huobiDifOKCoin) > 0 {
		difStr := ""
		for i := 0; i < len(huobiDifOKCoin); i++ {
			difStr += fmt.Sprintf("%d,", huobiDifOKCoin[i])
		}

		mgr.fileHUOBIToOKCoin.WriteString(difStr)
	}

	if len(okcoinDifHUOBI) > 0 {
		difStr := ""
		for i := 0; i < len(okcoinDifHUOBI); i++ {
			difStr += fmt.Sprintf("%d,", okcoinDifHUOBI[i])
		}

		mgr.fileOKCoinToHUOBI.WriteString(difStr)
	}
}
