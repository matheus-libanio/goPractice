package calcimposto

func CalcImposto(salario int) int {
	var imposto int
	switch {
	case salario > 150:
		imposto = 10
		fallthrough
	case salario > 50:
		imposto += 17
		return imposto
	default:
		return imposto

	}
}
