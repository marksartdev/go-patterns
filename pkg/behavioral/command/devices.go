package command

import "fmt"

const stereoDefaultVolume = 11

const (
	ceilingFanOff = iota
	ceilingFanLow
	ceilingFanMedium
	ceilingFanHigh
)

// Light Светильник.
type Light struct {
	Location string
}

// Включает свет.
func (l *Light) on() string {
	return fmt.Sprintf("%s light is on", l.Location)
}

// Выключает свет.
func (l *Light) off() string {
	return fmt.Sprintf("%s light is off", l.Location)
}

// GarageDoor Дверь гаража.
type GarageDoor struct{}

// Открывает дверь гаража.
func (g *GarageDoor) up() string {
	return "Garage door is open"
}

// Закрывает дверь гаража.
func (g *GarageDoor) down() string {
	return "Garage door is close"
}

// Включает свет в гараже.
func (g *GarageDoor) lightOn() string {
	return "Garage light is on"
}

// Выключает свет в гараже.
func (g *GarageDoor) lightOff() string {
	return "Garage light is off"
}

// Stereo Стереосистема.
type Stereo struct {
	Location string
	volume   int
}

// Включает стереосистему.
func (s *Stereo) on() string {
	return fmt.Sprintf("%s stereo is on", s.Location)
}

// Выключает стереосистему.
func (s *Stereo) off() string {
	return fmt.Sprintf("%s stereo is off", s.Location)
}

// Включает режим чтения с CD.
func (s *Stereo) setCd() string {
	return fmt.Sprintf("%s stereo is set CD", s.Location)
}

// Устанавливает уровень громкости.
func (s *Stereo) setVolume(volume int) string {
	s.volume = volume
	return fmt.Sprintf("%s stereo volume set to %d", s.Location, s.volume)
}

// CeilingFan Вентилятор.
type CeilingFan struct {
	Location string
	speed    int
}

// Включает вентилятор на высокую скорость.
func (c *CeilingFan) high() string {
	c.speed = ceilingFanHigh
	return fmt.Sprintf("%s ceiling fan is on high", c.Location)
}

// Включает вентилятор на среднюю скорость.
func (c *CeilingFan) medium() string {
	c.speed = ceilingFanMedium
	return fmt.Sprintf("%s ceiling fan is on medium", c.Location)
}

// Включает вентилятор на низкую скорость.
func (c *CeilingFan) low() string {
	c.speed = ceilingFanLow
	return fmt.Sprintf("%s ceiling fan is on low", c.Location)
}

// Выключает вентилятор.
func (c *CeilingFan) off() string {
	c.speed = ceilingFanOff
	return fmt.Sprintf("%s ceiling fan is off", c.Location)
}

// Возвращает текущую скорость.
func (c *CeilingFan) getSpeed() int {
	return c.speed
}

// TV Телевизор.
type TV struct {
	Location string
}

// Включает телевизор.
func (t *TV) on() string {
	return fmt.Sprintf("%s TV is on", t.Location)
}

// Выключает телевизор.
func (t *TV) off() string {
	return fmt.Sprintf("%s TV is off", t.Location)
}

// HotTub Джакузи.
type HotTub struct {
	Location string
}

// Включает джакузи.
func (h *HotTub) on() string {
	return fmt.Sprintf("%s hot tub is on", h.Location)
}

// Выключает джакузи.
func (h *HotTub) off() string {
	return fmt.Sprintf("%s hot tub is off", h.Location)
}
