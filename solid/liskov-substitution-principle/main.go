package liskovsubstitutionprinciple

type Area interface {
	GetArea() float64
}

type RetanguloType interface {
	Area

	GetWidth() float64
	GetHeight() float64

	SetWidth(width float64)
	SetHeight(height float64)
}

type QuadradoType interface {
	Area

	SetSize(size float64)
	GetSize() float64
}

func (r *Quadrado) GetArea() float64 {
	return r.size * r.size
}

func (r *Quadrado) SetSize(size float64) {
	r.size = size
}

func (r *Quadrado) GetSize() float64 {
	return r.size
}

type Quadrado struct {
	size float64
}

type Retangulo struct {
	width, height float64
}

func (r *Retangulo) GetArea() float64 {
	return r.width * r.height
}

func (r *Retangulo) GetWidth() float64 {
	return r.width
}

func (r *Retangulo) GetHeight() float64 {
	panic("not implemented") // TODO: Implement
}

func (r *Retangulo) SetWidth(width float64) {
	panic("not implemented") // TODO: Implement
}

func (r *Retangulo) SetHeight(height float64) {
	panic("not implemented") // TODO: Implement
}
