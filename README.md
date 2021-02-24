# logig

こんな感じで使う

```go
package main

import (
	"context"
	"fmt"

	"github.com/chiguhagu/logig"
	"github.com/chiguhagu/logig/logig_logrus"
	"github.com/chiguhagu/logig/logig_zap"
)

func main() {
	loglevel := "info"

	ctx := context.Background()

	fmt.Println("-default-")

	// 何もない場合は fmt を使用するように初期化される
	logig.Info(ctx, "sample!!")

	fmt.Println("-zap-")

	// Zap で初期化
	ctx = logig.ToContext(ctx, logig_zap.NewLogigZap(loglevel))
	logig.AddField(ctx, map[string]interface{}{"dd.trace_id": "123"})
	logig.AddField(ctx, map[string]interface{}{"dd.span_id": "456"})
	logig.Debug(ctx, "message") // 見えない
	logig.Info(ctx, "message")
	logig.Warn(ctx, "message")
	logig.Error(ctx, "message")
	//logig.Fatal(ctx, "fatal")

	fmt.Println("-logrus-")

	// Logrus で初期化
	ctx = logig.ToContext(ctx, logig_logrus.NewLogigLogrus(loglevel))
	logig.AddField(ctx, map[string]interface{}{"dd.trace_id": "123"})
	logig.AddField(ctx, map[string]interface{}{"dd.span_id": "456"})
	logig.Debug(ctx, "message") // 見えない
	logig.Info(ctx, "message")
	logig.Warn(ctx, "message")
	logig.Error(ctx, "message")
	//logig.Fatal(ctx, "message")

	fmt.Println("-end-")
}
```
