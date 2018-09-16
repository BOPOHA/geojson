package poly

import "testing"

func testIntersectsLinesA(t *testing.T, a, b, c, d Point, expect bool) {
	res := lineintersects(a, b, c, d)
	if res != expect {
		t.Fatalf("{%v,%v}, {%v,%v} = %t, expect %t", a, b, c, d, res, expect)
	}
	res = lineintersects(b, a, c, d)
	if res != expect {
		t.Fatalf("{%v,%v}, {%v,%v} = %t, expect %t", b, a, c, d, res, expect)
	}
	res = lineintersects(a, b, d, c)
	if res != expect {
		t.Fatalf("{%v,%v}, {%v,%v} = %t, expect %t", a, b, d, c, res, expect)
	}
	res = lineintersects(b, a, d, c)
	if res != expect {
		t.Fatalf("{%v,%v}, {%v,%v} = %t, expect %t", b, a, d, c, res, expect)
	}
}

func testIntersectsLines(t *testing.T, a, b, c, d Point, expect bool) {
	testIntersectsLinesA(t, a, b, c, d, expect)
	testIntersectsLinesA(t, c, d, a, b, expect)
}

func TestIntersectsLines(t *testing.T) {
	testIntersectsLines(t, P(0, 6), P(12, -6), P(0, 0), P(12, 0), true)
	testIntersectsLines(t, P(0, 0), P(5, 5), P(5, 5), P(0, 10), true)
	testIntersectsLines(t, P(0, 0), P(5, 5), P(5, 6), P(0, 10), false)
	testIntersectsLines(t, P(0, 0), P(5, 5), P(5, 4), P(0, 10), true)
	testIntersectsLines(t, P(0, 0), P(2, 2), P(0, 2), P(2, 0), true)
	testIntersectsLines(t, P(0, 0), P(2, 2), P(0, 2), P(1, 1), true)
	testIntersectsLines(t, P(0, 0), P(2, 2), P(2, 0), P(1, 1), true)
	testIntersectsLines(t, P(0, 0), P(0, 4), P(1, 4), P(4, 1), false)
	testIntersectsLines(t, P(0, 0), P(0, 4), P(1, 4), P(4, 4), false)
	testIntersectsLines(t, P(0, 0), P(0, 4), P(4, 1), P(4, 4), false)
	testIntersectsLines(t, P(0, 0), P(4, 0), P(1, 4), P(4, 1), false)
	testIntersectsLines(t, P(0, 0), P(4, 0), P(1, 4), P(4, 4), false)
	testIntersectsLines(t, P(0, 0), P(4, 0), P(4, 1), P(4, 4), false)
	testIntersectsLines(t, P(0, 4), P(4, 0), P(1, 4), P(4, 1), false)
	testIntersectsLines(t, P(0, 4), P(4, 0), P(1, 4), P(4, 4), false)
	testIntersectsLines(t, P(0, 4), P(4, 0), P(4, 1), P(4, 4), false)
}

func testIntersectsShapes(t *testing.T, exterior Polygon, holes []Polygon, shape Polygon, expect bool) {
	got := shape.Intersects(exterior, holes)
	if got != expect {
		t.Fatalf("%v intersects %v = %v, expect %v", shape, exterior, got, expect)
	}
	got = exterior.Intersects(shape, nil)
	if got != expect {
		t.Fatalf("%v intersects %v = %v, expect %v", exterior, shape, got, expect)
	}
}

func TestIntersectsShapes(t *testing.T) {

	testIntersectsShapes(t,
		Polygon{P(6, 0), P(12, 0), P(12, -6), P(6, 0)},
		nil,
		Polygon{P(0, 0), P(0, 6), P(6, 0), P(0, 0)},
		true)

	testIntersectsShapes(t,
		Polygon{P(7, 0), P(12, 0), P(12, -6), P(7, 0)},
		nil,
		Polygon{P(0, 0), P(0, 6), P(6, 0), P(0, 0)},
		false)

	testIntersectsShapes(t,
		Polygon{P(0.5, 0.5), P(0.5, 4.5), P(4.5, 0.5), P(0.5, 0.5)},
		nil,
		Polygon{P(0, 0), P(0, 6), P(6, 0), P(0, 0)},
		true)

	testIntersectsShapes(t,
		Polygon{P(0, 0), P(0, 6), P(6, 0), P(0, 0)},
		[]Polygon{{P(1, 1), P(1, 2), P(2, 2), P(2, 1), P(1, 1)}},
		Polygon{P(0.5, 0.5), P(0.5, 4.5), P(4.5, 0.5), P(0.5, 0.5)},
		true)

	testIntersectsShapes(t,
		Polygon{P(0, 0), P(0, 10), P(10, 10), P(10, 0), P(0, 0)},
		[]Polygon{{P(2, 2), P(2, 6), P(6, 6), P(6, 2), P(2, 2)}},
		Polygon{P(1, 1), P(1, 9), P(9, 9), P(9, 1), P(1, 1)},
		true)
}

