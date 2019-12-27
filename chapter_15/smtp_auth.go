// smtp_auth.go
package main

import (
	"log"
	"net/smtp"
)

/**
 译者注：
按照这个示例代码，发送邮件是会被拒信的。看了下 SendMail 的源码，它要求 msg 参数要符合 RFC 822 电子邮件的标准格式。
所以示例中的 "This is the email body." 要修改一下才能发送成功。修改后的示例如下：
"To: recipient@example.net\r\nFrom: sender@example.org\r\nSubject: 邮件主题\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\nHello World"
当然，也要把里面的寄件人、收件人等信息按照实际去修改一下。
 * @Author Liumm
 * @Date   2019-12-27
 * @return {[type]}   [description]
*/
func main() {
	// 设置认证信息。
	auth := smtp.PlainAuth(
		"",
		"user@example.com",
		"password",
		"mail.example.com",
	)
	// 连接到服务器, 认证, 设置发件人、收件人、发送的内容,
	// 然后发送邮件。
	err := smtp.SendMail(
		"mail.example.com:25",
		auth,
		"sender@example.org",
		[]string{"recipient@example.net"},
		[]byte("This is the email body."),
	)
	if err != nil {
		log.Fatal(err)
	}
}
