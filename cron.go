package main

import (
	"gin-blog/models"
	"github.com/robfig/cron"
	"log"
	"time"
)

func main() {
	log.Println("Starting...")

	c := cron.New()
	_ = c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		_ = models.CleanAllTag()
	})
	_ = c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		_ = models.CleanAllArticle()
	})

	c.Start()

	t1 := time.NewTimer(time.Second*10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second*10)
		}
	}
}
