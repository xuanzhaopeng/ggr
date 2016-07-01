package main

import (
	"testing"
	"time"

	. "github.com/aandryashin/matchers"
	"github.com/fsnotify/fsnotify"
)

func TestSingleTimer(t *testing.T) {
	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()
	call := false
	watch(watcher, 150*time.Millisecond, func() {
		call = true
	})
	watcher.Events <- fsnotify.Event{Op: fsnotify.Create}
	<-time.After(100 * time.Millisecond)
	AssertThat(t, call, Is{false})
	<-time.After(100 * time.Millisecond)
	AssertThat(t, call, Is{true})
}

func TestMultipleTimer(t *testing.T) {
	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()

	call := false
	watch(watcher, 150*time.Millisecond, func() {
		call = true
	})
	watcher.Events <- fsnotify.Event{Op: fsnotify.Create}

	<-time.After(100 * time.Millisecond)
	watcher.Events <- fsnotify.Event{Op: fsnotify.Create}

	<-time.After(100 * time.Millisecond)
	AssertThat(t, call, Is{false})

	<-time.After(100 * time.Millisecond)
	AssertThat(t, call, Is{true})
}

func TestTimerCalledOnce(t *testing.T) {
	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()

	call := 0
	watch(watcher, 150*time.Millisecond, func() {
		call++
	})
	watcher.Events <- fsnotify.Event{Op: fsnotify.Create}

	<-time.After(100 * time.Millisecond)
	watcher.Events <- fsnotify.Event{Op: fsnotify.Create}

	<-time.After(100 * time.Millisecond)
	watcher.Events <- fsnotify.Event{Op: fsnotify.Create}

	<-time.After(100 * time.Millisecond)
	watcher.Events <- fsnotify.Event{Op: fsnotify.Create}

	<-time.After(100 * time.Millisecond)
	AssertThat(t, call, EqualTo{0})

	<-time.After(100 * time.Millisecond)
	AssertThat(t, call, EqualTo{1})
}
