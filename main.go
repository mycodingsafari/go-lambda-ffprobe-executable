package main

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
)

//ProbeOutput
type ProbeOutput struct {
	Streams []struct {
		Index     string `json:"index"`
		CodecName string `json:"codec_name"`
		CodecType string `json:"codec_type"`
		Framerate string `json:"r_frame_rate"`
	} `json:"streams"`
	Format struct {
		Duration string `json:"duration"`
		Bitrate  string `json:"bit_rate"`
	} `json:"format"`
}

func ProbeCommand(mediaPath string) *exec.Cmd {
	// get path of the current executable
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	ffprobePath := filepath.Join(filepath.Dir(ex), "bin/ffprobe")
	command := exec.Command(
		ffprobePath,
		"-v", "quiet",
		"-print_format", "json",
		"-show_format",
		"-show_streams",
		"-i", mediaPath,
	)

	return command
}

func Probe(mediaPath string) (ProbeOutput, error) {
	command := ProbeCommand(mediaPath)
	var out bytes.Buffer
	var pb ProbeOutput
	command.Stdout = &out
	err := command.Run()

	if err != nil {
		return ProbeOutput{}, err
	}

	json.Unmarshal(out.Bytes(), &pb)

	return pb, nil
}

func handler() error {
	mediaPath := "https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4"
	p, _ := Probe(mediaPath)
	log.Info(p)
	return nil
}

func main() {
	lambda.Start(handler)
}
