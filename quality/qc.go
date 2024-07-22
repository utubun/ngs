package quality

import "github.com/utubun/ngs/quality/internal/core"

//type QC interface{}

func NewQC() *core.QC {
	return &core.QC{Bases: make(map[string][]int), Seq: make([]core.Sequence, 1, 1000000)}
}