func TestPointIntersectsLineString(t *testing.T) {
	poly := Polygon{P(0, 0), P(10, 10), P(20, 0)}
	if !P(0, 0).IntersectsLineString(poly) {
		t.Fatal("expected true")
	}
	if !P(10, 10).IntersectsLineString(poly) {
		t.Fatal("expected true")
	}
	if !P(20, 0).IntersectsLineString(poly) {
		t.Fatal("expected true")
	}
	if !P(5, 5).IntersectsLineString(poly) {
		t.Fatal("expected true")
	}
	if !P(15, 5).IntersectsLineString(poly) {
		t.Fatal("expected true")
	}
	if P(20, 5).IntersectsLineString(poly) {
		t.Fatal("expected false")
	}
}

func TestPointIntersects(t *testing.T) {
	poly := Polygon{P(0, 0), P(10, 0), P(10, 10), P(0, 10), P(0, 0)}

	if !P(0, 0).Intersects(poly, nil) {
		t.Fatal("expected true")
	}
	if !P(10, 0).Intersects(poly, nil) {
		t.Fatal("expected true")
	}
	if !P(0, 10).Intersects(poly, nil) {
		t.Fatal("expected true")
	}
	if !P(10, 10).Intersects(poly, nil) {
		t.Fatal("expected true")
	}
}
func TestRectIntersectsPolygon(t *testing.T) {
	poly := Polygon{P(0, 0), P(10, 0), P(10, 10), P(0, 10), P(0, 0)}
	if !R(0, 0, 5, 5).Intersects(poly, nil) {
		t.Fatal("expected true")
	}
	if R(15, 15, 20, 20).Intersects(poly, nil) {
		t.Fatal("expected true")
	}
}

func TestLineStringIntersectsLineString(t *testing.T) {
	line1 := Polygon{P(0, 0), P(10, 10), P(20, 0)}
	line2 := Polygon{P(0, 1), P(10, 11), P(20, 1)}
	line3 := Polygon{P(0, -1), P(10, 11), P(20, -1)}
	if line1.LineStringIntersectsLineString(line2) {
		t.Fatal("expected false")
	}
	if !line1.LineStringIntersectsLineString(line3) {
		t.Fatal("expected true")
	}
}

func TestLineStringIntersects(t *testing.T) {
	poly := Polygon{P(0, 0), P(10, 0), P(10, 10), P(0, 10), P(0, 0)}
	line1 := Polygon{P(0, 0), P(10, 10), P(20, 0)}

	if !line1.LineStringIntersects(poly, nil) {
		t.Fatal("expected true")
	}
}

func TestDoesIntersect(t *testing.T) {
	exterior := Polygon{P(0, 0), P(10, 0), P(10, 10), P(0, 10), P(0, 0)}
	holes := []Polygon{{P(2, 2), P(2, 8), P(8, 8), P(8, 2), P(2, 2)}}

	if (Polygon{}).doesIntersects(false, exterior, holes) {
		t.Fatal("expected false")
	}
	if (Polygon{P(5, 5)}).doesIntersects(false, exterior, holes) {
		t.Fatal("expected false")
	}
	if !(Polygon{P(1, 1)}).doesIntersects(false, exterior, holes) {
		t.Fatal("expected true")
	}
	if (Polygon{P(1, 1)}).doesIntersects(false, Polygon{}, nil) {
		t.Fatal("expected false")
	}
	if !(Polygon{P(1, 1)}).doesIntersects(false, Polygon{P(1, 1)}, nil) {
		t.Fatal("expected true")
	}
	if (Polygon{P(1, 1), P(2, 2)}).doesIntersects(false, Polygon{}, nil) {
		t.Fatal("expected false")
	}
	if !exterior.doesIntersects(false, Polygon{P(1, 1)}, nil) {
		t.Fatal("expected true")
	}

	inner := Polygon{P(3, 3), P(7, 3), P(7, 7), P(3, 7), P(3, 3)}
	if inner.doesIntersects(false, exterior, holes) {
		t.Fatal("expected false")
	}

	outside := Polygon{P(30, 30), P(60, 30), P(60, 60), P(30, 60), P(30, 30)}
	if outside.doesIntersects(false, exterior, holes) {
		t.Fatal("expected false")
	}

	// triangles

	tri1 := Polygon{P(0, 0), P(10, 0), P(5, 10), P(0, 0)}
	tri2 := Polygon{P(7, 9), P(17, 9), P(12, 19), P(7, 9)}

	if tri1.doesIntersects(false, tri2, nil) {
		t.Fatal("expected false")
	}

	// if !line1.doesIntersect(poly, nil) {
	// 	t.Fatal("expected true")
	// }
}
