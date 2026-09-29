package clock

import (
    "fmt"
)

// Define the Clock type here.

type Clock struct{
    min int
}

func New(h, m int) Clock {
    c := Clock{h*60 + m}
	return c.Normalize()
}

func (c Clock) Normalize() Clock {
    for c.min < 0 {
		c.min += 1440   
    }
    c.min %= 1440
    return c
}

func (c Clock) Add(m int) Clock {
	c.min += m
    return c.Normalize() 
}

func (c Clock) Subtract(m int) Clock {
	c.min -= m
    return c.Normalize()
}

func (c Clock) String() string {
	h := c.min / 60
    m := c.min % 60
    return fmt.Sprintf("%02d:%02d", h, m)

}
