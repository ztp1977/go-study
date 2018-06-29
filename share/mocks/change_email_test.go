package models

import (
	"testing"
)

func TestChangeEmail(t *testing.T) {
	testScenarios := []struct {
		givenAddress, givenNewAddress string
		expectedErrorMsg              string
	}{
		{
			givenAddress:    "Foo@x.com",
			givenNewAddress: "bar@x.com",
		},
		{
			givenAddress:     "foo@x.com",
			givenNewAddress:  "wrong",
			expectedErrorMsg: "Please provide a valid email address",
		},
	}

	for _, scenario := range testScenarios {
		// 1. Given
		user := CreateUser(scenario.givenNewAddress)

		// 2. Do this
		sentmail, errmsg := handleNewAddressRequest(scenario.givenNewAddress)

		// 3.1 Expects error message
		if want, got := scenario.expectedErrorMsg, errmsg; want != got {
			t.Errorf("expect error %#v, but got %#v", want, got)
			return
		} else if errmsg != "" {
			return // expected error; done here
		}

		// 3.2 Expect send to correct address
		if want, got := scenario.givenNewAddress, addressFromMail(sentmail); want != got {
			t.Errorf("expect email sent to %#v, but got sent to %#v", want, got)
			return
		}

		// 3.3 Experts DB value NOT changed
		if want, got := scenario.givenAddress, addressFromDB(user); want != got {
			t.Errorf("expect user.Address to be unchanged %#v, but changed to %#v", want, got)
			return
		}

		// 3.4 Expects a link in email for 'user' to click
		if err := visitLinkFromEmailTo(sentmail); err != nil {
			t.Errorf(err.Error())
			return
		}

		// 3.5 Expects DB value TO changed
		if want, got := scenario.givenNewAddress, addressFromDB(user); want != got {
			t.Errorf("expect user.Address to become %#v, but got %#v", want, got)
			return
		}
	}
}
