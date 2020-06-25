package reflection_test

import (
	"fmt"
	"testing"

	"wesionary.team/dipeshdulal/golang-test/reflection"
)

type Sample struct {
	Name string
}

func TestGetDataType(t *testing.T) {
	reflection.GetDataType(20)
	reflection.GetDataType(Sample{"Dipesh"})
	reflection.GetDataType(&Sample{"Dipesh"})
}

func TestField(t *testing.T) {
	reflection.GetFields(Sample{"Dipesh"})
}

func TestValueChange(t *testing.T) {
	dipesh := Sample{"Dipesh"}

	fmt.Println("before change: ", dipesh)

	reflection.ChangeValueWithReflection(&dipesh)

	fmt.Println("after change: ", dipesh)
}
