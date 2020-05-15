/**
  create by yy on 2019-08-23
*/

package main

import (
	"tpay_dock/app/ginServer"
	_ "tpay_dock/app/init"
	_ "tpay_dock/app/router"
)

func main() {
	ginServer.Run()
}
