package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"os"
	"runtime"
	"time"
)

const (
	FPS                       = 60
	windowWidth, windowHeight = 800, 600
)

var window *sdl.Window
var renderer *sdl.Renderer

func main() {
	sdlInitVideo()
	sdlInitImage()
	sdlInitAudio()
	window = createWindow()
	renderer = createRenderer()
	var _ = sdl.PollEvent() //MacOS won't draw the window without this line

	//Start intro
	//_ = showKickstart()
	//playFloppySounds()
	//time.Sleep(time.Second * 2)
	//
	//backgroundFill(255, 255, 255) //Fill bg with white
	//time.Sleep(time.Second * 9)
	//decrunch(100)

	playMusic()

	//wipeLeft(255, 0, 90)
	//wipeRight(0, 120, 128)
	//horizontalBars2(30, 0, 95, 30, 055, 200)
	//wipeLeft(0, 120, 128)
	//wipeRight(255, 0, 90)
	//wipeLeft(255, 0, 90)
	//
	//boingBall(255, 0, 90)
	//
	rasterBars()
	rainbowScroll()
	//
	wipeTopDown(0, 0, 0)
	drawBubbles()
	wipeTopDown(0, 0, 0)

	//wipeLeft(95, 95, 0)
	//wipeRight(0, 95, 0)

	horizontalBars(0, 0, 95, 0, 055, 200)
	//horizontalBars(30, 0, 95, 30, 055, 200)

	wipeTopDown(0, 0, 0)
	//wipeLeft(0, 0, 0)
	//wipeRight(0, 0, 0)

	_ = renderer.Destroy()
	_ = window.Destroy()
}
func sdlInitVideo() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to initialise video: %s\n", err)
		os.Exit(1)
	}

	defer sdl.Quit()
}
func sdlInitImage() {
	err := img.Init(img.INIT_PNG)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to initialise image lib: %s\n", err)
		os.Exit(1)
	}

	defer img.Quit()
}
func sdlInitAudio() {
	errSDLAudioInit := sdl.Init(sdl.INIT_AUDIO)
	if errSDLAudioInit != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to initialise audio: %s\n", errSDLAudioInit)
		os.Exit(1)
	}

	errOpeningAudioDevice := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096)
	if errOpeningAudioDevice != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to open audio device: %s\n", errOpeningAudioDevice)
		os.Exit(1)
	}

	errSDLMixerInit := mix.Init(mix.INIT_MOD)
	if errSDLMixerInit != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to initialise mixer: %s\n", errSDLMixerInit)
		os.Exit(1)
	}

}
func createWindow() *sdl.Window {
	window, errCreatingSDLWindow := sdl.CreateWindow("Gotro by Intuition",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		windowWidth,
		windowHeight,
		sdl.WINDOW_SHOWN|sdl.WINDOW_OPENGL)

	if errCreatingSDLWindow != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to create SDL window: %s\n", errCreatingSDLWindow)
		os.Exit(1)
	}
	return window
}
func createRenderer() *sdl.Renderer {

	var numDrivers, _ = sdl.GetNumRenderDrivers()
	fmt.Println(numDrivers)

	var errCreatingSDLRenderer error
	sdl.SetHint(sdl.HINT_RENDER_VSYNC, "1")
	if runtime.GOOS == "darwin" {
		sdl.SetHint(sdl.HINT_RENDER_DRIVER, "software")
	} else {
		sdl.SetHint(sdl.HINT_RENDER_DRIVER, "opengl")
	}
	renderer, errCreatingSDLRenderer = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC|sdl.RENDERER_TARGETTEXTURE)

	/*
		for i := 0; i < numDrivers; i++ {
			driverInfo, _ := renderer.GetInfo()
			fmt.Println("Driver name (", i, "): ", driverInfo.Name)
			//if (driverInfo.Name == "SDL_RENDERER_SOFTWARE") {fmt.Println(" the renderer is a software fallback")}
			//if (driverInfo.Name == "SDL_RENDERER_ACCELERATED") {fmt.Println(" the renderer uses hardware acceleration")}
			//if (driverInfo.Name == "SDL_RENDERER_PRESENTVSYNC") {fmt.Println(" present	is synchronized with the refresh rate")}
			//if (driverInfo.Name == "SDL_RENDERER_TARGETTEXTURE") {fmt.Println( " the renderer supports rendering to texture")}
		}
	*/

	/*var info sdl.RendererInfo
	info, _ =renderer.GetInfo()
	fmt.Println(info.Name)
	*/

	if errCreatingSDLRenderer != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", errCreatingSDLRenderer)
		os.Exit(1)
	}
	return renderer
}
func showKickstart() error {
	/*
		t, err := img.LoadTexture(renderer, "kick13.png")
		if err != nil {
			return fmt.Errorf("couldn't load image from disk")
		}
		if err := renderer.Copy(t, nil, nil); err != nil {
			return fmt.Errorf("couldn't copy texture: %v", err)
		}
		_ = renderer.SetDrawColor(255, 255, 255, 0)
		//updateScreen("r")
		renderer.Present()
		return err
	*/
	kickrect := sdl.Rect{W: 800, H: 600}
	s, _ := img.Load("kick13.png")
	t, _ := renderer.CreateTextureFromSurface(s)
	err := renderer.Copy(t, nil, &kickrect)
	if err != nil {
		return err
	}
	updateScreen()
	err = renderer.Clear()
	if err != nil {
		return err
	}
	return nil
}
func playFloppySounds() {
	if music, errLoadingMusic := mix.LoadMUS("./floppy.mp3"); errLoadingMusic != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to load music from disk: %s\n", errLoadingMusic)
		os.Exit(1)
	} else {
		_ = music.Play(1)
	}
}
func backgroundFill(R, G, B uint8) {
	drawFillRect(0, 0, windowWidth, windowHeight, R, G, B)
	updateScreen()
	_ = renderer.Clear()
}
func decrunch(loops int) {
	var startY, n int32 = 0, 0
	for x := 0; x <= loops; x++ {
		n = 0
		for i := 0; i <= windowHeight; i++ {
			drawFillRect(0, n+startY, windowWidth, 8, 46, 43, 95)
			drawFillRect(0, n+startY+8, windowWidth, 8, 255, 0, 0)
			drawFillRect(0, n+startY+16, windowWidth, 8, 139, 0, 255)
			drawFillRect(0, n+startY+24, windowWidth, 8, 255, 255, 0)
			drawFillRect(0, n+startY+32, windowWidth, 8, 0, 255, 0)
			n += 40
		}
		updateScreen()
		n = 0
		for i := 0; i <= windowHeight; i++ {
			drawFillRect(0, n+startY, windowWidth, 4, 0, 255, 0)
			drawFillRect(0, n+startY+6, windowWidth, 4, 255, 255, 0)
			drawFillRect(0, n+startY+14, windowWidth, 4, 46, 43, 95)
			drawFillRect(0, n+startY+22, windowWidth, 4, 255, 0, 0)
			drawFillRect(0, n+startY+30, windowWidth, 4, 139, 0, 255)
			n += 40
		}
		updateScreen()
		n = 0
		for i := 0; i <= windowHeight; i++ {
			drawFillRect(0, n+startY, windowWidth, 12, 70, 55, 0)
			drawFillRect(0, n+startY+11, windowWidth, 12, 55, 25, 70)
			drawFillRect(0, n+startY+19, windowWidth, 12, 46, 113, 5)
			drawFillRect(0, n+startY+27, windowWidth, 12, 25, 0, 70)
			drawFillRect(0, n+startY+35, windowWidth, 12, 13, 70, 55)
			n += 40
		}
		updateScreen()
	}
	backgroundFill(255, 255, 255) //Fill bg with white
	sdl.Delay(1000)
}
func playMusic() {
	if music, errLoadingMusic := mix.LoadMUS("echoing.mod"); errLoadingMusic != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to load music from disk: %s\n", errLoadingMusic)
		os.Exit(1)
	} else {
		_ = music.Play(1)
	}
}

