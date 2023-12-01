package main

import (
	"fmt"
	"strings"
)

type Mobil struct {
	Roda  [4]string
	Pintu [2]Pintu
}

type Pintu struct {
	BunyiKetuk string
	BunyiBuka  string
}

type Ban struct {
	Karet string
	Besi  string
	Kayu  string
}

func (m *Mobil) GantiBan(ban ...string) {
	copy(m.Roda[:], ban[:])
}

func (m *Mobil) KetukPintu(sisi string) {
	sisi = strings.ToLower(sisi)
	if sisi == "kanan" {
		fmt.Println("Ketuk Pintu Kanan:", m.Pintu[0].BunyiKetuk)
	} else if sisi == "kiri" {
		fmt.Println("Ketuk Pintu Kiri:", m.Pintu[1].BunyiKetuk)
	}
}

func (m *Mobil) BukaPintu(sisi string) {
	sisi = strings.ToLower(sisi)
	if sisi == "kanan" {
		fmt.Println("Buka Pintu Kanan:", m.Pintu[0].BunyiBuka)
	} else if sisi == "kiri" {
		fmt.Println("Buka Pintu Kiri:", m.Pintu[1].BunyiBuka)
	}
}

func main() {
	ban := Ban{
		Karet: "Ban Karet",
		Besi:  "Ban Besi",
		Kayu:  "Ban Kayu",
	}
	mobil := Mobil{
		Roda: [4]string{
			ban.Besi,
			ban.Besi,
			ban.Besi,
			ban.Besi,
		},
		Pintu: [2]Pintu{
			{
				BunyiKetuk: "Knock",
				BunyiBuka:  "Beep",
			},
			{
				BunyiKetuk: "Beep",
				BunyiBuka:  "Knock",
			},
		},
	}

	mobil.BukaPintu("KaNaN")
	mobil.BukaPintu("KIRI")
	mobil.KetukPintu("kanan")
	mobil.KetukPintu("kiRI")

	fmt.Println("Ban mobil sebelum diganti:")
	for i := 0; i < len(mobil.Roda); i++ {
		fmt.Println("Ban mobil roda ke-", i+1, ":", mobil.Roda[i])
	}
	mobil.GantiBan(ban.Karet, ban.Kayu, ban.Karet, ban.Kayu)
	fmt.Println("Ban mobil setelah diganti:")
	for i := 0; i < len(mobil.Roda); i++ {
		fmt.Println("Ban mobil roda ke-", i+1, ":", mobil.Roda[i])
	}
}
