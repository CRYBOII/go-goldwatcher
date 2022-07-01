package main

import "testing"

func TestConfig_getHoldings(t *testing.T) {
	all, err := testApp.currentHoldings()
	if err != nil {
		t.Error("failed to get holdings", err)
	}

	if len(all) != 2 {
		t.Error("wrong number of rows", err)
	}

}

func TestConfig_getHoldingsSlice(t *testing.T) {
	slice := testApp.getHoldingSlice()

	if len(slice) != 3 {
		t.Error("wrong number of slice")
	}

}
