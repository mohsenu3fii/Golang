package main

import "fmt"

func (o Order) allowedTransition(state OrderState) bool {

	for _, value := range valid_transition[o.State] {
		if state == value {
			return true
		}
	}
	return false
}

func (o *Order) ChangeState(newState OrderState) error {

	if ok := o == nil; ok {
		return fmt.Errorf("order is empty")
	}
	if o.State == newState {
		return fmt.Errorf("Operation failed, cannot change state to itself")
	}

	// - Only valid transitions allowed.
	// - Use switch.
	switch o.State {

	case Created, Paid, VendorAccepted, Shipped:
		if ok := o.allowedTransition(newState); ok {
			o.State = newState
			return nil
		}
		return fmt.Errorf("Transition not allowed")

		// - Cancelled is terminal.
	case Delivered, Cancelled:

		// - Return descriptive errors.
		return fmt.Errorf("Operation failed, status is invalid to change")

	default:
		return fmt.Errorf("State is not valid")
	}

}
