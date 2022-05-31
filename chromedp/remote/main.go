/**
 * @Author: sonic
 * @File:  main.go
 * @Date: 2022/5/30
 * @Description:
 * chromedp 的远程调用函数
 * - 远程的chrome服务器使用docker代替，使用的镜像是chrome官方推荐的：https://hub.docker.com/r/chromedp/headless-shell/。
 *   注意：
 *    	- headless-shell使用的是websockt协议，所以远程调用地址是ws开头，eg：ws://127.0.0.1:9222。
 * 		- 关于headless-shell的端口或启动配置，可以参考上面的docker hub地址有说明。
 *		- 一般使用这个命令启动满需求：docker run -d -p 9222:9222 --rm --name headless-shell chromedp/headless-shell
 * - chromedp 的远程调用函数是 NewRemoteAllocator(parent context.Context, url string) (context.Context, context.CancelFunc)
 * - 经测试，可以仿真浏览器获取网页的动态html
 */
package main

import (
	"context"
	"flag"
	"github.com/chromedp/chromedp"
	"log"
)

func main() {
	devtoolsWsURL := flag.String("devtools-ws-url", "ws://127.0.0.1:9222", "")
	flag.Parse()
	if *devtoolsWsURL == "" {
		log.Fatal("must specify -devtools-ws-url")
	}

	// create allocator context for use with creating a browser context later
	// 这里其实可以直接使用string的url代替构造devtoolsWsURL这个较为负责的变量类型
	allocatorContext, cancel := chromedp.NewRemoteAllocator(context.Background(), *devtoolsWsURL)
	defer cancel()

	// create context
	ctxt, cancel := chromedp.NewContext(allocatorContext)
	defer cancel()

	// run task list
	var body string
	if err := chromedp.Run(ctxt,
		chromedp.Navigate("https://bbs.ipapazao.com/2048/thread.php?fid-3.html"),
		//chromedp.WaitVisible("#logo_homepage_link"),
		chromedp.OuterHTML("html", &body),
	); err != nil {
		log.Fatalf("Failed getting body of duckduckgo.com: %v", err)
	}

	log.Println("Body of duckduckgo.com starts with:")
	log.Println(body)
}
