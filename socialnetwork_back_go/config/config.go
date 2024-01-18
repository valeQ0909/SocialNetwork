package config

// jwt令牌的Secret
var Secret = "WanderingEarth"

// 局域网ip
var IpUrl = "http://127.0.0.1:3000"

// redis地址
var RedisUrl = "192.168.112.130"

// Feed流每次获得的视频数量
var VideoCount = 5

// FmtUser相关
var DefaultAvatar = IpUrl + "/static/images/cat_1.jpg"
var DefaultBGI = IpUrl + "/static/images/background.jpg"
var DefaultSign = "这个人很懒，什么都没有留下"
