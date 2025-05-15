package neteasebox

import (
    "bytes"
    "context"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
    "netease-music-box-go/pkg/netease"
    githubapi "netease-music-box-go/pkg/github"
)

// Box 封装网易云和 Gist/Markdown 操作
//
type Box struct {
    UserID    string
    UserToken string
    GhToken   string
}

// NewBox 创建 Box 实例
//
// params:
//   - userID: 网易云用户ID
//   - userToken: 网易云用户Token
//   - ghToken: GitHub Token
// return: Box 实例
func NewBox(userID, userToken, ghToken string) *Box {
    return &Box{
        UserID:    userID,
        UserToken: userToken,
        GhToken:   ghToken,
    }
}

// GetStats 获取一周听歌排行，返回每一行内容
//
// return: []string，每行为一条排行
func (b *Box) GetStats(ctx context.Context) ([]string, error) {
    weekData, err := netease.GetUserRecord(b.UserID, b.UserToken)
    if err != nil {
        return nil, err
    }
    return b.GenerateGistLines(ctx, weekData)
}

// GenerateGistLines 生成排行榜每一行内容
//
// params:
//   - ctx: context
//   - weekData: 网易云一周听歌数据
// return: []string
func (b *Box) GenerateGistLines(ctx context.Context, weekData []netease.WeekData) ([]string, error) {
    icons := []string{"🥇", "🥈", "🥉", "🏅", "🏅"}
    var lines []string
    for i, d := range weekData {
        if i >= 5 {
            break
        }
        lines = append(lines, b.ConstructLine(ctx, i, d, icons))
    }
    if len(lines) == 0 {
        lines = append(lines, "Oh my God!\n~~~~~~\n我最近一周竟然没有听歌～\n~~~~~~")
    }
    return lines, nil
}

// ConstructLine 构造单行排行榜内容
//
// params:
//   - ctx: context
//   - idx: 排名索引
//   - d: WeekData
//   - icons: 奖牌icon数组
// return: string
func (b *Box) ConstructLine(ctx context.Context, idx int, d netease.WeekData, icons []string) string {
    artists := []string{}
    for _, a := range d.Song.Ar {
        artists = append(artists, a.Name)
    }
    name := d.Song.Name + " - " + strings.Join(artists, "/")
    flag := len([]rune(name)) > 15
    if flag {
        name = string([]rune(name)[:15]) + "..."
    }
    tab := "\t\t\t"
    if len([]rune(name)) <= 8 {
        tab = "\t\t\t\t"
    }
    return fmt.Sprintf("%s %s%s%d次    ", icons[idx], name, tab, d.PlayCount)
}

// UpdateGist 更新 Gist 内容
//
// params:
//   - ctx: context
//   - gistID: Gist ID
//   - filename: Gist 文件名
//   - lines: 内容行
// return: 错误
func (b *Box) UpdateGist(ctx context.Context, gistID, filename string, lines []string) error {
    content := strings.Join(lines, "\n")
    return githubapi.UpdateGist(b.GhToken, gistID, filename, content)
}

// UpdateMarkdown 更新 Markdown 文件注释块
//
// params:
//   - ctx: context
//   - title: 标题
//   - filename: Markdown 文件名
//   - lines: 内容行
//   - startTag: 注释块起始
//   - endTag: 注释块结束
// return: 错误
func (b *Box) UpdateMarkdown(ctx context.Context, title, filename string, lines []string, startTag, endTag string) error {
    md, err := ioutil.ReadFile(filename)
    if err != nil {
        return fmt.Errorf("neteasebox.UpdateMarkdown: 读取文件失败: %w", err)
    }
    start := []byte(startTag)
    before := md[:bytes.Index(md, start)+len(start)]
    end := []byte(endTag)
    after := md[bytes.Index(md, end):]

    newMd := bytes.NewBuffer(nil)
    newMd.Write(before)
    newMd.WriteString("\n" + title + "\n")
    newMd.WriteString("```text\n")
    newMd.WriteString(strings.Join(lines, "\n"))
    newMd.WriteString("\n```")
    newMd.WriteString("\n<!-- Powered by https://github.com/Viper373/netease-music-box-go . -->\n")
    newMd.Write(after)

    err = ioutil.WriteFile(filename, newMd.Bytes(), os.ModeAppend)
    if err != nil {
        return fmt.Errorf("neteasebox.UpdateMarkdown: 写入文件失败: %w", err)
    }
    return nil
} 