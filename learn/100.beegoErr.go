package main

import (
	"errors"
	"github.com/astaxie/beego"
)

func internalCalculationFunc(x, y int) (result int, err error) {
	beego.Debug("calculating z. x:", x, " y:", y)
	z := y
	switch {
	case x == 3:
		beego.Debug("x == 3")
		panic("Failure.")
	case y == 1:
		beego.Debug("y == 1")
		return 0, errors.New("Error!")
	case y == 2:
		beego.Debug("y == 2")
		z = x
	default:
		beego.Debug("default")
		z += x
	}
	retVal := z - 3
	beego.Debug("Returning ", retVal)

	return retVal, nil
}

func processInput(input int) {
	defer func() {
		if r := recover(); r != nil {
			beego.Error("Unexpected error occurred: ", r)
			outputs <- outputData{result: 0, error: true}
		}
	}()
	beego.Informational("Received input signal. x:", input.x, " y:", input.y)

	res, err := internalCalculationFunc(input.x, input.y)
	if err != nil {
		beego.Warning("Error in calculation:", err.Error())
	}

	beego.Informational("Returning result: ", res, " error: ", err)
	outputs <- outputData{result: res, error: err != nil}
}

func main() {
	inputs = make(chan int)
	outputs = make(chan int)
	criticalChan = make(chan int)
	beego.Informational("App started.")

	go consumeResults(outputs)
	beego.Informational("Started receiving results.")

	go generateInputs(inputs)
	beego.Informational("Started sending signals.")

	for {
		select {
		case input := <-inputs:
			processInput(input)
		case <-criticalChan:
			beego.Critical("Caught value from criticalChan: Go shut down.")
			panic("Shut down due to critical fault.")
		}
	}
}
