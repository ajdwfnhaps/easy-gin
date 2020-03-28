# easy-gin
gin插件封装,简单易用,方便在项目中集成应用

### 使用介绍
```

package sample

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"

	easygin "github.com/ajdwfnhaps/easy-gin"
	"github.com/ajdwfnhaps/easy-gin/conf"
)

func TestEasyGin(t *testing.T) {
	if err := easygin.UseEasyGin(func(opt *conf.Config) {
		opt.HTTP.Port = 8888
		opt.HTTP.ShutdownTimeout = 60
	}); err != nil {
		t.Error(err)
	}

	handleSignal()
}

func handleSignal() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case <-c:
		fmt.Println("服务退出")
	}
}


```

详细可参考[测试用例](sample/ginplus_test.go)

### todo列表：
. response封装
. 日志中间件完善(easy-logrus)
. auth中间件开发
. 集成swagger
. 提供统一注册路由方法等。。。