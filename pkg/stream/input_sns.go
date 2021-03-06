package stream

import (
	"github.com/applike/gosoline/pkg/cfg"
	"github.com/applike/gosoline/pkg/cloud"
	"github.com/applike/gosoline/pkg/mon"
	"github.com/applike/gosoline/pkg/sns"
	"github.com/applike/gosoline/pkg/sqs"
)

type SnsInputSettings struct {
	cfg.AppId
	QueueId           string                `cfg:"queue_id"`
	WaitTime          int64                 `cfg:"wait_time"`
	RedrivePolicy     sqs.RedrivePolicy     `cfg:"redrive_policy"`
	VisibilityTimeout int                   `cfg:"visibility_timeout"`
	RunnerCount       int                   `cfg:"runner_count"`
	Client            cloud.ClientSettings  `cfg:"client"`
	Backoff           cloud.BackoffSettings `cfg:"backoff"`
}

type SnsInputTarget struct {
	cfg.AppId
	TopicId string
}

type snsInput struct {
	*sqsInput
}

func NewSnsInput(config cfg.Config, logger mon.Logger, s SnsInputSettings, targets []SnsInputTarget) *snsInput {
	s.PadFromConfig(config)
	autoSubscribe := config.GetBool("aws_sns_autoSubscribe")

	sqsInput := NewSqsInput(config, logger, SqsInputSettings{
		AppId:             s.AppId,
		QueueId:           s.QueueId,
		WaitTime:          s.WaitTime,
		VisibilityTimeout: s.VisibilityTimeout,
		RunnerCount:       s.RunnerCount,
		RedrivePolicy:     s.RedrivePolicy,
		Client:            s.Client,
		Backoff:           s.Backoff,
		Unmarshaller:      UnmarshallerSns,
	})

	queueArn := sqsInput.GetQueueArn()

	if autoSubscribe {
		for _, t := range targets {
			t.PadFromConfig(config)

			topic := sns.NewTopic(config, logger, &sns.Settings{
				AppId:   t.AppId,
				TopicId: t.TopicId,
				Client:  s.Client,
				Backoff: s.Backoff,
			})

			err := topic.SubscribeSqs(queueArn)

			if err != nil {
				panic(err)
			}
		}
	}

	return NewSnsInputWithInterfaces(sqsInput)
}

func NewSnsInputWithInterfaces(sqsInput *sqsInput) *snsInput {
	return &snsInput{
		sqsInput,
	}
}
