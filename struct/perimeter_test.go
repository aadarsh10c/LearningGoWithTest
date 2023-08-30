package structs

import "testing"



func TestPremieter(t *testing.T) {
	rectangle:=Rectangle{10.0, 10.0}
	got:= Perimeter(rectangle)
	want:=40.0

	assertEqual(got, want, t)

	
}

func TestArea(t *testing.T){
	rectangle:=Rectangle{12.0,6.0}
	got:= Area(rectangle)
	want:= 72.0

	assertEqual(got, want, t)
}

func assertEqual(got float64, want float64, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}