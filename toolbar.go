package main

import (
	"goldwatcher/repository"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) getToolBar() *widget.Toolbar {
	toolBar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			app.addHoldingDialog()
		}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			app.refreshPriceContent()
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {

		}),
	)

	return toolBar
}

func (app *Config) addHoldingDialog() dialog.Dialog {
	addAmountEntry := widget.NewEntry()
	purChaseDateEntry := widget.NewEntry()
	purchasePriceEntry := widget.NewEntry()
	app.AddHoldingsPurChaseAmountEntry = addAmountEntry
	app.AddHoldingsPurChaseDateEntry = purChaseDateEntry
	app.AddHoldingsPurChasePriceEntry = purchasePriceEntry

	dateValidator := func(text string) error {
		if _, err := time.Parse("2006-01-02", text); err != nil {
			return err
		}

		return nil
	}
	purChaseDateEntry.Validator = dateValidator

	isIntValidator := func(text string) error {
		_, err := strconv.Atoi(text)
		if err != nil {
			return err
		}
		return nil
	}

	addAmountEntry.Validator = isIntValidator

	isFloatValidator := func(text string) error {
		_, err := strconv.ParseFloat(text, 32)
		if err != nil {
			return err
		}
		return nil
	}

	purchasePriceEntry.Validator = isFloatValidator

	purChaseDateEntry.PlaceHolder = "YYYY-MM-DD"

	addForm := dialog.NewForm(
		"Add gold",
		"Add",
		"Cancel",
		[]*widget.FormItem{
			{Text: "Amount om toz", Widget: addAmountEntry},
			{Text: "PurChase price", Widget: purchasePriceEntry},
			{Text: "PurChase date", Widget: purChaseDateEntry},
		},
		func(valid bool) {
			if valid {
				amount, _ := strconv.Atoi(addAmountEntry.Text)
				purchasePrice, _ := strconv.ParseFloat(purchasePriceEntry.Text, 32)
				purchaseDate, _ := time.Parse("2006-01-02", purChaseDateEntry.Text)
				purchasePrice = purchasePrice * 100.0

				_, err := app.DB.InsertHolding(repository.Holdings{
					Amount:        amount,
					PurchaseDate:  purchaseDate,
					PurchasePrice: int(purchasePrice),
				})

				if err != nil {
					app.Errorlog.Println(err)
					return
				}

				app.refreshHoldingsTable()

			}
		},
		app.MainWindow)

	addForm.Resize(fyne.Size{Width: 400})
	addForm.Show()

	return addForm
}
