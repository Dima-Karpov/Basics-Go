package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0x00, 0xff, 0x00, 0xff}, // Зеленый
	color.RGBA{0xff, 0x00, 0x00, 0xff}, // Красный
	color.RGBA{0x00, 0x00, 0xff, 0xff}, // Синий
	color.RGBA{0xff, 0xff, 0x00, 0xff}, // Желтый
	color.RGBA{0x00, 0xff, 0xff, 0xff}, // Голубой
	color.RGBA{0xff, 0x00, 0xff, 0xff}, // Пурпурный
	color.RGBA{0xff, 0xa5, 0x00, 0xff}, // Оранжевый
}

const (
	backgroundIndex = 0
)

// go run main.go > out.gif
func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Количество полных колебаний x
		res     = 0.001 // Угловое разрешение
		size    = 100   // Канва изображения охватывает [size..+size]
		nframes = 64    // Количество кадров анимации
		delay   = 8     // Задержка между кадрами (единица - 10мс)
	)

	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64()*2.0 + 1.0 // Избегаем слишком медленных
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			//Нетривиальный выбор цветаЖ зависит от t и i
			colorIndex := uint8(1 + (i+int(t*7))%(len(palette)-1)) // избегаем 0 (фон)

			img.SetColorIndex(
				size+int(x*size+0.5),
				size+int(y*size+0.5),
				colorIndex,
			)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
