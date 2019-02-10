package main

import (
	"github.com/peterhellberg/gfx"
)

type Creature struct {
	pos, v      gfx.Vec
	path        []gfx.Vec
	currentNode int
}

const maxSpeed = 2
const maxAcceleration = 0.5

func (c *Creature) move() {
	steering := c.followPath()

	// Limit acceleration
	if steering.Len() > maxAcceleration {
		steering = steering.Unit().Scaled(maxAcceleration)
	}

	c.v = c.v.Add(steering)

	// Limit velocity
	if c.v.Len() > maxSpeed {
		c.v = c.v.Unit().Scaled(maxSpeed)
	}

	// Move creature with velocity c.v
	c.pos = c.pos.Add(c.v)
}

// Find the vector to the next target
func (c *Creature) followPath() gfx.Vec {
	target := c.path[c.currentNode]

	// Calculate distance to target node
	distance := target.Sub(c.pos).Len()
	if distance < 20.0 {
		c.currentNode++
	}

	return c.pos.To(target)
}