var currentColor sdl.Color

func setDrawColor(r, g, b uint8) {
	color := sdl.Color{R: r, G: g, B: b, A: 255}
	if color != currentColor {
		renderer.SetDrawColor(r, g, b, 255)
		currentColor = color
	}
}

func drawPixel(x, y int32, r, g, b uint8) {
	setDrawColor(r, g, b)
	renderer.DrawPoint(x, y)
}

var currentFillColor sdl.Color

func setFillColor(r, g, b uint8) {
	color := sdl.Color{R: r, G: g, B: b, A: 255}
	if color != currentFillColor {
		renderer.SetDrawColor(r, g, b, 255)
		currentFillColor = color
	}
}

func drawFillRect(x, y, w, h int32, r, g, b uint8) {
	setFillColor(r, g, b)
	rect := sdl.Rect{X: x, Y: y, W: w, H: h}
	renderer.FillRect(&rect)
}

func wipeLeft(R, G, B uint8) {
	var i int32 = 0
	//Mid to left full length screen wipe
	renderer.Clear()
	for i = 0; i <= (windowWidth / 2); i++ {
		drawFillRect((windowWidth/2)-i, 0, i, windowHeight, R, G, B)
		// Update the screen periodically to maintain animation
		if i%2 == 0 {
			updateScreen()
		}
	}
}

