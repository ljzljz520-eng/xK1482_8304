# 职业培训分镜工作台

这是一个纯 Go 的职业培训课件分镜后端示例。培训师提交课程目标后，服务会生成章节卡、讲师台词、操作演示素材、练习提示和预计时长；章节重排会同步刷新项目总时长。用户中心提供最近课程与导出记录。

项目使用固定的真实业务风格 fixture，不依赖数据库、网络、系统时钟、随机数或 CGO。所有编号、排序和导出结果均可重复。

## 环境

- Go 1.24.13
- `CGO_ENABLED=0`

## 运行

```bash
CGO_ENABLED=0 go run ./cmd/storyboard
```

可以传入课程目标和导出格式：

```bash
CGO_ENABLED=0 go run ./cmd/storyboard \
  -goal '仓储组长能够完成一次叉车班前点检并正确升级制动异常' \
  -format pdf
```

命令会向标准输出写入生成项目、导出记录和用户中心的 JSON。

## 测试

```bash
CGO_ENABLED=0 go test -count=1 ./...
```

验收套件包含本项目要求保留的可编辑明细回归场景。当前基线中 `TestEditableDetailRoundTrip` 会稳定失败，用于呈现调用方修改返回数据后再次查询受到污染的行为，其余业务链路用例应通过。

## 构建

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /tmp/storyboard-amd64 ./cmd/storyboard
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /tmp/storyboard-arm64 ./cmd/storyboard
```

核心代码位于 `internal/workbench`，可执行入口位于 `cmd/storyboard`。项目仅使用标准库。
