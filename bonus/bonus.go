package bonus

//Joãozinho ganhou um kit de construção de torres. O kit é composto por várias barras de madeira, e seus comprimentos são
//conhecidos. As barras podem ser empilhadas umas sobre as outras, desde que seus comprimentos sejam iguais.
//
//Joãozinho quer construir o menor número possível de torres com as barras que tem. Você deve ajudar Joãozinho a usar as
//barras da melhor maneira possível, determinando a altura da torre mais alta e quantas torres podem ser construídas.

func CalculateTowers(barLengths []int) (int, int) {
	// Seu código aqui
	height := 1
	quantidade := len(barLengths)
	for i := 0; i < len(barLengths); i++ {
		if barLengths[i] == barLengths[i+1] {
			height++
			quantidade--
		}
	}
	return height, quantidade
}
