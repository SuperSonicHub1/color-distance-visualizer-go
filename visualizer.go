package main

import (
	"fmt"
	"github.com/AlexEidt/Vidio"
)

type PreviousCurrent struct {
	Previous []byte
	Current  []byte
}

/// A nice little iterator.
func previousCurrentFrames(video *vidio.Video) <-chan PreviousCurrent {
	ch := make(chan PreviousCurrent)

	width, height := video.Width(), video.Height()
	length := width * height * 3
	// The only thing I don't like is all this
	// awkwardness in order to copy an array
	go func() {
		var current []byte
		var previous []byte
		for video.Read() {
			if current != nil {
				previous = make([]byte, length)
				copy(previous, current)
			}
			current = video.FrameBuffer()
			ch <- PreviousCurrent{Previous: previous, Current: current}
		}
		close(ch)
	}()

	return ch
}

func visualize(
	input, output string,
	unchangedPixels bool,
	vcodec string,
) error {
	inputVideo, err := vidio.NewVideo(input)
	if err != nil {
		return err
	}

	width, height := inputVideo.Width(), inputVideo.Height()
	if vcodec == "" {
		vcodec = inputVideo.Codec()
	}

	var audioInput string
	if inputVideo.AudioCodec() != "" {
		audioInput = input
	}

	writer, err := vidio.NewVideoWriter(
		output,
		width,
		height,
		&vidio.Options{
			Codec: vcodec,
			Audio: audioInput,
		},
	)
	if err != nil {
		return err
	}
	defer writer.Close()

	i := 0
	totalFrames := inputVideo.Duration() * inputVideo.FPS()
	for pc := range previousCurrentFrames(inputVideo) {
		previous, current := pc.Previous, pc.Current

		length := width * height * 3
		frame := make([]byte, length)

		// If previous is nil, we'll write a black frame.
		// Otherwise, we'll create a frame.
		if previous != nil {
			for i := 0; i < length; i += 3 {
				// uint8s are always positive, so there's no need to use abs
				redDistance := current[i] - previous[i]
				greenDistance := current[i+1] - previous[i+1]
				blueDistance := current[i+2] - previous[i+2]

				// If our new pixel is (0, 0, 0),
				// that means the pixel in the previous
				// frame and in the current are the same
				if unchangedPixels && redDistance == 0 && blueDistance == 0 && greenDistance == 0 {
					frame[i] = current[i]
					frame[i+1] = current[i+1]
					frame[i+2] = current[i+2]
				} else {
					frame[i] = redDistance
					frame[i+1] = greenDistance
					frame[i+2] = blueDistance
				}
			}
		}

		writer.Write(frame)

		i++
		fmt.Printf("%d/%d (%f%%) \n", i, int(totalFrames), (float64(i)/totalFrames)*100)
	}

	return nil
}
