package netease

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

// Song 歌曲结构体
//
// 包含歌曲名和歌手信息
//
type Artist struct {
    Name string `json:"name"`
}
type Song struct {
    Name string   `json:"name"`
    Ar   []Artist `json:"ar"`
}

// WeekData 一周听歌数据结构体
//
// 包含播放次数和歌曲信息
//
type WeekData struct {
    PlayCount int  `json:"playCount"`
    Song      Song `json:"song"`
}

// RecordBody API返回的body部分
//
type RecordBody struct {
    WeekData []WeekData `json:"weekData"`
}

// RecordResp API完整返回结构体
//
type RecordResp struct {
    Code int        `json:"code"`
    Body RecordBody `json:"body"`
}

// GetUserRecord 获取网易云用户一周听歌记录
//
// params:
//   - userID: 网易云用户ID
//   - userToken: 网易云用户Token
// return: weekData列表, 错误
//
func GetUserRecord(userID, userToken string) ([]WeekData, error) {
    url := fmt.Sprintf("https://neteasecloudmusicapi.vercel.app/user/record?uid=%s&type=1", userID)
    client := &http.Client{}
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("Cookie", fmt.Sprintf("MUSIC_U=%s", userToken))
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Printf("网易云API返回内容: %s\n", string(body))
    var record RecordResp
    if err := json.Unmarshal(body, &record); err != nil {
        return nil, err
    }
    if record.Code != 200 {
        return nil, fmt.Errorf("netease api error: %s", string(body))
    }
    return record.Body.WeekData, nil
} 