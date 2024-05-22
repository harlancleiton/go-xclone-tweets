package events

import (
	"errors"
	"sync"
)

var (
	ErrHandlerAlreadyRegistered = errors.New("handler already registered")
	ErrHandlerNotRegistered     = errors.New("handler not registered")
	ErrHandlerNotFound          = errors.New("handler not found")
)

const GlobalHandler = "*"

type EventDispatcher interface {
	Dispatch(event Event) error
	RegisterHandler(eventName string, handler EventHandler) error
	UnregisterHandler(eventName string, handler EventHandler) error
	Has(eventName string, handler EventHandler) bool
	Clear()
}

type ConcreteEventDispatcher struct {
	handlers map[string][]EventHandler
}

func NewConcreteEventDispatcher() *ConcreteEventDispatcher {
	return &ConcreteEventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (d *ConcreteEventDispatcher) RegisterHandler(eventName string, handler EventHandler) error {
	if _, ok := d.handlers[eventName]; ok {
		for _, h := range d.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	d.handlers[eventName] = append(d.handlers[eventName], handler)
	return nil
}

func (d *ConcreteEventDispatcher) UnregisterHandler(eventName string, handler EventHandler) error {
	if _, ok := d.handlers[eventName]; !ok {
		return ErrHandlerNotRegistered
	}

	for i, h := range d.handlers[eventName] {
		if h == handler {
			d.handlers[eventName] = append(d.handlers[eventName][:i], d.handlers[eventName][i+1:]...)
			return nil
		}
	}

	return ErrHandlerNotRegistered
}

func (d *ConcreteEventDispatcher) Dispatch(event Event) error {
	_, containsGlobalHandlers := d.handlers[GlobalHandler]
	_, containsNamedHandlers := d.handlers[event.Name()]
	if !containsNamedHandlers && !containsGlobalHandlers {
		return ErrHandlerNotFound
	}

	wg := &sync.WaitGroup{}
	for _, handler := range d.handlers[GlobalHandler] {
		wg.Add(1)
		go handler.Handle(event, wg)
	}
	for _, handler := range d.handlers[event.Name()] {
		wg.Add(1)
		go handler.Handle(event, wg)
	}
	wg.Wait()
	return nil
}

func (d *ConcreteEventDispatcher) Has(eventName string, handler EventHandler) bool {
	if _, ok := d.handlers[eventName]; !ok {
		return false
	}

	for _, h := range d.handlers[eventName] {
		if h == handler {
			return true
		}
	}
	return false
}

func (d *ConcreteEventDispatcher) Clear() {
	d.handlers = make(map[string][]EventHandler)
}
