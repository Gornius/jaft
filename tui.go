package main

import (
	"time"

	"charm.land/bubbles/v2/timer"
	tea "charm.land/bubbletea/v2"
)

type teaModel struct {
	aborted bool
	timer   timer.Model
}

func (m teaModel) Init() tea.Cmd {
	return m.timer.Init()
}

func (m teaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd

	case timer.StartStopMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd

	case timer.TimeoutMsg:
		return m, tea.Quit

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.aborted = true
			return m, tea.Quit
		}

	}

	return m, nil
}

func (m teaModel) View() tea.View {
	return tea.NewView("Time left: " + m.timer.View())
}

// function displayTimerAndWait waits for a certain amount of time displaying a timer
// and then returns true if it successfully ends, otherwise returns false (e.g. when user aborted timer)
func displayTimerAndWait(duration time.Duration) (bool, error) {
	m := teaModel{
		timer: timer.New(duration, timer.WithInterval(100*time.Millisecond)),
	}
	um, err := tea.NewProgram(m).Run()
	if err != nil {
		return false, err
	}

	return !um.(teaModel).aborted, nil
}
