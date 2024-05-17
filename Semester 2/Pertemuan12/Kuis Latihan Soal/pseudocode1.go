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
	SelectionSort(arr, x)
	output(median(arr, x))
Endprogram

procedure SelectionSort(in/out A : tabInt, in T : integer)
Kamus
	minIdx, i : integer
Algoritma
	for i <- 0 to T-1 do
		minIdx <- i
		for x <- i to T-1 do 
			if A[j] > A[minIdx] then
				minIdx <- x 
			endif
		endfor
		A[i], A[minIdx] <- A[minIdx], A[i]
	endfor
endprocedure

function median(A : tabInt, T : integer) -> float64
Algoritma
	if T mod 2 == 0 {
		return float64(A[T/2-1]+A[T/2]) / 2.0
	}
	return float64(A[T/2])
Endfunction