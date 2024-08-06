package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"math"
	"strconv"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Golang-StrongholdFinder by Intrigant13")
	myWindow.Resize(fyne.NewSize(500, 200))
	myWindow.SetFixedSize(true)
	// Объявление некоторых переменных /Declaration of some variables
	Z := ""
	// Создание виджетов, кнопок и тд. /Creation of widgets, buttons, etc.
	Out := widget.NewLabel("")
	EntryX1 := widget.NewEntry()
	EntryX2 := widget.NewEntry()
	EntryZ1 := widget.NewEntry()
	EntryZ2 := widget.NewEntry()
	EntryF1 := widget.NewEntry()
	EntryF2 := widget.NewEntry()
	bt1 := widget.NewButton("Найти!", func() {
		// Преобразование всех значений из str в float64 /Converting all values from str to float64
		F1, err := strconv.ParseFloat(EntryF1.Text, 64)
		if err != nil {
			Out.SetText("Недопустимые символы F1")
			return
		}
		F2, err := strconv.ParseFloat(EntryF2.Text, 64)
		if err != nil {
			Out.SetText("Недопустимые символы F2")
			return
		}
		X1, err := strconv.ParseFloat(EntryX1.Text, 64)
		if err != nil {
			Out.SetText("Недопустимые символы X1")
			return
		}
		X2, err := strconv.ParseFloat(EntryX2.Text, 64)
		if err != nil {
			Out.SetText("Недопустимые символы X2")
			return
		}
		Z1, err := strconv.ParseFloat(EntryZ1.Text, 64)
		if err != nil {
			Out.SetText("Недопустимые символы Z1")
			return
		}
		Z2, err := strconv.ParseFloat(EntryZ2.Text, 64)
		if err != nil {
			Out.SetText("Недопустимые символы Z2")
			return
		}
		if F1 == F2 {
			Out.SetText("Углы F1 и F2 равны\nВычисление невозможно")
			return
		}
		F1 = math.Pi * (0.5 + F1/180)
		F2 = math.Pi * (0.5 + F2/180)
		// Основной код для расчета нахождение крепости с помощью триангуляции. /Main code for calculating the stronghold location using triangulation
		findX := (math.Cos(F2)*(math.Cos(F1)*Z1-math.Sin(F1)*X1) - math.Cos(F1)*(math.Cos(F2)*Z2-math.Sin(F2)*X2)) / math.Sin(F2-F1)
		findZ := (math.Sin(F2)*(math.Sin(F1)*X1-math.Cos(F1)*Z1) - math.Sin(F1)*(math.Sin(F2)*X2-math.Cos(F2)*Z2)) / math.Sin(F1-F2)
		Z = fmt.Sprintf("X:%d Z:%d", int(findX), int(findZ))
		Out.SetText(Z)
	})

	// Кординаты где будут размещены объекты / Coordinates where objects will be placed
	Out.Move(fyne.NewPos(50, 100))
	EntryX1.Move(fyne.NewPos(50, 20))
	EntryX2.Move(fyne.NewPos(50, 60))
	EntryZ1.Move(fyne.NewPos(200, 20))
	EntryZ2.Move(fyne.NewPos(200, 60))
	EntryF1.Move(fyne.NewPos(350, 20))
	EntryF2.Move(fyne.NewPos(350, 60))
	bt1.Move(fyne.NewPos(380, 100))

	// Размер объектов /Size of objects
	Out.Resize(fyne.NewSize(150, 25))
	EntryX1.Resize(fyne.NewSize(100, 30))
	EntryX2.Resize(fyne.NewSize(100, 30))
	EntryZ1.Resize(fyne.NewSize(100, 30))
	EntryZ2.Resize(fyne.NewSize(100, 30))
	EntryF1.Resize(fyne.NewSize(100, 30))
	EntryF2.Resize(fyne.NewSize(100, 30))
	bt1.Resize(fyne.NewSize(70, 25))
	//Подсказки к полям ввода /Hints for input fields
	EntryF1.SetPlaceHolder("Угол1")
	EntryF2.SetPlaceHolder("Угол2")
	EntryX1.SetPlaceHolder("X1")
	EntryX2.SetPlaceHolder("X2")
	EntryZ1.SetPlaceHolder("Z1")
	EntryZ2.SetPlaceHolder("Z2")

	// Контейнер предназначен для вывода тех элементов которые указаны /The container is designed to display the specified elements
	content := container.NewWithoutLayout(EntryX1, EntryX2, EntryZ1, EntryZ2, EntryF2, EntryF1, bt1, Out)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
