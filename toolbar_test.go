package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestApp_getToolBar(t *testing.T) {
	tb := testApp.getToolBar()

	if len(tb.Items) != 4 {
		t.Error("wrong number of items in toolbar")
	}
}

func TestApp_addHoldingsDialog(t *testing.T) {
	testApp.addHoldingDialog()
	test.Type(testApp.AddHoldingsPurChaseAmountEntry, "1")
	test.Type(testApp.AddHoldingsPurChaseDateEntry, "2020-01-01")
	test.Type(testApp.AddHoldingsPurChasePriceEntry, "1000")

	if testApp.AddHoldingsPurChaseDateEntry.Text != "2020-01-01" {
		t.Error("date not correct")
	}

	if testApp.AddHoldingsPurChaseAmountEntry.Text != "1" {
		t.Error("amount not correct")
	}

	if testApp.AddHoldingsPurChasePriceEntry.Text != "1000" {
		t.Error("price not correct")
	}

}
