package pps

type particleGrid struct {
	cellSize int
	cells    map[gridCell][]*Particle
}

type gridCell struct {
	x, y int
}

func (c gridCell) add(x, y int) gridCell {
	return gridCell{x: c.x + x, y: c.y + y}
}

func makeParticleGrid(cellSize int) particleGrid {
	return particleGrid{
		cellSize: cellSize,
		cells:    make(map[gridCell][]*Particle),
	}
}

func (g *particleGrid) addParticle(p *Particle) {
	g.add(g.cell(p.Pos), p)
}

func (g *particleGrid) add(c gridCell, p *Particle) {
	g.cells[c] = append(g.cells[c], p)
}

func (g *particleGrid) remove(c gridCell, p *Particle) {
	for i, cp := range g.cells[c] {
		if cp != p {
			continue
		}
		g.cells[c] = append(g.cells[c][:i], g.cells[c][i+1:]...)
		return
	}
}

func (g *particleGrid) updatePos(p *Particle, pos Vec2) {
	oldCell := g.cell(p.Pos)
	newCell := g.cell(pos)
	if newCell != oldCell {
		g.remove(oldCell, p)
		g.add(newCell, p)
	}
	p.Pos = pos
}

func (g *particleGrid) cell(pos Vec2) gridCell {
	return gridCell{x: int(pos.X) / g.cellSize, y: int(pos.Y) / g.cellSize}
}

// neighbours returns the number of neighbours within radius r on the left
// side (L) and on the right side (R) of a particle p, as well as the number
// of close neighbours NClose, which are within radius rClose.
// The left side and the right side are semicircles with the radius r.
// The sum N=L+R is the total number of neighbours within this radius.
func (g *particleGrid) neighbours(p *Particle, r, rClose float64) (L, R, NClose int) {
	center := g.cell(p.Pos)
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for _, np := range g.cells[center.add(dx, dy)] {
				if p == np {
					continue
				}
				pos := p.Pos
				nPos := np.Pos
				d := pos.dist(nPos)
				if d <= r {
					if isLeft(nPos, pos, pos.add(p.Dir())) {
						L++
					} else {
						R++
					}
					if d <= rClose {
						NClose++
					}
				}
			}
		}
	}
	return
}

// isLeft reports whether point p is to the left side of the line
// through a and b.
func isLeft(p, a, b Vec2) bool {
	return (p.X-a.X)*(b.Y-a.Y)-(p.Y-a.Y)*(b.X-a.X) > 0
}
