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
	"os"
	"time"

	"github.com/fzipp/canvas"
	"github.com/fzipp/pps"
)

var scaling = pps.Vec2{X: 7, Y: 7}

func main() {
	http := flag.String("http", ":8080", "HTTP service address (e.g., '127.0.0.1:8080' or just ':8080')")
	alpha := flag.Float64("alpha", 180.0, "Alpha angle (α) in degrees")
	beta := flag.Float64("beta", 17.0, "Beta angle (β) in degrees")
	velocity := flag.Float64("v", 0.67, "Velocity in space units per time step")
	radius := flag.Float64("r", 5.0, "Radius in space units")
	dpe := flag.Float64("dpe", 0.08, "Density in particles per space unit (p/su)")
	sizeFlag := flag.String("size", "100x100", "Size of universe as `WIDTHxHEIGHT in space units`")
	flag.Parse()

	var size pps.Vec2
	_, err := fmt.Sscanf(*sizeFlag, "%gx%g", &size.X, &size.Y)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Invalid value of -size flag.")
		flag.Usage()
		os.Exit(1)
	}

	params := pps.ParamSet{
		Alpha:    *alpha,
		Beta:     *beta,
		Velocity: *velocity,
		Radius:   *radius,
	}
	DPE := *dpe

	fmt.Println("Visit " + httpLink(*http) + " in a web browser")
	err = canvas.ListenAndServe(*http, func(ctx *canvas.Context) { run(ctx, params, size, DPE) },
		canvas.Title("Primordial Particle System"),
		canvas.Size(int(size.X*scaling.X), int(size.Y*scaling.Y)),
		canvas.ScaleFullPage(false, true),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func run(ctx *canvas.Context, params pps.ParamSet, size pps.Vec2, DPE float64) {
	particles := int(size.X * size.Y * DPE)
	rand.Seed(time.Now().UnixNano())
	u := pps.NewUniverse(size, particles, params)

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
