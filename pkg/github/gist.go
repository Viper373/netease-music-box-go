package github

import (
    "context"
    "github.com/google/go-github/v50/github"
    "golang.org/x/oauth2"
)

// UpdateGist 更新指定Gist内容
//
// params:
//   - token: GitHub Token
//   - gistID: Gist ID
//   - filename: Gist文件名
//   - content: 新内容
// return: 错误
//
func UpdateGist(token, gistID, filename, content string) error {
    ctx := context.Background()
    ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
    tc := oauth2.NewClient(ctx, ts)
    client := github.NewClient(tc)
    files := map[github.GistFilename]github.GistFile{
        github.GistFilename(filename): {Content: github.String(content)},
    }
    _, _, err := client.Gists.Edit(ctx, gistID, &github.Gist{Files: files})
    return err
} 