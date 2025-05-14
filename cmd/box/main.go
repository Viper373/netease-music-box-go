package main

import (
    "context"
    "fmt"
    "os"
    "netease-music-box/pkg/neteasebox"
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
    userID := os.Getenv("USER_ID")
    userToken := os.Getenv("USER_TOKEN")
    gistID := os.Getenv("GIST_ID")
    ghToken := os.Getenv("GH_TOKEN")
    updateOption := getEnv("UPDATE_OPTION", "GIST")
    gistFilename := getEnv("GIST_FILENAME", "netease-music-box.md")
    markdownFile := getEnv("MARKDOWN_FILE", "README.md")
    startTag := getEnv("MARKDOWN_START_TAG", "<!-- netease-music-box start -->")
    endTag := getEnv("MARKDOWN_END_TAG", "<!-- netease-music-box end -->")

    // 预留样式参数，便于后续扩展
    style := neteasebox.BoxStyle{}

    var updateGist, updateMarkdown bool
    if updateOption == "MARKDOWN" {
        updateMarkdown = true
    } else if updateOption == "GIST_AND_MARKDOWN" {
        updateGist = true
        updateMarkdown = true
    } else {
        updateGist = true
    }

    box := neteasebox.NewBox(userID, userToken, ghToken, style)
    ctx := context.Background()
    lines, err := box.GetStats(ctx)
    if err != nil {
        fmt.Println("获取网易云听歌记录失败:", err)
        os.Exit(1)
    }

    if updateGist {
        err := box.UpdateGist(ctx, gistID, gistFilename, lines)
        if err != nil {
            fmt.Println("更新Gist失败:", err)
        } else {
            fmt.Println("Gist已更新")
        }
    }

    if updateMarkdown && markdownFile != "" {
        // 标题：如果也更新gist则带gist链接，否则只显示文件名
        title := gistFilename
        if updateGist {
            title = fmt.Sprintf(`#### <a href=\"https://gist.github.com/%s\" target=\"_blank\">%s</a>`, gistID, gistFilename)
        } else {
            title = fmt.Sprintf("#### %s", gistFilename)
        }
        err := box.UpdateMarkdown(ctx, title, markdownFile, lines, startTag, endTag)
        if err != nil {
            fmt.Println("更新Markdown失败:", err)
        } else {
            fmt.Println("Markdown已更新")
        }
    }
} 