//go:generate statik -src=./img

package main

import (
	"flag"
	_ "github.com/CovenantSQL/notpro/statik"
	"github.com/disintegration/imaging"
	"github.com/rakyll/statik/fs"
	"image"
	"log"
)

var (
	opacity float64
	in      string
	out     string
)

func main() {
	flag.StringVar(&in, "input", "head.png", "Your input image")
	flag.StringVar(&out, "output", "newhead.png", "Output filename")
	flag.Float64Var(&opacity, "opacity", 0.95, "Opacity level")
	flag.Parse()

	// Open notpro.png
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	npf, err := statikFS.Open("/notpro.png")
	if err != nil {
		log.Fatalf("read from statikFS: %v", err)
	}
	notpro, err := imaging.Decode(npf)
	if err != nil {
		log.Fatalf("failed to open notpro.png: %v", err)
	}

	// Open input image.
	src, err := imaging.Open(in)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	width := src.Bounds().Max.X
	height := src.Bounds().Max.Y
	log.Printf("Width: %d, Height: %d", width, height)

	notpro = imaging.Resize(notpro, width, 0, imaging.Lanczos)

	output := imaging.Overlay(src, notpro,
		image.Pt(0, height-notpro.Bounds().Max.Y), opacity)
	err = imaging.Save(output, out)
	if err != nil {
		log.Fatalf("save new head failed: %v", err)
	}
	return
}
