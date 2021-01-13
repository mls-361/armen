/*
------------------------------------------------------------------------------------------------------------------------
####### bus ####### (c) 2020-2021 mls-361 ########################################################## MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package bus

import (
	"fmt"
	"regexp"
	"sync"

	"github.com/mls-361/armen-sdk/message"
	"github.com/mls-361/minikit"
	"github.com/mls-361/uuid"

	"github.com/mls-361/armen/internal/components"
)

const (
	_maxChannelCapacity = 10
	_maxConsumer        = 3
)

type (
	// Bus AFAIRE.
	Bus struct {
		*minikit.Base
		components  *components.Components
		subscribers map[*regexp.Regexp]func(*message.Message)
		rwMutex     sync.RWMutex
		waitGroup   sync.WaitGroup
	}
)

// New AFAIRE.
func New(components *components.Components) *Bus {
	cb := &Bus{
		Base:        minikit.NewBase("bus", "bus"),
		components:  components,
		subscribers: make(map[*regexp.Regexp]func(*message.Message)),
	}

	components.CBus = cb

	return cb
}

// Dependencies AFAIRE.
func (cb *Bus) Dependencies() []string {
	return []string{
		"application",
		"logger",
		"metrics",
	}
}

func (cb *Bus) goConsumer(publisher string, ch <-chan *message.Message) {
	cb.waitGroup.Add(1)

	go func() { //@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		defer cb.waitGroup.Done()

		logger := cb.components.CLogger.CreateLogger(uuid.New(), publisher)

		logger.Info(">>>Bus") //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

		mcsCounter := cb.components.CMetrics.NewCounter("bus.publishers." + publisher)

		for msg := range ch {
			msg.Host = cb.components.CApplication.Host()
			msg.Publisher = publisher

			logger.Debug("Publish message", "id", msg.ID, "topic", msg.Topic) //::::::::::::::::::::::::::::::::::::::::

			cb.rwMutex.RLock()

			for re, callback := range cb.subscribers {
				if re.MatchString(msg.Topic) {
					callback(msg)
				}
			}

			cb.rwMutex.RUnlock()

			mcsCounter.Inc()
		}

		logger.Info("<<<Bus") //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	}()
}

// AddPublisher AFAIRE.
func (cb *Bus) AddPublisher(name string, chCapacity, nbConsumer int) chan<- *message.Message {
	if chCapacity < 0 {
		chCapacity = 0
	} else if chCapacity > _maxChannelCapacity {
		chCapacity = _maxChannelCapacity
	}

	ch := make(chan *message.Message, chCapacity)

	if nbConsumer < 1 {
		nbConsumer = 1
	} else if nbConsumer > _maxConsumer {
		nbConsumer = _maxConsumer
	}

	for i := 0; i < nbConsumer; i++ {
		cb.goConsumer(name, ch)
	}

	return ch
}

// Subscribe AFAIRE.
func (cb *Bus) Subscribe(callback func(*message.Message), regexpList ...string) error {
	cb.rwMutex.Lock()
	defer cb.rwMutex.Unlock()

	for _, re := range regexpList {
		regExp, err := regexp.Compile(fmt.Sprintf(`^%s$`, re))
		if err != nil {
			return err
		}

		cb.subscribers[regExp] = callback
	}

	return nil
}

// Close AFAIRE.
func (cb *Bus) Close() {
	cb.waitGroup.Wait()
}

/*
######################################################################################################## @(°_°)@ #######
*/
