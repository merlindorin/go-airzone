package v1

type Option func(request *ZoneSetRequest)

func (receiver Option) Apply(request *ZoneSetRequest) {
	receiver(request)
}

func WithName(name string) Option {
	return func(request *ZoneSetRequest) {
		request.Name = &name
	}
}

type Sleep int

const (
	Sleep0  Sleep = 0
	Sleep30 Sleep = 30
	Sleep60 Sleep = 60
	Sleep90 Sleep = 90
)

func WithSleep(sleep Sleep) Option {
	return func(request *ZoneSetRequest) {
		request.Sleep = (*int)(&sleep)
	}
}

func WithSetpoint(temp float64) Option {
	return func(request *ZoneSetRequest) {
		request.SetPoint = &temp
	}
}

func WithSpeed(speed int) Option {
	return func(request *ZoneSetRequest) {
		request.Mode = &speed
	}
}

type Power int

const (
	PowerOFF Power = iota
	PowerON
)

func WithPower(power Power) Option {
	return func(request *ZoneSetRequest) {
		request.Power = (*int)(&power)
	}
}

func WithPowerON() Option {
	return WithPower(PowerON)
}

func WithPowerOFF() Option {
	return WithPower(PowerOFF)
}
