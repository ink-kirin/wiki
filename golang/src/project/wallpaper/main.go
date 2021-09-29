package main

import (
	"fmt"
	"project/wallpaper/tool"
	"sync"
)

var UrlMap map[string]bool
var wg sync.WaitGroup

func main() {
	u := "https://wall.alphacoders.com/by_sub_category.php?id=333944&name=Genshin+Impact+Wallpapers&quickload=3154&page=101"
	err, s := tool.GetListInfoUrl(u)
	if err != nil {
		panic(err.Error())
	}
	UrlMap = make(map[string]bool)

	for _, v := range s {
		fmt.Println(v)
		wg.Add(1)
		go func(v string) {
			fmt.Println(v)
			_ = tool.GetImagesDownload(v, UrlMap)
			wg.Done()
		}(v)
	}
	wg.Wait()
	// tool.GetImagesDownload("https://wall.alphacoders.com/big.php?i=1111822")

}
