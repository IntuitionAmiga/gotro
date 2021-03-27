package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

const windowWidth, windowHeight = 800, 600

func sdlInitVideo() {
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialise video: %s\n", err)
		os.Exit(1)
	}
}

func playMusic() {
	errSDLAudioInit := sdl.Init(sdl.INIT_AUDIO)
	if errSDLAudioInit != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialise audio: %s\n", errSDLAudioInit)
		os.Exit(1)
	}

	errOpeningAudioDevice := mix.OpenAudio(22050, mix.DEFAULT_FORMAT, 2, 4096)
	if errOpeningAudioDevice != nil {
		fmt.Fprintf(os.Stderr, "Failed to open audio device: %s\n", errOpeningAudioDevice)
		os.Exit(1)
	}

	errSDLMixerInit := mix.Init(mix.INIT_MP3)
	if errSDLMixerInit != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialise mixer: %s\n", errSDLMixerInit)
		os.Exit(1)
	}

	if music, errLoadingMusic := mix.LoadMUS("despair4mat.mp3"); errLoadingMusic != nil {
		fmt.Fprintf(os.Stderr, "Failed to load music from disk: %s\n", errLoadingMusic)
		os.Exit(1)
	} else {
		music.Play(1)
	}
}

func playFloppySounds() {
	errSDLAudioInit := sdl.Init(sdl.INIT_AUDIO)
	if errSDLAudioInit != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialise audio: %s\n", errSDLAudioInit)
		os.Exit(1)
	}

	errOpeningAudioDevice := mix.OpenAudio(22050, mix.DEFAULT_FORMAT, 2, 4096)
	if errOpeningAudioDevice != nil {
		fmt.Fprintf(os.Stderr, "Failed to open audio device: %s\n", errOpeningAudioDevice)
		os.Exit(1)
	}

	errSDLMixerInit := mix.Init(mix.INIT_MP3)
	if errSDLMixerInit != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialise mixer: %s\n", errSDLMixerInit)
		os.Exit(1)
	}

	if music, errLoadingMusic := mix.LoadMUS("./floppy.mp3"); errLoadingMusic != nil {
		fmt.Fprintf(os.Stderr, "Failed to load music from disk: %s\n", errLoadingMusic)
		os.Exit(1)
	} else {
		music.Play(1)
	}
}

func showKickstart(kickRenderer *sdl.Renderer) error {
	kickRenderer.Clear()

	t, err := img.LoadTexture(kickRenderer, "kick13.png")
	if err != nil {
		return fmt.Errorf("Couldn't load image from disk")
	}
	if err := kickRenderer.Copy(t, nil, nil); err != nil {
		return fmt.Errorf("Couldn't copy texture: %v", err)
	}
	kickRenderer.Present()
	return err
}

func main() {
	//Setup video and audio
	sdlInitVideo()

	window, errCreatingSDLWindow := sdl.CreateWindow("Gotro by Intuition",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		windowWidth,
		windowHeight,
		sdl.WINDOW_SHOWN)

	if errCreatingSDLWindow != nil {
		fmt.Fprintf(os.Stderr, "Failed to create SDL window: %s\n", errCreatingSDLWindow)
		os.Exit(1)
	}

	renderer, errCreatingSDLRenderer := sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
	if errCreatingSDLRenderer != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", errCreatingSDLRenderer)
		os.Exit(1)
	}

	surface, errCreatingSDLSurface := window.GetSurface()
	if errCreatingSDLSurface != nil {
		fmt.Fprint(os.Stderr, "Failed to create window surface: \n", errCreatingSDLSurface)
		os.Exit(1)
	}

	//MacOS won't draw the window with this line
	var _ sdl.Event = sdl.PollEvent()

	//Start intro
	showKickstart(renderer)

	playFloppySounds()
	time.Sleep(time.Second * 2)
	surface.FillRect(nil, sdl.MapRGB(surface.Format, 255, 255, 255)) // Fill bg with white
	window.UpdateSurface()
	time.Sleep(time.Second * 9)

	playMusic()

	var i int32

	//Mid to left full length screen wipe
	for i = 1; i <= (windowWidth / 2); i++ {
		surface.FillRect(&sdl.Rect{(windowWidth / 2) - i, 0, i, windowHeight}, sdl.MapRGB(surface.Format, 0, 120, 128))
		window.UpdateSurface()
		time.Sleep(time.Second / 360)
	}

	//Mid to right full length screen wipe
	for i = 1; i <= (windowWidth / 2); i++ {
		surface.FillRect(&sdl.Rect{((windowWidth / 2) - 1) + i, 0, i, windowHeight}, sdl.MapRGB(surface.Format, 0, 120, 128))
		window.UpdateSurface()
		time.Sleep(time.Second / 360)
	}

	// Horizontal bars
	for i = 1; i < windowWidth; i++ {
		//L2R
		surface.FillRect(&sdl.Rect{0 + i, 0, 60, 60}, sdl.MapRGB(surface.Format, 0, 255, 90))
		surface.FillRect(&sdl.Rect{0 + i, 120, 60, 60}, sdl.MapRGB(surface.Format, 0, 255, 90))
		surface.FillRect(&sdl.Rect{0 + i, 240, 60, 60}, sdl.MapRGB(surface.Format, 0, 255, 90))
		surface.FillRect(&sdl.Rect{0 + i, 360, 60, 60}, sdl.MapRGB(surface.Format, 0, 255, 90))
		surface.FillRect(&sdl.Rect{0 + i, 480, 60, 60}, sdl.MapRGB(surface.Format, 0, 255, 90))
		//R2L
		surface.FillRect(&sdl.Rect{windowWidth - i, 60, 60, 60}, sdl.MapRGB(surface.Format, 0, 255, 200))
		surface.FillRect(&sdl.Rect{windowWidth - i, 180, 60, 60}, sdl.MapRGB(surface.Format, 0, 255, 200))
		surface.FillRect(&sdl.Rect{windowWidth - i, 300, 60, 60}, sdl.MapRGB(surface.Format, 0, 255, 200))
		surface.FillRect(&sdl.Rect{windowWidth - i, 420, 60, 60}, sdl.MapRGB(surface.Format, 0, 255, 200))
		surface.FillRect(&sdl.Rect{windowWidth - i, 540, 60, 60}, sdl.MapRGB(surface.Format, 0, 255, 200))
		window.UpdateSurface()
		time.Sleep(time.Second / 360)
	}

	//Horizontal bars 2
	for i = 1; i < windowWidth; i++ {
		//L2R
		surface.FillRect(&sdl.Rect{windowWidth - i, 60, 60, 60}, sdl.MapRGB(surface.Format, 0, 255, 90))
		surface.FillRect(&sdl.Rect{windowWidth - i, 180, 60, 60}, sdl.MapRGB(surface.Format, 0, 255, 90))
		surface.FillRect(&sdl.Rect{windowWidth - i, 300, 60, 60}, sdl.MapRGB(surface.Format, 0, 255, 90))
		surface.FillRect(&sdl.Rect{windowWidth - i, 420, 60, 60}, sdl.MapRGB(surface.Format, 0, 255, 90))
		surface.FillRect(&sdl.Rect{windowWidth - i, 540, 60, 60}, sdl.MapRGB(surface.Format, 0, 255, 90))
		//R2L
		surface.FillRect(&sdl.Rect{0 + i, 0, 60, 60}, sdl.MapRGB(surface.Format, 0, 120, 128))
		surface.FillRect(&sdl.Rect{0 + i, 120, 60, 60}, sdl.MapRGB(surface.Format, 0, 120, 128))
		surface.FillRect(&sdl.Rect{0 + i, 240, 60, 60}, sdl.MapRGB(surface.Format, 0, 120, 128))
		surface.FillRect(&sdl.Rect{0 + i, 360, 60, 60}, sdl.MapRGB(surface.Format, 0, 120, 128))
		surface.FillRect(&sdl.Rect{0 + i, 480, 60, 60}, sdl.MapRGB(surface.Format, 0, 120, 128))
		window.UpdateSurface()
		time.Sleep(time.Second / 360)
	}

	//Clear top to bottom
	for i = 1; i <= windowHeight; i++ {
		surface.FillRect(&sdl.Rect{0, 0, windowWidth, 0 + i}, sdl.MapRGB(surface.Format, 0, 120, 128))
		window.UpdateSurface()
	}

	//drawCircle(window, surface, 100, 100, 80, 255, 255, 255)
	//drawCircle2(window, renderer, 80, 60, 80)
	drawCircleBres(window, renderer, 100, 100, 80)
	time.Sleep(time.Second * 3)
	window.Destroy()
	sdl.Quit()
}

