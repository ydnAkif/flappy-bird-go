package soundgen

import (
	"encoding/binary"
	"math"
	"os"
)

const (
	sampleRate = 44100
	duration   = 0.1 // seconds
)

// GenerateSounds generates all game sound effects
func GenerateSounds() {
	// Generate jump sound (higher pitch)
	generateSound("../assets/audio/jump.wav", 880.0, 0.1)

	// Generate score sound (medium pitch, longer)
	generateSound("../assets/audio/score.wav", 660.0, 0.15)

	// Generate hit sound (lower pitch, shorter)
	generateSound("../assets/audio/hit.wav", 220.0, 0.05)
}

func generateSound(filename string, frequency, duration float64) {
	// Create WAV header
	header := make([]byte, 44)
	
	// "RIFF" chunk descriptor
	copy(header[0:4], []byte("RIFF"))
	// Chunk size (file size - 8)
	binary.LittleEndian.PutUint32(header[4:8], uint32(duration*float64(sampleRate)*2+36))
	// Format ("WAVE")
	copy(header[8:12], []byte("WAVE"))
	
	// "fmt " sub-chunk
	copy(header[12:16], []byte("fmt "))
	// Sub-chunk size (16 for PCM)
	binary.LittleEndian.PutUint32(header[16:20], 16)
	// Audio format (1 for PCM)
	binary.LittleEndian.PutUint16(header[20:22], 1)
	// Number of channels (1 for mono)
	binary.LittleEndian.PutUint16(header[22:24], 1)
	// Sample rate
	binary.LittleEndian.PutUint32(header[24:28], sampleRate)
	// Byte rate
	binary.LittleEndian.PutUint32(header[28:32], sampleRate*2)
	// Block align
	binary.LittleEndian.PutUint16(header[32:34], 2)
	// Bits per sample
	binary.LittleEndian.PutUint16(header[34:36], 16)
	
	// "data" sub-chunk
	copy(header[36:40], []byte("data"))
	// Sub-chunk size
	binary.LittleEndian.PutUint32(header[40:44], uint32(duration*float64(sampleRate)*2))

	// Create the file
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Write header
	f.Write(header)

	// Generate samples
	numSamples := int(duration * float64(sampleRate))
	for i := 0; i < numSamples; i++ {
		t := float64(i) / float64(sampleRate)
		
		// Generate a simple sine wave with exponential decay
		amplitude := math.Exp(-t*8.0) * 32767.0
		sample := int16(amplitude * math.Sin(2.0*math.Pi*frequency*t))
		
		// Write sample
		binary.Write(f, binary.LittleEndian, sample)
	}
}
