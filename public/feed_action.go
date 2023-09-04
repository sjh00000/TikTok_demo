package public

import (
	"fmt"
	"os"
	"os/exec"
)

func GetFeedCover(videoPath string) string {
	outputImagePath := "./public/" + videoPath + "_image.jpg"
	fmt.Println(outputImagePath + "\n" + videoPath + "\n")
	// 检查视频文件是否存在
	if _, err := os.Stat("./public/" + videoPath); os.IsNotExist(err) {
		fmt.Printf("视频文件不存在: %s\n", videoPath)
	}

	// 使用 FFmpeg 命令行工具截取视频的第一帧
	cmd := exec.Command("ffmpeg", "-i", "./public/"+videoPath, "-vframes", "1", "-q:v", "2", outputImagePath)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("截取视频第一帧失败: %v", err)
	}

	return videoPath + "_image.jpg"
}

func CutFeed(inputFile string) {
	outputFile := "zip_" + inputFile // 输出压缩后的视频文件路径

	// 执行FFmpeg命令
	cmd := exec.Command("ffmpeg", "-i", "./public/"+inputFile, "-r", "10", "-vf", "scale=480:320", "-b:v", "0.5M", "./public/"+outputFile)
	err := cmd.Run()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	err = os.Remove("./public/" + inputFile)
	if err != nil {
		fmt.Println("Failed to remove backup file:", err)
		return
	}
}
