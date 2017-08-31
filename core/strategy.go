package core

// log "github.com/sirupsen/logrus"

var strategies StrategyContainer

func InitStrategy() {
	strategies.Init()
}

// Container is the getter of container
func Strategies() *StrategyContainer {
	return &strategies
}

// Strategy is the interface of strategy
type Strategy interface {
	Register(StrategyRegistry)
	Name() string
	GetDataLevels() []HistDataBarSize
	Init()

	// exection
	Condition()

	// data
	OnTick(ask, bid, last float64)
	OnHistBarSize30Min(HistoricalData)
	OnHistBarSize1Day(HistoricalData)

	// order
	Order()
}

/////////////////////////////////////////////////

type StrategyRegistry interface {
	Register(Strategy)
}

// StrategyContainer is the set of strategies
type StrategyContainer struct {
	strategies map[string]Strategy
}

// Init is to initialize the container
func (container *StrategyContainer) Init() {
	container.strategies = make(map[string]Strategy)
}

// Register is used to register strategy to the container
func (container *StrategyContainer) Register(strategy Strategy) {
	if strategy.Name() == "" {
		return
	}

	if _, ok := container.strategies[strategy.Name()]; ok {
		return
	}

	container.strategies[strategy.Name()] = strategy
}

func (container *StrategyContainer) Find(name string) Strategy {
	if name == "" {
		return nil
	}

	return container.strategies[name]
}
