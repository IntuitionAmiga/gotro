package main

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

const windowWidth, windowHeight = 800, 600

func main() {
	//Setup video and audio
	sdlInitVideo()
	var window = createWindow()
	var renderer = createRenderer(window)
	var surface = createSurface(window)

	var _ = sdl.PollEvent() //MacOS won't draw the window without this line

	//Start intro
	_ = showKickstart(renderer)
	playFloppySounds()
	time.Sleep(time.Second * 2)

	backgroundFill(window, surface, 255, 255, 255) //Fill bg with white
	time.Sleep(time.Second * 9)

	playMusic()

	wipeToLeft(window, surface)
	wipeToRight(window, surface)
	horizontalBars(window, surface)
	horizontalBars2(window, surface)

	wipeTopDown(window, surface)
	drawCircle(renderer, 200, 200, 80, 255, 255, 255)
	time.Sleep(time.Second * 3)

	_ = window.Destroy()
	sdl.Quit()
}

func createWindow() *sdl.Window {
	window, errCreatingSDLWindow := sdl.CreateWindow("Gotro by Intuition",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		windowWidth,
		windowHeight,
		sdl.WINDOW_SHOWN)

	if errCreatingSDLWindow != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to create SDL window: %s\n", errCreatingSDLWindow)
		os.Exit(1)
	}
	return window
}
func createRenderer(window *sdl.Window) *sdl.Renderer {
	renderer, errCreatingSDLRenderer := sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
	if errCreatingSDLRenderer != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", errCreatingSDLRenderer)
		os.Exit(1)
	}
	return renderer
}
func createSurface(window *sdl.Window) *sdl.Surface {
	surface, errCreatingSDLSurface := window.GetSurface()
	if errCreatingSDLSurface != nil {
		_, _ = fmt.Fprint(os.Stderr, "Failed to create window surface: \n", errCreatingSDLSurface)
		os.Exit(1)
	}
	return surface
}
func sdlInitVideo() {
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to initialise video: %s\n", err)
		os.Exit(1)
	}
}
func showKickstart(kickRenderer *sdl.Renderer) error {
	_ = kickRenderer.Clear()

	t, err := img.LoadTexture(kickRenderer, "kick13.png")
	if err != nil {
		return fmt.Errorf("couldn't load image from disk")
	}
	if err := kickRenderer.Copy(t, nil, nil); err != nil {
		return fmt.Errorf("couldn't copy texture: %v", err)
	}
	kickRenderer.Present()
	return err
}
func playFloppySounds() {
	errSDLAudioInit := sdl.Init(sdl.INIT_AUDIO)
	if errSDLAudioInit != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to initialise audio: %s\n", errSDLAudioInit)
		os.Exit(1)
	}

	errOpeningAudioDevice := mix.OpenAudio(22050, mix.DEFAULT_FORMAT, 2, 4096)
	if errOpeningAudioDevice != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to open audio device: %s\n", errOpeningAudioDevice)
		os.Exit(1)
	}

	errSDLMixerInit := mix.Init(mix.INIT_MP3)
	if errSDLMixerInit != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to initialise mixer: %s\n", errSDLMixerInit)
		os.Exit(1)
	}

	if music, errLoadingMusic := mix.LoadMUS("./floppy.mp3"); errLoadingMusic != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to load music from disk: %s\n", errLoadingMusic)
		os.Exit(1)
	} else {
		_ = music.Play(1)
	}
}
func backgroundFill(window *sdl.Window, surface *sdl.Surface, R, G, B uint8) {
	_ = surface.FillRect(nil, sdl.MapRGB(surface.Format, R, G, B))
	_ = window.UpdateSurface()
}
func playMusic() {
	errSDLAudioInit := sdl.Init(sdl.INIT_AUDIO)
	if errSDLAudioInit != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to initialise audio: %s\n", errSDLAudioInit)
		os.Exit(1)
	}

	errOpeningAudioDevice := mix.OpenAudio(22050, mix.DEFAULT_FORMAT, 2, 4096)
	if errOpeningAudioDevice != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to open audio device: %s\n", errOpeningAudioDevice)
		os.Exit(1)
	}

	errSDLMixerInit := mix.Init(mix.INIT_MP3)
	if errSDLMixerInit != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to initialise mixer: %s\n", errSDLMixerInit)
		os.Exit(1)
	}

	if music, errLoadingMusic := mix.LoadMUS("despair4mat.mp3"); errLoadingMusic != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to load music from disk: %s\n", errLoadingMusic)
		os.Exit(1)
	} else {
		_ = music.Play(1)
	}
}
func drawPixel(renderer *sdl.Renderer, xpos, ypos int32, R, G, B uint8) {
	_ = renderer.SetDrawColor(R, G, B, 255)
	_ = renderer.DrawPoint(xpos, ypos)
	//renderer.Present()
}
func wipeToLeft(window *sdl.Window, surface *sdl.Surface) {
	var i int32

	//Mid to left full length screen wipe
	for i = 1; i <= (windowWidth / 2); i++ {
		_ = surface.FillRect(&sdl.Rect{X: (windowWidth / 2) - i, W: i, H: windowHeight}, sdl.MapRGB(surface.Format, 0, 120, 128))
		_ = window.UpdateSurface()
		time.Sleep(time.Second / 360)
	}
}
func wipeToRight(window *sdl.Window, surface *sdl.Surface) {
	var i int32
	//Mid to right full length screen wipe
	for i = 1; i <= (windowWidth / 2); i++ {
		_ = surface.FillRect(&sdl.Rect{X: ((windowWidth / 2) - 1) + i, W: i, H: windowHeight}, sdl.MapRGB(surface.Format, 0, 120, 128))
		_ = window.UpdateSurface()
		time.Sleep(time.Second / 360)
	}
}
func horizontalBars(window *sdl.Window, surface *sdl.Surface) {
	var i int32
	// Horizontal bars
	for i = 1; i < windowWidth; i++ {
		//L2R
		_ = surface.FillRect(&sdl.Rect{X: 0 + i, W: 60, H: 60}, sdl.MapRGB(surface.Format, 0, 255, 90))
		_ = surface.FillRect(&sdl.Rect{X: 0 + i, Y: 120, W: 60, H: 60}, sdl.MapRGB(surface.Format, 0, 255, 90))
		_ = surface.FillRect(&sdl.Rect{X: 0 + i, Y: 240, W: 60, H: 60}, sdl.MapRGB(surface.Format, 0, 255, 90))
		_ = surface.FillRect(&sdl.Rect{X: 0 + i, Y: 360, W: 60, H: 60}, sdl.MapRGB(surface.Format, 0, 255, 90))
		_ = surface.FillRect(&sdl.Rect{X: 0 + i, Y: 480, W: 60, H: 60}, sdl.MapRGB(surface.Format, 0, 255, 90))
		//R2L
		_ = surface.FillRect(&sdl.Rect{X: windowWidth - i, Y: 60, W: 60, H: 60}, sdl.MapRGB(surface.Format, 0, 255, 200))
		_ = surface.FillRect(&sdl.Rect{X: windowWidth - i, Y: 180, W: 60, H: 60}, sdl.MapRGB(surface.Format, 0, 255, 200))
		_ = surface.FillRect(&sdl.Rect{X: windowWidth - i, Y: 300, W: 60, H: 60}, sdl.MapRGB(surface.Format, 0, 255, 200))
		_ = surface.FillRect(&sdl.Rect{X: windowWidth - i, Y: 420, W: 60, H: 60}, sdl.MapRGB(surface.Format, 0, 255, 200))
		_ = surface.FillRect(&sdl.Rect{X: windowWidth - i, Y: 540, W: 60, H: 60}, sdl.MapRGB(surface.Format, 0, 255, 200))
		_ = window.UpdateSurface()
		time.Sleep(time.Second / 360)
	}

}
func horizontalBars2(window *sdl.Window, surface *sdl.Surface) {
	var i int32
	//Horizontal bars 2
	for i = 1; i < windowWidth; i++ {
		//L2R
		_ = surface.FillRect(&sdl.Rect{X: windowWidth - i, Y: 60, W: 60, H: 60}, sdl.MapRGB(surface.Format, 0, 255, 90))
		_ = surface.FillRect(&sdl.Rect{X: windowWidth - i, Y: 180, W: 60, H: 60}, sdl.MapRGB(surface.Format, 0, 255, 90))
		_ = surface.FillRect(&sdl.Rect{X: windowWidth - i, Y: 300, W: 60, H: 60}, sdl.MapRGB(surface.Format, 0, 255, 90))
		_ = surface.FillRect(&sdl.Rect{X: windowWidth - i, Y: 420, W: 60, H: 60}, sdl.MapRGB(surface.Format, 0, 255, 90))
		_ = surface.FillRect(&sdl.Rect{X: windowWidth - i, Y: 540, W: 60, H: 60}, sdl.MapRGB(surface.Format, 0, 255, 90))
		//R2L
		_ = surface.FillRect(&sdl.Rect{X: 0 + i, W: 60, H: 60}, sdl.MapRGB(surface.Format, 0, 120, 128))
		_ = surface.FillRect(&sdl.Rect{X: 0 + i, Y: 120, W: 60, H: 60}, sdl.MapRGB(surface.Format, 0, 120, 128))
		_ = surface.FillRect(&sdl.Rect{X: 0 + i, Y: 240, W: 60, H: 60}, sdl.MapRGB(surface.Format, 0, 120, 128))
		_ = surface.FillRect(&sdl.Rect{X: 0 + i, Y: 360, W: 60, H: 60}, sdl.MapRGB(surface.Format, 0, 120, 128))
		_ = surface.FillRect(&sdl.Rect{X: 0 + i, Y: 480, W: 60, H: 60}, sdl.MapRGB(surface.Format, 0, 120, 128))
		_ = window.UpdateSurface()
		time.Sleep(time.Second / 360)
	}
}
func wipeTopDown(window *sdl.Window, surface *sdl.Surface) {
	var i int32
	//Clear top to bottom
	for i = 1; i <= windowHeight; i++ {
		_ = surface.FillRect(&sdl.Rect{W: windowWidth, H: 0 + i}, sdl.MapRGB(surface.Format, 0, 120, 128))
		_ = window.UpdateSurface()
	}
}
func drawCircle(renderer *sdl.Renderer, x0, y0, r int32, R, G, B uint8) {
	var x, y, dx, dy int32 = r - 1, 0, 1, 1
	var err = dx - (r * 2)

	for x > y {
		drawPixel(renderer, x0+x, y0+y, R, G, B)
		drawPixel(renderer, x0+y, y0+x, R, G, B)
		drawPixel(renderer, x0-y, y0+x, R, G, B)
		drawPixel(renderer, x0-x, y0+y, R, G, B)
		drawPixel(renderer, x0-x, y0-y, R, G, B)
		drawPixel(renderer, x0-y, y0-x, R, G, B)
		drawPixel(renderer, x0+y, y0-x, R, G, B)
		drawPixel(renderer, x0+x, y0-y, R, G, B)

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}

		if err > 0 {
			x--
			dx += 2
			err += dx - (r * 2)
		}
	}
	renderer.Present()
}
