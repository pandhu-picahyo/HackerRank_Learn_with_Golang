# Alur berpikir logika dalam penyelesaian Hackerrank Challange

# Day 0: Hello WOrld
1. butuh input dari user melalui __scanner__
2. menggunakan __bufio__ untuk scanner
3. bentuknya __bufio.NewScanner__
4. input masuk di __os__ melalui stdin
5. karena __scanner__ adalah variabel baru, maka menggunakan __:=__
6. __scanner.Scan__ merupakan saat user input data
7. karena datainput sudah di deklarasikan di __var__ maka nanti cukup pakai samadengan __=__
8. Karena input berupa string atau text, maka menggunakan __scanner.Text__

# Day 1: Data Type
1. deklarasi variabel nilai1, nilai2, nilai3
2. input data melalui scanner.Scan
3. menggunakan __nilai1, _ = strconv.ParseUint(scanner.Text(), 10, 64)__ untuk read dan save file scanner<br>
    karena data type berupa uint, maka digunakan ParseUint <br>
    Karena data dari scanner berupa text, dan ingin diubah ke bentuk data lain (utamanya angka) digunakan strconv <br>
    karena data scanner berupa text digunakan scanner.Text <br>
    angka 10 merupakan base desimal <br>
    angka 64 merupakan bite size of resulting in uint64 <br>
    data hasil scanner masuk ke variabel nilai1 <br>
    sedangkan _ menyatakan tidak ada error, bentuk umumnya adalah nilai1, err := strconv.ParseUint(scanner.Text(), 10, 64) <br>
4. menggunakan __nilai2, _ = strconv.ParseFloat(scanner.Text(), 64)__ <br>
    Parsefloat karena data type adalah desimal <br>
    tidak perlu base jadi langsung menggunakan 64 untuk bite size variabel <br>
5. menggunakan __nilai3 = scanner.Text()__ <br>
    karena data yang ingin disimpan adalah text, sehingga tidak perlu menggunakan Parse.... langsung menggunakan scanner.Text
6. menggunakan format __fmt.Printf("%.1f\n", d+nilai2)__ <br>
    karena hasil outpun ingin 1 angka dibelakang koma (%.1f), karena itu menggunakan Printf

# Day2: Operator
1. pada __func solve(meal_cost float64, tip_percent int32, tax_percent int32)__ <br>
    variabel meal cost, tip percent dan tax percent sudah di deklarasikan di sini beserta data type nya <br>
    sehingga tidak perlu dibuatkan lagi variabel baru <br>
2. pada __tip := meal_cost * (float64(tip_percent) / 100.0)__ <br>
    variabel tip dan juga variabel tak di baris selanjutnya menggunakan tanda := karena variabel baru dan ingin memasukkan nilai ke variabel baru <br>
    sebelah kanan variabel adalah rumus perhitungan yang disediakan di soal <br>
3. pada __decimal_part := total_cost - float64(int32(total_cost))__ <br>
    dicari nilai pecahan 0,...., digunakan untuk mencari nilai pembulatan, karena di sample test butuh bilangan bulat <br>
4. pada __rounded_up := math.Ceil(total_cost)__ <br>
    fungsi untuk pembulatan ke atas, untuk 0.5 ke atas <br>
5. pada __fmt.Println(int32(total_cost))__ <br>
    digunakan untuk pembulatan ke bawah

   
