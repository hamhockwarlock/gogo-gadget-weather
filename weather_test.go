package main

import (
	"fmt"
	"testing"
)

const(
  valueThatCannotBeCoerced = "hi!"
  valueOutsideRange = "1000"

)

func TestLatitudeThatIsUncoercable(t *testing.T) {
  expectedError := fmt.Sprintf("latitude %s", mustBeNumericMessage)

  if err := validateLatitude(valueThatCannotBeCoerced); err.Error() !=  expectedError {
    t.Errorf("returned the wrong message: got %s want %s", err.Error(), expectedError)
  }
}

func TestLongitudeThatIsUncoercable(t *testing.T) {
  expectedError := fmt.Sprintf("longitude %s", mustBeNumericMessage)

  if err := validateLongitude(valueThatCannotBeCoerced); err.Error() !=  expectedError {
    t.Errorf("returned the wrong message: got %s want %s", err.Error(), expectedError)
  }
}

func TestLatitudeThatIsEmpty(t *testing.T) {
  expectedError := fmt.Sprintf("latitude %s", isRequiredMessage)

  if err := validateLatitude(""); err.Error() !=  expectedError {
    t.Errorf("returned the wrong message: got %s want %s", err.Error(), expectedError)
  }
}

func TestLongitudeThatIsEmpty(t *testing.T) {
  expectedError := fmt.Sprintf("longitude %s", isRequiredMessage)

  if err := validateLongitude(""); err.Error() !=  expectedError {
    t.Errorf("returned the wrong message: got %s want %s", err.Error(), expectedError)
  }
}

func TestLongitudeOutsideRange(t *testing.T) {
  if err := validateLongitude(valueOutsideRange); err.Error() != longitudeRangeMessage {
    t.Errorf("returned the wrong message: got %s want %s", err.Error(), longitudeRangeMessage)
  }
}

func TestLatitudeOutsideRange(t *testing.T) {
  if err := validateLatitude(valueOutsideRange); err.Error() != latitudeRangeMessage {
    t.Errorf("returned the wrong message: got %s want %s", err.Error(), latitudeRangeMessage)
  }
}

func TestLongitudeWithinRange(t * testing.T) { 
  if err := validateLongitude("179"); err != nil {
    t.Errorf("returned the wrong message: got %s want nil", err.Error())
  }
}

func TestLatitudeWithinRange(t * testing.T) { 
  if err := validateLatitude("89"); err != nil {
    t.Errorf("returned the wrong message: got %s want nil", err.Error())
  }
}
