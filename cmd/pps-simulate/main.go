// Copyright 2021 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/fzipp/canvas"
	"github.com/fzipp/pps"
)

var (
	size    = pps.Vec2{X: 100, Y: 100}
	scaling = pps.Vec2{X: 7, Y: 7}
)

func main() {
	http := flag.String("http", ":8080", "HTTP service address (e.g., '127.0.0.1:8080' or just ':8080')")
	flag.Parse()

	fmt.Println("Listening on " + httpLink(*http))
	err := canvas.ListenAndServe(*http, run,
		canvas.Title("Primordial Particle System"),
		canvas.Size(int(size.X*scaling.X), int(size.Y*scaling.Y)),
		canvas.ScaleFullPage(false, true),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func run(ctx *canvas.Context) {
	DPE := 0.08
	particles := int(size.X * size.Y * DPE)
	rand.Seed(time.Now().UnixNano())
	u := pps.NewUniverse(size, particles, pps.ParamSet{
		Alpha:    180,
		Beta:     17,
		Velocity: 0.67,
		Radius:   5.0,
	})

	ctx.Scale(scaling.X, scaling.Y)
	for {
		select {
		case ev := <-ctx.Events():
			if _, ok := ev.(canvas.CloseEvent); ok {
				return
			}
		default:
			u.Step()
			draw(ctx, u)
			ctx.Flush()
		}
	}
}

func draw(ctx *canvas.Context, u *pps.Universe) {
	ctx.SetFillStyle(color.Black)
	ctx.FillRect(0, 0, u.Size.X, u.Size.Y)
	for _, p := range u.Particles {
		drawParticle(ctx, p)
	}
}

func drawParticle(ctx *canvas.Context, p pps.Particle) {
	ctx.SetFillStyle(p.Color)
	const particleSize = 0.5
	ctx.BeginPath()
	ctx.Arc(p.Pos.X, p.Pos.Y, particleSize, 0, 2*math.Pi, false)
	ctx.Fill()
}

func httpLink(addr string) string {
	if addr[0] == ':' {
		addr = "localhost" + addr
	}
	return "http://" + addr
}
