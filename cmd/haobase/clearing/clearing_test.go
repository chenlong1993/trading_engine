package clearing

import (
	"testing"
	"time"

	"github.com/yzimhao/trading_engine/cmd/haobase/assets"
	"github.com/yzimhao/trading_engine/cmd/haobase/base"
	"github.com/yzimhao/trading_engine/cmd/haobase/base/symbols"
	"github.com/yzimhao/trading_engine/cmd/haobase/orders"
	"github.com/yzimhao/trading_engine/trading_core"
	"github.com/yzimhao/trading_engine/utils"
	"github.com/yzimhao/trading_engine/utils/app"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	sellUser           = "user1"
	buyUser            = "user2"
	testSymbol         = "usdjpy"
	testTargetSymbol   = "usd"
	testStandardSymbol = "jpy"
)

func initdb(t *testing.T) {
	app.DatabaseInit("mysql", "root:root@tcp(localhost:3306)/test?charset=utf8&loc=Local", true)
	app.RedisInit("127.0.0.1:6379", "", 0)

	base.Init()

	cleanAssets(t)
	cleanOrders(t)

}

func initAssets(t *testing.T) {
	assets.Init()
	symbols.DemoData()

	assets.SysRecharge("user1", "usd", "10000.00", "C001")
	assets.SysRecharge("user1", "jpy", "10000.00", "C001")
	assets.SysRecharge("user2", "usd", "10000.00", "C001")
	assets.SysRecharge("user2", "jpy", "10000.00", "C001")
}

func cleanAssets(t *testing.T) {
	db := app.Database()
	db.DropIndexes(new(assets.Assets))
	db.DropIndexes("assets_freeze")
	db.DropIndexes("assets_log")
	err := db.DropTables(new(assets.Assets), "assets_freeze", "assets_log")
	if err != nil {
		t.Logf("mysql droptables: %s", err)
	}

}

func cleanOrders(t *testing.T) {
	db := app.Database()
	db.DropIndexes(orders.GetOrderTableName(testSymbol))
	db.DropIndexes(new(orders.UnfinishedOrder))
	db.DropIndexes(orders.GetTradelogTableName(testSymbol))

	db.DropTables(orders.GetOrderTableName(testSymbol))
	db.DropTables(new(orders.UnfinishedOrder))
	db.DropTables(orders.GetTradelogTableName(testSymbol))

}

func TestLimitOrder(t *testing.T) {
	initdb(t)
	Convey("限价单完全成交结算测试", t, func() {
		initAssets(t)
		defer cleanOrders(t)
		defer cleanAssets(t)

		sell, err := orders.NewLimitOrder(sellUser, testSymbol, trading_core.OrderSideSell, "1.00", "1")
		So(err, ShouldBeNil)

		buy, err := orders.NewLimitOrder(buyUser, testSymbol, trading_core.OrderSideBuy, "1.00", "1")
		So(err, ShouldBeNil)

		result := trading_core.TradeResult{
			Symbol:        testSymbol,
			AskOrderId:    sell.OrderId,
			BidOrderId:    buy.OrderId,
			TradePrice:    utils.D("1.00"),
			TradeQuantity: utils.D("1"),
			TradeTime:     time.Now().UnixNano(),
		}
		clearing_trade_order(testSymbol, result.Json())

		//检查资产
		sell_assets_target := assets.FindSymbol(sellUser, testTargetSymbol)
		sell_assets_standard := assets.FindSymbol(sellUser, testStandardSymbol)

		buy_assets_target := assets.FindSymbol(buyUser, testTargetSymbol)
		buy_assets_standard := assets.FindSymbol(buyUser, testStandardSymbol)
		So(utils.D(sell_assets_target.Total), ShouldEqual, utils.D("9999"))
		So(utils.D(sell_assets_standard.Total), ShouldEqual, utils.D("10000.995"))

		So(utils.D(buy_assets_target.Total), ShouldEqual, utils.D("10001"))
		So(utils.D(buy_assets_standard.Total), ShouldEqual, utils.D("9998.995"))

		//检查订单状态
		sell_order := orders.Find(testSymbol, sell.OrderId)
		So(sell_order.Status, ShouldEqual, orders.OrderStatusDone)
		buy_order := orders.Find(testSymbol, buy.OrderId)
		So(buy_order.Status, ShouldEqual, orders.OrderStatusDone)

		sell_unfinished := orders.FindUnfinished(testSymbol, sell.OrderId)
		So(sell_unfinished, ShouldBeNil)
		buy_unfinished := orders.FindUnfinished(testSymbol, buy.OrderId)
		So(buy_unfinished, ShouldBeNil)

	})
}

func TestMarket(t *testing.T) {
	initdb(t)
	Convey("市价买指定的数量", t, func() {
		initAssets(t)

		s1, err := orders.NewLimitOrder(sellUser, testSymbol, trading_core.OrderSideSell, "1.00", "1")
		So(err, ShouldBeNil)

		s2, _ := orders.NewLimitOrder(sellUser, testSymbol, trading_core.OrderSideSell, "2.00", "1")

		buy, err := orders.NewMarketOrderByQty(buyUser, testSymbol, trading_core.OrderSideBuy, "3")
		So(err, ShouldBeNil)

		result1 := trading_core.TradeResult{
			Symbol:        testSymbol,
			AskOrderId:    s1.OrderId,
			BidOrderId:    buy.OrderId,
			TradePrice:    utils.D("1.00"),
			TradeQuantity: utils.D("1"),
			TradeTime:     time.Now().UnixNano(),
		}
		result2 := trading_core.TradeResult{
			Symbol:        testSymbol,
			AskOrderId:    s2.OrderId,
			BidOrderId:    buy.OrderId,
			TradePrice:    utils.D("2.00"),
			TradeQuantity: utils.D("1"),
			TradeTime:     time.Now().UnixNano(),
			Last:          buy.OrderId,
		}

		clearing_trade_order(testSymbol, result1.Json())
		clearing_trade_order(testSymbol, result2.Json())
		time.Sleep(5 * time.Second)
	})
}
