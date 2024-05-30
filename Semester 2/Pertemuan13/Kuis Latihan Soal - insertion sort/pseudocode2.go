Program mencari_median
Kamus
	constanta NMAX : integer = 10
	type tabInt : array [0..NMAX-1] of integer
	x, i : integer
	arr : tabInt
Algoritma
	input(x)
	for i <- 0 to x-1 do
		input(arr[i])
	endfor
	InsertionSort(arr, x)
	output(median(arr, x))
Endprogram

procedure InsertionSort(in/out A : tabInt, in T : integer)
Kamus
	i, j : integer
	save : tabInt
Algoritma
	for i <- 1 to T do
		save <- A[i]
		j <- i - 1
		while j >= 0 && A[j] > save do
			A[j+1] <- A[j]
			j <- j - 1
		endwhile
		A[j+1] <- save
	endfor
endprocedure

function median(A : tabInt, T : integer) -> real
Algoritma
	if T mod 2 == 0 {
		return A[T/2-1]+A[T/2] / 2.0
	}
	return A[T/2]
Endfunction