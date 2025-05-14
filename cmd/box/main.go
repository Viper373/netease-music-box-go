package main

import (
	"context"
	"fmt"
	"os"
	"netease-music-box-go/pkg/neteasebox"
)

// 获取环境变量，若不存在则返回默认值
func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

func main() {
	userID := os.Getenv("NETEASE_USER_ID")
	userToken := os.Getenv("NETEASE_USER_TOKEN")
	ghToken := os.Getenv("GH_TOKEN")
	gistID := os.Getenv("GIST_ID")
	updateOption := getEnv("UPDATE_OPTION", "GIST_AND_MARKDOWN")
	markdownFile := getEnv("MARKDOWN_FILE", "README.md")
	startTag := getEnv("MARKDOWN_START_TAG", "<!-- netease-music-box start -->")
	endTag := getEnv("MARKDOWN_END_TAG", "<!-- netease-music-box end -->")

	var updateGist, updateMarkdown bool
	if updateOption == "MARKDOWN" {
		updateMarkdown = true
	} else if updateOption == "GIST_AND_MARKDOWN" {
		updateGist = true
		updateMarkdown = true
	} else {
		updateGist = true
	}
	
	box := neteasebox.NewBox(userID, userToken, ghToken)
	ctx := context.Background()
	lines, err := box.GetStats(ctx)
	if err != nil {
		fmt.Println("获取网易云听歌记录失败:", err)
		os.Exit(1)
	}
	filename := "🎵 Weekly Listen Music Rank"

	if updateGist {
		err := box.UpdateGist(ctx, gistID, filename, lines)
		if err != nil {
			fmt.Println("更新Gist失败:", err)
			os.Exit(1)
		}
		fmt.Println("Gist已更新")
	}

	if updateMarkdown && markdownFile != "" {
		title := filename
		if updateGist {
			title = fmt.Sprintf("#### <a href=\"https://gist.github.com/%s\" target=\"_blank\">%s</a>", gistID, filename)
		} else {
			title = fmt.Sprintf("#### %s", filename)
		}
		err := box.UpdateMarkdown(ctx, title, markdownFile, lines, startTag, endTag)
		if err != nil {
			fmt.Println("更新Markdown失败:", err)
		} else {
			fmt.Println("Markdown已更新")
		}
	}
} 