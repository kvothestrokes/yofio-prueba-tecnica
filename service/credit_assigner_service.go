package service

import "errors"

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

type credictAssigner struct{}

func NewCreditAssignerService() CreditAssigner {
	return &credictAssigner{}
}

func (s *credictAssigner) Assign(inversion int32) (int32, int32, int32, error) {

	credit_700, credit_500, credit_300 := CalcularInversion(inversion)
	if credit_700 == 0 && credit_500 == 0 && credit_300 == 0 {
		return 0, 0, 0, errors.New("no se pudo realizar la inversion")
	}

	return credit_700, credit_500, credit_300, nil
}

// Catalogo de creditos
const CREDIT_700 = 700
const CREDIT_500 = 500
const CREDIT_300 = 300

func CalcularInversion(inversion int32) (int32, int32, int32) {

	//Se calcula la maxima cantidad de inversiones de 700 posibles
	max_700 := inversion / CREDIT_700

	//Iteramos sobre la cantidad de inversiones de 700 posibles
	for i := max_700; i >= 0; i-- {
		//Calculamos la inversion restante despues de asignar las inversiones de 700
		inversion_restante_700 := inversion - (i * CREDIT_700)

		//Calculamos la maxima cantidad de inversiones de 500 posibles
		max_500 := inversion_restante_700 / CREDIT_500

		//Iteramos sobre la cantidad de inversiones de 500 posibles
		for j := max_500; j >= 0; j-- {
			//Calculamos la inversion restante despues de asignar las inversiones de 500
			inversion_restante_500 := inversion_restante_700 - (j * CREDIT_500)

			//Si la inversion es divisible entre 300, significa que no sobra nada y se encontro una combinacion de inversiones
			if inversion_restante_500%CREDIT_300 == 0 {
				//Calculamos la cantidad de inversiones de 300
				k := inversion_restante_500 / CREDIT_300
				return i, j, k
			}
		}
	}
	//Si no se encontro una combinacion de inversiones valida, se retorna 0, 0, 0
	return 0, 0, 0
}
