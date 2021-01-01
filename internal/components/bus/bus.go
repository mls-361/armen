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
	"github.com/mls-361/uuid"

	"github.com/mls-361/armen/internal/components"
)

const (
	_maxChannelCapacity = 10
	_maxConsumer        = 3
)

type (
	bus struct {
		components  *components.Components
		subscribers map[*regexp.Regexp]func(*message.Message)
		rwMutex     sync.RWMutex
		waitGroup   sync.WaitGroup
	}
)

func newBus(components *components.Components) *bus {
	return &bus{
		components:  components,
		subscribers: make(map[*regexp.Regexp]func(*message.Message)),
	}
}

func (cb *bus) goConsumer(publisher string, ch <-chan *message.Message) {
	cb.waitGroup.Add(1)

	go func() { //@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
		logger := cb.components.Logger.CreateLogger(uuid.New(), publisher)

		logger.Info(">>>Bus") //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

		for msg := range ch {
			msg.Host = cb.components.Application.Host()
			msg.Publisher = publisher

			logger.Debug("Publish message", "id", msg.ID, "topic", msg.Topic) //::::::::::::::::::::::::::::::::::::::::

			cb.rwMutex.RLock()

			for re, cb := range cb.subscribers {
				if re.MatchString(msg.Topic) {
					cb(msg)
				}
			}

			cb.rwMutex.RUnlock()
		}

		logger.Info("<<<Bus") //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

		logger.RemoveLogger("")

		cb.waitGroup.Done()
	}()
}

// AddPublisher AFAIRE.
func (cb *bus) AddPublisher(name string, chCapacity, nbConsumer int) chan<- *message.Message {
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
func (cb *bus) Subscribe(callback func(*message.Message), regexpList ...string) error {
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

func (cb *bus) Close() { cb.waitGroup.Wait() }

/*
######################################################################################################## @(°_°)@ #######
*/
