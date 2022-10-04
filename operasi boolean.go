/*
&& ==> dan
|| ==> atau
!  ==> kebalikan
*/
package main

import "fmt"

func booll() {
	var nilaiAkhir = 98
	var absensi = 88

	var lulusUjian = nilaiAkhir > 89
	var lulusAbsensi = absensi > 88

	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusAbsensi && lulusUjian
	fmt.Println(lulus)

}
