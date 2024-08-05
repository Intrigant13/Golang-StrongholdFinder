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
	myWindow.Resize(fyne.NewSize(380, 160))
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
		F1, _ := strconv.ParseFloat(EntryF1.Text, 64)
		F2, _ := strconv.ParseFloat(EntryF2.Text, 64)
		X1, _ := strconv.ParseFloat(EntryX1.Text, 64)
		X2, _ := strconv.ParseFloat(EntryX2.Text, 64)
		Z1, _ := strconv.ParseFloat(EntryZ1.Text, 64)
		Z2, _ := strconv.ParseFloat(EntryZ2.Text, 64)
		F1 = 3.141592653589793 * (0.5 + F1/180)
		F2 = 3.141592653589793 * (0.5 + F2/180)
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
	EntryZ1.Move(fyne.NewPos(150, 20))
	EntryZ2.Move(fyne.NewPos(150, 60))
	EntryF1.Move(fyne.NewPos(250, 20))
	EntryF2.Move(fyne.NewPos(250, 60))
	bt1.Move(fyne.NewPos(250, 100))

	// Размер объектов /Size of objects
	Out.Resize(fyne.NewSize(150, 25))
	EntryX1.Resize(fyne.NewSize(70, 25))
	EntryX2.Resize(fyne.NewSize(70, 25))
	EntryZ1.Resize(fyne.NewSize(70, 25))
	EntryZ2.Resize(fyne.NewSize(70, 25))
	EntryF1.Resize(fyne.NewSize(70, 25))
	EntryF2.Resize(fyne.NewSize(70, 25))
	bt1.Resize(fyne.NewSize(70, 25))
	// Контейнер предназначен для вывода тех элементов которые указаны /The container is designed to display the specified elements
	content := container.NewWithoutLayout(EntryX1, EntryX2, EntryZ1, EntryZ2, EntryF2, EntryF1, bt1, Out)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
