# go-bd09mc

百度经纬度坐标（bd09ll）与百度墨卡托米制坐标（bd09mc）互转

# 安装

```bash
go get -u github.com/spencer404/go-bd09mc
```

# 使用

```go
package main

import (
	"fmt"
	"github.com/spencer404/go-bd09mc"
)

func main() {
	var lng, lat float64
	var err error

	lng, lat, err = bd09mc.LL2MC(108.95344, 34.265657)
	fmt.Println(lng, lat, err)
	// output: 1.212877343e+07 4.04024901e+06 <nil>

	lng, lat, err = bd09mc.MC2LL(12128773.43, 4040249.00)
	fmt.Println(lng, lat, err)
	// output: 108.95344 34.265657 <nil>
}
```
