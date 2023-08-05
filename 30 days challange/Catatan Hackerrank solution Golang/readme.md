# Catatan 30 days code challange
# Day 0 Hello World
scanner := bufio.NewScanner(os.Stdin) <br>
var datainput string <br>
scanner.Scan() <br>
datainput = scanner.Text() <br>
<br>
Contoh penggunaan umum __bufio__ dalam Go termasuk: <br>

1. Membaca dan Menulis dari/ke File: Anda dapat menggunakan bufio untuk membaca dan menulis data dari dan ke file dengan buffering. Ini dapat meningkatkan efisiensi operasi I/O. <br>
2. Membaca dari Input Standar (stdin): Anda dapat menggunakan bufio untuk membaca input dari pengguna melalui terminal. <br>
3. Menulis ke Output Standar (stdout): Anda dapat menggunakan bufio untuk menulis data ke terminal atau output standar lainnya. <br>
4. Membaca dan Menulis dari/ke Koneksi Jaringan: Jika Anda mengembangkan program yang berkomunikasi melalui jaringan, Anda bisa menggunakan bufio untuk membaca dan menulis data dari dan ke koneksi jaringan. <br>

# Day 1 Data Type
Here are some of the common tasks that the __strconv__ package can help you with: <br>

1. Converting Numbers to Strings and Vice Versa: <br>
* __strconv.Itoa__ and __strconv.FormatInt__ for converting integers to strings.
* __strconv.FormatFloat__ for converting floating-point numbers to strings.
* __strconv.Atoi__ and __strconv.ParseInt__ for parsing strings into integers.
* __strconv.ParseFloat__ for parsing strings into floating-point numbers.
2. Boolean Conversions:

strconv.FormatBool for converting boolean values to strings ("true" or "false").
strconv.ParseBool for parsing strings into boolean values.
Base Conversions:

strconv.FormatInt and strconv.FormatUint for converting integers to strings with a specified base (e.g., decimal, binary, hexadecimal).
strconv.ParseInt and strconv.ParseUint for parsing strings in various bases into integers.
Error Handling:

The functions in the strconv package return an error as the second return value if the conversion fails. You should always check this error value to handle conversion errors gracefully.Here are some of the common tasks that the strconv package can help you with:

Converting Numbers to Strings and Vice Versa:

strconv.Itoa and strconv.FormatInt for converting integers to strings.
strconv.FormatFloat for converting floating-point numbers to strings.
strconv.Atoi and strconv.ParseInt for parsing strings into integers.
strconv.ParseFloat for parsing strings into floating-point numbers.
Boolean Conversions:

strconv.FormatBool for converting boolean values to strings ("true" or "false").
strconv.ParseBool for parsing strings into boolean values.
Base Conversions:

strconv.FormatInt and strconv.FormatUint for converting integers to strings with a specified base (e.g., decimal, binary, hexadecimal).
strconv.ParseInt and strconv.ParseUint for parsing strings in various bases into integers.
Error Handling:

The functions in the strconv package return an error as the second return value if the conversion fails. You should always check this error value to handle conversion errors gracefully.

