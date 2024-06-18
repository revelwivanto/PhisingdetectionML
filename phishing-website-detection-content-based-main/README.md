# pemfilteran-website-pemfilteran-content-based
Ini adalah proyek ML yang berfokus pada pemfilteran website pemfilteran. Khususnya, saya fokus pada fitur-fitur berdasarkan HTML tag seperti fitur-fitur konten. Anda dapat menemukan proses ekstraksi fitur, pengumpulan, dan persiapan data di sini. Juga, membangun model ML, menilai mereka tersedia di sini.

## input
- file csv dari URL website yang pemfilteran dan yang sehat
  - verified_online.csv --> URL website pemfilteran dari phishtank.org
  - tranco_list.csv --> URL website yang sehat dari tranco-list.eu
  
## alur umum
- Gunakan file csv untuk mendapatkan URL
- Kirim permintaan ke setiap URL dan menerima respons dengan menggunakan modul requests dari python
- Gunakan konten respons dan melalui modul BeautifulSoup untuk membaca konten
- Ekstraksi fitur dan membuat vektor yang berisi nilai numerik untuk setiap fitur
- Ulangi proses ekstraksi fitur untuk semua konten/website dan buat dataframe yang terstruktur
- Tambahkan label di akhir untuk dataframe | 1 untuk pemfilteran 0 untuk sehat
- Simpan dataframe sebagai file csv dan file structured_data sudah siap!
  - Periksa file "structured_data_sehat.csv" dan "structured_data_pemfilteran.csv"
- Setelah mendapatkan data yang terstruktur, Anda dapat menggabungkan keduanya dan menggunakannya sebagai data latih dan uji
- Anda dapat memisahkan data sebagai latih dan uji seperti dalam file machine_learning.py bagian pertama, atau Anda dapat menerapkan K-fold cross-validation seperti dalam bagian kedua dari file yang sama. Saya menerapkan K = 5.
- Setelah itu, saya menerapkan beberapa model ML:
  - Support Vector Machine
  - Naive Bayes Gaussian
  - Pohon Keputusan
  - Hutan Ramuan
  - AdaBoost
- Anda dapat mendapatkan matriks konfusi dan metrik performa: akurasi, presisi, recall
