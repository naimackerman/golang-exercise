package main

import "fmt"

type Mobil struct {
	Roda                  [4]string
	PintuKanan, PintuKiri Pintu
}

type Pintu struct {
	BunyiKetuk string
	BunyiBuka  string
}

func (m *Mobil) GantiBan(ban [4]string) {
	for i := 0; i < 4; i++ {
		m.Roda[i] = ban[i]
	}
}

func (m *Mobil) KetukDanBukaPintu() {
	fmt.Println("Ketuk Pintu Kanan:", m.PintuKanan.BunyiKetuk)
	fmt.Println("Buka Pintu Kanan:", m.PintuKanan.BunyiBuka)

	fmt.Println("Ketuk Pintu Kiri:", m.PintuKiri.BunyiKetuk)
	fmt.Println("Buka Pintu Kiri:", m.PintuKiri.BunyiBuka)
}

func main() {
	ban := [4]string{"Ban Karet", "Ban Karet", "Ban Karet", "Ban Karet"}
	mobil := Mobil{
		Roda: ban,
		PintuKanan: Pintu{
			BunyiKetuk: "Knock",
			BunyiBuka:  "beep",
		},
		PintuKiri: Pintu{
			BunyiKetuk: "beep",
			BunyiBuka:  "Knock",
		},
	}

	fmt.Println("Mobil Awal:")
	mobil.KetukDanBukaPintu()
	for i := 0; i < 4; i++ {
		fmt.Println("Roda ke-", i+1, ":", mobil.Roda[i])
	}

	// Ganti ban mobil
	ban[0] = "Ban Besi"
	ban[1] = "Ban Kayu"
	ban[2] = "Ban Karet"
	ban[3] = "Ban Besi"
	mobil.GantiBan(ban)

	fmt.Println("\nMobil Setelah Ganti Ban:")
	mobil.KetukDanBukaPintu()
	for i := 0; i < 4; i++ {
		fmt.Println("Roda ke-", i+1, ":", mobil.Roda[i])
	}
}
