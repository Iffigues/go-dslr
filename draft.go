
package data

import (
	"errors"
	"math"
)


type Meta struct {
	Theta_init  [] float64
	Epsilon     float64
	Seuil       float64
	L           Data
}


func (d Meta) logistic(theta, x [] float64) (result float64) {
	prod = tab_prod(theta, x)
	log = 1. / (1. + math.Exp(-prod))
	return [log, log, log ,log]
}


func (d Meta) Grad_theta_feature(theta, X, Y [] float64) (res float64) {

	res = 0
	for i := 0; i <= d.L.Size-1; i++ {
		res = res + d.tab_scalar( d.tab_elem_prod( d.L.X[i,j], Y[i] ), d.tab_add( d.logistic(theta,d.L.X[i]), d.tab_neg(Y[i]) )  )
	}
	return (1.0 / float64(d.L.Size) * res

}


func (d Meta) Grad_theta(theta, X, Y [] float64) (res [] float64) {
	res = []
	for j := 0; j <= d.theta.Size-1; j++ {
		res[j] = d.Grad_theta_feature(theta, X, Y)
	}
	return res
}


func (d Meta) Train() (theta, X, Y [] float64) {
	theta = d.Theta_init
	theta_next := d.tab_add( theta, tab_neg( d.Grad_theta(theta, X, Y) ) ) 

	for i := 0; i < 1000; i++ {
		
		diff := d.tab_abs( theta_next, theta )
		bool := True
		
		for j := 0; j <= d.theta.Size-1; j++ {
			bool = ( (diff[j] <= d.Seuil) && (bool) )
		}

		if bool {
			return theta
		}

		theta = theta_next
		theta_next = d.tab_add( theta, d.tab_neg(d.Grad_theta(theta, X, Y)) )

	}
	return
}


#############
### Specs ###
#############

# Y : matrice des labels des données | Y[i] -> tableau du style [1, 0, 0, 0] de la ième ligne
# X : matrice des données | X[i] -> tableau avec les données de la ligne i ; X[i,j] -> valeur de la jème feature de la ligne i
# tab_abs : effectue le produit scalaire de chaque élément du tableau | ex : tab_abs([-1,2]) -> [1,2]
# tab_scalar : effecture un produit scalaire | ex : tab_scalar([1,2], [3,4]) -> 1*3 + 2*4 = 11
# tab_elem_prod : effectue un produit avec un tableau par un elem | ex : tab_elem_prod(2,[1,3]) -> [2,6]
# tab_prod : effectue un produit de tableaux | ex : tab_prod([1,2],[3,4]) -> [3,8]
# tab_neg : négative le tableau | ex : tab_neg([1,1]) -> [-1,-1]
# tab_add : additionne deux tableaux valeur par valeur | ex : tab_add([1,1],[2,2]) -> [3,3]

