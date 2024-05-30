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
kamus
	i, j : integer
	save : arrayBunga
Algoritma
	for i <- 1 to N do
		save <- tabBunga[i]
		j <- i - 1
		while j >= 0 && panjang(tabBunga[j]) > panjang(save) do
			tabBunga[j+1] <- tabBunga[j]
			j <- j - 1
		endwhile
		tabBunga[j+1] <- save
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