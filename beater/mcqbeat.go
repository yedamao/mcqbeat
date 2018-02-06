package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/yedamao/gomemcacheq"
	"github.com/yedamao/mcqbeat/config"
)

type Mcqbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

type Stat struct {
	In   int
	Out  int
	Stay int
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Mcqbeat{
		done:   make(chan struct{}),
		config: config,
	}
	return bt, nil
}

func (bt *Mcqbeat) Run(b *beat.Beat) error {
	logp.Info("mcqbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)

	mcq, err := memcacheq.New(bt.config.Host)
	if err != nil {
		return err
	}

	if err := mcq.Dial(); err != nil {
		return err
	}
	defer mcq.Close()

	// the previous period queues statuses
	pre_stats := make(map[string]memcacheq.Stat)

	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		// current period queues statues
		stats, err := mcq.StatsQueue()
		if err != nil {
			return err
		}

		// to be publishing map
		queues := common.MapStr{}
		for _, queue := range *stats {
			pre_stat, ok := pre_stats[queue.QueueName]
			pre_stats[queue.QueueName] = queue
			if !ok {
				queues.Put(queue.QueueName, Stat{})
			} else {
				queues.Put(queue.QueueName, Stat{
					In:   queue.AllIn - pre_stat.AllIn,
					Out:  queue.AllOut - pre_stat.AllOut,
					Stay: queue.AllIn - queue.AllOut,
				})
			}
		}

		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"type":  b.Info.Name,
				"stats": queues,
			},
		}

		bt.client.Publish(event)
		logp.Info("Event sent")
	}
}

func (bt *Mcqbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
