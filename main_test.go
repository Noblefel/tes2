package main

import (
	"testing"
)

func TestHitungGaji(t *testing.T) {
	tests := []struct {
		name  string
		jam   float64
		tarif float64

		//hasil yg diinginkan
		gaji       float64
		gajiLembur float64
	}{
		{
			name:       "10 jam kerja dengan tarif 150rb/jam",
			jam:        10,
			tarif:      150_000,
			gaji:       1_500_000,
			gajiLembur: 0,
		},
		{
			name:       "25 jam kerja dengan tarif 200rb/jam",
			jam:        25,
			tarif:      200_000,
			gaji:       5_000_000,
			gajiLembur: 0,
		},
		{
			name:       "45 jam kerja (lembur) dengan tarif 100rb/jam",
			jam:        45,
			tarif:      100_000,
			gaji:       4_000_000, //gaji biasa yg dihitung selama 40 jam
			gajiLembur: 750_000,   //berarti 5 jam kerja lembur, 5 * 100rb * 1.5
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gaji, gajiLembur := hitungGaji(tt.jam, tt.tarif)

			if gaji != tt.gaji {
				t.Errorf("gaji seharusnya %.0f, malah dapet %.0f", tt.gaji, gaji)
			}

			if gajiLembur != tt.gajiLembur {
				t.Errorf("gaji lembur seharusnya %.0f, malah dapet %.0f", tt.gajiLembur, gajiLembur)
			}
		})
	}
}
