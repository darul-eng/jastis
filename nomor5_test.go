package jastis

import (
	"fmt"
	"testing"
)

type Kendaraan struct {
	suara string
}

func (kendaraan *Kendaraan) Akselerasi() {
	fmt.Println(kendaraan.suara)
}

type Sepeda struct {
	kendaraan Kendaraan
	rantai    string
}

func (sepeda Sepeda) Akselerasi() {
	sepeda.rantai = "Perlu Perbaikan"
	sepeda.kendaraan.Akselerasi()
}

type Mobil struct {
	kendaraan Kendaraan
	bensin    string
}

func (mobil Mobil) Akselerasi() {
	mobil.bensin = "Habis"
	mobil.kendaraan.Akselerasi()
}

func Vihicle() {
	sepeda := &Sepeda{
		kendaraan: Kendaraan{suara: "Swoosh"},
		rantai:    "Normal",
	}

	mobil := &Mobil{
		kendaraan: Kendaraan{suara: "Vroom"},
		bensin:    "Penuh",
	}

	sepeda.Akselerasi()
	fmt.Println(sepeda.rantai)

	mobil.Akselerasi()
	fmt.Println(mobil.bensin)
}
func TestAja(t *testing.T) {
	Vihicle()
}
