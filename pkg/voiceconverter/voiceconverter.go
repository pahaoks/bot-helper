package voiceconverter

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type VoiceConverter struct {
}

func NewVoice() *VoiceConverter {
	return &VoiceConverter{}
}

// OggToMp3 converts ogg data to mp3 and returns path to mp3 file
func (v *VoiceConverter) OggToMp3(data []byte) (*os.File, error) {
	ts := time.Now().UnixNano()
	oggFilePath := fmt.Sprintf("/tmp/voice-%d.ogg", ts)
	oggFile, err := os.Create(oggFilePath)
	if err != nil {
		return nil, err
	}

	oggFile.Write(data)
	oggFile.Close()

	mp3FilePath := fmt.Sprintf("/tmp/voice-%d.mp3", ts)
	cmd := exec.Command(
		"ffmpeg",
		"-y", // overwrite if output exists
		"-i", oggFilePath,
		"-codec:a", "libmp3lame",
		"-qscale:a", "2", // quality (0 = best, 9 = worst)
		mp3FilePath,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr // show ffmpeg logs
	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	mp3File, err := os.Open(mp3FilePath)
	if err != nil {
		return nil, err
	}

	return mp3File, nil
}
