package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	FPS                       = 1
	windowWidth, windowHeight = 800, 600
)

var window *sdl.Window
var renderer *sdl.Renderer
var surface *sdl.Surface

func main() {
	sdlInitVideo()
	window = createWindow()
	renderer = createRenderer()
	surface = createSurface()
	var _ = sdl.PollEvent() //MacOS won't draw the window without this line

	//Start intro
	_ = showKickstart()
	playFloppySounds()
	time.Sleep(time.Second * 2)

	backgroundFill(255, 255, 255) //Fill bg with white
	time.Sleep(time.Second * 9)

	playMusic()

	wipeToLeft(255, 0, 90)
	wipeToRight(0, 120, 128)
	wipeToLeft(0, 120, 128)
	wipeToRight(255, 0, 90)
	wipeToLeft(255, 0, 90)

	boingBall(255, 0, 90)
	copperBars()

	wipeTopDown(0, 0, 0)
	drawBubbles()

	wipeToLeft(0, 120, 128)
	wipeToRight(0, 120, 128)

	horizontalBars(255, 0, 90, 0, 255, 200)
	horizontalBars2(0, 120, 128, 255, 0, 128)

	wipeToLeft(0, 0, 0)
	wipeToRight(0, 0, 0)

	_ = renderer.Destroy()
	_ = window.Destroy()

}
func sdlInitVideo() {
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to initialise video: %s\n", err)
		os.Exit(1)
	}

	defer sdl.Quit()
}
func createWindow() *sdl.Window {
	window, errCreatingSDLWindow := sdl.CreateWindow("Gotro by Intuition",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		windowWidth,
		windowHeight,
		sdl.WINDOW_SHOWN)

	if errCreatingSDLWindow != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to create SDL window: %s\n", errCreatingSDLWindow)
		os.Exit(1)
	}
	return window
}
func createRenderer() *sdl.Renderer {
	//Disabled because MacOS SDL can't do hardware rendering
	var errCreatingSDLRenderer error
	if runtime.GOOS == "darwin" {
		renderer, errCreatingSDLRenderer = sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE|sdl.RENDERER_TARGETTEXTURE)
	} else {
		sdl.SetHint(sdl.HINT_RENDER_VSYNC, "1")
		renderer, errCreatingSDLRenderer = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC|sdl.RENDERER_TARGETTEXTURE)
	}

	if errCreatingSDLRenderer != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", errCreatingSDLRenderer)
		os.Exit(1)
	}
	return renderer
}
func createSurface() *sdl.Surface {
	surface, errCreatingSDLSurface := window.GetSurface()
	if errCreatingSDLSurface != nil {
		_, _ = fmt.Fprint(os.Stderr, "Failed to create window surface: \n", errCreatingSDLSurface)
		os.Exit(1)
	}
	return surface
}
func showKickstart() error {
	t, err := img.LoadTexture(renderer, "kick13.png")
	if err != nil {
		return fmt.Errorf("couldn't load image from disk")
	}
	if err := renderer.Copy(t, nil, nil); err != nil {
		return fmt.Errorf("couldn't copy texture: %v", err)
	}
	_ = renderer.SetDrawColor(255, 255, 255, 0)
	updateScreen("r")
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
func backgroundFill(R, G, B uint8) {
	_ = surface.FillRect(nil, sdl.MapRGB(surface.Format, R, G, B))
	updateScreen("r")
	_ = renderer.Clear()
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

	errSDLMixerInit := mix.Init(mix.INIT_MOD)
	if errSDLMixerInit != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to initialise mixer: %s\n", errSDLMixerInit)
		os.Exit(1)
	}

	if music, errLoadingMusic := mix.LoadMUS("echoing.mod"); errLoadingMusic != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to load music from disk: %s\n", errLoadingMusic)
		os.Exit(1)
	} else {
		_ = music.Play(1)
	}
}
func drawPixel(xpos, ypos int32, R, G, B uint8) {
	_ = renderer.SetDrawColor(R, G, B, 255)
	_ = renderer.DrawPoint(xpos, ypos)
}
func wipeToLeft(R, G, B uint8) {
	var i int32

	//Mid to left full length screen wipe
	for i = 1; i <= (windowWidth / 2); i++ {
		//time.Sleep(time.Second / 270)
		_ = surface.FillRect(&sdl.Rect{X: (windowWidth / 2) - i, W: i, H: windowHeight}, sdl.MapRGB(surface.Format, R, G, B))
		updateScreen("s")
	}
}
func wipeToRight(R, G, B uint8) {
	var i int32
	//Mid to right full length screen wipe
	for i = 1; i <= (windowWidth / 2); i++ {
		//time.Sleep(time.Second / 270)
		_ = surface.FillRect(&sdl.Rect{X: ((windowWidth / 2) - 1) + i, W: i, H: windowHeight}, sdl.MapRGB(surface.Format, R, G, B))
		updateScreen("s")
	}
}
func drawSprite(x, y int32, R, G, B uint8) {
	src := sdl.Rect{W: 455, H: 456}
	dst := sdl.Rect{X: x, Y: y, W: 128, H: 128}
	sprite, _ := img.Load("boingball.png")
	texture, _ := renderer.CreateTextureFromSurface(sprite)
	_ = renderer.SetDrawColor(R, G, B, 255)
	_ = renderer.Clear()
	_ = renderer.Copy(texture, &src, &dst)
	//renderer.Present()
	updateScreen("r")

}
func boingBall(R, G, B uint8) {
	_ = renderer.Clear()
	var xPos, yPos int
	for i := 0; i <= (windowHeight - 128); i++ {
		drawSprite(int32(i), int32(i), R, G, B)
		xPos = i
		yPos = i
		//fmt.Println("X:", i, "Y:", yPos)
	}
	for i := xPos; i <= (windowWidth - 128); i++ {
		drawSprite(int32(i), int32(yPos+10), R, G, B)
		yPos -= 1
		xPos = i
		//fmt.Println("X:", i, "Y:", yPos)
	}
	for i := yPos; i >= 0; i-- {
		drawSprite(int32(xPos+i), int32(i), R, G, B)
		xPos = i
		yPos -= 1
		//fmt.Println("X:", i, "Y:", yPos)
	}
}
func copperBars() {
	var startX, startY int32 = windowWidth, windowHeight/2 + 16
	var redY int32 = 0
	var greenY = int32(windowHeight / 3)
	var blueY = int32((windowHeight / 3) * 2)

	//_ = renderer.SetDrawColor(0, 0, 0, 0)

	redBar := func() {
		for i := int32(0); i <= 35; i++ {
			//Red
			redY += i
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: redY + 24, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 8, 0, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: redY + 20, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 16, 0, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: redY + 16, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 32, 0, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: redY + 12, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 63, 0, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: redY + 8, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 127, 0, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: redY + 4, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 191, 0, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: redY, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 255, 0, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: redY - 4, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 191, 0, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: redY - 8, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 127, 0, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: redY - 12, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 63, 0, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: redY - 16, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 32, 0, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: redY - 20, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 16, 0, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: redY - 24, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 8, 0, 0))
			//_ = window.UpdateSurface()
			updateScreen("s")
			time.Sleep(time.Second / 8)
		}
	}
	greenBar := func() {
		for i := int32(0); i <= 35; i++ {
			//Green
			greenY += i
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: greenY + 24, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 8, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: greenY + 20, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 16, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: greenY + 16, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 32, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: greenY + 12, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 63, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: greenY + 8, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 127, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: greenY + 4, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 191, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: greenY, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 255, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: greenY - 4, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 191, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: greenY - 8, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 127, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: greenY - 12, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 63, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: greenY - 16, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 32, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: greenY - 20, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 16, 0))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: greenY - 24, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 8, 0))
			//_ = window.UpdateSurface()
			updateScreen("s")
			time.Sleep(time.Second / 8)
		}
	}
	blueBar := func() {
		for i := int32(0); i <= 35; i++ {
			//Blue
			blueY += i
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: blueY + 24, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 0, 8))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: blueY + 20, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 0, 16))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: blueY + 16, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 0, 32))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: blueY + 12, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 0, 63))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: blueY + 8, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 0, 127))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: blueY + 4, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 0, 191))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: blueY, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 0, 255))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: blueY - 4, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 0, 191))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: blueY - 8, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 0, 127))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: blueY - 12, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 0, 63))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: blueY - 16, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 0, 32))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: blueY - 20, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 0, 16))
			_ = surface.FillRect(&sdl.Rect{X: startX - windowWidth, Y: blueY - 24, W: windowWidth, H: 4}, sdl.MapRGB(surface.Format, 0, 0, 8))
			//_ = window.UpdateSurface()
			updateScreen("s")
			time.Sleep(time.Second / 8)
		}
	}

	redBar()
	greenBar()
	blueBar()

	for i := 0; i < windowWidth+windowWidth/2; i++ {
		_ = surface.FillRect(&sdl.Rect{X: startX - int32(i), Y: startY - 48, W: 1, H: 16}, sdl.MapRGB(surface.Format, 255, 0, 0))
		_ = surface.FillRect(&sdl.Rect{X: startX - int32(i), Y: startY - 32, W: 1, H: 16}, sdl.MapRGB(surface.Format, 255, 127, 0))
		_ = surface.FillRect(&sdl.Rect{X: startX - int32(i), Y: startY - 16, W: 1, H: 16}, sdl.MapRGB(surface.Format, 255, 255, 0))
		_ = surface.FillRect(&sdl.Rect{X: startX - int32(i), Y: startY, W: 1, H: 16}, sdl.MapRGB(surface.Format, 0, 255, 0))
		_ = surface.FillRect(&sdl.Rect{X: startX - int32(i), Y: startY + 16, W: 1, H: 16}, sdl.MapRGB(surface.Format, 0, 0, 255))
		_ = surface.FillRect(&sdl.Rect{X: startX - int32(i), Y: startY + 32, W: 1, H: 16}, sdl.MapRGB(surface.Format, 46, 43, 95))
		_ = surface.FillRect(&sdl.Rect{X: startX - int32(i), Y: startY + 48, W: 1, H: 16}, sdl.MapRGB(surface.Format, 139, 0, 255))
		updateScreen("s")
	}
}
func wipeTopDown(R, G, B uint8) {
	var i int32
	//Clear top to bottom
	for i = 1; i <= windowHeight; i++ {
		_ = surface.FillRect(&sdl.Rect{W: windowWidth, H: 0 + i}, sdl.MapRGB(surface.Format, R, G, B))
		updateScreen("s")
	}
}
func drawCircle(x0, y0, r int32, R, G, B uint8) {
	var x, y, dx, dy int32 = r - 1, 0, 1, 1
	var err = dx - (r * 2)

	for x > y {
		drawPixel(x0+x, y0+y, R, G, B)
		drawPixel(x0+y, y0+x, R, G, B)
		drawPixel(x0-y, y0+x, R, G, B)
		drawPixel(x0-x, y0+y, R, G, B)
		drawPixel(x0-x, y0-y, R, G, B)
		drawPixel(x0-y, y0-x, R, G, B)
		drawPixel(x0+y, y0-x, R, G, B)
		drawPixel(x0+x, y0-y, R, G, B)

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
	//renderer.Present()
	updateScreen("r")

}
func drawBubbles() {
	_ = renderer.SetDrawColor(0, 0, 0, 0)
	for i := 0; i <= 300; i++ {
		drawCircle(int32(rand.Intn(800)), int32(rand.Intn(600)), int32(rand.Intn(80)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)))
	}
	time.Sleep(time.Second)
}
func horizontalBars(R1, G1, B1, R2, G2, B2 uint8) {
	var i int32
	// Horizontal bars
	//_ = renderer.SetDrawColor(0, 120, 128,0)
	//_ = renderer.Clear()
	for i = 1; i < windowWidth; i++ {
		//L2R
		_ = surface.FillRect(&sdl.Rect{X: 0 + i, W: 60, H: 60}, sdl.MapRGB(surface.Format, R1, G1, B1))
		_ = surface.FillRect(&sdl.Rect{X: 0 + i, Y: 120, W: 60, H: 60}, sdl.MapRGB(surface.Format, R1, G1, B1))
		_ = surface.FillRect(&sdl.Rect{X: 0 + i, Y: 240, W: 60, H: 60}, sdl.MapRGB(surface.Format, R1, G1, B1))
		_ = surface.FillRect(&sdl.Rect{X: 0 + i, Y: 360, W: 60, H: 60}, sdl.MapRGB(surface.Format, R1, G1, B1))
		_ = surface.FillRect(&sdl.Rect{X: 0 + i, Y: 480, W: 60, H: 60}, sdl.MapRGB(surface.Format, R1, G1, B1))
		//R2L
		_ = surface.FillRect(&sdl.Rect{X: windowWidth - i, Y: 60, W: 60, H: 60}, sdl.MapRGB(surface.Format, R2, G2, B2))
		_ = surface.FillRect(&sdl.Rect{X: windowWidth - i, Y: 180, W: 60, H: 60}, sdl.MapRGB(surface.Format, R2, G2, B2))
		_ = surface.FillRect(&sdl.Rect{X: windowWidth - i, Y: 300, W: 60, H: 60}, sdl.MapRGB(surface.Format, R2, G2, B2))
		_ = surface.FillRect(&sdl.Rect{X: windowWidth - i, Y: 420, W: 60, H: 60}, sdl.MapRGB(surface.Format, R2, G2, B2))
		_ = surface.FillRect(&sdl.Rect{X: windowWidth - i, Y: 540, W: 60, H: 60}, sdl.MapRGB(surface.Format, R2, G2, B2))
		updateScreen("s")
	}
}
func horizontalBars2(R1, G1, B1, R2, G2, B2 uint8) {
	var i int32
	//Horizontal bars 2
	//_ = renderer.SetDrawColor(0, 120, 128,0)
	//_ = renderer.Clear()
	for i = 1; i < windowWidth; i++ {
		//L2R
		_ = surface.FillRect(&sdl.Rect{X: windowWidth - i, Y: 60, W: 60, H: 60}, sdl.MapRGB(surface.Format, R1, G1, B1))
		_ = surface.FillRect(&sdl.Rect{X: windowWidth - i, Y: 180, W: 60, H: 60}, sdl.MapRGB(surface.Format, R1, G1, B1))
		_ = surface.FillRect(&sdl.Rect{X: windowWidth - i, Y: 300, W: 60, H: 60}, sdl.MapRGB(surface.Format, R1, G1, B1))
		_ = surface.FillRect(&sdl.Rect{X: windowWidth - i, Y: 420, W: 60, H: 60}, sdl.MapRGB(surface.Format, R1, G1, B1))
		_ = surface.FillRect(&sdl.Rect{X: windowWidth - i, Y: 540, W: 60, H: 60}, sdl.MapRGB(surface.Format, R1, G1, B1))
		//R2L
		_ = surface.FillRect(&sdl.Rect{X: 0 + i, W: 60, H: 60}, sdl.MapRGB(surface.Format, R2, G2, B2))
		_ = surface.FillRect(&sdl.Rect{X: 0 + i, Y: 120, W: 60, H: 60}, sdl.MapRGB(surface.Format, R2, G2, B2))
		_ = surface.FillRect(&sdl.Rect{X: 0 + i, Y: 240, W: 60, H: 60}, sdl.MapRGB(surface.Format, R2, G2, B2))
		_ = surface.FillRect(&sdl.Rect{X: 0 + i, Y: 360, W: 60, H: 60}, sdl.MapRGB(surface.Format, R2, G2, B2))
		_ = surface.FillRect(&sdl.Rect{X: 0 + i, Y: 480, W: 60, H: 60}, sdl.MapRGB(surface.Format, R2, G2, B2))
		updateScreen("s")
	}
}
func updateScreen(surfaceOrRenderer string) {
	var lastTime uint32 = 0
	const ticksForNextFrame = 1000 / FPS

	for lastTime-sdl.GetTicks() < ticksForNextFrame {
		sdl.Delay(1)
	}

	if surfaceOrRenderer == "s" {
		_ = window.UpdateSurface()
	}
	if surfaceOrRenderer == "r" {
		renderer.Present()
	}
	//fmt.Println(surfaceOrRenderer)
	fmt.Println("Last time: ", lastTime)
	fmt.Println("GetTicks:", sdl.GetTicks())
	fmt.Println("ticksForNextFrame", ticksForNextFrame)

	lastTime = sdl.GetTicks()
}
