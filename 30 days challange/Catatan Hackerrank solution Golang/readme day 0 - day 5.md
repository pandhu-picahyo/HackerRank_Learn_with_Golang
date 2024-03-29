# Catatan 30 days code challange
# Day 0 Hello World
Contoh penggunaan umum __bufio__ dalam Go termasuk: <br>

1. Membaca dan Menulis dari/ke File: Anda dapat menggunakan bufio untuk membaca dan menulis data dari dan ke file dengan buffering. Ini dapat meningkatkan efisiensi operasi I/O. <br>
2. Membaca dari Input Standar (stdin): Anda dapat menggunakan bufio untuk membaca input dari pengguna melalui terminal. <br>
3. Menulis ke Output Standar (stdout): Anda dapat menggunakan bufio untuk menulis data ke terminal atau output standar lainnya. <br>
4. Membaca dan Menulis dari/ke Koneksi Jaringan: Jika Anda mengembangkan program yang berkomunikasi melalui jaringan, Anda bisa menggunakan bufio untuk membaca dan menulis data dari dan ke koneksi jaringan. <br>

# Day 1 Data Type
Here are some of the common tasks that the __strconv__ package can help you with: <br>

1. Converting Numbers to Strings and Vice Versa:
    * __strconv.Itoa__ and __strconv.FormatInt__ for converting integers to strings.
    * __strconv.FormatFloat__ for converting floating-point numbers to strings.
    * __strconv.Atoi__ and __strconv.ParseInt__ for parsing strings into integers.
    * __strconv.ParseFloat__ for parsing strings into floating-point numbers. <br>
2. Boolean Conversions: 
    * __strconv.FormatBool__ for converting boolean values to strings ("true" or "false").
    * __strconv.ParseBool__ for parsing strings into boolean values. <br>
3. Base Conversions:
    * __strconv.FormatInt__ and __strconv.FormatUint__ for converting integers to strings with a specified base (e.g., decimal, binary, hexadecimal).
    * __strconv.ParseInt__ and __strconv.ParseUint__ for parsing strings in various bases into integers. <br>

# Day 2 Operators
Fungsi-fungsi lain yang sering dipakai dari paket math di Go meliputi:
1. math.Floor: Fungsi ini digunakan untuk melakukan pembulatan ke bawah (rounding down) dari sebuah angka desimal (float64) ke bilangan bulat terdekat yang lebih kecil atau sama.
2. math.Round: Fungsi ini digunakan untuk melakukan pembulatan ke bilangan bulat terdekat dari sebuah angka desimal (float64), dengan pembulatan ke atas jika bagian desimal lebih dari atau sama dengan 0.5.
3. math.Abs: Fungsi ini mengembalikan nilai absolut (nilai positif) dari sebuah angka desimal (float64).
4. math.Max dan math.Min: Fungsi-fungsi ini digunakan untuk mencari nilai maksimum dan minimum dari dua angka desimal.
5. math.Pow: Fungsi ini digunakan untuk menghitung nilai dari pangkat dari suatu angka.
6. math.Sqrt: Fungsi ini digunakan untuk menghitung akar kuadrat dari suatu angka desimal.
7. math.Sin, math.Cos, math.Tan: Fungsi-fungsi ini digunakan untuk menghitung sinus, kosinus, dan tangen dari suatu sudut dalam radian.
8. math.Log dan math.Exp: Fungsi-fungsi ini digunakan untuk menghitung logaritma alami dan eksponensial dari suatu angka.
9. math.Mod: Fungsi ini mengembalikan sisa dari pembagian dua angka desimal.
10. math.Ceil digunakan untuk melakukan pembulatan ke atas (rounding up) dari sebuah angka desimal (float64) ke bilangan bulat terdekat yang lebih besar atau sama.
