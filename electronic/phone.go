package electronic

type Phone interface {
	Brand() string
	Model() string
	Type() string
}
type StationPhone interface {
	ButtonsCount() int
}
type Smartphone interface {
	OS() string
}

type applePhone struct {
	model string
}

func (p *applePhone) Brand() string {
	return "Apple"
}
func (p *applePhone) Model() string {
	return p.model
}
func (p *applePhone) Type() string {
	return "smartphone"
}
func (p *applePhone) OS() string {
	return "iOS"
}

type androidPhone struct {
	brand string
	model string
}

func (p *androidPhone) Brand() string {
	return p.brand
}
func (p *androidPhone) Model() string {
	return p.model
}
func (p *androidPhone) Type() string {
	return "smartphone"
}
func (p *androidPhone) OS() string {
	return "Android"
}

type radioPhone struct {
	brand       string
	model       string
	buttonCount int
}

func (r *radioPhone) Brand() string {
	return r.brand
}
func (r *radioPhone) Model() string {
	return r.model
}
func (r *radioPhone) Type() string {
	return "station"
}
func (r *radioPhone) ButtonsCount() int {
	return r.buttonCount
}

func NewApplePhone(model string) *applePhone {
	return &applePhone{model: model}
}

func NewAndroidPhone(brand, model string) *androidPhone {
	return &androidPhone{brand: brand, model: model}
}

func NewRadioPhone(brand, model string, buttonsCount int) *radioPhone {
	return &radioPhone{brand: brand, model: model, buttonCount: buttonsCount}
}
