package public

import (
	"log"
	"os"
	"os/exec"
)

func GetFeedCover(videoPath string) string {
	outputImagePath := videoPath + "_image"

	// 检查视频文件是否存在
	if _, err := os.Stat(videoPath); os.IsNotExist(err) {
		log.Fatalf("视频文件不存在: %s", videoPath)
	}

	// 使用 FFmpeg 命令行工具截取视频的第一帧
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-vframes", "1", "-q:v", "2", outputImagePath)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("截取视频第一帧失败: %v", err)
	}

	return outputImagePath
}
