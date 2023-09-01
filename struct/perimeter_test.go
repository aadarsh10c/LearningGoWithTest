package structs

import "testing"



func TestPremieter(t *testing.T) {
	rectangle:=Rectangle{10.0, 10.0}
	got:= Perimeter(rectangle)
	want:=40.0

	if got != want{
		t.Errorf("got %g, want %g", got, want)
	}

	
}

func TestArea(t *testing.T){
	
	//check area
	checkArea:=func (t *testing.T, shape Shape, want float64){
		t.Helper()
		got:= shape.Area()
		if got != want{
			t.Errorf("got %g, want %g", got, want)
		}
	}
	t.Run("rectangles",func(t *testing.T) {
		rectangle:=Rectangle{12.0,6.0}
		want:= 72.0
		checkArea(t,rectangle,want)
	})

	t.Run("circles",func(t *testing.T) {
		circle:=Circle{10.0}
		want:=314.1592653589793
		checkArea(t,circle,want)
	})
	
}
