package task

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/redis"
)

func StartTask(db mysql.Repo, cache redis.Repo) {

	c := cron.New(cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)))

	_, err := c.AddFunc("*/1 * * * * *", func() {

		fmt.Print(`定时任务1`)

	})

	if err != nil {
		return
	}

	_, err2 := c.AddFunc("*/1 * * * * *", func() {

		fmt.Print(`定时任务2`)

	})

	if err2 != nil {
		return
	}

	_, err3 := c.AddFunc("*/1 * * * * *", func() {

		fmt.Print(`定时任务3`)

	})

	if err3 != nil {
		return
	}

	_, err4 := c.AddFunc("*/1 * * * * *", func() {

		fmt.Print(`定时任务4`)

	})

	if err4 != nil {
		return
	}

	c.Start()

}
