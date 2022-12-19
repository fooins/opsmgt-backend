package tests

import "time"

// 测试
func Test(query map[string][]string) (map[string]any, error) {
	rst := map[string]any{
		"time":  time.Now(),
		"query": query,
	}
	return rst, nil
}
