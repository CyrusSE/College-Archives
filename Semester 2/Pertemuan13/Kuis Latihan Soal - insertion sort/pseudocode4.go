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
endfunction

procedure sorting(in/out t : olimpiade, in n : integer)
kamus 
	pass, i : integer
	temp : peserta
Algoritma
	pass <- 1
	while pass <= n - 1 do
		i <- pass
		temp <- t[pass]
		while i > 0 AND poin(temp.g, temp.s, temp.s) > poin(t[i-1].g, t[i-1].s, t[i-1].b) do
			t[i] <- t[i-1]
			i <- i - 1
		endwhile
		t[i] <- temp
		pass <- pass + 1
	endwhile
endprocedure