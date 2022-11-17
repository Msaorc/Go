package module

type User struct {
	Name     string
	PassWord string
	Age      int
	Weight   float32
	Height   float32
	Sex      string
	BodyComposition
}

type BodyComposition struct {
	FatPerc float32
	Muscle  float32
	Water   float32
	Imc     float32
}

func (U *User) CalcIMC() {
	U.Imc = U.Weight / (U.Height * U.Height)
}
