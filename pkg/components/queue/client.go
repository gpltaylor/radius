/*
Copyright 2023 The Radius Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package queue

import (
	"context"
	"errors"
	"time"

	"github.com/radius-project/radius/pkg/ucp/ucplog"
)

var (
	// ErrDequeuedMessage represents the error when message has already been dequeued.
	ErrDequeuedMessage = errors.New("message was dequeued by the other client")

	// ErrMessageNotFound represents the error when queue is empty or all messages are leased by clients.
	ErrMessageNotFound = errors.New("queue is empty or messages are leased")

	// ErrInvalidMessage represents the error when the message has already been requeued.
	ErrInvalidMessage = errors.New("this message has been requeued or deleted")

	// ErrUnsupportedContentType represents the error when the content type is unsupported.
	ErrUnsupportedContentType = errors.New("this message content type is unsupported")

	// ErrEmptyMessage represents nil or empty Message.
	ErrEmptyMessage = errors.New("message must not be nil or message is empty")
)

//go:generate mockgen -typed -destination=./mock_client.go -package=queue -self_package github.com/radius-project/radius/pkg/components/queue github.com/radius-project/radius/pkg/components/queue Client

// Client is an interface to implement queue operations.
type Client interface {
	// Enqueue enqueues message to queue.
	Enqueue(ctx context.Context, msg *Message, opts ...EnqueueOptions) error

	// Dequeue dequeues message from queue.
	Dequeue(ctx context.Context, cfg QueueClientConfig) (*Message, error)

	// FinishMessage finishes or deletes the message in the queue.
	FinishMessage(ctx context.Context, msg *Message) error

	// ExtendMessage extends the message lock.
	ExtendMessage(ctx context.Context, msg *Message) error
}

// StartDequeuer starts a dequeuer to consume the message from the queue and return the output channel.
func StartDequeuer(ctx context.Context, cli Client, opts ...DequeueOptions) (<-chan *Message, error) {
	log := ucplog.FromContextOrDiscard(ctx)
	out := make(chan *Message, 1)

	go func() {
		for {
			var queueconfig QueueClientConfig
			if opts != nil {
				queueconfig = NewDequeueConfig(opts...)
			}
			msg, err := cli.Dequeue(ctx, queueconfig)
			if err == nil {
				out <- msg
			} else if !errors.Is(err, ErrMessageNotFound) {
				log.Error(err, "fails to dequeue the message")
			}

			select {
			case <-ctx.Done():
				close(out)
				return
			case <-time.After(queueconfig.DequeueIntervalDuration):
			}

		}
	}()

	return out, nil
}
