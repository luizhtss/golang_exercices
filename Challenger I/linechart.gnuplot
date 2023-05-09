set datafile separator ","
set title "Evolução do número de nascimentos"
set xlabel "Ano"
set ylabel "Número de nascimentos"
set key autotitle columnheader
set grid
set term pngcairo enhanced font "Helvetica,14"
set output "grafico_linha.png"

set style line 1 lc rgb "#1f77b4" lw 2
set style line 2 lc rgb "#ff7f0e" lw 2
set style line 3 lc rgb "#2ca02c" lw 2
set style line 4 lc rgb "#d62728" lw 2
set style line 5 lc rgb "#9467bd" lw 2
set style line 6 lc rgb "#8c564b" lw 2

set border linewidth 1.5
set grid linestyle 0 lc rgb "#dddddd" lt 1
set style line 11 lc rgb "#808080" lt 1
set style line 12 lc rgb "#808080" lt 0 lw 1

set xlabel font "Helvetica,14"
set ylabel font "Helvetica,14"

set key font "Helvetica,12"

set title font "Helvetica,16"

set size ratio 0.5
set terminal pngcairo size 1200, 600

set key right top outside
set xrange[1994:2020]
plot for [i=2:*] "nascimentos-alvos.dat" using 1:(column(i) == 0 ? NaN : column(i)) with lines linestyle i-1