func putPixel(window *sdl.Window, surface *sdl.Surface, xpos, ypos int32, R, G, B uint8) {
	surface.FillRect(&sdl.Rect{xpos, ypos, 1, 1}, sdl.MapRGB(surface.Format, R, G, B))
	window.UpdateSurface()
}

func drawCircle(window *sdl.Window, surface *sdl.Surface, x, y, radius float64, R, G, B uint8) {
	var i, angle, x1, y1 float64

	for i = 0; i < 360; i += 0.1 {
		angle = i
		x1 = radius * math.Cos(angle*math.Pi/180)
		y1 = radius * math.Sin(angle*math.Pi/180)
		putPixel(window, surface, int32(x)+int32(x1), int32(y)+int32(y1), R, G, B)
	}
}

func putPixel2(window *sdl.Window, renderer *sdl.Renderer, xpos, ypos int32, R, G, B uint8) {
	renderer.SetDrawColor(R, G, B, 255)
	renderer.DrawPoint(xpos, ypos)
	renderer.Present()
}

func drawCircle2(window *sdl.Window, renderer *sdl.Renderer, x, y, radius float64) {
	//const pi float64 = 3.1415926535
	var i, angle, x1, y1 float64

	for i = 0; i < 360; i += 0.1 {
		angle = i
		x1 = radius * math.Cos(angle*math.Pi/180)
		y1 = radius * math.Sin(angle*math.Pi/180)
		putPixel2(window, renderer, int32(x)+int32(x1), int32(y)+int32(y1), 255, 255, 255)
	}
}

func drawCircle3(window *sdl.Window, renderer *sdl.Renderer, xc, yc, x, y int32, R, G, B uint8) {
	putPixel2(window, renderer, xc+x, yc+y, R, G, B)
	putPixel2(window, renderer, xc-x, yc+y, R, G, B)
	putPixel2(window, renderer, xc+x, yc-y, R, G, B)
	putPixel2(window, renderer, xc-x, yc-y, R, G, B)
	putPixel2(window, renderer, xc+y, yc+x, R, G, B)
	putPixel2(window, renderer, xc-y, yc+x, R, G, B)
	putPixel2(window, renderer, xc+y, yc-x, R, G, B)
	putPixel2(window, renderer, xc-y, yc-x, R, G, B)
}

func drawCircleBres(window *sdl.Window, renderer *sdl.Renderer, xc, yc, r int32) {
	var x int32 = 0
	var y int32 = r
	var decision int32 = 3 - (2 * r)

	for y >= x {
		x++
		if decision > 0 {
			y--
			decision = decision + 4*(x-y) + 10
		} else {
			decision = decision + 4*x + 6
			drawCircle3(window, renderer, xc, yc, x, y, 255, 255, 255)
			time.Sleep(time.Second / 2880)
		}
	}
}
