package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Pedido struct {
	ID int32
	Cliente string
	Valor float32
	Status bool
}

func (p *Pedido) Processar() bool{
	var x bool = p.Validar()
	p.Status = x
	return  x
}

func (p *Pedido) Validar() bool {
	if(p.Valor < 0) {
		panic("Valor invalido")
	}
	if (p.Cliente == "" || p.Valor == 0) {
		return false
	} else {
		return true
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var chpedidos chan Pedido
	chpedidos = make(chan Pedido)
	sistem := func () {
		fmt.Println("Sistema encerrado!")
	}
	defer sistem()
	var pedidos [10]Pedido
	
	for i := 0; i < 10; i++ {
		pedidos[i].ID = int32(i+1)
		pedidos[i].Valor = rand.Float32()
		pedidos[i].Cliente = string(rand.Int())
	}

	go woker(chpedidos)

	for i := 0; i < 10; i++  {
		chpedidos <- pedidos[i]
	}


}

go woker(P Pedido)  {
	defer func() {
		if x := recover(); x != nil{
			fmt.Println("Erro capturado: ", x)
		}
	}()
	P.Processar()
}
