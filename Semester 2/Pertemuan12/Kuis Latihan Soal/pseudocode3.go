program olimpiade
kamus
	type peserta <
		t : 	string
		g, s, b int
	>
	type olimpiade : array [0..99] of peserta
	tab : olimpiade
	n : integer
Algoritma
	isiArray(tab, n)
	sorting(tab, n)
	tampilArray(tab, n)
Endprogram

procedure isiArray(in/out t : olimpiade, in/out n : integer)
kamus
	i : integer
Algoritma
	input(n)
	for i <- 0 to n-1 do
		input(t[i].t, t[i].g, t[i].s, t[i].b)
	endfor
endprocedure

procedure tampilArray(in t : olimpiade, in n : integer)
kamus
	i : integer
Algoritma
	for i <- 0 to n-1 do
		output(t[i].t, t[i].g, t[i].s, t[i].b)
	endfor
endprocedure

function poin(g, s, b : integer) -> integer
Algoritma
	return (4 * g) + (3 * s) + b
Endprogram

procedure sorting(in/out t : olimpiade, in n : integer)
kamus
	i, max, x : integer
Algoritma
	for i <- 0 to n-1 do
		max = i
		for x <- i to n-1 do 
			if poin(t[x].g, t[x].s, t[x].b) > poin(t[max].g, t[max].s, t[max].b) then
				max <- x
			endif
		endfor
		t[i], t[max] <- t[max], t[i]
	endfor
endprocedure