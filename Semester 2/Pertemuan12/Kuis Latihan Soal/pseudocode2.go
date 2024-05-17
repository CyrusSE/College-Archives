Program buanga_terurut
Kamus
	constanta NMAX : integer = 1000
	type arrayBunga : array [0..NMAX-1] of string
	A : arrayBunga
	N : integer
Algoritma
	input(N)
	isiArray(A, N)
	mengurutkan(A, N)
	tampilArray(A, N)
Endprogram

function panjang(bunga : string) -> integer
Algoritma
	return len(bunga)
Endfunction

procedure mengurutkan(in/out tabBunga : arrayBunga, in N : integer)
Kamus
	i, min, x : integer
Algoritma
	for i <- 0 to N-1 do
		min = i
		for x <- i to N-1 do
			if panjang(tabBunga[min]) > panjang(tabBunga[x]) then
				min = x
			endif
		endfor
		tabBunga[i], tabBunga[min] <- tabBunga[min], tabBunga[i]
	endfor
endprocedure

procedure isiArray(in/out A : arrayBunga, in N : integer)
Kamus
	i : integer
Algoritma
	for i <- 0 to N-1 do
		input(A[i])
	endfor
endprocedure

procedure tampilArray(in A : arrayBunga, in N : integer)
kamus 
	i : integer
Algoritma
	for i <- 0 to N-1 do
		output(A[i])
	endfor
Endprocedure