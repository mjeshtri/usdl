// Package ui provides the TUI for the application.
package ui

import (
	"context"
	"fmt"

	"github.com/ardanlabs/usdl/foundation/agents/ollamallm"
	"github.com/ardanlabs/usdl/foundation/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// =============================================================================

type TUI struct {
	tviewApp *tview.Application
	flex     *tview.Flex
	list     *tview.List
	textView *tview.TextView
	textArea *tview.TextArea
	button   *tview.Button
	app      *client.App
	agent    *ollamallm.Agent
	history  *history
}

func New(myAccountID common.Address, agent *ollamallm.Agent) *TUI {
	ui := TUI{
		agent:   agent,
		history: NewHistory(5),
	}

	tApp := tview.NewApplication()

	// -------------------------------------------------------------------------

	textView := tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetWordWrap(true).
		SetChangedFunc(func() {
			tApp.Draw()
		})

	textView.SetBorder(true)
	textView.SetTitle(fmt.Sprintf("*** %s ***", myAccountID))

	// -------------------------------------------------------------------------

	list := tview.NewList()
	list.SetBorder(true)
	list.SetTitle("Users")
	list.SetChangedFunc(func(idx int, name string, id string, shortcut rune) {
		textView.Clear()

		if ui.app == nil {
			return
		}

		addrID := common.HexToAddress(id)

		user, err := ui.app.QueryContactByID(addrID)
		if err != nil {
			textView.ScrollToEnd()
			fmt.Fprintln(textView, "-----")
			fmt.Fprintln(textView, err.Error())
			return
		}

		for i, msg := range user.Messages {
			fmt.Fprintf(textView, "%s: %s\n", msg.Name, string(msg.Content))
			if i < len(user.Messages)-1 {
				fmt.Fprintln(textView, "-----")
			}
		}

		list.SetItemText(idx, user.Name, user.ID.Hex())
	})

	// -------------------------------------------------------------------------

	button := tview.NewButton("SUBMIT")
	button.SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorGreen).Bold(true))
	button.SetActivatedStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorGreen).Bold(true))
	button.SetBorder(true)
	button.SetBorderColor(tcell.ColorGreen)

	// -------------------------------------------------------------------------

	textArea := tview.NewTextArea()
	textArea.SetWrap(false)
	textArea.SetPlaceholder("Enter message here...")
	textArea.SetBorder(true)
	textArea.SetBorderPadding(0, 0, 1, 0)

	// -------------------------------------------------------------------------

	flex := tview.NewFlex().
		AddItem(list, 20, 1, false).
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(textView, 0, 5, false).
			AddItem(tview.NewFlex().
				SetDirection(tview.FlexColumn).
				AddItem(textArea, 0, 90, false).
				AddItem(button, 0, 10, false),
				0, 1, false),
			0, 1, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape, tcell.KeyCtrlQ:
			tApp.Stop()
			return nil
		}

		return event
	})

	ui.tviewApp = tApp
	ui.flex = flex
	ui.list = list
	ui.textView = textView
	ui.textArea = textArea
	ui.button = button

	button.SetSelectedFunc(ui.buttonHandler)

	textArea.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			ui.buttonHandler()
			return nil
		}
		return event
	})

	return &ui
}

func (ui *TUI) SetApp(app *client.App) {
	ui.app = app

	for i, user := range app.Contacts() {
		shortcut := rune(i + 49)
		ui.list.AddItem(user.Name, user.ID.Hex(), shortcut, nil)
	}
}

func (ui *TUI) Run() error {
	return ui.tviewApp.SetRoot(ui.flex, true).EnableMouse(true).Run()
}

func (ui *TUI) WriteText(msg client.Message) {
	ui.textView.ScrollToEnd()

	switch msg.ID {
	case common.Address{}:
		fmt.Fprintln(ui.textView, "-----")
		fmt.Fprintf(ui.textView, "%s: %s\n", msg.Name, string(msg.Content))

	case ui.app.ID():
		fmt.Fprintln(ui.textView, "-----")
		fmt.Fprintf(ui.textView, "%s: %s\n", msg.Name, string(msg.Content))

	default:
		idx := ui.list.GetCurrentItem()

		_, currentID := ui.list.GetItemText(idx)
		if currentID == "" {
			fmt.Fprintln(ui.textView, "-----")
			fmt.Fprintln(ui.textView, "id not found: "+msg.ID.Hex())
			return
		}

		msgContent := fmt.Sprintf("%s: %s", msg.Name, string(msg.Content))

		ui.history.add(msg.ID, msgContent)

		if msg.ID.Hex() == currentID {
			fmt.Fprintln(ui.textView, "-----")
			fmt.Fprintf(ui.textView, "%s\n", msgContent)

			if ui.agent != nil {
				ui.agentResponse(msg.ID)
			}

			return
		}

		for i := range ui.list.GetItemCount() {
			name, idStr := ui.list.GetItemText(i)
			if msg.ID.Hex() == idStr {
				ui.list.SetItemText(i, "* "+name, idStr)
				ui.tviewApp.Draw()

				if ui.agent != nil {
					ui.agentResponse(msg.ID)
				}

				return
			}
		}
	}
}

func (ui *TUI) UpdateContact(id common.Address, name string) {
	shortcut := rune(ui.list.GetItemCount() + 49)
	ui.list.AddItem(name, id.Hex(), shortcut, nil)
}

// =============================================================================

func (ui *TUI) agentResponse(id common.Address) {
	ctx := context.TODO()

	msgs := ui.history.retrieve(id)

	input := msgs[len(msgs)-1]
	history := msgs[:len(msgs)-1]

	resp, err := ui.agent.Chat(ctx, input, history)
	if err != nil {
		fmt.Fprintln(ui.textView, "-----")
		fmt.Fprintln(ui.textView, "failed ollama response: "+err.Error())
	}

	// TODO: Create some artificial delay to simulate thinking

	ui.textArea.SetText(resp, true)
	ui.buttonHandler()
}

func (ui *TUI) buttonHandler() {
	_, to := ui.list.GetItemText(ui.list.GetCurrentItem())

	msg := ui.textArea.GetText()
	if msg == "" {
		return
	}

	id := common.HexToAddress(to)

	if err := ui.app.SendMessageHandler(id, []byte(msg)); err != nil {
		msg := client.Message{
			Name:    "system",
			Content: fmt.Appendf(nil, "Error sending message: %s", err),
		}
		ui.WriteText(msg)
		return
	}

	ui.history.add(id, msg)

	ui.textArea.SetText("", false)
}
