package shapedescription

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShapePerimeterCircleSimple(t *testing.T) {
	r := 9.0
	c := Circle{Radius: r}
	got, _ := Shape(c).Perimeter()
	want := 2 * math.Pi * r
	if got != want {
		t.Errorf("Shape(Circle{Radius: %.2f}).Perimeter= %.5f; want %.5f", r, got, want)
	}
}

func TestShapeAreaCircleSimple(t *testing.T) {
	r := 9.0
	c := Circle{Radius: r}
	got, _ := Shape(c).Area()
	want := math.Pi * math.Pow(r, 2)
	if got != want {
		t.Errorf("Shape(Circle{Radius: %.2f}).Area= %.5f; want %.5f", r, got, want)
	}
}

func TestShapePerimeterRectangle(t *testing.T) {
	h := 9.0
	w := 3.0
	c := Rectangle{
		Height: h,
		Width:  w,
	}
	got, _ := Shape(c).Perimeter()
	want := 2 * (h + w)
	if got != want {
		t.Errorf("Perimeter= %.5f; want %.5f", got, want)
	}
}

func TestShapePerimeterRectangleSimple(t *testing.T) {
	h := 9.0
	w := 3.0
	c := Rectangle{
		Height: h,
		Width:  w,
	}
	got, _ := Shape(c).Perimeter()
	want := 24.0
	if got != want {
		t.Errorf("Perimeter= %.5f; want %.5f", got, want)
	}
}

func TestShapeAreaRectangle(t *testing.T) {
	h := 9.0
	w := 3.0
	c := Rectangle{
		Height: h,
		Width:  w,
	}
	got, _ := Rectangle.Area(c)
	want := float64(27.00)
	assert.Equal(t, got, want)
}

func TestShapeAreaCircle(t *testing.T) {
	r := 8.0
	c := Circle{Radius: r}
	got, _ := Circle.Area(c)
	want := float64(201.06192982974676)
	assert.Equal(t, got, want)
}

func TestShapePerimeterCircle(t *testing.T) {
	r := 8.0
	c := Circle{Radius: r}
	got, _ := Circle.Perimeter(c)
	want := float64(50.26548245743669)
	assert.Equal(t, got, want)
}

func TestCircleShapestring(t *testing.T) {
	r := 8.0
	c := Circle{Radius: r}
	got := Circle.String(c)
	want := "\nCirlce: Radius 8.00"
	assert.Equal(t, got, want)
}

func TestRectangleShapestring(t *testing.T) {
	h := 9.0
	w := 3.0
	c := Rectangle{
		Height: h,
		Width:  w,
	}
	got := Rectangle.String(c)
	want := "\nRectangle with height 9.00 and width 3.00"
	assert.Equal(t, got, want)
}

func TestRectangle_Perimeter(t *testing.T) {
	tests := []struct {
		name    string
		r       Rectangle
		want    float64
		wantErr bool
	}{
		{
			name: "if value is zero return error",
			r: Rectangle{
				Height: 0.0,
				Width:  0.0,
			},
			want:    0.00,
			wantErr: true,
		},
		{
			name: "if value is negative return error",
			r: Rectangle{
				Height: -1.0,
				Width:  -2.0,
			},
			want:    0.00,
			wantErr: true,
		},
		{
			name: "if value is valid return Perimeter",
			r: Rectangle{
				Height: 9.0,
				Width:  3.0,
			},
			want:    24.00,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Perimeter()
			if (err != nil) != tt.wantErr {
				t.Errorf("Rectangle.Perimeter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rectangle.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Perimeter(t *testing.T) {
	tests := []struct {
		name    string
		c       Circle
		want    float64
		wantErr bool
	}{
		{name: "if value is zero return error",
			c: Circle{
				Radius: 0.0,
			},
			want:    0.00,
			wantErr: true,
		},
		{
			name: "if value is negative return error",
			c: Circle{
				Radius: -1.0,
			},
			want:    0.00,
			wantErr: true,
		},
		{
			name: "if value is valid return Perimeter",
			c: Circle{
				Radius: 8.0,
			},
			want:    50.26548245743669,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Perimeter()
			if (err != nil) != tt.wantErr {
				t.Errorf("Circle.Perimeter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Circle.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Area(t *testing.T) {
	tests := []struct {
		name    string
		c       Circle
		want    float64
		wantErr bool
	}{
		{name: "if value is zero return error",
			c: Circle{
				Radius: 0.0,
			},
			want:    0.00,
			wantErr: true,
		},
		{
			name: "if value is negative return error",
			c: Circle{
				Radius: -1.0,
			},
			want:    0.00,
			wantErr: true,
		},
		{
			name: "if value is valid return Area",
			c: Circle{
				Radius: 8.0,
			},
			want:    201.06192982974676,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Circle.Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Circle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Area(t *testing.T) {
	tests := []struct {
		name    string
		r       Rectangle
		want    float64
		wantErr bool
	}{
		{
			name: "if value is zero return error",
			r: Rectangle{
				Height: 0.0,
				Width:  0.0,
			},
			want:    0.00,
			wantErr: true,
		},
		{
			name: "if value is negative return error",
			r: Rectangle{
				Height: -1.0,
				Width:  -2.0,
			},
			want:    0.00,
			wantErr: true,
		},
		{
			name: "if value is valid return Area",
			r: Rectangle{
				Height: 9.0,
				Width:  3.0,
			},
			want:    27.00,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Rectangle.Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rectangle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}
