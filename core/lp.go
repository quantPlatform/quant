package core

var lps LpContainer

func InitLp() {
	lps.Init()
}

func Lps() *LpContainer {
	return &lps
}

//////////////////////////////////////////////////////
type Lp interface {
	Register(LpRegistry)
	Name() string
	Start(strategy string, contracts []Contract)
}

////////////////////////////////////////////////////
type LpRegistry interface {
	Register(lp Lp)
}

type LpContainer struct {
	lps map[string]Lp
}

func (container *LpContainer) Init() {
	container.lps = make(map[string]Lp)
}

func (container *LpContainer) Register(lp Lp) {
	if lp.Name() == "" {
		return
	}

	if _, ok := container.lps[lp.Name()]; ok {
		return
	}

	container.lps[lp.Name()] = lp
}

func (container *LpContainer) Find(name string) Lp {
	if name == "" {
		return nil
	}

	return container.lps[name]
}

func GetLp(name string) Lp {
	return lps.Find(name)
}
