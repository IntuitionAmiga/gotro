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

	time.Sleep(time.Second * 3)
	window.Destroy()
	sdl.Quit()
}