func wipeRight(R, G, B uint8) {
	var i int32 = 0
	renderer.Clear()
	//Mid to left full length screen wipe
	for i = 0; i <= (windowWidth / 2); i++ {
		drawFillRect(windowWidth/2, 0, i+1, windowHeight, R, G, B)
		// Update the screen periodically to maintain animation
		if i%2 == 0 {
			updateScreen()
		}
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
	updateScreen()
}

func boingBall(R, G, B uint8) {
	var xPos, yPos int
	for i := 0; i <= (windowHeight - 128); i++ {
		drawSprite(int32(i), int32(i), R, G, B)
		xPos = i
		yPos = i
	}
	for i := xPos; i <= (windowWidth - 128); i++ {
		drawSprite(int32(i), int32(yPos+10), R, G, B)
		yPos -= 1
		xPos = i
	}
	for i := yPos; i >= 0; i-- {
		drawSprite(int32(xPos+i), int32(i), R, G, B)
		xPos = i
		yPos = -i
	}
}

func rasterBars() {
	var startX int32 = windowWidth
	var redY int32 = 0
	var greenY = int32(windowHeight / 3)
	var blueY = int32((windowHeight / 3) * 2)

	redBar := func() {
		for i := int32(0); i <= 35; i++ {
			//Red
			redY += i
			drawFillRect(startX-windowWidth, redY+24, windowWidth, 4, 8, 0, 0)
			drawFillRect(startX-windowWidth, redY+20, windowWidth, 4, 16, 0, 0)
			drawFillRect(startX-windowWidth, redY+16, windowWidth, 4, 32, 0, 0)
			drawFillRect(startX-windowWidth, redY+12, windowWidth, 4, 63, 0, 0)
			drawFillRect(startX-windowWidth, redY+8, windowWidth, 4, 127, 0, 0)
			drawFillRect(startX-windowWidth, redY+4, windowWidth, 4, 191, 0, 0)
			drawFillRect(startX-windowWidth, redY, windowWidth, 4, 255, 0, 0)
			drawFillRect(startX-windowWidth, redY-4, windowWidth, 4, 191, 0, 0)
			drawFillRect(startX-windowWidth, redY-8, windowWidth, 4, 127, 0, 0)
			drawFillRect(startX-windowWidth, redY-12, windowWidth, 4, 63, 0, 0)
			drawFillRect(startX-windowWidth, redY-16, windowWidth, 4, 32, 0, 0)
			drawFillRect(startX-windowWidth, redY-20, windowWidth, 4, 16, 0, 0)
			drawFillRect(startX-windowWidth, redY-24, windowWidth, 4, 8, 0, 0)
			time.Sleep(time.Second / 8)
			updateScreen()
		}
	}

	greenBar := func() {
		for i := int32(0); i <= 35; i++ {
			//Green
			greenY += i
			drawFillRect(startX-windowWidth, greenY+24, windowWidth, 4, 0, 8, 0)
			drawFillRect(startX-windowWidth, greenY+20, windowWidth, 4, 0, 16, 0)
			drawFillRect(startX-windowWidth, greenY+16, windowWidth, 4, 0, 32, 0)
			drawFillRect(startX-windowWidth, greenY+12, windowWidth, 4, 0, 63, 0)
			drawFillRect(startX-windowWidth, greenY+8, windowWidth, 4, 0, 127, 0)
			drawFillRect(startX-windowWidth, greenY+4, windowWidth, 4, 0, 191, 0)
			drawFillRect(startX-windowWidth, greenY, windowWidth, 4, 0, 255, 0)
			drawFillRect(startX-windowWidth, greenY-4, windowWidth, 4, 0, 191, 0)
			drawFillRect(startX-windowWidth, greenY-8, windowWidth, 4, 0, 127, 0)
			drawFillRect(startX-windowWidth, greenY-12, windowWidth, 4, 0, 63, 0)
			drawFillRect(startX-windowWidth, greenY-16, windowWidth, 4, 0, 32, 0)
			drawFillRect(startX-windowWidth, greenY-20, windowWidth, 4, 0, 16, 0)
			drawFillRect(startX-windowWidth, greenY-24, windowWidth, 4, 0, 8, 0)
			time.Sleep(time.Second / 8)
			updateScreen()
		}
	}
	blueBar := func() {
		for i := int32(0); i <= 35; i++ {
			//Blue
			blueY += i
			drawFillRect(startX-windowWidth, blueY+24, windowWidth, 4, 0, 0, 8)
			drawFillRect(startX-windowWidth, blueY+20, windowWidth, 4, 0, 0, 16)
			drawFillRect(startX-windowWidth, blueY+16, windowWidth, 4, 0, 0, 32)
			drawFillRect(startX-windowWidth, blueY+12, windowWidth, 4, 0, 0, 63)
			drawFillRect(startX-windowWidth, blueY+8, windowWidth, 4, 0, 0, 127)
			drawFillRect(startX-windowWidth, blueY+4, windowWidth, 4, 0, 0, 191)
			drawFillRect(startX-windowWidth, blueY, windowWidth, 4, 0, 0, 255)
			drawFillRect(startX-windowWidth, blueY-4, windowWidth, 4, 0, 0, 191)
			drawFillRect(startX-windowWidth, blueY-8, windowWidth, 4, 0, 0, 127)
			drawFillRect(startX-windowWidth, blueY-12, windowWidth, 4, 0, 0, 63)
			drawFillRect(startX-windowWidth, blueY-16, windowWidth, 4, 0, 0, 32)
			drawFillRect(startX-windowWidth, blueY-20, windowWidth, 4, 0, 0, 16)
			drawFillRect(startX-windowWidth, blueY-24, windowWidth, 4, 0, 0, 8)
			time.Sleep(time.Second / 8)
			updateScreen()
		}
	}

	redBar()
	greenBar()
	blueBar()
}
func rainbowScroll() {
	var startY int32 = windowHeight/2 + 16
	renderer.Clear()
	for i := 0; i < windowWidth; i++ {
		drawFillRect(windowWidth-int32(i), startY-48, 30, 16, 255, 0, 0)
		drawFillRect(windowWidth-int32(i), startY-32, 30, 16, 255, 127, 0)
		drawFillRect(windowWidth-int32(i), startY-16, 30, 16, 255, 255, 0)
		drawFillRect(windowWidth-int32(i), startY, 30, 16, 0, 255, 0)
		drawFillRect(windowWidth-int32(i), startY+16, 30, 16, 0, 0, 255)
		drawFillRect(windowWidth-int32(i), startY+32, 30, 16, 46, 43, 95)
		drawFillRect(windowWidth-int32(i), startY+48, 30, 16, 139, 0, 255)
		// Update the screen periodically to maintain animation
		if i%2 == 0 {
			updateScreen()
		}
	}
}
func wipeTopDown(R, G, B uint8) {
	var i int32
	//Clear top to bottom
	renderer.Clear()
	for i = 1; i <= windowHeight; i++ {
		drawFillRect(0, 0, windowWidth, 0+i, R, G, B)
		// Update the screen periodically to maintain animation
		if i%2 == 0 {
			updateScreen()
		}
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
}
func drawBubbles() {
	_ = renderer.SetDrawColor(0, 0, 0, 0)
	_ = renderer.Clear()
	for i := 0; i <= 300; i++ {
		drawCircle(int32(rand.Intn(800)), int32(rand.Intn(600)), int32(rand.Intn(80)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)))
		// Update the screen periodically to maintain animation
		if i%2 == 0 {
			updateScreen()
		}
		time.Sleep(time.Second / 500)
	}
}
func horizontalBars(R1, G1, B1, R2, G2, B2 uint8) {
	var i int32
	// Horizontal bars
	for i = 1; i < windowWidth; i++ {
		//L2R
		drawFillRect(0+i, 0, 60, 60, R1, G1, B1)
		drawFillRect(0+i, 120, 60, 60, R1, G1, B1)
		drawFillRect(0+i, 240, 60, 60, R1, G1, B1)
		drawFillRect(0+i, 360, 60, 60, R1, G1, B1)
		drawFillRect(0+i, 480, 60, 60, R1, G1, B1)
		//R2L
		drawFillRect(windowWidth-i, 60, 60, 60, R2, G2, B2)
		drawFillRect(windowWidth-i, 180, 60, 60, R2, G2, B2)
		drawFillRect(windowWidth-i, 300, 60, 60, R2, G2, B2)
		drawFillRect(windowWidth-i, 420, 60, 60, R2, G2, B2)
		drawFillRect(windowWidth-i, 540, 60, 60, R2, G2, B2)
		// Update the screen periodically to maintain animation
		if i%2 == 0 {
			updateScreen()
		}
	}
}
func horizontalBars2(R1, G1, B1, R2, G2, B2 uint8) {
	var i int32
	//Horizontal bars 2
	//_ = renderer.SetDrawColor(0, 120, 128,0)
	//_ = renderer.Clear()
	for i = 1; i < windowWidth; i++ {
		//L2R
		drawFillRect(windowWidth-i, 60, 60, 60, R1, G1, B1)
		drawFillRect(windowWidth-i, 180, 60, 60, R1, G1, B1)
		drawFillRect(windowWidth-i, 300, 60, 60, R1, G1, B1)
		drawFillRect(windowWidth-i, 420, 60, 60, R1, G1, B1)
		drawFillRect(windowWidth-i, 540, 60, 60, R1, G1, B1)
		//R2L
		drawFillRect(0+i, 0, 60, 60, R2, G2, B2)
		drawFillRect(0+i, 120, 60, 60, R2, G2, B2)
		drawFillRect(0+i, 240, 60, 60, R2, G2, B2)
		drawFillRect(0+i, 360, 60, 60, R2, G2, B2)
		drawFillRect(0+i, 480, 60, 60, R2, G2, B2)
		// Update the screen periodically to maintain animation
		if i%2 == 0 {
			updateScreen()
		}
	}
}
func updateScreen() {
	const ticksForNextFrame uint32 = 1000 / FPS
	lastTime := sdl.GetTicks()
	if sdl.GetTicks()-lastTime < ticksForNextFrame {
		sdl.Delay(1)
		/*time.Sleep(time.Second / 1000)
		fmt.Println("\nlastTime: ", lastTime)
		fmt.Println("ticksForNextFrame: ", ticksForNextFrame)
		fmt.Println("sdl.GetTicks(): ", sdl.GetTicks())
		fmt.Println("GetTicks()-lastTime: ", sdl.GetTicks()-lastTime)
		*/
	}
	renderer.Present()
}
