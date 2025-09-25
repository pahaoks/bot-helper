package voiceconverter

import (
	_ "embed"
	"io"
	"testing"
)

//go:embed testdata/file_0.oga
var oggData []byte

func TestTransform(t *testing.T) {
	v := NewVoice()

	mp3File, err := v.OggToMp3(oggData)
	if err != nil {
		t.Fatal(err)
	}
	defer mp3File.Close()

	// Check if the mp3 file is created
	if mp3File.Name() == "" {
		t.Fatal("mp3 file not created")
	}

	b, err := io.ReadAll(mp3File)
	if err != nil {
		t.Fatal(err)
	}

	if len(b) == 0 {
		t.Fatal("mp3 file is empty")
	}
}
