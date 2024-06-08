# phishing-website-detection-content-based
Ini adalah Tugas Akhir Kelas Pembelajaran Mesin End-to-End yang berfokus pada situs web phishing untuk mengklasifikasikan situs web phishing dan situs web sah. Secara khusus, saya berfokus pada fitur berbasis konten seperti fitur berdasarkan tag HTML. 


## inputs
- csv files
  - verified_online.csv --> URL situs web phishing dari phishtank.org
  - tranco_list.csv --> URL situs web sah dari tranco-list.eu
  
## general flowGunakan file CSV untuk mendapatkan URL
- Kirim permintaan ke setiap URL dan terima respons menggunakan pustaka requests Python
- Gunakan konten respons dan analisis dengan modul BeautifulSoup
- Ekstrak fitur dan buat vektor yang berisi nilai numerik untuk setiap fitur
- Ulangi proses ekstraksi fitur untuk semua konten/situs web dan buat dataframe terstruktur
- Tambahkan label pada akhir dataframe | 1 untuk phishing, 0 untuk sah
- Simpan dataframe sebagai file CSV dan file structured_data sudah siap!
  - Periksa file “structured_data_legitimate.csv” dan “structured_data_phishing.csv”.
- Setelah mendapatkan data terstruktur, Anda dapat menggabungkannya dan menggunakannya sebagai data latih dan uji
- Membagi data sebagai data latih dan uji seperti pada bagian pertama file machine_learning.py, atau Anda  dapat mengimplementasikan validasi silang K-fold seperti pada bagian kedua file yang sama. Saya mengimplementasikan K-fold dengan K=5.
 - Kemudian saya mengimplementasikan lima model ML berbeda:
 - Mesin Vektor Pendukung (SVM)
 - Naive Bayes Gaussian
 - Pohon Keputusan
 - Hutan Acak
 - AdaBoost
- Dapat memperoleh confusion matrix dan metrik kinerja: akurasi, presisi, recall
Akhirnya, saya memvisualisasikan metrik kinerja untuk semua model.
Naive Bayes adalah yang terbaik untuk kasus saya.
