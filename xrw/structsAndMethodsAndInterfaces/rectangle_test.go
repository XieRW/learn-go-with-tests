package structsAndMethodsAndInterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	//checkArea := func(t *testing.T,shape Shape,want float64) {
	//	t.Helper()
	//	got:=shape.Area()
	//	if got != want {
	//		t.Errorf("got %.2f want %.2f",got,want)
	//	}
	//}
	//
	//t.Run("rectangles", func(t *testing.T) {
	//	rectangle := Rectangle{12.0,6.0}
	//	checkArea(t,rectangle,72.0)
	//})
	//
	//t.Run("circles", func(t *testing.T) {
	//	circle := Circle{10}
	//	checkArea(t,circle,314.1592653589793)
	//})

	areaTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{width: 12, height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{width: 12, height: 6}, want: 36.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %.2f want %.2f. %#v", tt.shape, got, tt.want, tt.shape)
			}
		})
	}
}
