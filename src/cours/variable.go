package cours

import "fmt"

var LeNomDeMaVariablePublic = "LeNomDeMaVariablePublic"
var leNomDeMaVariablePrivate = "LeNomDeMaVariablePrivate"

var LeNomDeMaVariablePublicAvecType int = 3
var leNomDeMaVariablePrivateAvecType int = 3

var LeNomDeMaVariablePublicAvecTypeSansValeur int  // nil
var LeNomDeMaVariablePrivateAvecTypeSansValeur int // nil

var LeNomDeMaConstantePublic = "LeNomDeMaConstantePublic"
var leNomDeMaConstantePrivate = "leNomDeMaConstantePrivate"

func DisplayVariable() {

		fmt.Println(leNomDeMaVariablePrivateAvecType)
		fmt.Println(leNomDeMaVariablePrivate)
		fmt.Println(LeNomDeMaVariablePrivateAvecTypeSansValeur)
		fmt.Println(leNomDeMaConstantePrivate)
	
}
