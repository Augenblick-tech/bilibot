package basetask

type BaseTask struct {
	Status
}

type Status int

//go:generate stringer -type=Status --linecomment
const (
	Running Status = iota // 运行中
	Warning				  // 发生错误
)
