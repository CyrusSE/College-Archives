program Himpunan
kamus
	constanta nMax : integer = 37
	type tHimpunan <
		anggota : array [1..nMax] of integer
		panjang : integer
	>
	set1, set2 : tHimpunan
Algoritma
	output("Anggota himpunan 1:")
	bacaMasukan(set1)
	output("Anggota himpunan 2:")
	bacaMasukan(set2)
	output("Himpunan 1 = Himpunan 2?", sama(set1, set2))
Endprogram

procedure bacaMasukan(in/out set : tHimpunan)
kamus
	angka : integer
Algoritma
	repeat 
		input(angka)
		set.anggota[set.panjang] = angka
		set.panjang++
	untill ada(set, angka)
	urut(set)
endprocedure

function ada(set : tHimpunan, x : integer) -> boolean
kamus
	found, i : integer
Algoritma
	for i <- 0 to set.panjang do
		if set.anggota[i] == x then
			ada <- true
		endif
	endfor
	return ada
endfunction

procedure urut(in/out set : tHimpunan)
kamus
	i, j : integer
	save : tHimpunan
Algoritma
	for i <- 1 to set.panjang do
		save <- set.anggota[i]
		j <- i
		while j > 0 AND save > set.anggota[j-1] do
			set.anggota[j] <- set.anggota[j-1]
			j <- j - 1
		endwhile
		set.anggota[j] <- save
	endfor
endprocedure

function sama(set1, set2 : tHimpunan) -> boolean do
kamus
	same : boolean
Algoritma
	if set1.panjang == set2.panjang then
		if set1.anggota == set2.anggota then
			same <- true
		endif
	endif
	return same
endfunction