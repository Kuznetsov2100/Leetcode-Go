/*
 * @lc app=leetcode.cn id=535 lang=golang
 *
 * [535] TinyURL 的加密与解密
 */

// @lc code=start
type Codec struct {
    m map[int]string
}


func Constructor() Codec {
    return Codec{m: make(map[int]string)}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	hash := hashcode(longUrl)
	this.m[hash] = longUrl
	return "http://tinyurl.com/" + strconv.Itoa(hash)
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	hash, _ := strconv.Atoi(shortUrl[19:])
	return this.m[hash]
}

func hashcode(s string) int {
	R, hash := 31, 0
	for _, runeValue := range s {
		hash = R*hash + int(runeValue)
	}
	return hash
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */

// @lc code=end

