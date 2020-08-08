/**
 * @Time: 2020/8/8 17:32
 * @Author: solacowa@gmail.com
 * @File: migrate_all
 * @Software: GoLand
 */

package migrate

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	migrateRedisAllCmd = &cobra.Command{
		Use:               `all <args> [flags]`,
		Short:             "迁移所有",
		SilenceErrors:     false,
		DisableAutoGenTag: false,
		Example: `
redis-tool migrate all {key} --source-hosts 127.0.0.1:6379 --source-auth 123456 --target-redis-cluster true --target-hosts 127.0.0.1:6379,127.0.0.1:7379 --target-auth 123456
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// 关闭资源连接
			defer func() {
				log.Printf("source redis close err: %v", sourceRedis.Close())
				log.Printf("target redis close err: %v", targetRedis.Close())
			}()
			var key string
			if len(args) == 1 {
				key = args[0]
			}
			return migrateRedisAll(key)
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err = prepare(); err != nil {
				fmt.Println(fmt.Sprintf("prepare error: %s", err.Error()))
				return err
			}
			return nil
		},
	}
)

func migrateRedisAll(key string) error {
	begin := time.Now()

	keys, err := sourceRedis.Keys(key)
	if err != nil {
		err = errors.Wrap(err, "sourceRedis.Keys")
		return err
	}
	var total = len(keys)

	fmt.Println(fmt.Sprintf("Key: [%s] 总数: [%d]", key, total))
	var base float64 = 50000

	step := math.Ceil(float64(total) / base)

	fmt.Println(step)

	for _, v := range keys {
		// 判断类型
		// 只支持 set, hash, list, sortedSet 其他的跳过

		fmt.Println(v)
	}

	fmt.Println(fmt.Sprintf("迁移完成, 用时 [%v]", time.Since(begin)))

	return nil
}